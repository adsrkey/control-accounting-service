package types

import (
	"github.com/google/uuid"
)

type ReqID struct {
	ID string `uri:"id" binding:"required,uuid"`
}

func (r ReqID) Valid() (uuid.UUID, error) {
	return uuid.Parse(r.ID)
}
