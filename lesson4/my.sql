
CREATE TABLE `blog_app`.`post_list` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `category` TEXT NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `id_UNIQUE` (`id` ASC) VISIBLE);

CREATE TABLE `blog_app`.`posts` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `category_id` INT NOT NULL,
  `title` TEXT NOT NULL,
  `author` TEXT NOT NULL,
  `text` TEXT NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `id_UNIQUE` (`id` ASC) VISIBLE);

insert into blog_app.post_list (category) values ("common"), ("news"), ("article"), ("personal");

insert into blog_app.posts (category_id, title, author, text) values (1, "Title 1", "Author 1", "post text1"), (1, "Title 2", "Author 2", "post text2"), (2, "Title 3", "Author 3", "post text3"), (4, "Title 4", "Author 4", "post text4"), (4, "Title 5", "Author 5", "post text5");