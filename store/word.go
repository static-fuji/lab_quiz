/*
*wordsテーブルに関する処理を行うパッケージ
 */

package store

import (
	"context"

	"github.com/static-fuji/lab_quiz/entity"
)

func (r *Repository) ListWords(
	ctx context.Context, db Queryer,
) (entity.Words, error) {
	words := entity.Words{}
	sql := `SELECT
		id, title, description, created, modified
		FROM words`
	if err := db.SelectContext(ctx, &words, sql); err != nil {
		return nil, err
	}
	return words, nil
}

func (r *Repository) AddWord(
	ctx context.Context, db Execer, t *entity.Word,
) error {
	t.Created = r.Clocker.Now()
	t.Modified = r.Clocker.Now()
	sql := `INSERT INTO words
		(title, description, lab, created, modified)
		VALUES (?, ?, ?, ?, ?)`

	results, err := db.ExecContext(
		ctx, sql, t.Title, t.Desc, t.Lab, t.Created, t.Modified,
	)

	if err != nil {
		return err
	}

	id, err := results.LastInsertId()
	if err != nil {
		return err
	}

	t.ID = entity.WordID(id)
	return nil
}
