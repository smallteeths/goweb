package main

import (
	"tool-backend/config"
	"github.com/gin-gonic/gin"
	"log"
	"tool-backend/router"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

func main() {
	if err := config.Init();err != nil {
		panic(err)
	}
	//g := gin.Default()

	// Set gin mode.
	// gin.SetMode(viper.GetString("runmode"))

	// Create the Gin engine.
	g := gin.New()

	router.InitRouter(g)
	log.Printf("Start to listening the incoming requests on http address: %s\n", viper.GetString("addr"))
	//log.Println(http.ListenAndServe(viper.GetString("addr"), g).Error())
	// if err := g.Run(viper.GetString("addr"));err != nil {log.Fatal("ListenAndServe:", err)}
	s := &http.Server{
		Addr:           viper.GetString("addr"),
		Handler:        g,
		ReadTimeout:    200 * time.Second,
		WriteTimeout:   500 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()

}
