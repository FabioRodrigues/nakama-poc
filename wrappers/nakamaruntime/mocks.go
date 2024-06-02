package nakamaruntime

import (
	"context"
	"github.com/heroiclabs/nakama-common/api"
	"github.com/heroiclabs/nakama-common/runtime"
)

type Mocks struct {
	StorageWriteFn func(ctx context.Context, writes []*runtime.StorageWrite) ([]*api.StorageObjectAck, error)
}

func (m Mocks) StorageWrite(ctx context.Context, writes []*runtime.StorageWrite) ([]*api.StorageObjectAck, error) {
	if m.StorageWriteFn != nil {
		return m.StorageWriteFn(ctx, writes)
	}
	return nil, nil
}
