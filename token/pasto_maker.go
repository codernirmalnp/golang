package token

import (
	"fmt"
	"time"

	"github.com/o1egl/paseto"
	"golang.org/x/crypto/chacha20poly1305"
)

type PasetoMakes struct {
	pasteo       *paseto.V2
	symmetrickey []byte
}

func NewPasetoMaker(symmetrickey string) (Maker, error) {
	if len(symmetrickey) < chacha20poly1305.KeySize {
		return nil, fmt.Errorf("Invalid Key %v", minSecretKeySize)

	}
	maker := &PasetoMakes{pasteo: paseto.NewV2(), symmetrickey: []byte(symmetrickey)}
	return maker, nil

}

func (maker *PasetoMakes) CreateToken(username string, duration time.Duration) (string, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", err
	}

	return maker.pasteo.Encrypt(maker.symmetrickey, payload, nil)

}

func (maker *PasetoMakes) VerifyToken(token string) (*Payload, error) {
	payload := &Payload{}
	err := maker.pasteo.Decrypt(token, maker.symmetrickey, payload, nil)
	if err != nil {
		return nil, errorInvalidToken
	}
	err = payload.Valid()
	if err != nil {
		return nil, err

	}
	return payload, nil

}
