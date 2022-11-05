package mock

import "context"

type Person struct {
	ID      string
	Name    string
	Company string
	Website string
}

//go:generate moq -out mocks_test.go . PersonStore

type PersonStore interface {
	Get(ctx context.Context, id string) (*Person, error)
	Create(ctx context.Context, person *Person, confirm bool) error
}
