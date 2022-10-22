package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {

		cmd := exec.Command("env")
		envs, err := cmd.Output()
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		return c.String(http.StatusOK, string(envs))
	})
	e.GET("/version", func(c echo.Context) error {
		return c.String(http.StatusOK, fmt.Sprintf("v1.0.0   => %s", os.Getenv("HOSTNAME")))
	})
	e.Logger.Fatal(e.Start(":8000"))
}
