package main

import (
	"embed"
	_ "embed"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
)

//go:embed resources/views/*.gohtml
var templates embed.FS
var myTemplate = template.Must(template.ParseFS(templates, "resources/views/*.gohtml"))

func main() {
	router := httprouter.New()

	router.GET("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		_, err := fmt.Fprint(writer, "Hello world")
		if err != nil {
			panic(err)
		}
	})

	router.GET("/login", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		err := myTemplate.ExecuteTemplate(writer, "login.gohtml", map[string]interface{}{
			"Title": "Login Page",
		})
		if err != nil {
			panic(err)
		}
	})

	router.GET("/register", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		err := myTemplate.ExecuteTemplate(writer, "register.gohtml", map[string]interface{}{
			"Title": "Register Page",
		})
		if err != nil {
			panic(err)
		}

	})

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
