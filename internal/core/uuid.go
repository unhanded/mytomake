package core

import "github.com/google/uuid"

type MytoUUID struct{}

func (tbu MytoUUID) Value() string {
	id := uuid.NewString()
	return id
}
