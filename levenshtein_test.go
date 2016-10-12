package levenshtein

import "testing"

type WordPair struct {
	x    []rune
	y    []rune
	dist uint
}

// There are 20 pairs
var pairs = []WordPair{
	WordPair{[]rune("apple"), []rune("snapple"), 2},
	WordPair{[]rune("book"), []rune("burn"), 3},
	WordPair{[]rune("clown"), []rune("brown"), 2},
	WordPair{[]rune("kitten"), []rune("sitting"), 3},
	WordPair{[]rune("computer"), []rune("maple"), 6},
	WordPair{[]rune("keyboard"), []rune("bread"), 6},
	WordPair{[]rune("pencil"), []rune("telephone"), 8},
	WordPair{[]rune("paper"), []rune("pipes"), 2},
	WordPair{[]rune("display"), []rune("yardsale"), 7},
	WordPair{[]rune("cable"), []rune("stable"), 2},
	WordPair{[]rune("lecture"), []rune("college"), 6},
	WordPair{[]rune("bottle"), []rune("throttle"), 3},
	WordPair{[]rune("desk"), []rune("drawing"), 6},
	WordPair{[]rune("microwave"), []rune("waffles"), 9},
	WordPair{[]rune("memory"), []rune("morty"), 3},
	WordPair{[]rune("flowers"), []rune("wolf"), 6},
	WordPair{[]rune("carpet"), []rune("caterer"), 4},
	WordPair{[]rune("soup"), []rune("parsley"), 6},
	WordPair{[]rune("porridge"), []rune("prestige"), 5},
	WordPair{[]rune("hickory"), []rune("crabapple"), 9},
}

func TestMin(t *testing.T) {
	x := uint(0)
	y := uint(1)
	z := uint(2)

	for i := 0; i < 25; i++ {
		minimum := min(x, y, z)
		if minimum != uint(i) {
			t.Errorf("min = %d,  x=%d, y=%d, z=%d", minimum, x, y, z)
		}
		x++
		y++
		z++
	}
}

func TestNewDMatrix(t *testing.T) {
	result := NewDMatrix(5, 5)
	if len(result) != 6 {
		t.Errorf("size of d: %d", len(result))
	}
}

func TestCreateDMatrix(t *testing.T) {
	result := createDMatrix([]rune("apple"), []rune("snapple"))
	if len(result) != 6 {
		t.Errorf("len(result) == %d,   supposed to be 6", len(result))
	}
}

func TestEditDistanceAppleSnapple(t *testing.T) {
	w := Weights{1, 1, 1}
	result := EditDistance([]rune("apple"), []rune("snapple"), w)
	if result != 2 {
		t.Errorf("d(apple, snapple) != 2,   result was: %d", result)
	}
}

func TestEditDistanceClownBrown(t *testing.T) {
	w := Weights{1, 1, 1}
	result := EditDistance([]rune("clown"), []rune("brown"), w)
	if result != 2 {
		t.Errorf("d(clown, brown) != 2,   result was: %d", result)
	}
}

func TestEditDistanceBookBurn(t *testing.T) {
	w := Weights{1, 1, 1}
	result := EditDistance([]rune("book"), []rune("burn"), w)
	if result != 3 {
		t.Errorf("d(book, burn) != 3,   result was: %d", result)
	}
}

func TestEditDistanceKittenSitting(t *testing.T) {
	w := Weights{1, 1, 1}
	result := EditDistance([]rune("kitten"), []rune("sitting"), w)
	if result != 3 {
		t.Errorf("d(kitten, sitting) != 3,   result was: %d", result)
	}

}

func TestBufferedEditDistance(t *testing.T) {
	d := NewDMatrix(10, 10)
	w := Weights{1, 1, 1}

	for _, pair := range pairs {
		result := BufferedEditDistance(pair.x, pair.y, w, d)
		if result != pair.dist {
			t.Errorf("\nresult : %d\ncorrect: %d\npair   : %s -- %s", result, pair.dist, string(pair.x), string(pair.y))
		}
	}
}

func BenchmarkUnBufferedDistance(b *testing.B) {
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

func BenchmarkBufferedDistance(b *testing.B) {
	buffer := NewDMatrix(10, 10)
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
