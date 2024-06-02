package handlers

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/fabiorodrigues/nakama-poc/dtos"
	"github.com/fabiorodrigues/nakama-poc/services/fileseeker"
	"github.com/fabiorodrigues/nakama-poc/wrappers/logger"
	"github.com/heroiclabs/nakama-common/runtime"
)

type Handler struct {
	fileSeekerService fileseeker.Provider
	logger            logger.Provider
}

func New(
	fileSeekerService fileseeker.Provider,
	logger logger.Provider) Handler {
	return Handler{
		fileSeekerService: fileSeekerService,
		logger:            logger,
	}
}

func (h Handler) HandleSeekFile(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, payload string) (string, error) {
	var request dtos.FileSeekerRequest
	if err := json.Unmarshal([]byte(payload), &request); err != nil {
		return h.handleError(err)
	}

	if isNullOrEmpty(request.Type) {
		request.Type = convertToPointer("core")
	}

	if isNullOrEmpty(request.Version) {
		request.Version = convertToPointer("1.0.0")
	}

	response, err := h.fileSeekerService.Seek(ctx, request)
	if err != nil {
		return h.handleError(err)
	}

	return h.handleSuccess(response)
}
