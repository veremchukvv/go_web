{{define "Post"}}
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
                {{range .}}
                <h1>{{.Title}}</h1>
                <div class="post-meta">
                    <span class="author">Автор: {{.Author}}</span>
                    <span class="category">Категория: {{.Category}}</span>
                </div>
            </div>
            <p>{{.Text}}
            </p>
            <p><a href="/post/edit/{{.ID}}">Изменить</a></p>
        </div>
    </div>
</div>
{{end}}
</body>
</html>
{{end}}