# GRRT (Go Request RouTer)

GRRT (Go Request RouTer) is a direct replacement for the archived gorilla/mux.
It has built-in CORS and Method based routing.


## Replaces gorilla/mux with one line of code

[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=GolangToolKits_grrt&metric=alert_status)](https://sonarcloud.io/dashboard?id=GolangToolKits_grrt)
[![CircleCI](https://circleci.com/gh/GolangToolKits/grrt.svg?style=svg)](https://circleci.com/gh/GolangToolKits/grrt)
[![Go Report Card](https://goreportcard.com/badge/github.com/GolangToolKits/grrt)](https://goreportcard.com/report/github.com/GolangToolKits/grrt)


## Features

- Request Routing
- Method Based Routing
- CORS

#### [Full REST Service Example Project](https://github.com/GolangToolKits/grrtRouterRestExample)

#### [Full Web Example Project](https://github.com/GolangToolKits/grrtRouterWebSiteExample)

Package `GolangToolKits/grrt` implements a request router and dispatcher for handling incoming requests to their associated handler.

The name mux stands for "HTTP request multiplexer". Like the standard `http.ServeMux`, `grrt.Router` matches incoming requests against a list of registered routes and calls a handler for the route that matches the URL. The main features are:


* Replaces gorilla/mux with one line of code
* It implements the `http.Handler` interface so it is compatible with the standard `http.ServeMux`.
* URL hosts, paths and query values can be handled.
* Path variable can be used instead of query parameters with ease.
* Method base routing is easy
* CORS is built in with no need for additional modules.


---

* [Install](#install)
* [Examples REST Service](#RestExample)
* [Examples Web Server](#WebExample)


---


## Install


```sh
go get -u github.com/GolangToolKits/grrt

```


## RestExample

#### [REST Service Full Example](https://github.com/GolangToolKits/grrtRouterRestExample)

```go
import(

    "fmt"
    "net/http"
    "os"
    "strconv"
    ph "github.com/GolangToolKits/grrtRouterRestExample/handlers"
    mux "github.com/GolangToolKits/grrt"
)


func main() {

	var sh ph.StoreHandler //see the example project for the full code

	h := sh.New()

	router := mux.NewRouter()
	router.EnableCORS()
	router.CORSAllowCredentials()
	router.SetCorsAllowedHeaders("X-Requested-With, Content-Type, api-key, customer-key, Origin")
	router.SetCorsAllowedOrigins("*")
	router.SetCorsAllowedMethods("GET, DELETE, POST, PUT")

	port := "3000"
	envPort := os.Getenv("PORT")
	if envPort != "" {
		portInt, _ := strconv.Atoi(envPort)
		if portInt != 0 {
			port = envPort
		}
	}

	router.HandleFunc("/rs/product/get/{id}", h.GetProduct).Methods("GET")
	router.HandleFunc("/rs/product/get/{id}/{sku}", h.GetProductWithIDAndSku).Methods("GET")
	router.HandleFunc("/rs/products", h.GetProducts).Methods("GET")
	router.HandleFunc("/rs/product/add", h.AddProduct).Methods("POST")
	router.HandleFunc("/rs/product/update", h.UpdateProduct).Methods("PUT")
	fmt.Println("running on Port:", port)
	http.ListenAndServe(":"+port, (router))

}
```


## WebExample

#### [Web Full Example](https://github.com/GolangToolKits/grrtRouterWebSiteExample)

```go

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"

	mux "github.com/GolangToolKits/grrt"
	hd "github.com/GolangToolKits/grrtRouterWebSiteExample/handlers"
)

func main() {

	var sh hd.SiteHandler

	sh.Templates = template.Must(template.ParseFiles("./static/index.html",
		"./static/product.html", "./static/addProduct.html"))

	router := mux.NewRouter()

	h := sh.New()

	router.HandleFunc("/", h.Index).Methods("GET")
	router.HandleFunc("/product/{id}/{sku}", h.ViewProduct).Methods("GET")
	router.HandleFunc("/addProduct", h.AddProduct).Methods("POST")

	port := "8080"
	envPort := os.Getenv("PORT")
	if envPort != "" {
		portInt, _ := strconv.Atoi(envPort)
		if portInt != 0 {
			port = envPort
		}
	}

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	fmt.Println("Web UI is running on port 8080!")

	http.ListenAndServe(":"+port, (router))
}

```