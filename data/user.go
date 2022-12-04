package data

import "github.com/uptrace/bun"

type User struct {
	bun.BaseModel `bun:"table:users"`

	Id int ``
}
