package jetcloud

import (
	"context"

	"go.jetpack.io/pkg/auth/session"
)

func GetAccessToken(
	ctx context.Context,
	tok *session.Token,
) (string, error) {
	return newClient().getAccessToken(ctx, tok)
}
