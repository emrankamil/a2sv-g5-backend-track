package models

type student struct {
	ID    int
	Name  string
	BorrowedBooks []Book
}