package main

import (
	"net/http"
	"log"
	"html/template"
)

type buildController struct {
}

func (this *buildController)IndexAction(w http.ResponseWriter, r *http.Request) {
	log.Println("react web build result")
	t, err := template.ParseFiles("template/html/build/index.html")
	if (err != nil) {
		log.Println(err)
		log.Println("buildHandler error")
	}else {
		log.Println("buildHandler success")
	}
	log.Println(w)
	t.Execute(w, nil)
}
