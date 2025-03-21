package fixedwidth

import (
	"bytes"
	"testing"

	"github.com/indece-official/go-ebcdic"
)

func TestEbcdicString_UnmarshalText(t *testing.T) {
	// Create a set of test cases for EBCDIC -> string conversion
	tests := []struct {
		name    string
		input   []byte
		want    string
		wantErr bool
	}{
		{
			name:    "ASCII characters",
			input:   []byte{0xC1, 0xC2, 0xC3}, // ABC in EBCDIC
			want:    "ABC",
			wantErr: false,
		},
		{
			name:    "Numbers",
			input:   []byte{0xF1, 0xF2, 0xF3}, // 123 in EBCDIC
			want:    "123",
			wantErr: false,
		},
		{
			name:    "Mixed characters",
			input:   []byte{0xC1, 0xF1, 0x5B}, // A1$ in EBCDIC ($ is 0x5B in EBCDIC037)
			want:    "A1$",
			wantErr: false,
		},
		{
			name:    "Empty input",
			input:   []byte{},
			want:    "",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &EbcdicString{}
			if err := s.UnmarshalText(tt.input); (err != nil) != tt.wantErr {
				t.Errorf("EbcdicString.UnmarshalText() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if s.S != tt.want {
				t.Errorf("EbcdicString.UnmarshalText() got = %q, want %q", s.S, tt.want)
			}
		})
	}
}

func TestEbcdicString_MarshalText(t *testing.T) {
	// Create a set of test cases for string -> EBCDIC conversion
	tests := []struct {
		name    string
		value   string
		want    []byte
		wantErr bool
	}{
		{
			name:    "ASCII characters",
			value:   "ABC",
			want:    []byte{0xC1, 0xC2, 0xC3}, // ABC in EBCDIC
			wantErr: false,
		},
		{
			name:    "Numbers",
			value:   "123",
			want:    []byte{0xF1, 0xF2, 0xF3}, // 123 in EBCDIC
			wantErr: false,
		},
		{
			name:    "Mixed characters",
			value:   "A1$",
			want:    []byte{0xC1, 0xF1, 0x5B}, // A1$ in EBCDIC ($ is 0x5B in EBCDIC037)
			wantErr: false,
		},
		{
			name:    "Empty string",
			value:   "",
			want:    []byte{},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := EbcdicString{S: tt.value}
			got, err := s.MarshalText()
			if (err != nil) != tt.wantErr {
				t.Errorf("EbcdicString.MarshalText() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !bytes.Equal(got, tt.want) {
				t.Errorf("EbcdicString.MarshalText() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEbcdicString_RoundTrip(t *testing.T) {
	// Test round-trip conversion (string -> EBCDIC -> string)
	testStrings := []string{
		"Hello, World!",
		"ABC123",
		"Special $#@",
		"",
	}

	for _, str := range testStrings {
		t.Run(str, func(t *testing.T) {
			original := EbcdicString{S: str}

			// Convert to EBCDIC
			ebcdicBytes, err := original.MarshalText()
			if err != nil {
				t.Fatalf("Failed to marshal: %v", err)
			}

			// Convert back to string
			decoded := &EbcdicString{}
			if err := decoded.UnmarshalText(ebcdicBytes); err != nil {
				t.Fatalf("Failed to unmarshal: %v", err)
			}

			// Compare result
			if decoded.S != str {
				t.Errorf("Round-trip conversion failed: got %q, want %q", decoded.S, str)
			}
		})
	}
}

func TestEbcdicString_WithFixedWidth(t *testing.T) {
	// Test EbcdicString with fixedwidth encoding/decoding
	t.Run("Marshal", func(t *testing.T) {
		type TestStruct struct {
			ID   int          `fixed:"1,5"`
			Name EbcdicString `fixed:"6,15,left,@"` // @ in Unicode is equivalent to space in EBCDIC
		}

		testData := TestStruct{
			ID:   12345,
			Name: EbcdicString{S: "Test"},
		}

		data, err := Marshal(testData)
		if err != nil {
			t.Fatalf("Failed to marshal: %v", err)
		}

		// Convert expected EBCDIC value for comparison
		expectedName, _ := ebcdic.Encode("Test", ebcdic.EBCDIC037)

		// The first 5 bytes should be "12345"
		firstField := data[:5]
		if !bytes.Equal(firstField, []byte("12345")) {
			t.Errorf("ID field not correctly encoded: got %v", firstField)
		}
		// The next 10 bytes should be the EBCDIC encoding of "Test" followed by spaces
		secondField := data[5:15]
		expectedNameWithSpaces := append(expectedName, bytes.Repeat([]byte{0x40}, 10-len(expectedName))...) // 0x40 is space in EBCDIC
		if !bytes.Equal(secondField, expectedNameWithSpaces) {
			t.Errorf("Name field not correctly encoded: got %v, want %v", secondField, expectedNameWithSpaces)
		}
		// Only check the first 4 bytes of the Name field (which should contain "Test" in EBCDIC)
		// We can't directly verify the EBCDIC in the output because Marshal encodes the result of MarshalText
		if len(data) < 9 {
			t.Fatalf("Output too short: %v", data)
		}
	})

	t.Run("Unmarshal", func(t *testing.T) {
		type TestStruct struct {
			ID   int          `fixed:"1,5"`
			Name EbcdicString `fixed:"6,15,left,@"` // without "left,@"" => padding will not be removed
		}

		// Create input with ASCII ID and EBCDIC name
		nameEbcdic, _ := ebcdic.Encode("Test", ebcdic.EBCDIC037)
		input := append([]byte("12345"), nameEbcdic...)
		input = append(input, bytes.Repeat([]byte{0x40}, 6)...) // 0x40 is space in EBCDIC

		var result TestStruct
		err := Unmarshal(input, &result)
		if err != nil {
			t.Fatalf("Failed to unmarshal: %v", err)
		}

		if result.ID != 12345 {
			t.Errorf("ID not correctly decoded: got %v, want %v", result.ID, 12345)
		}
		expectedString := "Test"
		if result.Name.S != expectedString {
			t.Errorf("Name not correctly decoded: got %q, want %q", result.Name.S, expectedString)
		}
	})
}
