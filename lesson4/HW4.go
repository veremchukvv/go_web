package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"text/template"
)

var database *sql.DB
var DSN = "root:1234@tcp(localhost:3306)/blog_app?charset=utf8"

const port = "8080"

type Server struct {
	db *sql.DB
}

type Post struct {
	ID       int
	Title    string
	Author   string
	Category string
	Text     string
}

var tmplBlog = template.Must(template.New("MyBlogTemplate").ParseFiles("blog.html"))
var tmplPost = template.Must(template.New("MyPostTemplate").ParseFiles("post.html"))

func main() {
	db, err := sql.Open("mysql", DSN)
	if err != nil {
		log.Fatal(err)
	}
	database = db
	defer database.Close()

	s := Server{
		db: db,
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	} else {
		log.Println("Database connection OK")
	}

	router := http.NewServeMux()

	router.HandleFunc("/", s.viewAllPostsForMain)
	// router.HandleFunc("/post", s.viewOnePost)
	// router.HandleFunc("/edit", s.editPost)

	log.Printf("server start at port: %v", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func (server *Server) viewAllPostsForMain(wr http.ResponseWriter, r *http.Request) {
	posts, err := getAllPosts(server.db)
	log.Println(posts)
	if err != nil {
		log.Print(err)
		wr.WriteHeader(500)
		return
	}

	if err := tmplBlog.ExecuteTemplate(wr, "Blog", posts); err != nil {
		log.Println(err)
	}
}

// func (server *Server) viewOnePost(wr http.ResponseWriter, r *http.Request) {
// 	post, err := getOnePost(server.db, r.URL.Query().Get("id"))
// 	if err != nil {
// 		log.Print(err)
// 		wr.WriteHeader(404)
// 		return
// 	}

// 	if err := tmplPost.ExecuteTemplate(wr, "Post", post); err != nil {
// 		log.Println(err)
// 	}
// }

func getAllPosts(db *sql.DB) ([]Post, error) {
	res := make([]Post, 0, 1)
	rows, err := db.Query("select * from blog_app.posts")
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		blog := Post{}
		if err := rows.Scan(&blog.ID, &blog.Title, &blog.Author, &blog.Category, &blog.Text); err != nil {
			log.Println(err)
			continue
		}
		res = append(res, blog)
	}

	return res, nil
}

// func getOnePost(db *sql.DB, id string) ([]PostPage, error) {
// 	return
// }
