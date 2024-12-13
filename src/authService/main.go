package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	limit "github.com/aviddiviner/gin-limit"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.Use(limit.MaxAllowed(200))
	router.Use(cors.Default())
	startServer(router)
}

func startServer(router http.Handler) {
	s := &http.Server{
		Addr:           ":6060",
		Handler:        router,
		ReadTimeout:    18000 * time.Second,
		WriteTimeout:   18000 * time.Second,
		MaxHeaderBytes: 1 << 20, // 1 MB
	}

	fmt.Print("Server running on http://127.0.0.1:6060")

	if err := s.ListenAndServe(); err != nil {
		_ = fmt.Errorf("fatal error description: %s", strings.ToLower(err.Error()))
		panic(err)
	}

}
