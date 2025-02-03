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
