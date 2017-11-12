package main

import (
	"net/http"
	"time"
	"html/template"
	"log"
)

type HomePageStruct struct{
	Date string
	Time string
}


func HomePageFunction(w http.ResponseWriter, r *http.Request){
	now := time.Now()
	homePage := HomePageStruct{
		Date:now.Format("01-01-1990"),
		Time:now.Format("01:01:01"),
	}
	log.Println(now)

	t,err := template.ParseFiles("/Users/egozi/GoglandProjects/own-server/html/homePage.html")
	if err != nil {
		log.Println("template parsing error: ", err)
	}

	err = t.Execute(w,homePage)
	if err != nil {
		log.Println("executaion error: ", err)
	}

}

func main() {
	http.HandleFunc("/",HomePageFunction)
	http.ListenAndServe(":50001",nil)
}