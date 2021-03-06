// Package aescrypter provides a sekrat.Crypter implementation that uses the
// AES-256-GCM algorithm to encrypt/decrypt data.
package aescrypter

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"io"

	"github.com/sekrat/sekrat"
)

const (
	nonceSize = 16
)

// Crypter is a sekrat.Crypter implementation.
type Crypter struct{}

// Encrypt takes a key and a chunk of data, encrypts it, and returns an
// encrypted payload (as a byte array) and an error. If there are any issues
// along the way, the payload is nil and the error is populated. Otherwise,
// the payload is populated and the error is nil.
func (crypter *Crypter) Encrypt(key string, data []byte) ([]byte, error) {
	hashKey := normalize(key)

	block, err := aes.NewCipher(hashKey)
	if err != nil {
		return nil, errors.New("could not set up cipher")
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, errors.New("could not set up gcm")
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, errors.New("could not set up nonce")
	}

	ciphertext := gcm.Seal(nonce, nonce, data, nil)

	return ciphertext, nil
}

// Decrypt takes an encryption key and an encrypted payload, decrypts the
// payload, and returns the decrypted data and an error. If there are any
// issues along the way, the data is nil and the error is populated. Otherwise,
// the data is populated and the error is nil.
func (crypter *Crypter) Decrypt(key string, data []byte) ([]byte, error) {
	hashKey := normalize(key)

	block, err := aes.NewCipher(hashKey)
	if err != nil {
		return nil, errors.New("could not set up cipher")
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, errors.New("could not set up gcm")
	}

	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]

	decrypted, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, errors.New("could not decrypt")
	}

	return decrypted, nil
}

func normalize(key string) []byte {
	sum := make([]byte, 0)

	for _, item := range sha256.Sum256([]byte(key)) {
		sum = append(sum, item)
	}

	return sum
}

// New creates a new instance of Crypter and returns it.
func New() sekrat.Crypter {
	return &Crypter{}
}

/*
Copyright 2019 Dennis Walters

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
