package main

import "net/http"

func main() {
	/* WEB */
	http.HandleFunc("/login/", loginHandler)
	http.HandleFunc("/build/", buildHandler)
	http.HandleFunc("/admin/", adminHandler)
	http.HandleFunc("/ajax/", ajaxHandler)

	/* API */
	http.HandleFunc("/users", usersHandler)
	http.HandleFunc("/post", postHander)
	http.HandleFunc("/put", putHander)
	http.HandleFunc("/delete", deleteHander)

	http.ListenAndServe(":1111", nil)
}
