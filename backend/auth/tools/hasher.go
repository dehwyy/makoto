package tools

import "golang.org/x/crypto/bcrypt"

var (
	hash_salt = 10
)

type Hasher struct{}

func NewHasher() *Hasher {
	return new(Hasher)
}

func (h *Hasher) Hash(s string) (string, error) {
	hashed_string_bytes, err := bcrypt.GenerateFromPassword([]byte(s), hash_salt)
	return string(hashed_string_bytes), err
}

func (h *Hasher) Compare(s, hashed_s string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed_s), []byte(s))
	return err == nil
}
