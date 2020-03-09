<!DOCTYPE html>
<html lang="ru">
  <head>
    <meta charset="utf-8"> 
    <title>БЛОГ</title>
  </head>
  <body>
    <article>
      <header>
        <ul>
          {{range .}}
          <li>
               <h1><a href="/post?id={{.ID}}">{{.Title}}</a></h1>
              <p>Автор: {{.Author}}</p>
              <p>Категория поста: {{.Category}}</p>
          </li>
          {{end}}
        </ul>
      </header>
    </article>
    <h1><a href="/post/new">Добавить пост</a></h1>
  </body>
</html>