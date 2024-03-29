// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package sqlc

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type Access string

const (
	AccessPUBLIC  Access = "PUBLIC"
	AccessPRIVATE Access = "PRIVATE"
)

func (e *Access) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Access(s)
	case string:
		*e = Access(s)
	default:
		return fmt.Errorf("unsupported scan type for Access: %T", src)
	}
	return nil
}

type NullAccess struct {
	Access Access `json:"access"`
	Valid  bool   `json:"valid"` // Valid is true if Access is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullAccess) Scan(value interface{}) error {
	if value == nil {
		ns.Access, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.Access.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullAccess) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.Access), nil
}

func (e Access) Valid() bool {
	switch e {
	case AccessPUBLIC,
		AccessPRIVATE:
		return true
	}
	return false
}

func AllAccessValues() []Access {
	return []Access{
		AccessPUBLIC,
		AccessPRIVATE,
	}
}

type File struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Access    Access    `json:"access"`
	Path      string    `json:"path"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Ext       string    `json:"ext"`
}
