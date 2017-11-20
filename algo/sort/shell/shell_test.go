package shell

import (
	"sort"
	"testing"
	"math/rand"
	"time"
)

func TestSort(t *testing.T) {
	var tests = [][]int{
		{3, 5, 6, 8, 19},
		{6, 23, 12, 6},
	}
	for _, test := range tests {
		Sort(test)
		if !sort.IntsAreSorted(test) {
			t.Errorf("sort output %v", test)
		}
	}
}

func randomArray(rng *rand.Rand, size int) []int {
	array := make([]int, size)
	for i := 0; i < size; i++ {
		array[i] = rng.Intn(10000000)
	}
	return array
}

func TestSortRandom(t *testing.T) {
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < 100; i++ {
		a := randomArray(rng, 1000)
		Sort(a)
		if !sort.IntsAreSorted(a) {
			t.Errorf("sort output %v", a)
		}
	}
}

func BenchmarkSort(b *testing.B) {
	b.StopTimer()
	var seed = time.Now().UTC().UnixNano()
	var rng = rand.New(rand.NewSource(seed))
	a := randomArray(rng, 10000)
	data := make([]int, len(a))
	//sort.Sort(sort.Reverse(sort.IntSlice(a)))
	//sort.Ints(a)

	for i := 0; i < b.N; i++ {
		copy(data, a)
		b.StartTimer()
		Sort(data)
		b.StopTimer()
	}
}

func BenchmarkInsertSort(b *testing.B) {
	b.StopTimer()
	var seed = time.Now().UTC().UnixNano()
	var rng = rand.New(rand.NewSource(seed))
	a := randomArray(rng, 10000)
	data := make([]int, len(a))
	//sort.Sort(sort.Reverse(sort.IntSlice(a)))
	//sort.Ints(a)

	for i := 0; i < b.N; i++ {
		copy(data, a)
		b.StartTimer()
		InsertSort(data)
		b.StopTimer()
	}
}
