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
           {{range $id, $item := .Post}}
                <h1>{{$item.Title}}</h1>
                <div class="post-meta">
                    <span class="author">Автор: {{$item.Author}}</span>
                    <span class="category">Категория: {{$item.Category}}</span>
                </div>
            </div>
            <p>{{$item.Text}}
            </p>
            <p><a href="/post/edit/{{$item.ID}}">Изменить</a></p>
        </div>
    </div>
</div>
{{end}}
</body>
</html>