package handlers

import (
	"context"
	"errors"
	"github.com/fabiorodrigues/nakama-poc/constants"
	"github.com/fabiorodrigues/nakama-poc/dtos"
	"github.com/fabiorodrigues/nakama-poc/services/fileseeker"
	"github.com/fabiorodrigues/nakama-poc/wrappers/logger"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandler_HandleSeekFile(t *testing.T) {
	t.Run("should set default params when not found in the payload", func(t *testing.T) {
		calledType := ""
		calledVersion := ""

		mockLogger := logger.Mock{}

		mockSvc := fileseeker.Mocks{
			SeekFn: func(ctx context.Context, request dtos.FileSeekerRequest) (*dtos.FileSeekerResponse, error) {
				calledType = *request.Type
				calledVersion = *request.Version
				return nil, nil
			},
		}

		handler := New(mockSvc, mockLogger)

		result, err := handler.HandleSeekFile(context.Background(), nil, nil, nil, "{}")

		assert.Equal(t, "core", calledType)
		assert.Equal(t, "1.0.0", calledVersion)
		assert.Equal(t, "{\"data\":null}", result)
		assert.NoError(t, err)
	})

	t.Run("should return data when success", func(t *testing.T) {

		mockLogger := logger.Mock{}
		mockSvc := fileseeker.Mocks{
			SeekFn: func(ctx context.Context, request dtos.FileSeekerRequest) (*dtos.FileSeekerResponse, error) {
				return &dtos.FileSeekerResponse{
					Type:    "core",
					Version: "1.0.0",
					Hash:    "123",
					Content: "{}",
				}, nil
			},
		}

		handler := New(mockSvc, mockLogger)

		result, err := handler.HandleSeekFile(context.Background(), nil, nil, nil, "{}")

		assert.Equal(t, `{"data":{"type":"core","version":"1.0.0","hash":"123","content":"{}"}}`, result)
		assert.NoError(t, err)
	})

	t.Run("should return error when failure", func(t *testing.T) {

		mockLogger := logger.Mock{}
		mockSvc := fileseeker.Mocks{
			SeekFn: func(ctx context.Context, request dtos.FileSeekerRequest) (*dtos.FileSeekerResponse, error) {
				return nil, errors.New("some error")
			},
		}

		handler := New(mockSvc, mockLogger)

		result, err := handler.HandleSeekFile(context.Background(), nil, nil, nil, "{}")

		assert.Empty(t, result)
		assert.Equal(t, err, constants.ErrInternalError)
	})

}
