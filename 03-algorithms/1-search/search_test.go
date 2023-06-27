package bsearch

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

func TestSearch(t *testing.T) {
	tests := []struct {
		name string
		data []int
		item int
		want int
	}{
		{
			name: "#1",
			data: []int{1, 2, 4, 6, 10},
			item: 6,
			want: 3,
		},
		{
			name: "#2",
			data: []int{1, 2, 4, 5, 6, 10},
			item: 6,
			want: 4,
		},
		{
			name: "#3",
			data: []int{1, 2, 4, 5, 6, 10},
			item: 16,
			want: -1,
		},
		{
			name: "#4",
			data: []int{1, 2, 4, 5, 6, 10},
			item: 2,
			want: 1,
		},
		{
			name: "#5",
			data: []int{1, 2, 4, 5, 6, 10, 100},
			item: 4,
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Binary(tt.data, tt.item); got != tt.want {
				t.Errorf("Binary() = %v, want %v", got, tt.want)
			}
			if got := Simple(tt.data, tt.item); got != tt.want {
				t.Errorf("Simple() = %v, want %v", got, tt.want)
			}
		})
	}
}

func sampleData(n int) []int {
	rand.Seed(time.Now().UnixNano())
	var data []int
	for i := 0; i < n; i++ {
		data = append(data, rand.Intn(1000))
	}

	sort.Slice(data, func(i, j int) bool { return data[i] < data[j] })
	return data
}

func BenchmarkSimple(b *testing.B) {
	k := 1_000_000
	data := sampleData(k)
	for i := 0; i < b.N; i++ {
		n := rand.Intn(k)
		Simple(data, n)
	}
}

func BenchmarkBinary(b *testing.B) {
	k := 1_000_000
	data := sampleData(k)
	for i := 0; i < b.N; i++ {
		n := rand.Intn(k)
		Binary(data, n)
	}
}
