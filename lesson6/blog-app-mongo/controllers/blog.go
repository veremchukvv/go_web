package controllers

import (
	"context"
	"github.com/astaxie/beego"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go_web/lesson6/blog-app-mongo/models"
	"log"
)

type BlogController struct {
	beego.Controller
	Explorer Explorer
}

type Explorer struct {
	Db     *mongo.Client
	DbName string
}

func (blog *BlogController) Get() { // цикломатическая сложность функции = 2
	posts, err := blog.Explorer.getAllPosts()

	if err != nil {
		blog.Ctx.ResponseWriter.WriteHeader(404)
		return
	}
	blog.Data["Post"] = posts
	blog.TplName = "blog.tpl"
}

func (e Explorer) getAllPosts() ([]models.Post, error) { // цикломатическая сложность функции = 3
	posts := e.Db.Database(e.DbName).Collection("posts")
	cur, err := posts.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	res := make([]models.Post, 0, 1)
	if err := cur.All(context.Background(), &res); err != nil {
		log.Println(err)
		return nil, err
	}
	return res, nil
}
