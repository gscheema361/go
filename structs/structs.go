package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser *user.User

	// appUser = user{
	// 	firstName: userFirstName,
	// 	lastName:  userLastName,
	// 	birthdate: userBirthdate,
	// 	createdAt: time.Now(),
	// }
	// appUser, err := newUser(userFirstName, userLastName, userBirthdate)
	appUser, err := user.New(userFirstName, userLastName, userBirthdate)
	// appUser = &user.User{
	// 	FirstName: "Max",
	// }

	if err != nil {
		fmt.Println(err)
		return
	}

	// ... do something awesome with that gathered data!

	// outputUserDetails(&appUser)
	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
