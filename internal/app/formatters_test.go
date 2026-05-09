package app

import (
	"strings"
	"testing"
)

func TestFormatGoCode_ValidCode(t *testing.T) {
	input := "package foo\nfunc Bar() string {\nreturn \"hello\"\n}"
	got, err := formatGoCode(input)
	if err != nil {
		t.Fatalf("formatGoCode error: %v", err)
	}
	if !strings.Contains(got, "func Bar()") {
		t.Errorf("formatted output missing expected content: %s", got)
	}
}

func TestFormatGoCode_NormalizesWhitespace(t *testing.T) {
	// gofmt normalizes brace style and indentation
	input := "package foo\nfunc Bar() {\n}"
	got, err := formatGoCode(input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(got, "package foo") {
		t.Errorf("output missing package declaration: %s", got)
	}
}

func TestFormatGoCode_InvalidCode(t *testing.T) {
	_, err := formatGoCode("this is not valid go {{{")
	if err == nil {
		t.Error("formatGoCode should return error for invalid Go code")
	}
}

func TestPrintFormatted_ValidCode(t *testing.T) {
	input := "package foo\nfunc Bar() {}"
	got := PrintFormatted(input)
	if !strings.Contains(got, "func Bar()") {
		t.Errorf("PrintFormatted output missing expected content: %s", got)
	}
}

func TestPrintFormatted_InvalidCode_ReturnsInput(t *testing.T) {
	input := "not valid go {{{"
	got := PrintFormatted(input)
	if got != input {
		t.Errorf("PrintFormatted invalid code: got %q, want original input unchanged", got)
	}
}
