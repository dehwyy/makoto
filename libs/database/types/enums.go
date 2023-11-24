package types

import (
	"database/sql/driver"
)

type UserPrivacyType uint8

const (
	NOBODY UserPrivacyType = iota + 1
	FRIENDS
	EVERYONE
)

func (u *UserPrivacyType) Scan(value interface{}) error {
	*u = UserPrivacyType(value.(uint8))
	return nil
}

func (u UserPrivacyType) Value() (driver.Value, error) {
	return int8(u), nil
}
