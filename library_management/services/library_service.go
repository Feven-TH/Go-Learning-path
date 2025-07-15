package services

import (
	"library_management/models" 
	"fmt"
)

type LibraryManager interface{
	AddBook(book models.Book)
	RemoveBook(bookID int)
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int,memberID int) error
	ListAvailableBooks()[]models.Book
	ListBorrowedBooks(memberID int) []models.Book
	AddMember(member models.Member) error

} 
type Library struct{
	Books map[int]*models.Book
	Members  map[int]*models.Member
}

func NewLibrary() *Library{
	return &Library{
		Books : make(map[int]*models.Book),
		Members : make(map[int]*models.Member),
	}
} 

func (l *Library) AddBook(book models.Book) error{
	if _,exists := l.Books[book.ID] ; exists{
		return fmt.Errorf("book already exists")
	}
	book.Status = "Available"
	l.Books[book.ID] = &book
	return nil
}
func (l *Library) AddMember(member models.Member) error {
	if _, exists := l.Members[member.ID]; exists {
		return fmt.Errorf("member with the given ID already exists")
	}
	l.Members[member.ID] = &member
	return nil
}


func (l *Library) RemoveBook(bookID int) error{
	if _,exists := l.Books[bookID] ; exists{
		delete(l.Books, bookID)
		return nil
	}
	return fmt.Errorf("book not fund")
}

func (l *Library) BorrowBook(bookID int, memberID int) error{
	book, bookExists := l.Books[bookID]
	member,member_exists := l.Members[memberID]
	if !bookExists{
		return fmt.Errorf("book doesn't exist")
	}
	if !member_exists{
		return fmt.Errorf("is not a member of the library ")
	}
	if l.Books[bookID].Status == "Borrowed"{
		return fmt.Errorf("sorry the book is currentl rented")
	}
	
	member.BorrowedBooks = append(member.BorrowedBooks, book)
	book.Status = "Borrowed" 
	return nil
}

func (l *Library) ReturnBook(bookID int,memberID int) error{
	member, member_exists := l.Members[memberID]
	if !member_exists {
		return fmt.Errorf("member does not exist")
	}
	found := -1
	for i,b := range(member.BorrowedBooks){
		if b.ID == bookID{
			found = i
			break
		}
	}
	if found == -1{
		return fmt.Errorf("member didn't borrow the book")
	}
	member.BorrowedBooks = append(member.BorrowedBooks[:found] ,member.BorrowedBooks[found+1:]...)
	book,book_exists := l.Books[bookID]
	if !book_exists{
		return fmt.Errorf("book doesn't exist")
	}
	book.Status = "Available"
	return nil
}

func (l *Library) ListAvailableBooks()[]models.Book{
	available := []models.Book{}
	for _, b := range l.Books{
		if b.Status == "Available"{
			available = append(available, *b)
		}
	}
	return available
}
func (l *Library) ListBorrowedBooks(memberID int) []models.Book {
	member, member_Exists := l.Members[memberID]
	if !member_Exists {
		return []models.Book{}
	}
	borrowed:= make([]models.Book, len(member.BorrowedBooks))
	for i, b := range member.BorrowedBooks{
		borrowed[i] = *b
	}
	return borrowed
}
