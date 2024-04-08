package common

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type Hasher struct {
}

func (h *Hasher) RandomString(length int) (string, error) {
	var bytes = make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func (h *Hasher) HashPassword(salt string, password string) (string, error) {
	spStr := fmt.Sprintf("%s%s", salt, password)
	hash, err := bcrypt.GenerateFromPassword([]byte(spStr), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (h *Hasher) CheckPassword(salt string, password string, hashPassword string) bool {
	spStr := fmt.Sprintf("%s%s", salt, password)
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(spStr)) == nil
}
