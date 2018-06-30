package uidutils

import (
	"github.com/satori/go.uuid"
	"fmt"
)

func GenerateNewUniqueId() string {

	id, err := uuid.NewV4()

	if err != nil {
		fmt.Println("Something went wrong while generating a new unique id")
		return ""
	}

	return id.String()
}