package main

import (
	"github.com/charlitoro/go-simple-api/src/handlers"
	"github.com/charlitoro/go-simple-api/src/middlewares"
	"github.com/charlitoro/go-simple-api/src/server"
)


func main() {
	server := server.NewServer(":4400")
	server.Handle("/", "GET", server.AddMiddleware(handlers.HandleRoot, middlewares.Logging(), middlewares.CheckAuth()))
	server.Handle("/create", "POST", handlers.PostRequest)
	server.Handle("/user", "POST", server.AddMiddleware(handlers.CreateRequest, middlewares.CheckCreateRequest()))
	server.Listen()
}
