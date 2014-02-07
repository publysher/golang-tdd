package main

import (
	"bytes"
	"errors"
	"image/png"
	"testing"
)

func TestGenerateQRCodeGeneratesPNG(t *testing.T) {
	buffer := new(bytes.Buffer)
	GenerateQRCode(buffer, "555-2368", Version(1))

	if buffer.Len() == 0 {
		t.Errorf("No QRCode generated")
	}

	_, err := png.Decode(buffer)

	if err != nil {
		t.Errorf("Generated QRCode is not a PNG: %s", err)
	}
}

type ErrorWriter struct{}

func (e *ErrorWriter) Write(b []byte) (int, error) {
	return 0, errors.New("Expected error")
}

func TestGenerateQRCodePropagatesErrors(t *testing.T) {
	w := new(ErrorWriter)
	err := GenerateQRCode(w, "555-2368", Version(1))

	if err == nil || err.Error() != "Expected error" {
		t.Errorf("Error not propagated correctly, got %v", err)
	}
}

func TestVersionDeterminesSize(t *testing.T) {
	table := []struct {
		version  int
		expected int
	}{
		{1, 21},
		{2, 25},
		{6, 41},
		{7, 45},
		{14, 73},
		{40, 177},
	}

	for _, test := range table {
		size := Version(test.version).PatternSize()
		if size != test.expected {
			t.Errorf("Version %2d, expected %3d but got %3d",
				test.version, test.expected, size)
		}
	}
}
