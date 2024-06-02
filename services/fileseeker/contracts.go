package fileseeker

import (
	"context"
	"github.com/fabiorodrigues/nakama-poc/dtos"
)

type Provider interface {
	Seek(ctx context.Context, request dtos.FileSeekerRequest) (*dtos.FileSeekerResponse, error)
}
