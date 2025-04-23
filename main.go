package main

import (
	"shopping/router"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	router.RegisterApiRouter(r)

	r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
