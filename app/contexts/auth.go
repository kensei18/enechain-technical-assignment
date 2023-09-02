package contexts

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

const userIDKeyName = "userId"

type userIDKey string

func GetUserID(ctx context.Context) (*uuid.UUID, error) {
	err := errors.New("failed to get userID from context")
	var key userIDKey = userIDKeyName
	value := ctx.Value(key)
	if value == nil {
		return nil, err
	}
	userID, ok := value.(uuid.UUID)
	if !ok {
		return nil, err
	}
	return &userID, nil
}

func WithUserID(ctx context.Context, userID uuid.UUID) context.Context {
	var key userIDKey = userIDKeyName
	return context.WithValue(ctx, key, userID)
}
