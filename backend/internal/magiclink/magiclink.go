package magiclink

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"os"

	"github.com/valkey-io/valkey-go"
)

func generateToken() (string, error) {
	bytes := make([]byte, 16)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", bytes), nil
}

func createMagicLink(email string, valkey *valkey.Client, ctx context.Context) (string, error) {
	token, err := generateToken()
	if err != nil {
		return "", err
	}

	err = (*valkey).Do(ctx, (*valkey).B().Set().Key(token).Value(email).Build()).Error()
	if err != nil {
		return "", err
	}

	err = (*valkey).Do(ctx, (*valkey).B().Expire().Key(token).Seconds(300).Build()).Error()
	if err != nil {
		return "", err
	}

	host := os.Getenv("MAGICLINK_BASE")
	if host == "" {
		return "", errors.New("magic link base not set")
	}

	return fmt.Sprintf(host+"/auth?token%s", token), nil
}

func validateMagicLink(token string, valkey *valkey.Client, ctx context.Context) (string, error) {

	email, err := (*valkey).Do(ctx, (*valkey).B().Get().Key(token).Build()).ToString()
	if err != nil {
		return "", err
	}

	err = (*valkey).Do(ctx, (*valkey).B().Del().Key(token).Build()).Error()
	if err != nil {
		return "", err
	}

	return email, nil
}
