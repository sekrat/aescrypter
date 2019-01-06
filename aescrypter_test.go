package aescrypter

import (
	"testing"
)

func TestNew(t *testing.T) {
	crypter := New()

	t.Run("it is an aescrypter.Crypter", func(t *testing.T) {
		_, ok := crypter.(*Crypter)
		if !ok {
			t.Errorf("Expected an aescrypter.Crypter")
		}
	})
}

func TestCrypter_Encrypt(t *testing.T) {
	key := "The chump punk with a skeleton key"
	data := []byte("Hey! That chump is me")

	crypter := New()

	encrypted, err := crypter.Encrypt(key, data)

	t.Run("it returns an encrypted payload", func(t *testing.T) {
		if encrypted == nil || string(encrypted) == string(data) {
			t.Errorf("Expected encrypted payload, got '%s'", string(encrypted))
		}
	})

	t.Run("it returns no error", func(t *testing.T) {
		if err != nil {
			t.Errorf("Expected no error, got '%s'", err.Error())
		}
	})
}

func TestCrypter_Decrypt(t *testing.T) {
	key := "sausages"
	data := []byte("Oh what a lovely bunch of bananas!")
	crypter := New()
	encrypted, _ := crypter.Encrypt(key, data)

	t.Run("when the correct key is used", func(t *testing.T) {
		decrypted, err := crypter.Decrypt(key, encrypted)

		t.Run("it returns the proper data", func(t *testing.T) {
			if string(decrypted) != string(data) {
				t.Errorf("Expected decrypted to be '%s', got '%s'", string(data), string(decrypted))
			}
		})

		t.Run("it returns no error", func(t *testing.T) {
			if err != nil {
				t.Errorf("Expected no error, got '%s'", err.Error())
			}
		})
	})

	t.Run("when a bad key is used", func(t *testing.T) {
		key = key + key

		decrypted, err := crypter.Decrypt(key, encrypted)

		t.Run("it returns no data", func(t *testing.T) {
			if decrypted != nil {
				t.Errorf("Expected no data, got '%s'", string(decrypted))
			}
		})

		t.Run("it returns an error", func(t *testing.T) {
			if err == nil {
				t.Errorf("Expected an error")
			}
		})
	})
}
