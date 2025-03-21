package fixedwidth

import (
	"github.com/indece-official/go-ebcdic"
)

type EbcdicString struct {
	S string
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *EbcdicString) UnmarshalText(text []byte) error {
	decoded, err := ebcdic.Decode(text, ebcdic.EBCDIC037)
	if err != nil {
		return err
	}
	s.S = decoded
	return nil
}

// MarshalText implements encoding.TextUnmarshaler.
func (s EbcdicString) MarshalText() ([]byte, error) {
	encoded, err := ebcdic.Encode(s.S, ebcdic.EBCDIC037)
	if err != nil {
		return nil, err
	}
	return encoded, nil
}
