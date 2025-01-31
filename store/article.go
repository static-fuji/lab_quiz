package store

import (
	"context"

	"github.com/static-fuji/lab_quiz/entity"
)

func (r *Repository) AddArticle(
	ctx context.Context, db Execer, t *entity.Article,
) error {
	t.Created = r.Clocker.Now()
	t.Modified = r.Clocker.Now()
	sql := `INSERT INTO articles
		(title, author, created, modified)
		VALUES (?, ?, ?, ?)`

	results, err := db.ExecContext(
		ctx, sql, t.Title, t.Author, t.Created, t.Modified,
	)

	if err != nil {
		return err
	}

	id, err := results.LastInsertId()
	if err != nil {
		return err
	}

	t.ID = entity.ArticleID(id)
	return nil
}

func (r *Repository) SearchArticleID(
	ctx context.Context, db Queryer, id int,
) error {
	sql := `SELECT id FROM articles WHERE id = ?`
	if err := db.QueryRowContext(ctx, sql, id).Scan(&id); err != nil {
		return err
	}
	return nil
}
