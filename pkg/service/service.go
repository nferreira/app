package service

import "context"

const (
	HeadersField = "Headers"
	BodyField = "Body"
)

type Service interface {
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
	CheckHealth(ctx context.Context) error
}
