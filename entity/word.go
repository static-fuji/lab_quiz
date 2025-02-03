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
