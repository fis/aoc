package day25

import "testing"

var algos = []struct {
	name string
	f    func(b, a, m int) int
}{
	{name: "trialMultiplication", f: trialMultiplication},
	{name: "babyStep", f: babyStep},
	{name: "pohligHellman", f: pohligHellman},
}

func TestFindKey(t *testing.T) {
	pub1, pub2 := 5764801, 17807724
	want := 14897079
	for _, algo := range algos {
		got := findKey(pub1, pub2, algo.f)
		if got != want {
			t.Errorf("findKey(%d, %d, %s) = %d, want %d", pub1, pub2, algo.name, got, want)
		}
	}
}

func BenchmarkAlgos(b *testing.B) {
	inB, inA, inM := 7, 9232416, 20201227
	for _, algo := range algos {
		b.Run(algo.name, func(b *testing.B) {
			for range b.N {
				algo.f(inB, inA, inM)
			}
		})
	}
}
