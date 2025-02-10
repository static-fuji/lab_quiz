/*
*論文情報を表すエンティティ
*ID: 論文ID
*Title: 論文タイトル
*Author: 著者
*Created: 作成日時
*Modified: 更新日時
 */

package entity

import (
	"time"
)

type ArticleID int64

type Article struct {
	ID       ArticleID `json:"id" db:"id"`
	Title    string    `json:"title" db:"title"`
	Author   string    `json:"author" db:"author"`
	Created  time.Time `json:"created" db:"created"`
	Modified time.Time `json:"modified" db:"modified"`
}

type Articles []*Article
