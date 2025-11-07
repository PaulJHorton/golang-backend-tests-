package main

import (
	"net/http"
	"paul/imagesaver"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// Add your own CORS headers before sending the response
	e.Pre(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			res := c.Response()
			req := c.Request()

			res.Header().Set("Access-Control-Allow-Origin", "*")
			res.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,PATCH,OPTIONS")
			res.Header().Set("Access-Control-Allow-Headers", "*")

			// Handle OPTIONS (preflight) requests directly
			if req.Method == http.MethodOptions {
				return c.NoContent(http.StatusNoContent)
			}

			return next(c)
		}
	})

	e.POST("/resume", func(c echo.Context) error {
		err := imagesaver.SaveImage(c)
	
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
	
		return c.String(http.StatusOK, "Blake is gay")
	})


	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Blake is gay")
	})


	e.Logger.Fatal(e.Start(":8080"))
}

