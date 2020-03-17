package controllers

import (
	"context"
	"github.com/astaxie/beego"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

type NewController struct {
	beego.Controller
	Explorer Explorer
}

func (post *NewController) Get() {
	post.TplName = "new.tpl"
}

func (post *NewController) Post() {
	title := post.GetString("title")
	author := post.GetString("author")
	category := post.GetString("category")
	text := post.GetString("text")

	update := bson.D{{Key: "Title", Value: title}, {Key: "Author", Value: author}, {Key: "Category", Value: category}, {Key: "Text", Value: text}}
	c := post.Explorer.Db.Database(post.Explorer.DbName).Collection("posts")
	_, err := c.InsertOne(context.Background(), update)

	if err != nil {
		log.Println(err)
	}

	post.Redirect("/", 302)
}
