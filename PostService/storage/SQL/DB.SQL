
CREATE TABLE posts(
    id UUID PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    description VARCHAR(50) NOT NULL,
    user_id UUID NOT NULL
);
CREATE TABLE medias(
    id UUID NOT NULL PRIMARY KEY,
    post_id UUID NOT NULL,
    type TEXT NOT NULL,
    link TEXT NOT NULL,
    FOREIGN KEY(post_id)
      REFERENCES posts(id)
      ON DELETE CASCADE
);
SELECT * FROM posts;
SELECT * FROM medias;
SELECT id,name,description FROM posts WHERE user_id = '042072bf-7a5f-40c4-82da-e1734165f8fb';