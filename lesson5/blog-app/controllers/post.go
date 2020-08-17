package controllers

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
	"go_web/lesson5/blog-app/models"
	"log"
)

type PostController struct {
	beego.Controller
	Db *sql.DB
}

func (post *PostController) Get() {
	id := post.Ctx.Request.URL.Query().Get("id")

	if len(id) == 0 {
		post.Ctx.ResponseWriter.WriteHeader(404)
		_, _ = post.Ctx.ResponseWriter.Write([]byte("empty id"))
		return
	}

	getPost, err := getOnePost(post.Db, id)

	if err != nil {
		post.Ctx.ResponseWriter.WriteHeader(404)
		return
	}

	post.Data["Post"] = getPost
	post.TplName = "post.tpl"

}

func getOnePost(db *sql.DB, id string) ([]models.Post, error) {
	res := make([]models.Post, 0, 1)
	rows, err := db.Query(fmt.Sprintf("select * from blog_app.posts WHERE ID= %v", id))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		post := models.Post{}
		if err := rows.Scan(&post.ID, &post.Category, &post.Title, &post.Author, &post.Text); err != nil {
			log.Println(err)
			continue
		}
		res = append(res, post)
	}

	return res, nil
}
