package exercises

import (
	"errors"
	"fmt"
	"os"
)

func Check(a, b int) error {
	if a == 0 && b == 0 {
		return errors.New("this is a custom error message")
	}

	return nil
}

func FormattedError(a, b int) error {
	if a == 0 && b == 0 {
		return fmt.Errorf("a %d and b %d. UserID %d", a, b, os.Geteuid())
	}

	return nil
}
