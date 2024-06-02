package ioadapter

type Mock struct {
	ReadFileFn func(path string) ([]byte, error)
}

func (m Mock) ReadFile(path string) ([]byte, error) {
	return m.ReadFileFn(path)
}
