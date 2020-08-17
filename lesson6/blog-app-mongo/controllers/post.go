package controllers

import (
	"context"
	"github.com/astaxie/beego"
	"go.mongodb.org/mongo-driver/bson"
	"go_web/lesson6/blog-app-mongo/models"
)

type PostController struct {
	beego.Controller
	Explorer Explorer
}

func (post *PostController) Get() {
	id := post.Ctx.Request.URL.Query().Get("id")

	if len(id) == 0 {
		post.Ctx.ResponseWriter.WriteHeader(404)
		_, _ = post.Ctx.ResponseWriter.Write([]byte("empty id"))
		return
	}

	getPost, err := post.Explorer.getOnePost(id)

	if err != nil {
		post.Ctx.ResponseWriter.WriteHeader(404)
		return
	}

	post.Data["Post"] = getPost
	post.TplName = "post.tpl"

}

func (e Explorer) getOnePost(id string) (models.Post, error) {
	c := e.Db.Database(e.DbName).Collection("posts")
	filter := bson.D{{Key: "ID", Value: id}}
	res := c.FindOne(context.Background(), filter)

	post := new(models.Post)
	if err := res.Decode(post); err != nil {
		return models.Post{}, err
	}

	return *post, nil
}
