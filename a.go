import (
	"bufio"
	"os"
)

func Open(name string) (*Reader, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	return &Reader{f: f, Reader: bufio.NewReader(f)}, nil
}
