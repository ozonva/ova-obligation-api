package entity

import "fmt"

type Obligation struct {
	ID          uint
	Title       string
	Description string
}

func (o Obligation) String() string {
	return fmt.Sprintf("ID: %d, Title: %s, Description: %s", o.ID, o.Title, o.Description)
}
