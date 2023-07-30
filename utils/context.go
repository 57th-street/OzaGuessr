package utils

import "context"

type userIDKey struct{}

func SetUserID(ctx context.Context, userID int) context.Context {
	return context.WithValue(ctx, userIDKey{}, userID)
}

func GetUserID(ctx context.Context) int {
	userID := ctx.Value(userIDKey{})

	if userID, ok := userID.(int); ok {
		return userID
	}
	return 0
}
