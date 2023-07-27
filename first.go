package main
import (
	"github.com/kataras/iris/v12"
)
func list(ctx iris.Context){
	ctx.WriteString("Hello, Iris!")
}
func main() {
	app:=iris.New()//create iris app
	app.Get("/",list)
	app.Listen(":8080")
}