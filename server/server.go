package main

import (
	"net/http"
	"html/template"
	"log"
)

type Button struct{
	Name     		string
	TextView 		string
	Value 	 		string
	IsDisabled  	bool
}

type PageVars struct{
	SelectButtons  []Button
	Title 		   string
	Response	   string
}

func DisplayFunction(w http.ResponseWriter, r *http.Request){
	selects := []Button{
		Button{
			Name:"animal",
			TextView:"Choose Dog",
			Value:"Dog",
		},
		Button{
			Name:"animal",
			TextView:"Choose Cat",
			Value:"Cat",
		},
	}

	pageVars := PageVars{
		SelectButtons:selects,
		Title:"Selection",
	}
	t,err := template.ParseFiles("/Users/egozi/GoglandProjects/own-server/html/homePage.html")
	if err != nil {
		log.Println("template parsing error: ", err)
	}
	err = t.Execute(w,pageVars)
	if err != nil {
		log.Println("executaion error: ", err)
	}

}

func DisplayResponse(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	chosen := r.Form.Get("animal")

	pageVars := PageVars{
		Title:"Your Selection is:",
		Response:chosen,
	}

	t,err := template.ParseFiles("/Users/egozi/GoglandProjects/own-server/html/homePage.html")
	if err != nil {
		log.Println("template parsing error - ",err)
	}
	err = t.Execute(w,pageVars)
	if err != nil {
		log.Println("execution error", err)
	}
}

func main() {
	http.HandleFunc("/",DisplayFunction)
	http.HandleFunc("/silence",DisplayResponse)
	http.ListenAndServe(":7785",nil)
}