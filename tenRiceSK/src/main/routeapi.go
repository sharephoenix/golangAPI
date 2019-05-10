package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"strings"
)

func usersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Print("-------")
	pathInfo := strings.Trim(r.URL.Path, "/")
	parts := strings.Split(pathInfo, "/")
	var action = ""
	if len(parts) > 1 {
		action = strings.Title(parts[1]) + "Action"
	}

	ajax := &usersController{}
	controller := reflect.ValueOf(ajax)
	method := controller.MethodByName(action)
	if !method.IsValid() {
		method = controller.MethodByName(strings.Title("users") + "Action")
	}
	requestValue := reflect.ValueOf(r)
	responseValue := reflect.ValueOf(w)
	method.Call([]reflect.Value{responseValue, requestValue})
}

func postHander(w http.ResponseWriter, r *http.Request) {
	// 表单中的数据
	pathInfo := strings.Trim(r.URL.Path, "/")
	r.ParseForm()
	for k, v := range r.PostForm {
		fmt.Printf("k:%v\n", k)
		fmt.Printf("v:%v\n", v)
	}

	// 解析 body 中的数据
	bodyss, _ := ioutil.ReadAll(r.Body)
	fmt.Print(bodyss)
	var receiver map[string]interface{}
	json.Unmarshal(bodyss, &receiver)
	log.Println(receiver["token"])

	parts := strings.Split(pathInfo, "/")
	w.Header()
	fmt.Print(parts)
	var action = ""
	if len(parts) > 1 {
		action = strings.Title(parts[1]) + "Action"
	}

	r.ParseForm()
	//_, found1 := r.Form["token"]
	//if !(found1) {
	//	fmt.Println("非法访问aaal")
	//	return
	//}

	ajax := &usersController{}
	controller := reflect.ValueOf(ajax)
	method := controller.MethodByName(action)
	if !method.IsValid() {
		method = controller.MethodByName("UsersAction")
	}
	requestValue := reflect.ValueOf(r)
	responseValue := reflect.ValueOf(w)
	method.Call([]reflect.Value{responseValue, requestValue})
}

func putHander(w http.ResponseWriter, r *http.Request) {
	pathInfo := strings.Trim(r.URL.Path, "/")
	r.ParseForm()
	for k, v := range r.PostForm {
		fmt.Printf("k:%v\n", k)
		fmt.Printf("v:%v\n", v)
	}

	parts := strings.Split(pathInfo, "/")
	w.Header()
	fmt.Print(parts)
	var action = ""
	if len(parts) > 1 {
		action = strings.Title(parts[1]) + "Action"
	}

	r.ParseForm()
	//_, found1 := r.Form["token"]
	//if !(found1) {
	//	fmt.Println("非法访问aaal")
	//	return
	//}

	ajax := &usersController{}
	controller := reflect.ValueOf(ajax)
	method := controller.MethodByName(action)
	if !method.IsValid() {
		method = controller.MethodByName("UsersAction")
	}
	requestValue := reflect.ValueOf(r)
	responseValue := reflect.ValueOf(w)
	method.Call([]reflect.Value{responseValue, requestValue})
	}


func deleteHander(w http.ResponseWriter, r *http.Request) {
	pathInfo := strings.Trim(r.URL.Path, "/")
	r.ParseForm()
	for k, v := range r.PostForm {
		fmt.Printf("k:%v\n", k)
		fmt.Printf("v:%v\n", v)
	}

	parts := strings.Split(pathInfo, "/")
	w.Header()
	fmt.Print(parts)
	var action = ""
	if len(parts) > 1 {
		action = strings.Title(parts[1]) + "Action"
	}

	r.ParseForm()
	//_, found1 := r.Form["token"]
	//if !(found1) {s
	//	fmt.Println("非法访问aaal")
	//	return
	//}

	ajax := &usersController{}
	controller := reflect.ValueOf(ajax)
	method := controller.MethodByName(action)
	if !method.IsValid() {
		method = controller.MethodByName("UsersAction")
	}
	requestValue := reflect.ValueOf(r)
	responseValue := reflect.ValueOf(w)
	method.Call([]reflect.Value{responseValue, requestValue})
}