package main

import (
	"fmt"
	"time"
)

type PathError struct {
	Op   string // "open", "unlink", etc.
	Path string // The associated file.
	Err  error  // Returned by the system call.
}

func (e *PathError) Error() string {
	return e.Op + " " + e.Path + ": " + e.Err.Error()
}

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

// func main() {
// 	err := fmt.Errorf("error occurred")
// 	fmt.Println(err)
// 	fmt.Println(err.Error())

// 	if err := run(); err != nil {
// 		fmt.Println(err)
// 	}
// }
