package main

import "fmt"

// Custom type created using struct

type Contact struct {
	ID    int
	Name  string
	Email string
	Phone string
}

var contactList []Contact
var contactIndexByName map[string]int
var nextID int = 1

func init() {
	contactList = make([]Contact, 0)
	contactIndexByName = make(map[string]int)
}

func addContact(name, email, phone string) {
	if _, exists := contactIndexByName[name]; exists {
		fmt.Printf("Contact already exists: %v\n", name)
		return
	}
	newContact := Contact{
		ID:    nextID,
		Name:  name,
		Email: email,
		Phone: phone,
	}
	nextID++
	contactList = append(contactList, newContact)
	contactIndexByName[name] = len(contactList) - 1
	fmt.Printf("Contact added: %v\n", name)
}

func findContactByName(name string) *Contact {
	if index, exists := contactIndexByName[name]; exists {
		return &contactList[index]
	}
	return nil
}

func listContacts() {
	fmt.Println("---Listing Contacts---")
	if len(contactList) == 0 {
		fmt.Println("No contact found")
	}
	for i, contact := range contactList {
		fmt.Printf("%d\nID: %d\nName: %s\nEmail: %s\nPhone: %s\n", i+1, contact.ID, contact.Name, contact.Email, contact.Phone)
		fmt.Println()
	}
}

func main() {
	fmt.Println("---Contact Management System---")
	addContact("Alice Wonderland", "alice@example.com", "111-2222")
	addContact("Bob The Builder", "bob@example.com", "333-4444")
	addContact("Charlie Brown", "charlie@example.com", "555-6666")
	addContact("Alice Wonderland", "alice.new@example.com", "777-8888") // Attempt to add duplicate

	listContacts()

	bob := findContactByName("Bob")
	if bob == nil {
		fmt.Println("No Bob contact found.")
	} else {
		fmt.Println("Bob contact found.", bob.Name)
	}
}
