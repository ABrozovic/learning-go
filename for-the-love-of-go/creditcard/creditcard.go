package creditcard

import "errors"

type card struct {
	number string
}

func New(num string) (card, error) {
	if len(num) < 1 {
		return card{}, errors.New("credit card number can't be empty")
	}

	return card{number: num}, nil
}

func (c *card) Number() string {
	return c.number
}
