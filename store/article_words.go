/*
*article_wordsテーブルに関する処理を行う
 */

package store

import (
	"context"

	"github.com/static-fuji/lab_quiz/entity"
)

func (r *Repository) BindArticleToWords(
	ctx context.Context, db Execer, t *entity.Word,
) error {
	t.Created = r.Clocker.Now()
	t.Modified = r.Clocker.Now()
	sql := `INSERT INTO article_words
		(article_id, word_id, created, modified)
		VALUES (?, ?, ?, ?)`

	for _, articleID := range t.ArticleIDs {
		_, err := db.ExecContext(
			ctx, sql, articleID, t.ID, t.Created, t.Modified,
		)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *Repository) ListBind(
	ctx context.Context, db Queryer, id int,
) (entity.Words, error) {
	sql := `SELECT words.id, words.title, words.description, words.lab
		FROM words
		INNER JOIN article_words
		ON words.id = article_words.word_id
		WHERE article_words.article_id = ?`

	rows, err := db.QueryContext(ctx, sql, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var words entity.Words
	for rows.Next() {
		var w entity.Word
		if err := rows.Scan(
			&w.ID,
			&w.Title,
			&w.Desc,
			&w.Lab,
		); err != nil {
			return nil, err
		}
		words = append(words, &w)
	}

	return words, nil
}
