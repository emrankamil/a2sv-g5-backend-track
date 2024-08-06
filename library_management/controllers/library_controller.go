package controllers

import (
    "fmt"
    "library_management/models"
    "library_management/services"
    "os"
)

func Console() {
    library := services.NewLibrary()

    // Adding some initial data
    library.AddBook(models.Book{ID: 1, Title: "1984", Author: "George Orwell", Status: "Available"})
    library.AddBook(models.Book{ID: 2, Title: "To Kill a Mockingbird", Author: "Harper Lee", Status: "Available"})
    library.AddMember(models.Member{ID: 1, Name: "Emran "})
    library.AddMember(models.Member{ID: 2, Name: "Kamil"})

    for {
        fmt.Println("\nLibrary Management System: For better user experience, increase your console height.")
        fmt.Println("1. Add a new book")
        fmt.Println("2. Remove an existing book")
        fmt.Println("3. Borrow a book")
        fmt.Println("4. Return a book")
        fmt.Println("5. List all available books")
        fmt.Println("6. List all borrowed books by a member")
        fmt.Println("7. Exit")
        fmt.Print("Enter your choice: ")

        var choice int
        fmt.Scan(&choice)

        switch choice {
        case 1:
            var id int
            var title, author string
            fmt.Print("Enter book ID: ")
            fmt.Scan(&id)
            fmt.Print("Enter book title: ")
            fmt.Scan(&title)
            fmt.Print("Enter book author: ")
            fmt.Scan(&author)

			fmt.Println("\n Result: ")
            addbook:= library.AddBook(models.Book{ID: id, Title: title, Author: author, Status: "Available"})
			if addbook != nil{
				fmt.Println("Book cannot be added: id already exists")
			}else{
				fmt.Println("Book added successfully.")
			}

        case 2:
            var id int
            fmt.Print("Enter book ID to remove: ")
            fmt.Scan(&id)
            err := library.RemoveBook(id)
			fmt.Println("\n Result: ")
            if err != nil {
                fmt.Println(err)
            } else {
                fmt.Println("Book removed successfully.")
            }

        case 3:
            var memberID, bookID int
            fmt.Print("Enter member ID: ")
            fmt.Scan(&memberID)
            fmt.Print("Enter book ID: ")
            fmt.Scan(&bookID)
            err := library.BorrowBook(memberID, bookID)
            if err != nil {
                fmt.Println(err)
            } else {
                fmt.Println("Book borrowed successfully.")
            }

        case 4:
            var memberID, bookID int
            fmt.Print("Enter member ID: ")
            fmt.Scan(&memberID)
            fmt.Print("Enter book ID: ")
            fmt.Scan(&bookID)
            err := library.ReturnBook(memberID, bookID)
            if err != nil {
                fmt.Println(err)
            } else {
                fmt.Println("Book returned successfully.")
            }

        case 5:
            available_books := library.ListAvailableBooks()
			fmt.Println("\n Available books: ")
			fmt.Println(available_books)

        case 6:
            var memberID int
            fmt.Print("Enter member ID: ")
            fmt.Scan(&memberID)
			fmt.Println("\n Result: ")
            err := library.ListBorrowedBooks(memberID)
            fmt.Println(err)

        case 7:
            fmt.Println("Exiting...")
            os.Exit(0)

        default:
            fmt.Println("Invalid choice, please try again.")
        }
    }
}
