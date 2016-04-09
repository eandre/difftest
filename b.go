import (
	"bufio"
	"os"
)

func Open(name string) (*Reader, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	return w.w.WriteString(p.r.r)
}
