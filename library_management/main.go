package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"library_management/controllers"
	"library_management/services"
)

func main() {
	lib := services.NewLibrary()
	controller := controllers.NewLibraryController(lib)
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to the Library System!")
	for{
		fmt.Println("\nSelect an option:")
		fmt.Println("1. Add Book")
		fmt.Println("2. Add Member")
		fmt.Println("3. Borrow Book")
		fmt.Println("4. Return Book")
		fmt.Println("5. Show Available Books")
		fmt.Println("6. Show Borrowed Books")
		fmt.Println("7. Exit")
		fmt.Print("Enter choice (1-7): ")

		input, _ := reader.ReadString('\n')
		choice := strings.TrimSpace(input)

		switch choice {
		case "1":
			controller.AddBook()
		case "2":
			controller.AddMember()
		case "3":
			controller.BorrowBook()
		case "4":
			controller.ReturnBook()
		case "5":
			controller.ShowAvailableBooks()
		case "6":
			controller.ShowBorrowedBooks()
		case "7":
			fmt.Println("Exiting program. Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please enter a number from 1 to 7.")
		}
	}
}
