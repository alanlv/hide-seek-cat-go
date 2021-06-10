package main

import (
	"HideSeekCatGo/router"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

/**
躲猫猫后端Go
@author yinlei
*/
func main() {
	// create the gin engine.
	g := gin.New()

	// gin middlewares.
	middlewares := []gin.HandlerFunc{}

	// routes.
	router.LoadRouter(
		g,
		middlewares...,
	)

	// ping the server to make sure the router is working.
	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("The router has no response, or it might took long to start up.", err)
		}
		log.Print("The router has been deployed successfully.")
	}()

	log.Printf("Start to listening the incoming requests on http address: %s", ":8080")
	log.Printf(http.ListenAndServe(":8080", g).Error())

}

// pings the http server to make sure the router is working.
func pingServer() error {
	for i := 0; i < 2; i++ {
		// ping the server by sending a GET request to `/health`.
		resp, err := http.Get("http://127.0.0.1:8080" + "/sd/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}
		log.Print("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}
	return errors.New("Cannot connect to the router.")
}
