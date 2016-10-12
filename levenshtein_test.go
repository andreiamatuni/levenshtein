package levenshtein

import "testing"

func TestNewDMatrix(t *testing.T) {
	result := NewDMatrix(5, 5)
	if len(result) != 6 {
		t.Errorf("size of d: %d", len(result))
	}
}

func TestEditDistance(t *testing.T) {
	w := Weights{1, 1, 1}

	for _, pair := range pairs[:25] {
		result := EditDistance(pair.x, pair.y, w)
		if result != pair.dist {
			t.Errorf("\nresult : %d\ncorrect: %d\npair   : %s -- %s", result, pair.dist, string(pair.x), string(pair.y))
		}
	}
}

func TestCompactEditDistance(t *testing.T) {
	w := Weights{1, 1, 1}

	for _, pair := range pairs[:25] {
		result := CompactEditDistance(pair.x, pair.y, w)
		if result != pair.dist {
			t.Errorf("\nresult : %d\ncorrect: %d\npair   : %s -- %s", result, pair.dist, string(pair.x), string(pair.y))
		}
	}
}

func TestBufferedEditDistance(t *testing.T) {
	d := NewDMatrix(11, 11)
	w := Weights{1, 1, 1}

	for _, pair := range pairs[:25] {
		result := BufferedEditDistance(pair.x, pair.y, w, d)
		if result != pair.dist {
			t.Errorf("\nresult : %d\ncorrect: %d\npair   : %s -- %s", result, pair.dist, string(pair.x), string(pair.y))
		}
	}
}

func TestBufferedCompactDistance(t *testing.T) {
	w := Weights{1, 1, 1}

	prevRow := make([]uint, 16, 16)
	currRow := make([]uint, 16, 16)

	for _, pair := range pairs[:25] {
		result := BufferedCompactDist(pair.x, pair.y, w, prevRow, currRow)
		if result != pair.dist {
			t.Errorf("\nresult : %d\ncorrect: %d\npair   : %s -- %s", result, pair.dist, string(pair.x), string(pair.y))
		}
	}
}

func BenchmarkUnbuffered(b *testing.B) {
	w := Weights{1, 1, 1}

	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		for _, pair := range pairs {
			result := EditDistance(pair.x, pair.y, w)
			if result != pair.dist {

				b.Errorf("\nresult : %d\ncorrect: %d\npair   : %s -- %s", result, pair.dist, string(pair.x), string(pair.y))
			}
		}
	}
}

func BenchmarkBuffered(b *testing.B) {
	buffer := NewDMatrix(11, 11)
	w := Weights{1, 1, 1}

	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		for _, pair := range pairs {
			result := BufferedEditDistance(pair.x, pair.y, w, buffer)
			if result != pair.dist {

				b.Errorf("\nresult : %d\ncorrect: %d\npair   : %s -- %s", result, pair.dist, string(pair.x), string(pair.y))
			}
		}
	}
}

func BenchmarkUnbufferedCompact(b *testing.B) {
	w := Weights{1, 1, 1}

	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		for _, pair := range pairs {
			result := CompactEditDistance(pair.x, pair.y, w)
			if result != pair.dist {
				b.Errorf("\nresult : %d\ncorrect: %d\npair   : %s -- %s", result, pair.dist, string(pair.x), string(pair.y))
			}
		}
	}
}

func BenchmarkBufferedCompact(b *testing.B) {
	w := Weights{1, 1, 1}

	prevRow := make([]uint, 16, 16)
	currRow := make([]uint, 16, 16)

	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		for _, pair := range pairs {
			result := BufferedCompactDist(pair.x, pair.y, w, prevRow, currRow)
			if result != pair.dist {
				b.Errorf("\nresult : %d\ncorrect: %d\npair   : %s -- %s", result, pair.dist, string(pair.x), string(pair.y))
			}
		}
	}
}

func BenchmarkBasePairsUnbuffered(b *testing.B) {
	w := Weights{1, 1, 1}

	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		for _, pair := range BASEpairs {
			result := EditDistance(pair.x, pair.y, w)
			if result != pair.dist {

				b.Errorf("\nresult : %d\ncorrect: %d\npair   : %s -- %s", result, pair.dist, string(pair.x), string(pair.y))
			}
		}
	}
}

func BenchmarkBasePairsBuffered(b *testing.B) {
	buffer := NewDMatrix(36, 36)
	w := Weights{1, 1, 1}

	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		for _, pair := range BASEpairs {
			result := BufferedEditDistance(pair.x, pair.y, w, buffer)
			if result != pair.dist {

				b.Errorf("\nresult : %d\ncorrect: %d\npair   : %s -- %s", result, pair.dist, string(pair.x), string(pair.y))
			}
		}
	}
}

func BenchmarkBasePairsUnbufferedCompact(b *testing.B) {
	w := Weights{1, 1, 1}

	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		for _, pair := range BASEpairs {
			result := CompactEditDistance(pair.x, pair.y, w)
			if result != pair.dist {
				b.Errorf("\nresult : %d\ncorrect: %d\npair   : %s -- %s", result, pair.dist, string(pair.x), string(pair.y))
			}
		}
	}
}

func BenchmarkBasePairsBufferedCompact(b *testing.B) {
	w := Weights{1, 1, 1}

	prevRow := make([]uint, 37, 37)
	currRow := make([]uint, 37, 37)

	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		for _, pair := range BASEpairs {
			result := BufferedCompactDist(pair.x, pair.y, w, prevRow, currRow)
			if result != pair.dist {
				b.Errorf("\nresult : %d\ncorrect: %d\npair   : %s -- %s", result, pair.dist, string(pair.x), string(pair.y))
			}
		}
	}
}
