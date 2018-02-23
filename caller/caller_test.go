package caller

import (
	"testing"
	"fmt"
	"runtime"
)

func TestLine(t *testing.T) {
	fmt.Println(runtime.GOOS, runtime.GOARCH, runtime.Version())

	Line()
}

func BenchmarkLine(b *testing.B) {
	Line()
}