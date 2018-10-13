package main

import (
	"fmt"

	"github.com/SlyMarbo/gmail"
)

func main() {
	email := gmail.Compose("Test sending email", "Email arrived!")
	email.From = "email@gmail.com"
	email.Password = "password"
	email.ContentType = "text/html; charset=utf-8"
	email.AddRecipient("email@gmail.com")
	err := email.Send()
	if err != nil {
		fmt.Printf("erro")
	}
	fmt.Printf("Email enviado com sucesso!")
}
