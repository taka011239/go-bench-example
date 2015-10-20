package closure

import (
	"testing"
)

func testHelper() []*Request {
	reqs := make([]*Request, 10000)
	for i := 0; i < len(reqs); i++ {
		reqs[i] = &Request{}
	}
	return reqs
}

func BenchmarkClosure1(b *testing.B) {
	reqs := testHelper()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Closure1(reqs)
	}
}

func BenchmarkClosure2(b *testing.B) {
	reqs := testHelper()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Closure2(reqs)
	}
}
