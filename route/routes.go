package route

import (
	"github.com/kataras/iris"
	"github.com/mohuishou/scuplus-go/api/jwc"
)

// Routes 路由
func Routes(app *iris.Application) {
	app.Get("/user/grade", jwc.GetGrade)
}
