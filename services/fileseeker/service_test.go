package fileseeker

import (
	"context"
	"errors"
	"github.com/fabiorodrigues/nakama-poc/dtos"
	"github.com/fabiorodrigues/nakama-poc/wrappers/ioadapter"
	"github.com/fabiorodrigues/nakama-poc/wrappers/nakamaruntime"
	"github.com/heroiclabs/nakama-common/api"
	"github.com/heroiclabs/nakama-common/runtime"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestService_Seek(t *testing.T) {
	t.Run("should return error when nakama could not save the data", func(t *testing.T) {
		mockIoAdapter := ioadapter.Mock{
			ReadFileFn: func(path string) ([]byte, error) {
				return []byte("hello world"), nil
			},
		}

		mockNakamaRuntime := nakamaruntime.Mocks{
			StorageWriteFn: func(ctx context.Context, writes []*runtime.StorageWrite) ([]*api.StorageObjectAck, error) {
				return nil, errors.New("some error")
			},
		}

		svc := New(mockIoAdapter, mockNakamaRuntime)

		_, err := svc.Seek(context.Background(), dtos.FileSeekerRequest{
			Version: convertToPointer("1.0.0"),
			Type:    convertToPointer("core"),
		})

		assert.Equal(t, "error writing file: some error", err.Error())
	})

	t.Run("should return no error when nakama saved the data correctly", func(t *testing.T) {
		mockIoAdapter := ioadapter.Mock{
			ReadFileFn: func(path string) ([]byte, error) {
				return []byte("hello world"), nil
			},
		}

		mockNakamaRuntime := nakamaruntime.Mocks{
			StorageWriteFn: func(ctx context.Context, writes []*runtime.StorageWrite) ([]*api.StorageObjectAck, error) {
				return nil, nil
			},
		}

		svc := New(mockIoAdapter, mockNakamaRuntime)

		_, err := svc.Seek(context.Background(), dtos.FileSeekerRequest{
			Version: convertToPointer("1.0.0"),
			Type:    convertToPointer("core"),
		})

		assert.NoError(t, err)
	})

	t.Run("should send data to nakama storage correctly", func(t *testing.T) {
		mockIoAdapter := ioadapter.Mock{
			ReadFileFn: func(path string) ([]byte, error) {
				return []byte("hello world"), nil
			},
		}

		calledData := ""
		mockNakamaRuntime := nakamaruntime.Mocks{
			StorageWriteFn: func(ctx context.Context, writes []*runtime.StorageWrite) ([]*api.StorageObjectAck, error) {
				if writes != nil {
					calledData = writes[0].Value
				}
				return nil, nil
			},
		}

		svc := New(mockIoAdapter, mockNakamaRuntime)

		_, err := svc.Seek(context.Background(), dtos.FileSeekerRequest{
			Version: convertToPointer("2.0.0"),
			Type:    convertToPointer("user"),
		})

		assert.NoError(t, err)
		assert.Equal(t, `{"type":"user","version":"2.0.0"}`, calledData)
	})

	t.Run("should return content when hashes match", func(t *testing.T) {
		mockIoAdapter := ioadapter.Mock{
			ReadFileFn: func(path string) ([]byte, error) {
				return []byte(`{"test": "123"}`), nil
			},
		}

		mockNakamaRuntime := nakamaruntime.Mocks{
			StorageWriteFn: func(ctx context.Context, writes []*runtime.StorageWrite) ([]*api.StorageObjectAck, error) {
				return nil, nil
			},
		}

		svc := New(mockIoAdapter, mockNakamaRuntime)
		expectedHash := "154fe4d3f4c5f5aa4ea0ffb77d18d78352b17b6304e5d0fbbd539080e8a9dcb0"

		result, err := svc.Seek(context.Background(), dtos.FileSeekerRequest{
			Version: convertToPointer("2.0.0"),
			Type:    convertToPointer("user"),
			Hash:    convertToPointer(expectedHash),
		})

		assert.NoError(t, err)
		assert.Equal(t, `{"test": "123"}`, result.Content)
	})

	t.Run("should not return content when hashes match", func(t *testing.T) {
		mockIoAdapter := ioadapter.Mock{
			ReadFileFn: func(path string) ([]byte, error) {
				return []byte(`{"test": "123"}`), nil
			},
		}

		mockNakamaRuntime := nakamaruntime.Mocks{
			StorageWriteFn: func(ctx context.Context, writes []*runtime.StorageWrite) ([]*api.StorageObjectAck, error) {
				return nil, nil
			},
		}

		svc := New(mockIoAdapter, mockNakamaRuntime)
		expectedHash := "abc"

		result, err := svc.Seek(context.Background(), dtos.FileSeekerRequest{
			Version: convertToPointer("2.0.0"),
			Type:    convertToPointer("user"),
			Hash:    convertToPointer(expectedHash),
		})

		assert.NoError(t, err)
		assert.Empty(t, result.Content)
	})
}

func convertToPointer[T any](t T) *T {
	return &t
}
