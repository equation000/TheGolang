package main
import (
	"testing"
)

func setup() {
//	os.Args = []string{"cmd", "hello", "world", "love"}
}
//only "go test -bench . -benchmem -v" this cmd only show echov3's result ,but echov1 and echov2 no show
func BenchmarkEchov1(b *testing.B) {
	setup()
	for i := 0; i < b.N; i++ {
		Echov1()
	}
}
func BenchmarkEchov2(b *testing.B) {
	setup()
	for i := 0; i < b.N; i++ {
		Echov2()
	}
}
func BenchmarkEchov3(b *testing.B) {
	setup()
	for i := 0; i < b.N; i++ {
		Echov3()
	}
}
