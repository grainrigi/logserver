package data

import (
	"strings"
	"time"

	"github.com/uptrace/bun"
)

type ContestType int

const (
	SingleOp ContestType = 1
	MultiOp  ContestType = 2
)

type Contest struct {
	bun.BaseModel `bun:"table:contests"`

	ID int `bun:"id,pk,autoincrement" json:"id" param:"id"`

	Name      string      `bun:"name,notnull" json:"name"`
	StartTime time.Time   `bun:"start_time" json:"startTime"`
	EndTime   time.Time   `bun:"end_time" json:"endTime"`
	Type      ContestType `bun:"type,notnull" json:"type" validate:"required,oneof=1 2"`
	Cfg       string      `bun:"cfg,notnull" json:"cfg" validate:"required"`
	Call      string      `bun:"call,notnull" json:"call" validate:"required"`
}

// 入力データやDBからロードしたデータで不完全な部分を修正する
func (c *Contest) Normalize() error {
	// コールサインは大文字小文字関係ないので大文字にしておく
	c.Call = strings.ToUpper(c.Call)

	return nil
}
