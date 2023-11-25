package tools

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"github.com/PyMarcus/codes_download/schema"
)


func ReadJsonFile(filepath string) ([]schema.TbTopics, []schema.TbOwner, []schema.TbLicense, []schema.TbItems, schema.TbImports) {
	var items []schema.TbItems
	var licenses []schema.TbLicense
	var owners []schema.TbOwner
	var topics []schema.TbTopics

	jsonData, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	var imp schema.TbImports
	err = json.Unmarshal(jsonData, &imp)
	if err != nil {
		log.Println("fail to decode tbimports ", err)
	}
	log.Println("running ", imp.TotalCount, " items.")
	for i := 0; i < len(imp.Items); i++ {
	    
		var item schema.TbItems
		itemBytes, err := json.Marshal(imp.Items[i])
		if err != nil {
			log.Fatal("fail to convert item to bytes: ", err)
		}
		err = json.Unmarshal(itemBytes, &item)
		if err != nil {
			log.Fatal("fail to decode tbitem ", err)
		}
	
		items = append(items, item)
			
	
		var license schema.TbLicense
		licenseBytes, err := json.Marshal(imp.Items[i].License)
		if err != nil {
			log.Fatal("fail to convert license to bytes: ", err)
		}
		err = json.Unmarshal(licenseBytes, &license)
		if err != nil {
			log.Fatal("fail to decode tblicense ", err)
		}
	
		licenses = append(licenses, license)
		
	
		var owner schema.TbOwner

		ownerBytes, err := json.Marshal(imp.Items[i].Owner)
		if err != nil {
			log.Fatal("fail to convert owner to bytes: ", err)
		}
		err = json.Unmarshal(ownerBytes, &owner)
		if err != nil {
			log.Fatal("fail to decode tbowner ", err)
		}
	
		owners = append(owners, owner)
		
		
		if len(imp.Items[i].Topics) > 0 {
			var topic schema.TbTopics

			for _, topicName := range imp.Items[i].Topics {
				topic = schema.TbTopics{
					IdItem:    int64(imp.Items[i].ID),
					TopicName: topicName,
				}
				topics = append(topics, topic)
			}
		}
		
	}
	
	log.Println("To insert ", len(topics), " topics")
	log.Println("To insert ", len(owners), " owners")
	log.Println("To insert ", len(items), " items")
	log.Println("To insert ", len(licenses), " licenses")
	return topics, owners, licenses, items, imp
}
