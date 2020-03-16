package routers

import (
	"context"
	"github.com/astaxie/beego"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go_web/lesson6/blog-app-mongo/controllers"
	"log"
)

const dbName = "blog_app"

func init() {
	db, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Connect(context.Background()); err != nil {
		log.Fatal(err)
	}

	log.Println("Database connection OK")

	beego.Router("/", &controllers.BlogController{
		Controller: beego.Controller{},
		Explorer: controllers.Explorer{
			Db:     db,
			DbName: dbName,
		},
	})

	// beego.Router("/post", &controllers.PostController{
	// 	Controller: beego.Controller{},
	// 	Db:         db,
	// })

	// beego.Router("/post/edit/", &controllers.EditController{
	// 	Controller: beego.Controller{},
	// 	Db:         db,
	// })

	// beego.Router("/post/new", &controllers.NewController{
	// 	Controller: beego.Controller{},
	// 	Db:         db,
	// })

	// beego.Router("/post/delete/", &controllers.DeleteController{
	// 	Controller: beego.Controller{},
	// 	Db:         db,
	// })

}
