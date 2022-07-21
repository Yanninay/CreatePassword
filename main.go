package main

import (
	"fmt"
	"math/rand"
	"time"
)

// main function
func main() {
	fmt.Println("Welcome to the password generator!")
	info() // print info
	rand.Seed(time.Now().UnixNano())
	// generate password
	password := generatePassword()
	// print password
	fmt.Println(password)
}

// generatePassword function
func generatePassword() string {
	var passwordLength int
	fmt.Print("Enter password length: ")
	fmt.Scan(&passwordLength)
	if passwordLength < 15 {
		fmt.Println("Please note, that the password length less than 15 characters is not recommended, because it's unsafe.")
	}
	password := randomString(passwordLength)
	return password
}

// randomString function
func randomString(n int) string {
	// create random string
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()_+{}|<>?-=[]\\;':,./~`"
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

// info function
func info() {
	fmt.Println("This is a simple password generator. It generates a random password with a length that you specify.")
	fmt.Println("If you've often wondered whether sites collect information about the passwords they generate, you've come to the right place. Here passwords are generated only on your device and are not sent to third-party sources")
	fmt.Println("Please note, that the password length less than 15 characters is not recommended, because it's unsafe.")
}
