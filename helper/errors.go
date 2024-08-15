package helper

import (
	"fmt"
	"os"
)

// HandleErrorExit prints an error message and exits the program if an error is not nil
func HandleErrorExit(msg string, err error) {
	if err != nil {
		errorMsgWithError(msg, err)
		os.Exit(1)
	}
}

// HandleErrorExitStr prints an error message and exits the program if a string is not empty
func HandleErrorExitStr(msg string, str string) {
	if str != "" {
		errorMsgWithStr(msg, str)
		os.Exit(1)
	}
}

// HandleError prints an error message if an error is not nil
func HandleError(msg string, err error) {
	if err != nil {
		errorMsgWithError(msg, err)
	}
}

// HandleErrorStr prints an error message if a string is not empty
func HandleErrorStr(msg string, str string) {
	if str != "" {
		errorMsgWithStr(msg, str)
	}
}

// errorMsgWithError prints an error message with an error
func errorMsgWithError(msg string, err error) {
	if msg == "" {
		msg = "Error"
	}

	if err != nil {
		fmt.Printf("%s: %v\n", msg, err)
	}
}

// errorMsgWithStr prints an error message with a string
func errorMsgWithStr(msg string, str string) {
	if msg == "" {
		msg = "Error"
	}

	if str != "" {
		fmt.Printf("\n%s: %s\n", msg, str)
	}
}
