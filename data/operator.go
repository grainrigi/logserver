package data

import "github.com/uptrace/bun"

type LicenseGrade int

const (
	Grade1st LicenseGrade = 1
	Grade2nd LicenseGrade = 2
	Grade3rd LicenseGrade = 3
	Grade4th LicenseGrade = 4
)

type Operator struct {
	bun.BaseModel `bun:"table:operators"`

	ID int `bun:"id,pk,autoincrement" json:"id" param:"id"`

	Name    string       `bun:"name,notnull" json:"name" validate:"required"`
	License LicenseGrade `bun:"license,notnull" json:"license" validate:"required"`
}
