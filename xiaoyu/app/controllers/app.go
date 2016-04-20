package controllers

import (
	"fmt"
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	fmt.Println("I am app index")
	return c.Render()
}
func (c App) Fuck() revel.Result {
	fmt.Println(" Fuck")
	return c.RenderJson("FUCK")
}
