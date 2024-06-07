package techpalace

import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

	return strings.Repeat("*", numStarsPerLine) + "\n" + welcomeMsg + "\n" + strings.Repeat("*", numStarsPerLine)
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	panic("Please implement the AddBorder() function")
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	panic("Please implement the CleanupMessage() function")
}
