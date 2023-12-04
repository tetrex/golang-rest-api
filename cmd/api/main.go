package main

import (
	"fmt"
	"golang-rest-api/api/router"
	"golang-rest-api/config"
	"io"
	"log"
	"net/http"
)

//	@title			MYAPP API
//	@version		1.0
//	@description	This is a sample RESTful API with a CRUD

//	@contact.name	Dumindu Madunuwan
//	@contact.url	https://learning-cloud-native-go.github.io

//	@license.name	MIT License
//	@license.url	https://github.com/learning-cloud-native-go/myapp/blob/master/LICENSE

//	@host		localhost:8080
//	@basePath	/v1
func main() {

	c := config.New()
	r := router.New()

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hello)

	s := &http.Server{
		Addr:         fmt.Sprintf(":%d", c.Server.Port),
		Handler:      r,
		ReadTimeout:  c.Server.TimeoutRead,
		WriteTimeout: c.Server.TimeoutWrite,
		IdleTimeout:  c.Server.TimeoutIdle,
	}

	log.Println("Starting server " + s.Addr)
	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("Server startup failed")
	}
}

func hello(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, world!")
}
