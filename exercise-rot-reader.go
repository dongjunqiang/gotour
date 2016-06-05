import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot13 (a byte) byte {
	if (a >= 'a' && a <= 'm') || (a >= 'A' && a <= 'M') {
		a += 13
	} else if (a >= 'n' && a <= 'z') || (a >= 'N' && a <= 'Z') {
		a -= 13
	}
	return a
}
func (p rot13Reader) Read(b []byte) (int, error) {
	n, err := p.r.Read(b)
	for i:= 0; i < len(b); i += 1 {
		b[i] = rot13(b[i])
	}

	return n, err
}
func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
