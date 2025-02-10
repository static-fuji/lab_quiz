/*
*専門用語を表すエンティティ
*WordID: 専門用語のID
*Title: 専門用語
*Desc: 専門用語の説明
*Lab: 専門用語のラベル
*ArticleIDs: 関連する論文ID
*Created: 作成日時
*Modified: 更新日時
 */

package entity

import (
	"time"
)

type WordID int64

type Word struct {
	ID         WordID    `json:"id" db:"id"`
	Title      string    `json:"title" db:"title"`
	Desc       string    `json:"desc" db:"description"`
	Lab        string    `json:"lab" db:"lab"`
	ArticleIDs []int     `json:"article_ids"`
	Created    time.Time `json:"created" db:"created"`
	Modified   time.Time `json:"modified" db:"modified"`
}

type Words []*Word
