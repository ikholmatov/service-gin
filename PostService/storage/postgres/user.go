package postgres

import (
	"github.com/jmoiron/sqlx"
	pb "github.com/venomuz/project4/PostService/genproto"
)

type postRepo struct {
	db *sqlx.DB
}

//NewPostRepo ...
func NewPostRepo(db *sqlx.DB) *postRepo {
	return &postRepo{db: db}
}

func (r *postRepo) PostCreate(post *pb.Post) (*pb.Post, error) {
	PostQuery := `INSERT INTO posts(id,name,description,user_id) VALUES($1,$2,$3,$4) RETURNING id,name,description,user_id`
	Post := pb.Post{}
	err := r.db.QueryRow(PostQuery, post.Id, post.Name, post.Description, post.UserId).Scan(&Post.Id, &Post.Name, &Post.Description, &Post.UserId)
	if err != nil {
		return nil, err
	}
	for _, media := range post.Medias {
		Media := pb.Media{}
		PostQuery := `INSERT INTO medias(id,post_id,type,link) VALUES($1,$2,$3,$4) RETURNING id,post_id,type,link`
		err := r.db.QueryRow(PostQuery, media.Id, post.Id, media.Type, media.Link).Scan(&Media.Id, &Media.PostId, &Media.Type, &Media.Link)
		if err != nil {
			return nil, err
		}
		Post.Medias = append(Post.Medias, &Media)
	}

	return &Post, nil
}
func (r *postRepo) PostGetByID(ID string) (*pb.Post, error) {
	post := pb.Post{}
	GetPostQuery := `SELECT id,name,description FROM posts WHERE user_id = $1`
	err := r.db.QueryRow(GetPostQuery, ID).Scan(&post.Id, &post.Name, &post.Description)
	if err != nil {
		return nil, err
	}
	var medias []*pb.Media
	GetMediaQuery := `SELECT id,post_id,type,link FROM medias WHERE post_id = $1`
	rows, err := r.db.Query(GetMediaQuery, post.Id)
	for rows.Next() {
		media := pb.Media{}
		err := rows.Scan(&media.Id, &media.PostId, &media.Type, &media.Link)
		if err != nil {
			return nil, err
		}
		medias = append(medias, &media)
	}
	post.Medias = medias

	return &post, nil
}
func (r *postRepo) PostDeleteByID(ID string) (*pb.OkBOOL, error) {
	_, err := r.db.Exec(`DELETE  FROM posts WHERE user_id = $1`, ID)
	if err != nil {
		return nil, err
	}

	return &pb.OkBOOL{Status: true}, nil
}
func (r *postRepo) PostGetAllPosts(ID string) (*pb.AllPost, error) {
	allpost := pb.AllPost{}
	posts := allpost.Posts
	GetPostQuery := `SELECT id,name,description FROM posts WHERE user_id = $1`
	rows, err := r.db.Query(GetPostQuery, ID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		post := pb.Post{}
		err := rows.Scan(&post.Id, &post.Name, &post.Description)
		if err != nil {
			return nil, err
		}
		var medias []*pb.Media
		GetMediaQuery := `SELECT id,post_id,type,link FROM medias WHERE post_id = $1`
		rows, err := r.db.Query(GetMediaQuery, post.Id)
		for rows.Next() {
			media := pb.Media{}
			err := rows.Scan(&media.Id, &media.PostId, &media.Type, &media.Link)
			if err != nil {
				return nil, err
			}
			medias = append(medias, &media)
		}
		post.Medias = medias
		posts = append(posts, &post)
	}

	allpost.Posts = posts
	return &allpost, nil
}
