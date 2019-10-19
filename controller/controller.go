package controller

import "github.com/labstack/echo"

type Controller interface {
	Set(g *echo.Group)
	Init() error
}
