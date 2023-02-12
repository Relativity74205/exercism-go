package techpalace

import (
	"fmt"
	"strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return fmt.Sprintf("Welcome to the Tech Palace, %s", strings.ToUpper(customer))
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	returnString := strings.Repeat("*", numStarsPerLine)
	returnString += fmt.Sprintf("\n%s\n", welcomeMsg)
	returnString += strings.Repeat("*", numStarsPerLine)
	return returnString
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	cleanedMsg := strings.Replace(oldMsg, "*", "", -1)
	return strings.TrimSpace(cleanedMsg)
}
