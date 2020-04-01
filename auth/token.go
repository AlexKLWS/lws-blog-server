package auth

import uuid "github.com/satori/go.uuid"

func NewToken() string {
	uuidToken := uuid.Must(uuid.NewV4())
	t := uuidToken.String()

	go AddTokenToStorage(t)

	return t
}
