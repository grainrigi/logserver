package data

import (
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

	ID        int         `bun:"id,pk,autoincrement" json:"id"`
	StartTime time.Time   `bun:"start_time" json:"startTime"`
	EndTime   time.Time   `bun:"end_time" json:"endTime"`
	Type      ContestType `bun:"type,notnull" json:"type"`
	Cfg       string      `bun:"cfg,notnull" json:"cfg"`
	Call      string      `bun:"call,notnull" json:"call"`
}
