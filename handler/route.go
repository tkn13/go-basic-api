package handler

import (
	"net/http"

	"github.com/ThaksinCK/go-basic-api.git/middlewere"
	adduser "github.com/ThaksinCK/go-basic-api.git/service/addUser"
	"github.com/ThaksinCK/go-basic-api.git/service/central"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type route struct {
	Name        string
	Description string
	Method      string
	Pattern     string
	Endpoint    gin.HandlerFunc
	Validtion   gin.HandlerFunc
}

type Routes struct {
	router         *gin.Engine
	centralService []route
	userService    []route
}

func (r Routes) InitRouter() http.Handler {
	centralEndpoint := central.NewEndpoint()
	userEndpoint := adduser.NewEndpoint()

	r.centralService = []route{
		{
			Name:        "GET health method",
			Description: "check health",
			Method:      http.MethodGet,
			Pattern:     "helath",
			Endpoint:    centralEndpoint.Health,
			Validtion:   middlewere.GeneralValidtion,
		},
	}

	r.userService = []route{
		{
			Name:        "POST add user",
			Description: "add user to database",
			Method:      http.MethodPost,
			Pattern:     "adduser",
			Endpoint:    userEndpoint.AddUser,
			Validtion:   middlewere.GeneralValidtion,
		},
	}

	ro := gin.Default()
	ro.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Accept", "Content-type"},
	}))
	mainRoute := ro.Group("/api")

	for _, e := range r.centralService {
		mainRoute.Handle(e.Method, e.Pattern, e.Endpoint, e.Validtion)
	}

	for _, e := range r.userService {
		mainRoute.Handle(e.Method, e.Pattern, e.Endpoint, e.Validtion)
	}
	return ro
}
