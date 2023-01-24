package main

import (
	"net/http"
	"snowflake_id_generator/internal/seq"
	"snowflake_id_generator/internal/snowflake"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	seq := seq.Init()
	sn := snowflake.Init(seq)
	r.GET("id", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"id": sn.Get()})
	})

	r.Run()
}
