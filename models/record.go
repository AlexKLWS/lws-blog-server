package models

import "time"

type MaterialRecord interface {
	GetID() uint
	GetCreatedAt() time.Time
}
