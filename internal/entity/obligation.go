package entity

import "fmt"

type Obligation struct {
	ID          uint   `db:"id"`
	Title       string `db:"title"`
	Description string `db:"description"`
}

func (o Obligation) String() string {
	return fmt.Sprintf("ID: %d, Title: %s, Description: %s", o.ID, o.Title, o.Description)
}
