package models

import (
	"context"
	"dropCms/db"
	"dropCms/utils"
	"errors"
	"strings"
	"time"
)

type Tag struct {
	Id      uint
	Name    string `json:"name"`
	Slug    string `json:"slug"`
	Content string `json:"content"`
}

func (t Tag) Create() (interface{}, error) {
	if len(t.Name) < 3 {
		return nil, errors.New("tag's name is too short.")
	}
	if len(t.Slug) == 0 {
		slug := strings.ReplaceAll(t.Name, "", "-")
		t.Slug = strings.ToLower(slug)
	}

	query := "INSERT INTO tags(name, slug, content) VALUES (?,?,?)"
	ctx, cancelCtx := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancelCtx()
	stmt, err := db.Db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	res, execErr := stmt.ExecContext(ctx, t.Name, t.Slug, utils.NullString(t.Content))
	if execErr != nil {
		return nil, execErr
	}
	id, eror := res.LastInsertId()
	if eror != nil {
		return nil, eror
	}
	return id, nil
}

func MakeAssociationWithTag(postId, tagId uint) error {
	query := "INSERT INTO post_tags(post_id, tag_id) VALUES (?,?)"
	ctx, cancelCtx := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancelCtx()
  stmt, err := db.Db.PrepareContext(ctx, query)
  if err != nil {
    return err
  }
  defer stmt.Close()
  _, errorr := stmt.ExecContext(ctx, postId, tagId)
  if errorr != nil {
    return errorr
  }
  return nil
}
