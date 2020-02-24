package main

import (
	// "bytes"
	// "encoding/json"
	// "fmt"
	// "io/ioutil"
	"log"
	"net/http"
	// "os"
	// "strings"
	"text/template"
	// "time"
)

const port = "8080"

type BlogMainPage struct {
	Title    string
	PostList []Post
}

type Post struct {
	PostName   string
	PostAuthor string
}

type PostPage struct {
	Title    string
	Author   string
	Category string
	Date     string
	Text     string
}

var tmplBlog = template.Must(template.New("MyBlogTemplate").ParseFiles("blog.html"))
var tmplPost = template.Must(template.New("MyPostTemplate").ParseFiles("post.html"))

// var blog = map[string]Article{
// 	"Статья 1": {"Статья 1", "Автор 1"},
// 	"Статья 2": {"Статья 2", "Автор 2"},
// 	"Статья 3": {"Статья 3", "Автор 3"},
// }

var blog = BlogMainPage{
	Title: "Блог. Страница 1",
	PostList: []Post{
		{PostName: "Статья 1", PostAuthor: "Автор 1"},
		{PostName: "Статья 2", PostAuthor: "Автор 2"},
		{PostName: "Статья 3", PostAuthor: "Автор 3"},
	},
}

var post = PostPage{
	Title:    "Пост 1",
	Author:   "Автор 1",
	Category: "Категория 1",
	Date:     "11.1.2011",
	Text:     "1111111111111111",
}

func main() {
	router := http.NewServeMux()

	router.HandleFunc("/", viewAllPosts)
	router.HandleFunc("/post", viewPost)
	router.HandleFunc("/edit", viewPost)

	log.Printf("server start at port: %v", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func viewAllPosts(wr http.ResponseWriter, req *http.Request) {
	if err := tmplBlog.ExecuteTemplate(wr, "Blog", blog); err != nil {
		log.Println(err)
	}
}

func viewPost(wr http.ResponseWriter, req *http.Request) {
	if err := tmplPost.ExecuteTemplate(wr, "Post", post); err != nil {
		log.Println(err)
	}

}
