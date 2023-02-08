package main

import (
	"fmt"
	"net/http"

	"github.com/GRTheory/management-system/service"
	mm "github.com/GRTheory/my-model"
	"github.com/gin-gonic/gin"
)

type CentralServer struct {
	Service service.ServiceInterface
}

func NewServer(username, password, host, port, database string) *CentralServer {
	server := &CentralServer{}
	client, err := mm.GetClient(username, password, host, port, database)
	if err != nil {
		panic(err)
	}
	server.Service = service.NewService(client)
	return server
}

func main() {
	username := "root"
	password := "pass_123"
	host := "192.168.228.128"
	port := "3306"
	database := "base"

	server := NewServer(username, password, host, port, database)

	r := gin.Default()

	r.GET("/teacher/teachers", func(ctx *gin.Context) {
		teachers, err := server.Service.GetAllTeachers(ctx)
		if err != nil {
			fmt.Println("failed getting teachers", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "failed getting teahcers",
			})
		}
		ctx.JSON(http.StatusOK, teachers)
		fmt.Println("successful")
	})

	if err := r.Run(":9876"); err != nil {
		panic(err)
	}
}
