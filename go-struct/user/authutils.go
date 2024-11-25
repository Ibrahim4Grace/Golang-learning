package user

import (
	"fmt"
	"time"
)


type user struct {
	firstName string
	lastName string
	birthdate string
	createAt time.Time
}

//Struct embedding 
type admin struct {
	email string
	password string
	user 
}

func NewAdmin(email, password string) admin {
	return admin{
		email: email,
		password: password,
		user: user{
			firstName: "Admin",
			lastName: "Admin",
			birthdate: "11/11/1990", 
			createAt: time.Now(),
		 },
	}
}


//Creation || constructor function best to use
func New(firstName, lastName, birthdate string) user {
	// we can add check here if we want
	return user{
		firstName,
		lastName,
		birthdate,
		time.Now(),
	}
}

func OutPutUserDetails(u user){
	fmt.Println(u.firstName, u.lastName, u.birthdate, u.createAt)
}
