package db

import "github.com/joaoh82/housinganywhere/pkg/domain"

// Database interface that could be used to implement many different data readers
type Database interface {
	ReadList() (domain.Locations, error)
}
