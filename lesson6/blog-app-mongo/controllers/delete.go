package controllers

import (
	"database/sql"
	"github.com/astaxie/beego"
	"log"
)

type DeleteController struct {
	beego.Controller
	Db *sql.DB
}

func (post *DeleteController) Get() { // цикломатическая сложность функции = 2
	id := post.Ctx.Request.URL.Query().Get("id")

	_, err := post.Db.Exec("delete from blog_app.posts where id = ?", id)
	if err != nil {
		log.Println(err)
	}
	post.Redirect("/", 302)
}
