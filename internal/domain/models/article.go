package models

import (
	"time"
	"unicode/utf8"
)

type Article struct {
	ArticleID     string
	Title         string `validate:"required"`
	Content       string `validate:"required"`
	Author        string
	DatePublished time.Time
	ImageURL      string
}

func (a *Article) Validate() error {
	if err := validate.Struct(a); err != nil {
		return err
	}

	if utf8.RuneCountInString(a.Title) > 70 || utf8.RuneCountInString(a.Title) < 8 {
		return ErrWrongTitleLength
	}

	return nil
}
