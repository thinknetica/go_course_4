package bsearch

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

func TestSearches(t *testing.T) {
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

func sampleData() []int {
	rand.Seed(time.Now().UnixNano())
	var data []int
	for i := 0; i < 1_000_000; i++ {
		data = append(data, rand.Intn(1000))
	}

	sort.Slice(data, func(i, j int) bool { return data[i] < data[j] })
	return data
}

func BenchmarkBinary(b *testing.B) {
	data := sampleData()
	for i := 0; i < b.N; i++ {
		n := rand.Intn(1000)
		res := Binary(data, n)
		_ = res
	}
}

func BenchmarkSimple(b *testing.B) {
	data := sampleData()
	for i := 0; i < b.N; i++ {
		n := rand.Intn(1000)
		res := Simple(data, n)
		_ = res
	}
}
