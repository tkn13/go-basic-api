package handler

import (
	"net/http"

	"github.com/ThaksinCK/go-basic-api.git/middlewere"
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
}

func (r Routes) InitRouter() http.Handler {
	centralEndpoint := central.NewEndpoint()

	r.centralService = []route{
		{
			Name:        "GET helath method",
			Description: "check health",
			Method:      http.MethodGet,
			Pattern:     "helath",
			Endpoint:    centralEndpoint.Helath,
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
	return ro
}
