package matrix

import (
	"testing"
)

func TestNew(t *testing.T) {
	// replaced new assertions helper

	m := New(10, 5)
	rows, cols := m.Size()
	gotest.AssertEqual(t, 10, rows)
	gotest.AssertEqual(t, 5, cols)
	gotest.AssertZero(t, m.Get(0, 0))
	gotest.AssertZero(t, m.Get(9, 4))
}

func TestNewWithValues(t *testing.T) {
	// replaced new assertions helper

	m := New(5, 2, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	rows, cols := m.Size()
	gotest.AssertEqual(t, 5, rows)
	gotest.AssertEqual(t, 2, cols)
	gotest.AssertEqual(t, 1, m.Get(0, 0))
	gotest.AssertEqual(t, 10, m.Get(4, 1))
}

func TestIdentitiy(t *testing.T) {
	// replaced new assertions helper

	id := Identity(5)
	rows, cols := id.Size()
	gotest.AssertEqual(t, 5, rows)
	gotest.AssertEqual(t, 5, cols)
	gotest.AssertEqual(t, 1, id.Get(0, 0))
	gotest.AssertEqual(t, 1, id.Get(1, 1))
	gotest.AssertEqual(t, 1, id.Get(2, 2))
	gotest.AssertEqual(t, 1, id.Get(3, 3))
	gotest.AssertEqual(t, 1, id.Get(4, 4))
	gotest.AssertEqual(t, 0, id.Get(0, 1))
	gotest.AssertEqual(t, 0, id.Get(1, 0))
	gotest.AssertEqual(t, 0, id.Get(4, 0))
	gotest.AssertEqual(t, 0, id.Get(0, 4))
}

func TestNewFromArrays(t *testing.T) {
	// replaced new assertions helper

	m := NewFromArrays([][]float64{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
	})
	gotest.AssertNotNil(t, m)

	rows, cols := m.Size()
	gotest.AssertEqual(t, 2, rows)
	gotest.AssertEqual(t, 4, cols)
}

func TestOnes(t *testing.T) {
	// replaced new assertions helper

	ones := Ones(5, 10)
	rows, cols := ones.Size()
	gotest.AssertEqual(t, 5, rows)
	gotest.AssertEqual(t, 10, cols)

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			gotest.AssertEqual(t, 1, ones.Get(row, col))
		}
	}
}

func TestMatrixEpsilon(t *testing.T) {
	// replaced new assertions helper

	ones := Ones(2, 2)
	ones = ones.WithEpsilon(0.001)
	gotest.AssertEqual(t, 0.001, ones.Epsilon())
}

func TestMatrixArrays(t *testing.T) {
	// replaced new assertions helper

	m := NewFromArrays([][]float64{
		{1, 2, 3},
		{4, 5, 6},
	})

	gotest.AssertNotNil(t, m)

	arrays := m.Arrays()

	gotest.AssertEqual(t, arrays, [][]float64{
		{1, 2, 3},
		{4, 5, 6},
	})
}

func TestMatrixIsSquare(t *testing.T) {
	// replaced new assertions helper

	gotest.AssertFalse(t, NewFromArrays([][]float64{
		{1, 2, 3},
		{4, 5, 6},
	}).IsSquare())

	gotest.AssertFalse(t, NewFromArrays([][]float64{
		{1, 2},
		{3, 4},
		{5, 6},
	}).IsSquare())

	gotest.AssertTrue(t, NewFromArrays([][]float64{
		{1, 2},
		{3, 4},
	}).IsSquare())
}

func TestMatrixIsSymmetric(t *testing.T) {
	// replaced new assertions helper

	gotest.AssertFalse(t, NewFromArrays([][]float64{
		{1, 2, 3},
		{2, 1, 2},
	}).IsSymmetric())

	gotest.AssertFalse(t, NewFromArrays([][]float64{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}).IsSymmetric())

	gotest.AssertTrue(t, NewFromArrays([][]float64{
		{1, 2, 3},
		{2, 1, 2},
		{3, 2, 1},
	}).IsSymmetric())

}

func TestMatrixGet(t *testing.T) {
	// replaced new assertions helper

	m := NewFromArrays([][]float64{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	})

	gotest.AssertEqual(t, 1, m.Get(0, 0))
	gotest.AssertEqual(t, 2, m.Get(0, 1))
	gotest.AssertEqual(t, 3, m.Get(0, 2))
	gotest.AssertEqual(t, 4, m.Get(1, 0))
	gotest.AssertEqual(t, 5, m.Get(1, 1))
	gotest.AssertEqual(t, 6, m.Get(1, 2))
	gotest.AssertEqual(t, 7, m.Get(2, 0))
	gotest.AssertEqual(t, 8, m.Get(2, 1))
	gotest.AssertEqual(t, 9, m.Get(2, 2))
}

func TestMatrixSet(t *testing.T) {
	// replaced new assertions helper

	m := NewFromArrays([][]float64{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	})

	m.Set(1, 1, 99)
	gotest.AssertEqual(t, 99, m.Get(1, 1))
}

func TestMatrixCol(t *testing.T) {
	// replaced new assertions helper

	m := NewFromArrays([][]float64{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	})

	gotest.AssertEqual(t, []float64{1, 4, 7}, m.Col(0))
	gotest.AssertEqual(t, []float64{2, 5, 8}, m.Col(1))
	gotest.AssertEqual(t, []float64{3, 6, 9}, m.Col(2))
}

func TestMatrixRow(t *testing.T) {
	// replaced new assertions helper

	m := NewFromArrays([][]float64{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	})

	gotest.AssertEqual(t, []float64{1, 2, 3}, m.Row(0))
	gotest.AssertEqual(t, []float64{4, 5, 6}, m.Row(1))
	gotest.AssertEqual(t, []float64{7, 8, 9}, m.Row(2))
}

func TestMatrixSwapRows(t *testing.T) {
	// replaced new assertions helper

	m := NewFromArrays([][]float64{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	})

	m.SwapRows(0, 1)

	gotest.AssertEqual(t, []float64{4, 5, 6}, m.Row(0))
	gotest.AssertEqual(t, []float64{1, 2, 3}, m.Row(1))
	gotest.AssertEqual(t, []float64{7, 8, 9}, m.Row(2))
}

func TestMatrixCopy(t *testing.T) {
	// replaced new assertions helper

	m := NewFromArrays([][]float64{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	})

	m2 := m.Copy()
	gotest.AssertFalse(t, m == m2)
	gotest.AssertTrue(t, m.Equals(m2))
}

func TestMatrixDiagonalVector(t *testing.T) {
	// replaced new assertions helper

	m := NewFromArrays([][]float64{
		{1, 4, 7},
		{4, 2, 8},
		{7, 8, 3},
	})

	diag := m.DiagonalVector()
	gotest.AssertEqual(t, []float64{1, 2, 3}, diag)
}

func TestMatrixDiagonalVectorLandscape(t *testing.T) {
	// replaced new assertions helper

	m := NewFromArrays([][]float64{
		{1, 4, 7, 99},
		{4, 2, 8, 99},
	})

	diag := m.DiagonalVector()
	gotest.AssertEqual(t, []float64{1, 2}, diag)
}

func TestMatrixDiagonalVectorPortrait(t *testing.T) {
	// replaced new assertions helper

	m := NewFromArrays([][]float64{
		{1, 4},
		{4, 2},
		{99, 99},
	})

	diag := m.DiagonalVector()
	gotest.AssertEqual(t, []float64{1, 2}, diag)
}

func TestMatrixDiagonal(t *testing.T) {
	// replaced new assertions helper

	m := NewFromArrays([][]float64{
		{1, 4, 7},
		{4, 2, 8},
		{7, 8, 3},
	})

	m2 := NewFromArrays([][]float64{
		{1, 0, 0},
		{0, 2, 0},
		{0, 0, 3},
	})

	gotest.AssertTrue(t, m.Diagonal().Equals(m2))
}

func TestMatrixEquals(t *testing.T) {
	// replaced new assertions helper

	m := NewFromArrays([][]float64{
		{1, 4, 7},
		{4, 2, 8},
		{7, 8, 3},
	})

	gotest.AssertFalse(t, m.Equals(nil))
	var nilMatrix *Matrix
	gotest.AssertTrue(t, nilMatrix.Equals(nil))
	gotest.AssertFalse(t, m.Equals(New(1, 1)))
	gotest.AssertFalse(t, m.Equals(New(3, 3)))
	gotest.AssertTrue(t, m.Equals(New(3, 3, 1, 4, 7, 4, 2, 8, 7, 8, 3)))
}

func TestMatrixL(t *testing.T) {
	// replaced new assertions helper

	m := NewFromArrays([][]float64{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	})

	l := m.L()
	gotest.AssertTrue(t, l.Equals(New(3, 3, 1, 2, 3, 0, 5, 6, 0, 0, 9)))
}

func TestMatrixU(t *testing.T) {
	// replaced new assertions helper

	m := NewFromArrays([][]float64{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	})

	u := m.U()
	gotest.AssertTrue(t, u.Equals(New(3, 3, 0, 0, 0, 4, 0, 0, 7, 8, 0)))
}

func TestMatrixString(t *testing.T) {
	// replaced new assertions helper

	m := NewFromArrays([][]float64{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	})

	gotest.AssertEqual(t, "1 2 3 \n4 5 6 \n7 8 9 \n", m.String())
}

func TestMatrixLU(t *testing.T) {
	// replaced new assertions helper

	m := NewFromArrays([][]float64{
		{1, 3, 5},
		{2, 4, 7},
		{1, 1, 0},
	})

	l, u, p := m.LU()
	gotest.AssertNotNil(t, l)
	gotest.AssertNotNil(t, u)
	gotest.AssertNotNil(t, p)
}

func TestMatrixQR(t *testing.T) {
	// replaced new assertions helper

	m := NewFromArrays([][]float64{
		{12, -51, 4},
		{6, 167, -68},
		{-4, 24, -41},
	})

	q, r := m.QR()
	gotest.AssertNotNil(t, q)
	gotest.AssertNotNil(t, r)
}

func TestMatrixTranspose(t *testing.T) {
	// replaced new assertions helper

	m := NewFromArrays([][]float64{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
		{10, 11, 12},
	})

	m2 := m.Transpose()

	rows, cols := m2.Size()
	gotest.AssertEqual(t, 3, rows)
	gotest.AssertEqual(t, 4, cols)

	gotest.AssertEqual(t, 1, m2.Get(0, 0))
	gotest.AssertEqual(t, 10, m2.Get(0, 3))
	gotest.AssertEqual(t, 3, m2.Get(2, 0))
}
