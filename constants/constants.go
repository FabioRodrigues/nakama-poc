package constants

import "github.com/heroiclabs/nakama-common/runtime"

const FILESEEKERCOLLECTION = "files_seeker_collection"

// below defined here
// https://heroiclabs.com/docs/nakama/server-framework/go-runtime/#returning-errors-to-the-client
const (
	INTERNAL = 13
)

var (
	ErrInternalError = runtime.NewError("internal server error", INTERNAL)
)
