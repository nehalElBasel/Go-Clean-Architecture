package serviceErrors

import "errors"

type Error struct {
	Message string
}

var (
	ErrNotFound = errors.New("Resquesed item is not found!")
)
