// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package db

import (
	"context"
)

type Querier interface {
	CreateURL(ctx context.Context, arg CreateURLParams) (URL, error)
}

var _ Querier = (*Queries)(nil)
