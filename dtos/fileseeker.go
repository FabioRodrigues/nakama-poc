package dtos

type FileSeekerRequest struct {
	Type    *string `json:"type"`
	Version *string `json:"version"`
	Hash    *string `json:"hash"`
}

type FileSeekerResponse struct {
	Type    string `json:"type"`
	Version string `json:"version"`
	Hash    string `json:"hash"`
	Content string `json:"content"`
}
