package hasher

import (
	"crypto/sha1"
	"fmt"
	"github.com/nordew/ArcticArticles/internal/domain/models"
)

type passwordHasher struct {
	salt string
}

func NewPasswordHasher(salt string) Hasher {
	return &passwordHasher{
		salt: salt,
	}
}

func (p *passwordHasher) Hash(password string) (string, error) {
	hash := sha1.New()

	if _, err := hash.Write([]byte(password)); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", hash.Sum([]byte(p.salt))), nil
}

func (p *passwordHasher) Compare(hashedPassword, password string) error {
	newHashedPassword, err := p.Hash(password)
	if err != nil {
		return err
	}

	if hashedPassword != newHashedPassword {
		return models.ErrWrongEmailOrPassword
	}

	return nil
}
