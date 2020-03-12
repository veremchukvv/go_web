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
          {{range $id, $item := .Post}}
          <li>
               <h1><a href="/post?id={{$item.ID}}">{{$item.Title}}</a></h1>
              <p>Автор: {{$item.Author}}</p>
              <p>Категория поста: {{$item.Category}}</p>
          </li>
          {{end}}
        </ul>
      </header>
    </article>
    <h1><a href="/post/new">Добавить пост</a></h1>
  </body>
</html>