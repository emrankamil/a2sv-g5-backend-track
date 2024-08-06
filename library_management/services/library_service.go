package services

import (
	"errors"
	"library_management/models"
)

type Library struct{
	books map[int]models.Book
	members map[int]models.Member
}

type LibraryManager interface {
	AddBook(book models.Book) error
	RemoveBook(bookID int) error
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() ([]models.Book, error)
	ListBorrowedBooks(memberID int) ([]models.Book, error)
}

func NewLibrary() *Library {
    return &Library{
        books:   make(map[int]models.Book),
        members: make(map[int]models.Member),
    }
}

func (l *Library) AddMember(member models.Member){
	id := member.ID
	l.members[id] = member
}

func (l *Library) AddBook(book models.Book) error{
	id := book.ID
	for i := range l.books{
		if i == id{
			return errors.New("ID ALREADY EXISTS")
		}
	}
	l.books[id] = book

	return nil
}

func (l *Library) RemoveBook(bookID int) error {
    if _, exists := l.books[bookID]; !exists {
        return errors.New("book not found")
    }
    delete(l.books, bookID)
    return nil
}

func (l *Library) BorrowBook(bookID int, memberID int) error{
	book, bookExists := l.books[bookID]
	if !bookExists {
		return errors.New("BOOK NOT FOUND")
	}
	if book.Status != "Available"{
		return errors.New("THE REQUIRED BOOK IS NOT AVAILABLE")
	}
	member, memberExists := l.members[memberID]
    if !memberExists {
        return errors.New("MEMBER NOT FOUND")
    }

	book.Status = "Borrowed"
	l.books[bookID] = book

	member.BorrowedBooks = append(member.BorrowedBooks, book)
    l.members[memberID] = member

	return nil
}

func (l *Library) ReturnBook(bookID int, memberID int) error{
	book, bookExists := l.books[bookID]
	if !bookExists {
		return errors.New("BOOK NOT FOUND")
	}
	if book.Status == "Available"{
		return errors.New("THIS BOOK IS ALREADY IN THE LIBRARY")
	}
	member, memberExists := l.members[memberID]
    if !memberExists {
        return errors.New("MEMBER NOT FOUND")
    }

	book.Status = "Available"
	l.books[bookID] = book

	for i, b := range member.BorrowedBooks {
		if b.ID == bookID {
			member.BorrowedBooks = append(member.BorrowedBooks[:i], member.BorrowedBooks[i+1:]...)
			break
		}
	} 
    l.members[memberID] = member

	return nil
}

func (l *Library) ListAvailableBooks() []models.Book{
	var available_books = make([]models.Book, 0)
	for _, book := range l.books {
		if book.Status == "Available" {
			available_books = append(available_books, book)
		}
	}
	return available_books
}

func (l *Library) ListBorrowedBooks(memberID int) []models.Book{
	member := l.members[memberID]
	return member.BorrowedBooks
}