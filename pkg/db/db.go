package db

// Database interface that could be used to implement many different data readers
type Database interface {
	ReadList() error
}
