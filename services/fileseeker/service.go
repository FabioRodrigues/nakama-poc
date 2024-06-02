package fileseeker

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/fabiorodrigues/nakama-poc/constants"
	"github.com/fabiorodrigues/nakama-poc/dtos"
	"github.com/fabiorodrigues/nakama-poc/entities"
	"github.com/fabiorodrigues/nakama-poc/wrappers/ioadapter"
	"github.com/fabiorodrigues/nakama-poc/wrappers/nakamaruntime"
	"github.com/google/uuid"
	"github.com/heroiclabs/nakama-common/runtime"
	"os"
	"path/filepath"
)

type Service struct {
	ioAdapter     ioadapter.Provider
	nakamaRuntime nakamaruntime.Provider
}

func New(
	ioAdapter ioadapter.Provider,
	nakamaRuntime nakamaruntime.Provider) Provider {
	return Service{
		ioAdapter:     ioAdapter,
		nakamaRuntime: nakamaRuntime,
	}
}

func (s Service) Seek(ctx context.Context, request dtos.FileSeekerRequest) (*dtos.FileSeekerResponse, error) {
	filePath := filepath.Join(*request.Type, *request.Version+".json")
	fileContent, err := s.ioAdapter.ReadFile(filePath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, fmt.Errorf("file does not exist in path %s", filePath)
		}

		return nil, err
	}

	entity := entities.FileDetails{
		Type:    *request.Type,
		Version: *request.Version,
	}

	fileDetails, err := json.Marshal(entity)
	if err != nil {
		return nil, err
	}

	_, err = s.nakamaRuntime.StorageWrite(ctx, []*runtime.StorageWrite{
		{
			Collection: constants.FILESEEKERCOLLECTION,
			Key:        uuid.New().String(),
			Value:      string(fileDetails),
		},
	})
	if err != nil {
		return nil, fmt.Errorf("error writing file: %w", err)
	}

	content := ""
	var hash string

	if request.Hash != nil {
		hash = *request.Hash
		fileHash := sha256.Sum256(fileContent)
		calculatedHash := hex.EncodeToString(fileHash[:])
		if calculatedHash == *request.Hash {
			content = string(fileContent)
		}

	}

	return &dtos.FileSeekerResponse{
		Type:    *request.Type,
		Version: *request.Version,
		Content: content,
		Hash:    hash,
	}, nil

}
