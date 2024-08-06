package models

type Book struct{
	ID int
	Title string
	Author string
	Status string // can be "Available" or "Borrowed"
}