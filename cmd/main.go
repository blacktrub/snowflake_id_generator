package main

import (
	"fmt"
	"net/http"
	"os"
	"snowflake_id_generator/internal/snowflake"

	"github.com/gin-gonic/gin"
)

func main() {
	sn, err := snowflake.New()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Wrong configuration: %v\n", err.Error())
		os.Exit(1)
		return
	}

	r := gin.Default()
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
