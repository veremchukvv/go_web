<!DOCTYPE html>
<html lang="ru">
  <head>
    <meta charset="utf-8"> 
  </head>
  <body>
<div class="container">
    <div class="post">
        <div class="post-image">
        </div>
        <div class="post-content">
            <div class="post-header">
                <h1>{{.Post.Title}}</h1>
                <div class="post-meta">
                    <span class="author">Автор: {{.Post.Author}}</span>
                    <span class="category">Категория: {{.Post.Category}}</span>
                </div>
            </div>
            <p>{{.Post.Text}}
            </p>
            <p><a href="/post/edit?id={{.Post.ID}}">Изменить</a></p>
            <p><a href="/post/delete?id={{.Post.ID}}">Удалить</a></p>
        </div>
    </div>
</div>
</body>
</html>