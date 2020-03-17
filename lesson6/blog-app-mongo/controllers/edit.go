package controllers

import (
	"context"
	// "fmt"
	"github.com/astaxie/beego"
	"go.mongodb.org/mongo-driver/bson"
	"go_web/lesson6/blog-app-mongo/models"
	"log"
)

type EditController struct {
	beego.Controller
	Explorer Explorer
}

func (post *EditController) Get() {
	id := post.Ctx.Request.URL.Query().Get("id")

	if len(id) == 0 {
		post.Ctx.ResponseWriter.WriteHeader(404)
		_, _ = post.Ctx.ResponseWriter.Write([]byte("empty id"))
		return
	}

	editPost, err := post.Explorer.getPostForEdit(id)

	if err != nil {
		post.Ctx.ResponseWriter.WriteHeader(404)
		return
	}

	post.Data["Post"] = editPost
	post.TplName = "edit.tpl"

}

func (post *EditController) Post() {
	id := post.GetString("id")
	title := post.GetString("title")
	author := post.GetString("author")
	category := post.GetString("category")
	text := post.GetString("text")

	filter := bson.D{{Key: "ID", Value: id}}
	update := bson.D{{Key: "Title", Value: title}, {Key: "Author", Value: author}, {Key: "Category", Value: category}, {Key: "Text", Value: text}}

	update = bson.D{{Key: "$set", Value: update}}

	c := post.Explorer.Db.Database(post.Explorer.DbName).Collection("posts")
	_, err := c.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Println(err)
	}

	post.Redirect("/", 302)
}

func (e Explorer) getPostForEdit(id string) (models.Post, error) {
	c := e.Db.Database(e.DbName).Collection("posts")
	filter := bson.D{{Key: "ID", Value: id}}
	res := c.FindOne(context.Background(), filter)

	post := new(models.Post)
	if err := res.Decode(post); err != nil {
		return models.Post{}, err
	}

	return *post, nil
}
