package controllers

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"library_management/models" 
	"library_management/services"
)

type LibraryController struct{
	Library *services.Library
}

func NewLibraryController(lib *services.Library) *LibraryController{
	return &LibraryController{Library: lib}
}
func (lc *LibraryController) AddBook(){
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter book ID: ")
	_id, _ := reader.ReadString('\n')
	id,err := strconv.Atoi(strings.TrimSpace(_id))
	if err != nil {
		fmt.Println("Invalid book ID.")
		return
	}

	fmt.Print("Enter book title: ")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)

	fmt.Print("Enter author name: ")
	author, _ := reader.ReadString('\n')
	author = strings.TrimSpace(author)

	book := models.Book{
		ID:     id,
		Title:  title,
		Author: author,
		Status: "Available",
	}

	err = lc.Library.AddBook(book)
	if err != nil {
		fmt.Println("Error adding book:", err)
	} else {
		fmt.Println("Book added successfully.")
	}
}

func (lc *LibraryController) AddMember() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter member ID: ")
	idStr, _ := reader.ReadString('\n')
	id, _ := strconv.Atoi(strings.TrimSpace(idStr))

	fmt.Print("Enter member name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	member := models.Member{ID: id, Name: name}
	err := lc.Library.AddMember(member)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Member added successfully.")
	}
}

func (lc *LibraryController) BorrowBook(){
	reader:= bufio.NewReader(os.Stdin)
	fmt.Print("Enter book ID: ")
	_id, _ := reader.ReadString('\n')
	bookID,err := strconv.Atoi(strings.TrimSpace(_id))
	
	if err != nil{
		fmt.Print("Invalid Book ID")
		return 
	}

	fmt.Print("Enter member ID: ")
	m_ID, _ := reader.ReadString('\n')
	memberID, err := strconv.Atoi(strings.TrimSpace(m_ID))
	if err != nil{
		fmt.Print("Invalid member ID")
		return
	}
	err = lc.Library.BorrowBook(bookID, memberID)
	if err != nil{
		fmt.Println("Error:", err)
	} else{
		fmt.Println("Book borrowed successfully.")
	}
}

func (lc *LibraryController) ReturnBook() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter book ID: ")
	bookIDStr, _ := reader.ReadString('\n')
	bookID, _ := strconv.Atoi(strings.TrimSpace(bookIDStr))

	fmt.Print("Enter member ID: ")
	memberIDStr, _ := reader.ReadString('\n')
	memberID, _ := strconv.Atoi(strings.TrimSpace(memberIDStr))

	err := lc.Library.ReturnBook(bookID, memberID)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Book returned successfully.")
	}
}
func (lc *LibraryController) ShowBorrowedBooks() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter member ID: ")
	idStr, _ := reader.ReadString('\n')
	id, _ := strconv.Atoi(strings.TrimSpace(idStr))

	books := lc.Library.ListBorrowedBooks(id)
	if len(books) == 0 {
		fmt.Println("No borrowed books.")
		return
	}
	fmt.Println("Borrowed Books:")
	for _, b := range books {
		fmt.Printf("ID: %d | Title: %s | Author: %s\n", b.ID, b.Title, b.Author)
	}
}

func (lc *LibraryController) ShowAvailableBooks() {
	books := lc.Library.ListAvailableBooks()
	if len(books) == 0 {
		fmt.Println("No books available.")
		return
	}
	fmt.Println("Available Books:")
	for _, b := range books {
		fmt.Printf("ID: %d | Title: %s | Author: %s\n", b.ID, b.Title, b.Author)
	}
}

