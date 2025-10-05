package controllers

import (
	"github.com/go-crypt/crypt"
	"github.com/go-crypt/crypt/algorithm"
	"github.com/go-crypt/crypt/algorithm/argon2"
)

func encryptPassword(password string) (string, error) {
	var (
		hasher *argon2.Hasher
		err    error
		digest algorithm.Digest
	)
	if hasher, err = argon2.New(
		argon2.WithProfileRFC9106LowMemory(),
	); err != nil {
		return "", err
	}

	if digest, err = hasher.Hash(password); err != nil {
		return "", err
	}
	return digest.Encode(), nil
}

func verifyPassword(password, hashedPassword string) (bool, error) {
	var (
		valid bool
		err   error
	)

	if valid, err = crypt.CheckPassword(password, hashedPassword); err != nil {
		return false, err
	}

	return valid, nil
}
