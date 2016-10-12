package levenshtein

import (
	"fmt"
	"testing"
)

type WordPair struct {
	x    []rune
	y    []rune
	dist uint
}

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
	fmt.Println("\ntesting \"apple\" and \"snapple\"")
	w := Weights{1, 1, 1}
	result := EditDistance([]rune("apple"), []rune("snapple"), w)
	if result != 2 {
		t.Errorf("d(apple, snapple) != 2,   result was: %d", result)
	}
}

func TestEditDistanceClownBrown(t *testing.T) {
	fmt.Println("\ntesting \"clown\" and \"brown\"")
	w := Weights{1, 1, 1}
	result := EditDistance([]rune("clown"), []rune("brown"), w)
	if result != 2 {
		t.Errorf("d(clown, brown) != 2,   result was: %d", result)
	}
}

func TestEditDistanceBookBurn(t *testing.T) {
	fmt.Println("\ntesting \"book\" and \"burn\"")
	w := Weights{1, 1, 1}
	result := EditDistance([]rune("book"), []rune("burn"), w)
	if result != 3 {
		t.Errorf("d(book, burn) != 3,   result was: %d", result)
	}
}

func TestEditDistanceKittenSitting(t *testing.T) {
	fmt.Println("\ntesting \"kitten\" and \"sitting\"")
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
		fmt.Printf("result  : %d\n\n", result)
		fmt.Printf("correct : %d\n\n", pair.dist)
	}
}
