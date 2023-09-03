package contexts

import (
	"bytes"
	"context"
	"errors"

	"github.com/google/uuid"
)

const userIDKeyName = "userId"

type userIDKey string

func GetUserID(ctx context.Context) (uuid.UUID, error) {
	zeroUUID := uuid.UUID(bytes.Repeat([]byte("0"), 16))
	err := errors.New("failed to get userID from context")

	var key userIDKey = userIDKeyName
	value := ctx.Value(key)
	if value == nil {
		return zeroUUID, err
	}
	userID, ok := value.(uuid.UUID)
	if !ok {
		return zeroUUID, err
	}
	return userID, nil
}

func WithUserID(ctx context.Context, userID uuid.UUID) context.Context {
	var key userIDKey = userIDKeyName
	return context.WithValue(ctx, key, userID)
}
