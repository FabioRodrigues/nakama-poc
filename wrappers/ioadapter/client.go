package ioadapter

import "os"

type Client struct {
}

func New() Provider {
	return Client{}
}

func (c Client) ReadFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}
