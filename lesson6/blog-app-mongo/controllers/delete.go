package controllers

import (
	"context"
	"github.com/astaxie/beego"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

type DeleteController struct {
	beego.Controller
	Explorer Explorer
}

func (post *DeleteController) Delete() {
	id := post.Ctx.Request.URL.Query().Get("id")
	if len(id) == 0 {
		post.Ctx.ResponseWriter.WriteHeader(404)
		_, _ = post.Ctx.ResponseWriter.Write([]byte("empty id"))
		return
	}
	err := post.Explorer.deletePost(id)
	if err != nil {
		post.Ctx.ResponseWriter.WriteHeader(404)
		return
	}
	post.Redirect("/", 302)
}

func (e Explorer) deletePost(id string) error {
	c := e.Db.Database(e.DbName).Collection("posts")
	filter := bson.D{{Key: "ID", Value: id}}
	_, err := c.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Println(err)
	}
	return err
}
