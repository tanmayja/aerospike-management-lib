package lib

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type InfoParser struct {
	*bufio.Reader
}

func NewInfoParser(s string) *InfoParser {
	return &InfoParser{bufio.NewReader(strings.NewReader(s))}
}

// PeekAndExpect checks if the expected value is present without advancing the reader
func (ip *InfoParser) PeekAndExpect(s string) error {
	bytes, err := ip.Peek(len(s))
	if err != nil {
		return err
	}

	v := string(bytes)
	if v != s {
		return fmt.Errorf("InfoParser: Wrong value. Peek expected %s, but found %s", s, v)
	}

	return nil
}

func (ip *InfoParser) Expect(s string) error {
	buf := make([]byte, len(s))
	if _, err := ip.Read(buf); err != nil {
		return err
	}

	if sbuf := string(buf); sbuf != s {
		return fmt.Errorf("expected value %q found %q", s, sbuf)
	}

	return nil
}

func (ip *InfoParser) ReadUntil(delim byte) (string, error) {
	v, err := ip.ReadBytes(delim)

	switch len(v) {
	case 0:
		return string(v), err
	case 1:
		if v[0] == delim {
			return "", err
		}

		return string(v), err
	}

	return string(v[:len(v)-1]), err
}

func (ip *InfoParser) ReadFloat(delim byte) (float64, error) {
	s, err := ip.ReadUntil(delim)
	if err != nil {
		return 0, err
	}

	return strconv.ParseFloat(s, 64)
}
