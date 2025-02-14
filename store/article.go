/*
* articlesテーブルに関する処理を行う
 */
package store

import (
	"context"

	"github.com/static-fuji/lab_quiz/entity"
)

func (r *Repository) ListArticle(
	ctx context.Context, db Queryer,
) (entity.Articles, error) {
	sql := `SELECT id, title, author, created, modified FROM articles`

	rows, err := db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var articles entity.Articles
	for rows.Next() {
		var t entity.Article
		if err := rows.Scan(
			&t.ID,
			&t.Title,
			&t.Author,
			&t.Created,
			&t.Modified,
		); err != nil {
			return nil, err
		}
		articles = append(articles, &t)
	}

	return articles, nil
}

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
