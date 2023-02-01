# GRRT (Go Request RouTer)

GRRT (Go Request RouTer) is a direct replacement for the archived gorilla/mux.
It has built-in CORS and Method based routing.


[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=GolangToolKits_grrt&metric=alert_status)](https://sonarcloud.io/dashboard?id=GolangToolKits_grrt)
[![CircleCI](https://circleci.com/gh/GolangToolKits/grrt.svg?style=svg)](https://circleci.com/gh/GolangToolKits/grrt)
[![Go Report Card](https://goreportcard.com/badge/github.com/GolangToolKits/grrt)](https://goreportcard.com/report/github.com/GolangToolKits/grrt)


## Features

- Request Routing
- Method Based Routing
- CORS

#### [REST Service Example](https://github.com/GolangToolKits/grrtRouterRestExample)

Package `GolangToolKits/grrt` implements a request router and dispatcher for handling incoming requests to their associated handler.

The name mux stands for "HTTP request multiplexer". Like the standard `http.ServeMux`, `grrt.Router` matches incoming requests against a list of registered routes and calls a handler for the route that matches the URL. The main features are:

* It implements the `http.Handler` interface so it is compatible with the standard `http.ServeMux`.
* URL hosts, paths and query values can be handled.
* Path variable can be used instead of query parameters.
* Method base routing is easy
* CORS is built in with no need for additional modules.


