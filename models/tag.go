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

func (t Tag) Create() error {
	if len(t.Name) < 3 {
		return errors.New("tag's name is too short.")
	}
	if len(t.Slug) == 0 {
		slug := strings.ReplaceAll(t.Name, "", "-")
		t.Slug = strings.ToLower(slug)
	}

	query := "INSERT INTO tags(name, slug, content) VALUES (?,?,?)"
	ctx, cancelCtx := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelCtx()
	stmt, err := db.Db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, execErr := stmt.ExecContext(ctx, t.Name, t.Slug, utils.NullString(t.Content))
	if execErr != nil {
		return execErr
	}
	return nil
}
