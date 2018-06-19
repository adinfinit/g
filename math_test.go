package g

import (
	"math"
	"math/rand"
	"testing"
)

func absBits(v float32) float32 {
	return math.Float32frombits(math.Float32bits(v) &^ (1 << 31))
}

func absComp(v float32) float32 {
	if v < 0 {
		return -v
	}
	return v
}

func TestAbs(t *testing.T) {
	errcount := 0
	for i := 0; i < 10000; i++ {
		v := rand.Float32()*1e6 - 1e6/2
		bits := absBits(v)
		comp := absComp(v)
		if bits != comp {
			t.Errorf("different result %v: %v %v", v, bits, comp)
			errcount++
			if errcount >= 10 {
				t.Fail()
			}
		}
	}
}

const numberCount = 1024

var numbers = make([]float32, numberCount)
var positive = make([]float32, numberCount)

func init() {
	for i := range numbers {
		numbers[i] = rand.Float32()*1e6 - 1e6/2
		positive[i] = rand.Float32() * 1e6
	}
	numbers[0] = 0
}

func BenchmarkAbsBits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = absBits(numbers[i%numberCount])
	}
}

func BenchmarkAbsComp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = absComp(numbers[i%numberCount])
	}
}

func BenchmarkAbsBitsPositive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = absBits(positive[i%numberCount])
	}
}

func BenchmarkAbsCompPositive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = absComp(positive[i%numberCount])
	}
}
