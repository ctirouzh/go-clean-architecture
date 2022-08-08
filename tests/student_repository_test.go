package main

import (
	"testing"

	"github.com/tahadostifam/go-clean-architecture/database"
)

func TestCreateStudent(t *testing.T) {
	database.GetSqliteInstance()
}
