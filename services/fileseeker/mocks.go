package fileseeker

import (
	"context"
	"github.com/fabiorodrigues/nakama-poc/dtos"
)

type Mocks struct {
	SeekFn func(ctx context.Context, request dtos.FileSeekerRequest) (*dtos.FileSeekerResponse, error)
}

func (m Mocks) Seek(ctx context.Context, request dtos.FileSeekerRequest) (*dtos.FileSeekerResponse, error) {
	return m.SeekFn(ctx, request)
}
