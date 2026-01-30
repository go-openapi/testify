package leak

import (
	"bufio"
	"io"
)

// scanner provides a bufio.Scanner the ability to Unscan,
// which allows the current token to be read again
// after the next Scan.
type scanner struct {
	*bufio.Scanner

	unscanned bool
}

func newScanner(r io.Reader) *scanner {
	return &scanner{Scanner: bufio.NewScanner(r)}
}

func (s *scanner) Scan() bool {
	if s.unscanned {
		s.unscanned = false
		return true
	}
	return s.Scanner.Scan()
}

// Unscan stops the scanner from advancing its position
// for the next Scan.
//
// Bytes and Text will return the same token after next Scan
// that they do right now.
func (s *scanner) Unscan() {
	s.unscanned = true
}
