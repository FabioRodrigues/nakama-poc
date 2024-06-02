package ioadapter

type Provider interface {
	ReadFile(path string) ([]byte, error)
}
