package main

import (
	"github.com/kataras/iris/v12"
)

var a = map[string]string{
	"1":    "prash",
	"2": "kornu",
}

func second(ctx iris.Context) {
	id := ctx.Params().Get("id")

	if value, ok := a[id]; ok {
		ctx.JSON(iris.Map{"id": id, "value": value})
		return
	}

	ctx.StatusCode(iris.StatusNotFound)
	ctx.JSON(iris.Map{"error": "Data not found"})

}
func main() {
	app := iris.New()
	app.Get("/second/{id:string}", second)
	app.Listen(":8080")
}
