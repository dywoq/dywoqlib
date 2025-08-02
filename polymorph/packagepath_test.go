package polymorph

import (
	internal_testing "github.com/dywoq/dywoqlib/internal/testing"
	"io"
	"net/http"
	go_testing "testing"
)

func TestPackagePath(t *go_testing.T) {
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

func BenchmarkPackagePath(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	b.ReportAllocs()
	for b.Loop() {
		_ = PackagePath[http.Server]()
	}
}
