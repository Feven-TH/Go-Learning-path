<!-- Library Management System — Documentation -->

This is a simple CLI-based Library Management System written in Go. It allows managing books and members, borrowing and returning books, and viewing available or borrowed books.

<!-- Setup -->
Initialize Module: go mod init libraryapp
Run: go run main.go
Imports: Used libraryapp/... instead of relative imports for internal packages.

<!-- Structure -->
library_management/
├── controllers/    // User interaction (CLI)
├── services/       // Business logic
├── models/         // Data structures
├── main.go         // Entry point
└── go.mod          // Module definition


<!-- Use of Pointers -->
Structs use pointers (*) for shared state and efficiency instead of passing by value.


<!--Implemented Features -->

CLI Option                   Corresponding Method
1. Add Book             AddBook(book models.Book)
2. Add Member           AddMember(member models.Member)
3. Borrow Book          BorrowBook(bookID, memberID)
4. Return Book          ReturnBook(bookID, memberID)
5. Show Borrowed Books  ListBorrowedBooks(memberID)
6. Exit                 Program terminates

<!-- Flow of execution -->
User → Controller → Service → Model
               ↑          ↓
        (feedback)   (data structs)


<!-- Maintenance -->
Utilize pointers for shared state.
Isolate business logic in services.
Maintain clear CLI command mappings.
Extend features by adding service methods.