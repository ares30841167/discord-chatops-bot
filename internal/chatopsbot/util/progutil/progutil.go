package progutil

import "log"

// Check for errors
func CheckErrorOccurred(err error) {
	// If an error occurs, print the error and exit
	if err != nil {
		log.Fatal(err)
	}
}

// Check for errors
func ThrowPanicIfErrorOccured(err error) {
	// If an error occurs, throw panic
	if err != nil {
		panic(err)
	}
}

// Catch panic
func LogErrorMessageIfPanic() {
	// If a panic occurs, print the error message and recover
	if x := recover(); x != nil {
		log.Println(x)
	}
}
