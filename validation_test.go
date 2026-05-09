package goqr

import (
	"fmt"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/KhaledSaeed004/goqr/render"
)

func TestDeterministic(t *testing.T) {
	opts := DefaultOptions()

	g1, _ := Generate("the cake is a lie", opts)
	g2, _ := Generate("the cake is a lie", opts)

	if !reflect.DeepEqual(g1, g2) {
		t.Fatal("QR generation is not deterministic")
	}
}

func TestNonEmpty(t *testing.T) {
	qrCode, _ := Generate("test", DefaultOptions())

	if len(qrCode.Grid) == 0 {
		t.Fatal("empty grid")
	}
}

func TestVersionScaling(t *testing.T) {
	short, _ := Generate("123", DefaultOptions())
	long, _ := Generate(strings.Repeat("A", 100), DefaultOptions())

	if len(long.Grid) <= len(short.Grid) {
		t.Fatal("version scaling failed")
	}
}

func TestModes(t *testing.T) {
	cases := []string{
		"1234567890",
		"HELLO123",
		"hello world",
	}

	for _, c := range cases {
		_, err := Generate(c, DefaultOptions())
		if err != nil {
			t.Fatalf("failed on %s", c)
		}
	}
}

func TestASCIIVisual(t *testing.T) {
	qrCode, _ := Generate("HELLO", DefaultOptions())
	out, _ := render.RenderASCII(qrCode.Grid, render.DefaultASCIIOptions(false))

	fmt.Println(out)
}

func TestASCIIOutput(t *testing.T) {
	qrCode, _ := Generate("HELLO", DefaultOptions())
	out, _ := render.RenderASCII(qrCode.Grid, render.DefaultASCIIOptions(false))

	expected, _ := os.ReadFile("testdata/ascii_hello.txt")

	if string(expected) != out {
		t.Fatal("output changed unexpectedly")
	}
}
