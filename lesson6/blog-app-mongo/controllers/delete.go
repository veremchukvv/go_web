package controllers

import (
	"context"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
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

	c := post.Explorer.Db.Database(post.Explorer.DbName).Collection("posts")
	filter := bson.D{{Key: "ID", Value: id}}

	_, err := c.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Println(err)
	}
	post.Redirect("/", 302)
}
