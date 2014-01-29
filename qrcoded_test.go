package main

import (
	"bytes"
	"image/png"
	"testing"
)

func TestGenerateQRCodeReturnsValue(t *testing.T) {
	result := GenerateQRCode("555-2368")

	if result == nil {
		t.Errorf("Generated QRCode is nil")
	}
	if len(result) == 0 {
		t.Errorf("Generated QRCode has no data")
	}
}

func TestGenerateQRCodeGeneratesPNG(t *testing.T) {
	result := GenerateQRCode("555-2368")
	buffer := bytes.NewBuffer(result)
	_, err := png.Decode(buffer)

	if err != nil {
		t.Errorf("Generated QRCode is not a PNG: %s", err)
	}
}
