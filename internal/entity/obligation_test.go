package entity

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestObligation_ToString(t *testing.T) {
	obligation := Obligation{
		ID:          1,
		Title:       "Test title",
		Description: "Description",
	}

	assert.Equal(
		t,
		fmt.Sprintf("ID: %d, Title: %s, Description: %s", obligation.ID, obligation.Title, obligation.Description),
		obligation.ToString(),
	)
}
