package gapp

import "context"

type TransportServer interface {
	Start(context.Context) error
	Stop(context.Context) error
}
