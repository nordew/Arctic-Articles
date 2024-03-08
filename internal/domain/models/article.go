package models

import (
	"time"
	"unicode/utf8"
)

type Article struct {
	ArticleID     string    `json:"id"`
	Title         string    `json:"title" validate:"required"`
	Content       string    `json:"content" validate:"required"`
	Author        string    `json:"author"`
	DatePublished time.Time `json:"date_published"`
	ImageURL      string    `json:"image_url"`
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
