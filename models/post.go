package models

import (
	"context"
	"dropCms/db"
	"dropCms/utils"
	"errors"
	"time"
)

type Post struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Slug        string `json:"slug"`
	UserId      uint
	ParentId    uint
	Published   bool     `json:"published"`
	Content     string   `json:"content"`
	Tags        []string `json:"tags"`
	Categories  []string `json:"categories"`
}

func (p Post) Create() error {
	if len(p.Title) == 0 {
		return errors.New("Invalid input.")
	}
	if len(p.Slug) == 0 {
		return errors.New("Invalid input.")
	}
	if p.UserId == 0 {
		return errors.New("Invalid input.")
	}
	query := "INSERT INTO posts(title, description, slug, userId, parentId, published, content, publishedAt, updatedAt) VALUES (?,?,?,?,?,?,?,?,?)"
	tim := new(Time)
	if p.Published {
		tim.PublishedAt = time.Now()
		tim.UpdatedAt = time.Now()
	}
	ctx, cancelCtx := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelCtx()
	stmt, err := db.Db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, errr := stmt.ExecContext(ctx, p.Title, utils.NullString(p.Description), p.Slug, p.UserId, utils.NullInt(p.ParentId), p.Published, utils.NullString(p.Content), utils.NullTime(tim.PublishedAt), utils.NullTime(tim.UpdatedAt))
	if errr != nil {
		return errr
	}
	return nil
}

func GetPostById(slug string) (Post, error) {
	query := "SELECT * FROM posts WHERE slug = ? LIMIT 1"
	ctx, cancelCtx := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelCtx()
	stmt, err := db.Db.PrepareContext(ctx, query)
	if err != nil {
		return Post{}, err
	}
	defer stmt.Close()
	row := stmt.QueryRowContext(ctx, slug)
	eror := row.Scan()
	if eror != nil {
		return Post{}, eror
	}
	//Add tag support ???
	return Post{}, nil
}
