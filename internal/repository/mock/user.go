package mock

import (
	"context"

	"github.com/stretchr/testify/mock"
	"github.com/witchs-lounge_backend/ent"
)

type UserRepository struct {
	mock.Mock
}

func (m *UserRepository) Create(ctx context.Context, email, password, name string) (*ent.User, error) {
	args := m.Called(ctx, email, password, name)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*ent.User), args.Error(1)
}

func (m *UserRepository) FindAll(ctx context.Context) ([]*ent.User, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*ent.User), args.Error(1)
}

func (m *UserRepository) FindByEmail(ctx context.Context, email string) (*ent.User, error) {
	args := m.Called(ctx, email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*ent.User), args.Error(1)
}
