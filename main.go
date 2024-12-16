package main

import (
	"dip/cmd"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func main() {
	log.Print("dip is starting \n")
	ginEngine := gin.Default()
	cmd.AddHttpRouter(ginEngine)
	svc := &http.Server{
		Addr:           ":8080",
		Handler:        ginEngine,
		ReadTimeout:    15 * time.Second,
		WriteTimeout:   15 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Printf("dip started listen in %s \n", 8080)
	err := svc.ListenAndServe()
	if err != nil {
		log.Printf("dip running err %s \n", err.Error())
	}
}
