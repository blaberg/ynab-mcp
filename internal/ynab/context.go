package ynab

import "context"

type contextKey struct{}

// NewContext returns a new context with the given Client attached.
func NewContext(ctx context.Context, client *Client) context.Context {
	return context.WithValue(ctx, contextKey{}, client)
}

// FromContext extracts the Client from the context.
func FromContext(ctx context.Context) *Client {
	client, _ := ctx.Value(contextKey{}).(*Client)
	return client
}
