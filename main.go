package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/zinirun/go-restful-api/elasticsearch"
)

const _port = "8080"

func main() {
	elasticsearch.Initialize()

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Echo!")
	})
	e.Start(fmt.Sprintf(":%v", _port))
}
