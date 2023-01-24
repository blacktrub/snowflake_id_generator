package main

import (
	"net/http"
	"snowflake_id_generator/internal/snowflake"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("id", func(ctx *gin.Context) {
		sn := snowflake.Snowflake{}
		ctx.JSON(http.StatusOK, gin.H{"id": sn.Get()})
	})

	r.Run()
}
