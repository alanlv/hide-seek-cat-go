package main

import (
	"HideSeekCatGo/config"
	"HideSeekCatGo/model"
	"HideSeekCatGo/router"
	"HideSeekCatGo/router/middleware"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"time"
)

/**
躲猫猫后端Go
@author yinlei
*/

var (
	cfg = pflag.StringP("config", "c", "", "apiserver config file path.")
)

func main() {
	pflag.Parse()

	// init config.
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}

	// init db.
	model.DB.Init()
	defer model.DB.Close()

	// set gin mode.
	gin.SetMode(viper.GetString("runmode"))

	// create the gin engine.
	g := gin.New()

	// routes.
	router.LoadRouter(
		g,
		middleware.RequestId(),
	)

	// ping the server to make sure the router is working.
	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("The router has no response, or it might took long to start up.", err)
		}
		log.Print("The router has been deployed successfully.")
	}()

	log.Printf("Start to listening the incoming requests on http address: %s", viper.GetString("addr"))
	log.Printf(http.ListenAndServe(viper.GetString("addr"), g).Error())

}

// pings the http server to make sure the router is working.
func pingServer() error {
	for i := 0; i < viper.GetInt("max_ping_count"); i++ {
		// ping the server by sending a GET request to `/health`.
		resp, err := http.Get(viper.GetString("url") + "/sd/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}
		log.Print("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}
	return errors.New("Cannot connect to the router.")
}
