package utils

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

type ArgonParams struct {
	Memory      uint32
	Iterations  uint32
	Parallelism uint8
	SaltLength  uint32
	KeyLength   uint32
}

var params = &ArgonParams{
	Memory:      128 * 1024, // 128 MB
	Iterations:  3,
	Parallelism: 4,
	SaltLength:  16,
	KeyLength:   32,
}

func HashPassword(password string) (string, error) {

	// generate random salt
	salt := make([]byte, params.SaltLength)

	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}

	// generate hash
	hash := argon2.IDKey(
		[]byte(password),
		salt,
		params.Iterations,
		params.Memory,
		params.Parallelism,
		params.KeyLength,
	)

	// encode
	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	// production format
	encodedHash := fmt.Sprintf(
		"$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		argon2.Version,
		params.Memory,
		params.Iterations,
		params.Parallelism,
		b64Salt,
		b64Hash,
	)

	return encodedHash, nil
}

func VerifyPassword(password string, encodedHash string) error {

	parts := strings.Split(encodedHash, "$")

	if len(parts) != 6 {
		return errors.New("invalid hash format")
	}

	var memory uint32
	var iterations uint32
	var parallelism uint8

	_, err := fmt.Sscanf(
		parts[3],
		"m=%d,t=%d,p=%d",
		&memory,
		&iterations,
		&parallelism,
	)

	if err != nil {
		return err
	}

	// decode salt
	salt, err := base64.RawStdEncoding.DecodeString(parts[4])
	if err != nil {
		return err
	}

	// decode hash
	originalHash, err := base64.RawStdEncoding.DecodeString(parts[5])
	if err != nil {
		return err
	}

	// generate comparison hash
	comparisonHash := argon2.IDKey(
		[]byte(password),
		salt,
		iterations,
		memory,
		parallelism,
		uint32(len(originalHash)),
	)

	// secure compare
	if subtle.ConstantTimeCompare(originalHash, comparisonHash) == 1 {
		return nil
	}

	return errors.New("invalid password")
}
