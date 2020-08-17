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

func (post *DeleteController) Get() {
	id := post.Ctx.Request.URL.Query().Get("id")

	if len(id) == 0 {
		post.Ctx.ResponseWriter.WriteHeader(404)
		_, _ = post.Ctx.ResponseWriter.Write([]byte("empty id"))
		return
	}

	_, err := post.Db.Exec("delete from blog_app.posts where id = ?", id)
	if err != nil {
		log.Println(err)
	}
	post.Redirect("/", 302)

}
