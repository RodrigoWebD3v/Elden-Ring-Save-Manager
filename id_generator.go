package main

import (
	"github.com/google/uuid"
)

func GerarId() string {
	id := uuid.NewString()
	return id
}
