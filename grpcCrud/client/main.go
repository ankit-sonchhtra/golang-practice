package main

import (
	"go-tutorials/all-others/grpcCrud/proto"
	"log"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())

	if err != nil {
		panic(err)
	}

	client := proto.NewArithServiceClient(conn)

	g := gin.Default()
	g.GET("/add/:a/:b", func(ctx *gin.Context) {
		a, err := strconv.ParseUint(ctx.Param("a"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter A"})
			return
		}

		b, err := strconv.ParseUint(ctx.Param("b"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"errors": "Invalid parameter B"})
			return
		}

		request := &proto.Request{A: int64(a), B: int64(b)}
		response, err := client.Add(ctx, request)
		if err == nil {
			ctx.JSON(http.StatusOK, gin.H{"result": response.Result})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		}
	})

	g.GET("/mul/:a/:b", func(ctx *gin.Context) {
		a, err := strconv.ParseUint(ctx.Param("a"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter A"})
			return
		}

		b, err := strconv.ParseUint(ctx.Param("b"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"errors": "Invalid parameter B"})
			return
		}

		request := &proto.Request{A: int64(a), B: int64(b)}
		response, err := client.Multiply(ctx, request)
		if err == nil {
			ctx.JSON(http.StatusOK, gin.H{"result": response.Result})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		}
	})

	if err := g.Run(":8081"); err != nil {
		log.Fatalf("Failed to run on server: %v", err)
	}

}
