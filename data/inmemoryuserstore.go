package data

import "github.com/dgravesa/WaterLogger-UserServer/model"

// InMemoryUserStore is an in-memory container for user data.
type InMemoryUserStore struct {
	users []model.User
}

// NewInMemoryUserStore creates a new in-memory container for user data.
func NewInMemoryUserStore() *InMemoryUserStore {
	return new(InMemoryUserStore)
}

// Insert adds a new user to the data.
func (s *InMemoryUserStore) Insert(u model.User) {
	u.ID = uint64(len(s.users))
	s.users = append(s.users, u)
}
