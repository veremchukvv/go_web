package test

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go_web/lesson6/blog-app-mongo/controllers"
	"log"
	"testing"
)

func TestNewPostsOK(t *testing.T) {
	title := "test title"
	author := "test author"
	category := "test category"
	text := "test text"

	e, err := initDb()
	if err != nil {
		t.Error(err)
		return
	}
	controllers.Explorer.SaveNewPost(e, title, author, category, text)
}

func initDb() (controllers.Explorer, error) {
	db, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return controllers.Explorer{}, err
	}

	if err = db.Connect(context.Background()); err != nil {
		return controllers.Explorer{}, err
	}
	log.Print("mongo connected")

	e := controllers.Explorer{
		Db:     db,
		DbName: "blog_app",
	}
	return e, nil
}
