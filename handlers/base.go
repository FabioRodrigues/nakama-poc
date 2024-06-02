package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/fabiorodrigues/nakama-poc/constants"
)

func (h Handler) handleError(err error) (string, error) {
	errorMessage := fmt.Sprintf("there was an error processing your request: %s", err.Error())
	h.logger.Error(errorMessage)
	return "", constants.ErrInternalError
}

func (h Handler) handleSuccess(data any) (string, error) {
	if data == nil {
		return "", nil
	}

	response, err := json.Marshal(data)
	if err != nil {
		return h.handleError(err)
	}

	return fmt.Sprintf(`{"data":%s}`, response), nil
}

func isNullOrEmpty(s *string) bool {
	return s == nil || *s == ""
}

func convertToPointer[T any](t T) *T {
	return &t
}
