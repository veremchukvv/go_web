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
            <form method="POST">
              <label>ID</label><br>
              <input type="number" name="id" value="{{.Post.ID}}" disabled /><br><br>
              <label>Title</label><br>
              <input type="text" name="title" value="{{.Post.Title}}" /><br><br>
              <label>Author</label><br>
              <input type="text" name="author" value="{{.Post.Author}}" /><br><br>
              <label>Category</label><br>
              <input type="text" name="category" value="{{.Post.Category}}" /><br><br>
              <label>Text</label><br>
              <input type="text" name="text" value="{{.Post.Text}}" /><br><br>
              <input type="submit" value="Сохранить" />
          </form>
        </div>
    </div>
</div>
</body>
</html>