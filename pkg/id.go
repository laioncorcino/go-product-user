package pkg

import "github.com/google/uuid"

func NewID() string {
	return uuid.New().String()
}

func IsUUID(input string) bool {
	_, err := uuid.Parse(input)
	return err == nil
}
