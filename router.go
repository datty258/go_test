package main

import (
	. "github.com/datty258/go_test/apis"
	"github.com/gin-gonic/gin"
)

func initRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", Index)

	router.POST("/person", AddPerson)

	router.GET("/persons", GetPersons)

	router.GET("/person/:id", GetPerson)

	router.PUT("/person/:id", UpdatePerson)

	router.DELETE("/person/:id", DelPerson)

	return router
}
