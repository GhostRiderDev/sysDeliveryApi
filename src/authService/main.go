package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	limit "github.com/aviddiviner/gin-limit"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.Use(limit.MaxAllowed(200))
	router.Use(cors.Default())
	startServer(router)
}

func startServer(router http.Handler) {
	viper.SetConfigFile("config.json")

	if err := viper.ReadInConfig(); err != nil {
		_ = fmt.Errorf("fatal error in config file: %s", err.Error())
		panic(err)
	}

	serverPort := fmt.Sprintf(":%s", viper.GetString("ServerPort"))

	s := &http.Server{
		Addr:           serverPort,

		Handler:        router,
		ReadTimeout:    18000 * time.Second,
		WriteTimeout:   18000 * time.Second,
		MaxHeaderBytes: 1 << 20, // 1 MB
	}

	fmt.Printf("Server running on http:127.0.0.1%s", serverPort)


	if err := s.ListenAndServe(); err != nil {
		_ = fmt.Errorf("fatal error description: %s", strings.ToLower(err.Error()))
		panic(err)
	}

}
