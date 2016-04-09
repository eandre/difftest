import (
	"bufio"
	"os"
)

func Open(name string) (*Reader, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	return &Reader{f: f, r: bufio.NewReader(f)}, nil
}

func (w *Writer) Write(p []byte) (int, error) {
	return w.w.Write(p)
}

func (w *Writer) WriteString(p string) (int, error) {
	return w.w.WriteString(p)
}
