## Library Management System Console Interface

The Library Management System allows users to manage books and members in a library through a simple console interface. The following operations are supported:

### 1. Add a New Book

**Description**: Adds a new book to the library.

**Steps**:
1. Select option `1` from the menu.
2. Enter the book ID when prompted.
3. Enter the book title when prompted (multi-word titles are supported).
4. Enter the book author when prompted (multi-word author names are supported).

**Example**:
```
Enter book ID: 4
Enter book title: The Great Gatsby
Enter book author: F. Scott Fitzgerald
```

### 2. Remove an Existing Book

**Description**: Removes a book from the library using its ID.

**Steps**:
1. Select option `2` from the menu.
2. Enter the book ID when prompted.

**Example**:
```
Enter book ID to remove: 2
```

### 3. Borrow a Book

**Description**: Allows a member to borrow a book from the library.

**Steps**:
1. Select option `3` from the menu.
2. Enter the member ID when prompted.
3. Enter the book ID when prompted.

**Example**:
```
Enter member ID: 1
Enter book ID: 3
```

### 4. Return a Book

**Description**: Allows a member to return a borrowed book to the library.

**Steps**:
1. Select option `4` from the menu.
2. Enter the member ID when prompted.
3. Enter the book ID when prompted.

**Example**:
```
Enter member ID: 1
Enter book ID: 3
```

### 5. List All Available Books

**Description**: Lists all books in the library that are currently available.

**Steps**:
1. Select option `5` from the menu.

**Example Output**:
```
Available Books:
ID: 1, Title: 1984, Author: George Orwell
ID: 2, Title: To Kill a Mockingbird, Author: Harper Lee
ID: 3, Title: Harry Potter, Author: J.K. Rowling
```

### 6. List All Borrowed Books by a Member

**Description**: Lists all books currently borrowed by a specific member.

**Steps**:
1. Select option `6` from the menu.
2. Enter the member ID when prompted.

**Example Output**:
```
Enter member ID: 1
Borrowed Books by Member ID 1:
ID: 3, Title: Harry Potter, Author: J.K. Rowling
```

### 7. Exit

**Description**: Exits the Library Management System.

**Steps**:
1. Select option `7` from the menu.

**Example**:
```
Exiting...
```

### Example Session

Below is an example of how a typical session might look:

```
Library Management System:
For a better user experience, please increase the height of your console.
1. Add a new book
2. Remove an existing book
3. Borrow a book
4. Return a book
5. List all available books
6. List all borrowed books by a member
7. Exit
Enter your choice: 1

Enter book ID: 4
Enter book title: The Great Gatsby
Enter book author: F. Scott Fitzgerald
Book added successfully.

Enter your choice: 5

Available Books:
ID: 1, Title: 1984, Author: George Orwell
ID: 2, Title: To Kill a Mockingbird, Author: Harper Lee
ID: 3, Title: Harry Potter, Author: J.K. Rowling
ID: 4, Title: The Great Gatsby, Author: F. Scott Fitzgerald

Enter your choice: 7

Exiting...