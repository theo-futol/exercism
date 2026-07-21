package techpalace
import (
	"strings"
	"regexp"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
    return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	return strings.Repeat("*", numStarsPerLine) + "\n" + welcomeMsg + "\n" + strings.Repeat("*", numStarsPerLine)
}

var cleanupRE = regexp.MustCompile(`[\t\n*]+`)

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	return strings.TrimSpace(cleanupRE.ReplaceAllString(oldMsg, ""))
}
