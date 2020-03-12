package controllers

import (
	"database/sql"
	"github.com/astaxie/beego"
	"go_web/lesson5/blog-app/models"
	"log"
)

type BlogController struct {
	beego.Controller
	Db *sql.DB
}

func (blog *BlogController) Get() {

	posts, err := getAllPosts(blog.Db)

	if err != nil {
		blog.Ctx.ResponseWriter.WriteHeader(404)
		return
	}

	blog.Data["Post"] = posts
	blog.TplName = "blog.tpl"

}

func getAllPosts(db *sql.DB) ([]models.Post, error) {
	res := make([]models.Post, 0, 1)
	rows, err := db.Query("select * from blog_app.posts")
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		blog := models.Post{}
		if err := rows.Scan(&blog.ID, &blog.Category, &blog.Title, &blog.Author, &blog.Text); err != nil {
			log.Println(err)
			continue
		}
		res = append(res, blog)
	}

	return res, nil
}
