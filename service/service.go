package service

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"

	c "github.com/PyMarcus/codes_download/constants"
	tools "github.com/PyMarcus/codes_download/tools"
	rep "github.com/PyMarcus/codes_download/repository"

)

var wg sync.WaitGroup
var wgdb sync.WaitGroup

/*
	Repository receives Language: python,

go, ruby, c++ etc. OrderByStars: false
*/
type Repository struct {
	Language     string `json:"language"`
	OrderByStars bool   `json:"order_by_stars"`
	page         int
	perPage      int
	httpClient   *http.Client
}

// NewRepository create a new object with info to request
func NewRepository(language string, orderByStars bool) *Repository {
	return &Repository{
		Language:     language,
		OrderByStars: orderByStars,
		page:         1,
		perPage:      100,
		httpClient:   &http.Client{},
	}
}

func (r Repository) getUrl(page, perPage int) string {
	if r.OrderByStars {
		return fmt.Sprintf("https://api.github.com/search/repositories?q=language:%s&sort=stars&order=desc&page=%d&per_page=%d", r.Language, page, perPage)
	}
	return fmt.Sprintf("https://api.github.com/search/repositories?q=language:%s&sort=stars&order=asc&page=%d&per_page=%d", r.Language, page, perPage)
}

func (r Repository) createRequest() (*http.Request, error) {
	url := r.getUrl(r.page, r.perPage)

	log.Println(c.GREEN + url + c.RESET)

	request, err := http.NewRequest(c.GET_METHOD, url, nil)

	if err != nil {
		log.Println(fmt.Sprintf("%s - Error %v - URL: %s", c.RED, err, url))
		return nil, err
	}

	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3")
	request.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	return request, nil
}

func (r Repository) fetchGet() *http.Response {
	request, error := r.createRequest()

	if error != nil {
		log.Fatal(nil)
	}

	response, err := r.httpClient.Do(request)

	if err != nil {
		log.Fatal(err)
	}

	if response.StatusCode != http.StatusOK {
		log.Fatal("FAIL TO GET API RESPONSE")
	}

	return response
}

// insert data into database
func (r Repository) insertIntoDatabase(jsonPath string){
	defer wgdb.Done()
	
	rep.Insert(jsonPath)
	
}

func (r *Repository) fetchData() {
	data2 := r.fetchGet()
	r.saveJsonFile(data2)
	
	wgdb.Add(1)
	go r.insertIntoDatabase("json/" + r.Language + ".json")
	
	for {
		data := r.fetchGet()
		
		defer data.Body.Close()

		var result map[string]interface{}

		err := json.NewDecoder(data.Body).Decode(&result)
		

		if err != nil {
			log.Fatal("Fail to decode json response from API")
		}
		items, _ := result["items"].([]interface{})
		total := result["total_count"].(float64)

		for _, item := range items {
			repo := item.(map[string]interface{})
			owner := strings.Split(repo["full_name"].(string), "/")[0]
			log.Println("\n\n", c.GREEN)
			log.Printf("Owner: %s\nRepository: %s\nDescription: %s\nURL: %s\n\n", owner, repo["name"], repo["description"], repo["html_url"])
			log.Println(c.RESET)
			wg.Add(1)
			r.codeDownloadLikeZip(owner, repo["full_name"].(string))
		}
		
		wg.Wait()
		r.page++

		if r.page*r.perPage >= int(total) {
			break
		}
	}
	
	wgdb.Wait()
}

func (r Repository) codeDownloadLikeZip(owner string, repoFullName string) {
	defer wg.Done()

	url := fmt.Sprintf("https://github.com/%s/archive/%s.zip", repoFullName, "main")

	log.Println("FETCH ", url)

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(c.RED+"Fail to create request:", err)
		return
	}

	request.Header.Add("Authorization", "token "+tools.GetGithubWebToken())

	response, err := r.httpClient.Do(request)

	if err != nil {
		log.Println(c.RED+"Fail to get download response:", err)
		return
	}

	defer response.Body.Close()
	
	if response.StatusCode == http.StatusSeeOther || response.StatusCode == http.StatusFound {
		redirectURL := response.Header.Get("Location")
		log.Println("Redirected to:", redirectURL)

		// Fazer uma nova solicitação para a URL redirecionada
		response, err = r.httpClient.Get(redirectURL)
		if err != nil {
			log.Println(c.RED + "Fail to get redirected response:", err)
			return
		}
	}

	if response.StatusCode != http.StatusOK {
		log.Println(c.RED+" ERR status code ", response.StatusCode)
		if response.StatusCode == http.StatusNotFound{
			url = fmt.Sprintf("https://github.com/%s/archive/%s.zip", repoFullName, "master")

			log.Println(c.YELLOW + "RETRY FETCH ", url)

			request, err := http.NewRequest("GET", url, nil)
			if err != nil {
				log.Println(c.RED+"Fail to create request:", err)
				return
			}

			request.Header.Add("Authorization", "token "+tools.GetGithubWebToken())

			response, err = r.httpClient.Do(request)

			if err != nil {
				log.Println(c.RED+"Fail to get download response:", err)
				return
			}
		}
	}

	fout, err := os.Create(fmt.Sprintf("data/%s.zip", strings.Split(repoFullName, "/")[1]))

	if err != nil {
		log.Println(c.RED+"Fail to save file", err)
		return
	}

	defer fout.Close()

	io.Copy(fout, response.Body)

	log.Println(c.GREEN+"[OK] to download "+url, c.RESET)

}

// StartDownloads get sync downloads from repository
func (r Repository) StartDownloads() {
	r.fetchData()
}

func (r Repository) saveJsonFile(response *http.Response){
	body, err := io.ReadAll(response.Body)
	if err != nil{
		log.Println("Fail to save json file")
		return 
	}
	
	err = ioutil.WriteFile("json/"+r.Language + ".json", body, 0644)
	if err != nil{
		log.Println("fail to write json file")
		return
	}
	
	defer response.Body.Close()
	
	log.Println("json file saved")
}
