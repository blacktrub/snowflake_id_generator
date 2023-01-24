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
		id, err := sn.Get()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"id": id})
	})

	r.Run()
}
