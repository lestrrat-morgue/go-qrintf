package qrintf

import (
	"fmt"
	"testing"
)

func TestByte(t *testing.T) {
	buf := make([]byte, 3)
	(&ctx{buf, 0}).Byte('a').Byte('b').Byte('c').Finalize()

	if string(buf) != "abc" {
		t.Errorf("Expected '%#v', got '%#v'", []byte("abc"), buf)
	}
}

func TestPercentUInt(t *testing.T) {
	buf := make([]byte, 3)
	(&ctx{buf, 0}).UInt(100).Finalize()

	if string(buf) != "100" {
		t.Errorf("Expected '%#v', got '%#v'", []byte("100"), buf)
	}
}

const iter = 10000000
func BenchmarkIPv4Addr_qrintf(t *testing.B) {
	var buf = make([]byte, len("255.255.255.255"))
	var addr uint = 1234567890

	for i := 0; i < iter; i++ {
		(&ctx{buf, 0}).UInt((addr >> 24) & 0xff).Byte('.').UInt((addr >> 16) & 0xff).Byte('.').UInt((addr >> 8) & 0xff).Byte('.').UInt((addr) & 0xff).Finalize()
	}
}

func BenchmarkIPv4Addr_sprintf(t *testing.B) {
	var addr uint = 1234567890

	for i := 0; i < iter; i++ {
		x := fmt.Sprintf("%d.%d.%d.%d",
			(addr>>24)&0xff,
			(addr>>16)&0xff,
			(addr>>8)&0xff,
			(addr)&0xff,
		)
		_ = x
	}
}
