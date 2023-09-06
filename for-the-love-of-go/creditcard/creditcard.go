package creditcard

import "errors"

type Card struct {
	number string
}

func New(num string) (Card, error) {
	if len(num) < 1 {
		return Card{}, errors.New("credit card number can't be empty")
	}

	return Card{number: num}, nil
}

func (c *Card) Number() string {
	return c.number
}
