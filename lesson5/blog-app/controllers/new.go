package controllers

import (
	"database/sql"
	"github.com/astaxie/beego"
	"log"
)

type NewController struct {
	beego.Controller
	Db *sql.DB
}

func (post *NewController) Get() {
	post.TplName = "new.tpl"
}

func (post *NewController) Post() {
	title := post.GetString("title")
	author := post.GetString("author")
	category := post.GetString("category")
	text := post.GetString("text")

	_, err := post.Db.Exec("insert into blog_app.posts (category_id, title, author, text) values (?, ?, ?, ?)", category, title, author, text)
	if err != nil {
		log.Println(err)
	}
	post.Redirect("/", 302)
}
