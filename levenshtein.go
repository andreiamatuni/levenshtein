package levenshtein

import "fmt"

/*
Weights are the costs of making a particular edit.
This struct functions as a configuration parameter
for the EditDistance functions
*/
type Weights struct {
	Insert     uint
	Delete     uint
	Substitute uint
}

/*
NewDMatrix returns a new distance matrix, given
the dimensions as paramter. All the memory is
preallocated.
*/
func NewDMatrix(x, y uint) [][]uint {
	d := make([][]uint, x+1, x+1)

	for i := range d {
		d[i] = make([]uint, y+1, y+1)
	}
	return d
}

/*
PrintMatrix prints the distance matrix
with all the rows lined up with each other
*/
func PrintMatrix(d [][]uint) {
	for _, row := range d {
		fmt.Println(row)
	}
}

func min(x, y, z uint) uint {
	min := x
	if y < min {
		min = y
	} else if z < min {
		min = z
	}
	return min
}

/*
EditDistance is the main function provided by this
package. It calculates the Levenshtein edit distance
between two strings provided as argument.
*/
func EditDistance(x, y []rune, w Weights) uint {
	d := NewDMatrix(uint(len(x)), uint(len(y)))

	for i := 0; i <= len(x); i++ {
		d[i][0] = uint(i)
	}

	for j := 0; j <= len(y); j++ {
		d[0][j] = uint(j)
	}
	for i := 1; i <= len(x); i++ {
		for j := 1; j <= len(y); j++ {
			if x[i-1] == y[j-1] {
				d[i][j] = d[i-1][j-1]
			} else {
				d[i][j] = min(d[i-1][j]+w.Delete,
					d[i][j-1]+w.Insert,
					d[i-1][j-1]+w.Substitute,
				)
			}
		}
	}
	// PrintMatrix(d)
	return d[len(x)][len(y)]
}

/*
CompactEditDistance calculates distance with O(m)
memory rather than O(mn).
*/
func CompactEditDistance(x, y []rune, w Weights) uint {
	prevRow := make([]uint, len(y)+1, len(y)+1)
	currRow := make([]uint, len(y)+1, len(y)+1)

	for i := 0; i <= len(y); i++ {
		prevRow[i] = uint(i)
	}
	for i := 1; i <= len(x); i++ {
		currRow[0] = uint(i)
		for j := 1; j <= len(y); j++ {
			if x[i-1] == y[j-1] {
				currRow[j] = prevRow[j-1]
			} else {
				currRow[j] = min(prevRow[j]+w.Delete,
					currRow[j-1]+w.Insert,
					prevRow[j-1]+w.Substitute,
				)
			}
		}
		copy(prevRow, currRow)
	}
	return currRow[len(y)]
}

/*
BufferedCompactDist computes the distance using only
two rows, rather than a whole distance matrix. These two
rows are passed in as parameter, so the function itself does
not allocate.
*/
func BufferedCompactDist(x, y []rune, w Weights, prevRow []uint, currRow []uint) uint {
	for i := 0; i <= len(y); i++ {
		prevRow[i] = uint(i)
	}
	for i := 1; i <= len(x); i++ {
		currRow[0] = uint(i)
		for j := 1; j <= len(y); j++ {

			if x[i-1] == y[j-1] {
				currRow[j] = prevRow[j-1]
			} else {
				currRow[j] = min(prevRow[j]+w.Delete,
					currRow[j-1]+w.Insert,
					prevRow[j-1]+w.Substitute,
				)
			}
		}
		copy(prevRow, currRow)
	}
	return currRow[len(y)]
}

/*
BufferedEditDistance functions the same as EditDistance,
except that it takes the matrix as parameter, acting as a
preallocated buffer. This is useful for hot loops where you're
calculating edit distance of many string pairs. You can reuse
the buffer instead of allocating/deallocating on every function call.

Be sure to preallocate a buffer large enough to accomodate the
size range of all the pairs whose edit distances you'll be computing
*/
func BufferedEditDistance(x, y []rune, w Weights, d [][]uint) uint {
	for i := 0; i <= len(x); i++ {
		d[i][0] = uint(i)
	}
	for j := 0; j <= len(y); j++ {
		d[0][j] = uint(j)
	}
	for i := 1; i <= len(x); i++ {
		for j := 1; j <= len(y); j++ {
			if x[i-1] == y[j-1] {
				d[i][j] = d[i-1][j-1]
			} else {
				d[i][j] = min(d[i-1][j]+w.Delete,
					d[i][j-1]+w.Insert,
					d[i-1][j-1]+w.Substitute,
				)
			}
		}
	}
	return d[len(x)][len(y)]
}
