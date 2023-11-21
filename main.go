package main

import (
	"log"
	"os"

	c "github.com/PyMarcus/codes_download/constants"
	s "github.com/PyMarcus/codes_download/service"
)

func main() {
	log.Println(c.YELLOW + "Starting ...")

	os.MkdirAll("data", os.ModePerm)
	os.MkdirAll("json", os.ModePerm)

	args := os.Args[1:]
	if len(args) < 2{
		log.Println(c.YELLOW + "Missing arguments! Use ./main.go [language] [true/false]" + c.RESET)
	}

	if args[1] == "true"{
		repository := s.NewRepository(args[0], true)
		repository.StartDownloads()
	}else if args[1] == "false"{
		repository := s.NewRepository(args[0], false)
		repository.StartDownloads()
	}else{
		log.Println(c.YELLOW + "Missing arguments! Use ./main.go [language] [true/false]" + c.RESET)
	}
	
}
