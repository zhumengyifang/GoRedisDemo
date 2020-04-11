package internal

import (
	"github.com/typa01/go-utils"
)

func NewGUID() string {
	return tsgutils.GUID()
}

func NewUUID() string {
	return tsgutils.UUID()
}
