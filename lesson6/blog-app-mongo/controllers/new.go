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

func (post *NewController) Get() { // цикломатическая сложность функции = 2
	post.TplName = "new.tpl"
	title := post.GetString("title")
	author := post.GetString("author")
	category := post.GetString("category")
	text := post.GetString("text")
	err := post.Explorer.SaveNewPost(title, author, category, text)
	if err != nil {
		log.Println(err)
	}
	post.Redirect("/", 302)
}

func (e Explorer) SaveNewPost(title, author, category, text string) error { // цикломатическая сложность функции = 2
	update := bson.D{{Key: "Title", Value: title}, {Key: "Author", Value: author}, {Key: "Category", Value: category}, {Key: "Text", Value: text}}
	c := e.Db.Database(e.DbName).Collection("posts")
	_, err := c.InsertOne(context.Background(), update)
	if err != nil {
		log.Println(err)
	}
	return nil
}
