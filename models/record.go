package models

import "time"

type MaterialRecord interface {
	getID() uint
	getCreatedAt() time.Time
}
