package polymorph

import (
	"io"
	"net/http"
	"testing"
)

func TestPackagePath(t *testing.T) {
	tests := []struct {
		body      string
		got, want string
	}{
		{"PackagePath[http.Server]()", PackagePath[http.Server](), "net/http"},
		{"PackagePath[io.Writer]()", PackagePath[io.Writer](), "io"},
	}

	for _, test := range tests {
		if test.got != test.want {
			t.Errorf("%s = %s, want %s", test.body, test.got, test.want)
		}
	}
}

func BenchmarkPackagePath(b *testing.B) {
	for b.Loop() {
		_ = PackagePath[http.Server]()
	}
}
