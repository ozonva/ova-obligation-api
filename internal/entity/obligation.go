package entity

import "fmt"

type Obligation struct {
	ID          uint
	Title       string
	Description string
}

func (o Obligation) ToString() string {
	return fmt.Sprintf("ID: %d, Title: %s, Description: %s", o.ID, o.Title, o.Description)
}
