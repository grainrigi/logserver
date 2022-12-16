package data

import (
	"time"

	"github.com/uptrace/bun"
)

// 通信モード
type CommMode string

const (
	SSB CommMode = "SSB"
	CW  CommMode = "CW"
	FM  CommMode = "FM"
	AM  CommMode = "AM"
)

type Log struct {
	bun.BaseModel `bun:"table:logs"`

	ID        int `bun:"id,pk,autoincrement" json:"id" param:"id"`
	ContestID int `bun:"contest,notnull" json:"-" param:"cid"`

	Time       time.Time `bun:"time,notnull" json:"time"`
	Call       string    `bun:"call,notnull" json:"call" validate:"required"`
	RST        string    `bun:"rst,notnull" json:"rst" validate:"required"`
	Rcvd       string    `bun:"rcvd,notnull" json:"rcvd"`
	Band       BandFreq  `bun:"band,notnull" json:"band" validate:"required,oneof=1.9 3.5 7 11 14 18 21 24 28 50 144 430 1200 2400 5600 10G"`
	Mode       CommMode  `bun:"mode,notnull" json:"mode" validate:"required,oneof=SSB CW FM AM"`
	Pwr        *string   `bun:"pwr" json:"pwr"`
	Operator   *Operator `bun:"rel:belongs-to,join:op=id" json:"op" validate:"-"`
	OperatorID *int      `bun:"op" json:"-"`
	Note       string    `bun:"note,notnull" json:"note"`

	TxRST *string `bun:"txrst" json:"txrst,omitempty"`
	Txd   *string `bun:"txd" json:"txd,omitempty"`
}

// 入力データやDBからロードしたデータで不完全な部分を修正する
func (l *Log) Normalize() error {
	// デフォルトの交信時刻は現時刻
	if l.Time.Equal(time.Time{}) {
		l.Time = time.Now()
	}
	// OperatorIDをOperatorと同期
	if l.Operator != nil {
		l.OperatorID = &l.Operator.ID
	} else {
		l.OperatorID = nil
	}

	return nil
}
