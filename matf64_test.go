package matrix

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewf64(t *testing.T) {
	t.Helper()
	rows := 13
	cols := 7
	m := Newf64()
	assert.Equal(t, 0, m.r, "should be zero")
	assert.Equal(t, 0, m.c, "should be zero")
	assert.NotNil(t, m.vals, "should not be nil")
	assert.Equal(t, 0, len(m.vals), "should be zero")
	assert.Equal(t, 0, cap(m.vals), "should be zero")

	m = Newf64(rows)
	assert.Equal(t, rows, m.r, "should be equal")
	assert.Equal(t, rows, m.c, "should be equal")
	assert.NotNil(t, m.vals, "should not be nil")
	assert.Equal(t, rows*rows, len(m.vals), "should be equal")
	assert.Equal(t, 2*rows*rows, cap(m.vals), "should have twice the capacity")

	m = Newf64(rows, cols)
	assert.Equal(t, rows, m.r, "should be equal")
	assert.Equal(t, cols, m.c, "should be equal")
	assert.NotNil(t, m.vals, "should not be nil")
	assert.Equal(t, rows*cols, len(m.vals), "should be equal")
	assert.Equal(t, 2*rows*cols, cap(m.vals), "should have twice the capacity")

	assert.Panics(t, func() { Newf64(1, 2, 3) }, "should panic with 3+ args")
	assert.Panics(t, func() { Newf64(1, 2, 3, 4) }, "should panic with 3+ args")
}

func TestMatf64FromData(t *testing.T) {
	t.Helper()
	rows := 50
	cols := 2

	assert.Panics(t, func() { Matf64FromData(1.0) }, "should panic with wrong arg")

	v := make([]float64, rows*cols)
	for i := range v {
		v[i] = float64(i * i)
	}

	m := Matf64FromData(v)
	assert.Equal(t, 1, m.r, "should have one row")
	assert.Equal(t, len(v), len(m.vals), "should have the same # of elements")
	for i := range v {
		assert.Equal(t, v[i], m.vals[i], "should be equal")
	}
	v[0] = 1321.0
	assert.NotEqual(t, v[0], m.vals[0], "changing data should not effect mat")
	m.vals[0] = 1201.0
	assert.NotEqual(t, m.vals[0], v[0], "changing mat should not effect data")

	v[0] = 0.0
	m = Matf64FromData(v, rows*cols)
	assert.Equal(t, rows*cols, m.r, "should be equal")
	assert.Equal(t, 1, m.c, "should have one col")
	assert.Equal(t, len(v), len(m.vals), "should have the same # of elements")
	for i := range v {
		assert.Equal(t, v[i], m.vals[i], "should be equal")
	}
	v[0] = 1321.0
	assert.NotEqual(t, v[0], m.vals[0], "changing data should not effect mat")
	m.vals[0] = 1201.0
	assert.NotEqual(t, m.vals[0], v[0], "changing mat should not effect data")

	v[0] = 0.0
	m = Matf64FromData(v, rows, cols)
	assert.Equal(t, rows, m.r, "should be equal")
	assert.Equal(t, cols, m.c, "should be equal")
	assert.Equal(t, len(v), len(m.vals), "should have the same # of elements")
	for i := range v {
		assert.Equal(t, v[i], m.vals[i], "should be equal")
	}
	v[0] = 1321.0
	assert.NotEqual(t, v[0], m.vals[0], "changing data should not effect mat")
	m.vals[0] = 1201.0
	assert.NotEqual(t, m.vals[0], v[0], "changing mat should not effect data")

	assert.Panics(t, func() { Matf64FromData(v, 12) }, "wrong expected size")
	assert.Panics(t, func() { Matf64FromData(v, 11, 2) }, "wrong expected size")
	assert.Panics(t, func() { Matf64FromData(v, 1, 2, 3) }, "too many args")

	s := make([][]float64, rows)
	for i := range s {
		s[i] = make([]float64, cols)
	}
	for i := range s {
		for j := range s[i] {
			s[i][j] = float64(i + j)
		}
	}
	m = Matf64FromData(s)
	assert.Equal(t, rows*cols, len(m.vals), "should be equal")
	assert.Equal(t, 2*rows*cols, cap(m.vals), "should be equal")
	idx := 0
	for i := range s {
		for j := range s[i] {
			assert.Equal(t, s[i][j], m.vals[idx], "should be equal")
			idx++
		}
	}
	s[0][0] = 1021.0
	assert.NotEqual(t, s[0][0], m.vals[0], "changing data should not effect mat")
	m.vals[0] = 1201.0
	assert.NotEqual(t, m.vals[0], s[0][0], "changing mat should not effect data")

	s[0][0] = 0.0
	m = Matf64FromData(s, 10)
	assert.Equal(t, 10, m.r, "should be equal")
	assert.Equal(t, 10, m.c, "should be equal")
	assert.Equal(t, 100, len(m.vals), "should be equal")
	assert.Equal(t, 200, cap(m.vals), "should be equal")
	idx = 0
	for i := range s {
		for j := range s[i] {
			assert.Equal(t, s[i][j], m.vals[idx], "should be equal")
			idx++
		}
	}
	s[0][0] = 1021.0
	assert.NotEqual(t, s[0][0], m.vals[0], "changing data should not effect mat")
	m.vals[0] = 1201.0
	assert.NotEqual(t, m.vals[0], s[0][0], "changing mat should not effect data")

	s[0][0] = 0.0
	m = Matf64FromData(s, rows, cols)
	assert.Equal(t, rows, m.r, "should be equal")
	assert.Equal(t, cols, m.c, "should be equal")
	assert.Equal(t, rows*cols, len(m.vals), "should be equal")
	assert.Equal(t, 2*rows*cols, cap(m.vals), "should be equal")
	idx = 0
	for i := range s {
		for j := range s[i] {
			assert.Equal(t, s[i][j], m.vals[idx], "should be equal")
			idx++
		}
	}
	s[0][0] = 1021.0
	assert.NotEqual(t, s[0][0], m.vals[0], "changing data should not effect mat")
	m.vals[0] = 1201.0
	assert.NotEqual(t, m.vals[0], s[0][0], "changing mat should not effect data")

	assert.Panics(t, func() { Matf64FromData(s, 15) }, "wrong expected size")
	assert.Panics(t, func() { Matf64FromData(s, 1, 2) }, "wrong expected size")
	assert.Panics(t, func() { Matf64FromData(s, 12, 12, 4) }, "too many args")
}

func TestMatf64FromCSV(t *testing.T) {
	t.Helper()
	rows := 3
	cols := 4

	filename := "non-exitant-file"

	assert.Panics(t, func() { Matf64FromCSV(filename) }, "should panic")

	filename = "test.csv"
	str := "1.0,1.0,2.0,3.0\n5.0,8.0,13.0,21.0\n34.0,55.0,89.0,144.0"
	if _, err := os.Stat(filename); err == nil {
		err = os.Remove(filename)
		if err != nil {
			log.Fatal(err)
		}
	}
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	_, err = f.Write([]byte(str))
	if err != nil {
		log.Fatal(err)
	}
	err = f.Close()
	if err != nil {
		log.Fatal(err)
	}

	m := Matf64FromCSV(filename)
	assert.Equal(t, rows*cols, len(m.vals), "should be equal")
	assert.Equal(t, 1.0, m.vals[0], "should be equal")
	assert.Equal(t, 1.0, m.vals[1], "should be equal")
	for i := 2; i < m.r*m.c; i++ {
		assert.Equal(t, (m.vals[i-1] + m.vals[i-2]), m.vals[i], "should be equal")
	}
	err = os.Remove(filename)
	if err != nil {
		log.Fatal(err)
	}
}

func TestRand(t *testing.T) {
	t.Helper()
	rows := 31
	cols := 42

	m := RandMatf64(rows, cols)
	for i := 0; i < rows*cols; i++ {
		if m.vals[i] < 0.0 || m.vals[i] >= 1.0 {
			t.Errorf("at index %d, expected [0, 1.0), got %f", i, m.vals[i])
		}
	}
	m = RandMatf64(rows, cols, 100.0)
	for i := 0; i < rows*cols; i++ {
		if m.vals[i] < 0.0 || m.vals[i] >= 100.0 {
			t.Errorf("at index %d, expected [0, 100.0), got %f", i, m.vals[i])
		}
	}
	m = RandMatf64(rows, cols, -12.0, 2.0)
	for i := 0; i < rows*cols; i++ {
		if m.vals[i] < -12.0 || m.vals[i] >= 2.0 {
			t.Errorf("at index %d, expected [-12.0, 2.0), got %f", i, m.vals[i])
		}
	}

	assert.Panics(t, func() { RandMatf64(rows, cols, 12.0, 2.0, 13.0) }, "should panic")
	assert.Panics(t, func() { RandMatf64(rows, cols, 12.0, 2.0) }, "should panic")
}

func TestReshape(t *testing.T) {
	t.Helper()
	rows, cols := 10, 12
	s := make([]float64, 120)
	for i := 0; i < len(s); i++ {
		s[i] = float64(i * 3)
	}
	m := Matf64FromData(s).Reshape(rows, cols)
	assert.Equal(t, rows, m.r, "should be equal")
	assert.Equal(t, cols, m.c, "should be equal")
	for i := 0; i < len(s); i++ {
		assert.Equal(t, s[i], m.vals[i], "should be equal")
	}

	assert.Panics(t, func() { m.Reshape(rows, rows) }, "should panic")
}

func TestShape(t *testing.T) {
	t.Helper()
	m := Newf64(11, 10)
	r, c := m.Shape()
	assert.Equal(t, r, m.r, "should be equal")
	assert.Equal(t, c, m.c, "should be equal")
}

func TestVals(t *testing.T) {
	t.Helper()
	rows, cols := 22, 22
	m := Newf64(rows, cols)
	m.SetAll(1.0)
	assert.Equal(t, rows*cols, len(m.vals), "should be equal")
	for i := range m.vals {
		assert.Equal(t, 1.0, m.vals[i], "should be equal")
	}
}

func TestToSlice(t *testing.T) {
	t.Helper()
	rows := 13
	cols := 21
	m := Newf64(rows, cols)
	for i := 0; i < m.r*m.c; i++ {
		m.vals[i] = float64(i)
	}

	s := m.ToSlice()
	assert.Equal(t, m.r, len(s), "should be equal")
	assert.Equal(t, m.c, len(s[0]), "should be equal")
	idx := 0
	for i := range s {
		for j := range s[i] {
			assert.Equal(t, s[i][j], m.vals[idx], "should be equal")
			idx++
		}
	}
	s[0][0] = 1021.0
	assert.NotEqual(t, s[0][0], m.vals[0], "changing data should not effect mat")
	m.vals[0] = 1201.0
	assert.NotEqual(t, m.vals[0], s[0][0], "changing mat should not effect data")
}

func TestToCSV(t *testing.T) {
	t.Helper()
	m := Newf64(23, 17)
	for i := range m.vals {
		m.vals[i] = float64(i)
	}
	filename := "tocsv_test.csv"
	m.ToCSV(filename)
	n := Matf64FromCSV(filename)
	if !n.Equals(m) {
		t.Errorf("m and n are not equal")
	}
	os.Remove(filename)
}

func TestGet(t *testing.T) {
	t.Helper()
	rows := 17
	cols := 13
	m := Newf64(rows, cols)
	for i := range m.vals {
		m.vals[i] = float64(i)
	}
	idx := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			assert.Equal(t, m.vals[idx], m.Get(i, j), "should be equal")
			idx++
		}
	}
}

func TestMap(t *testing.T) {
	t.Helper()
	rows := 132
	cols := 24
	f := func(i *float64) {
		*i = 1.0
		return
	}
	m := Newf64(rows, cols).Map(f)
	for i := 0; i < rows*cols; i++ {
		assert.Equal(t, 1.0, m.vals[i], "should be equal")
	}
}

func BenchmarkMap(b *testing.B) {
	m := Newf64(1721, 311)
	for i := range m.vals {
		m.vals[i] = float64(i)
	}
	f := func(i *float64) {
		*i = 1.0
		return
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = m.Map(f)
	}
}

func TestSetAll(t *testing.T) {
	t.Helper()
	row := 3
	col := 4
	val := 11.0
	m := Newf64(row, col).SetAll(val)
	for i := 0; i < row*col; i++ {
		assert.Equal(t, val, m.vals[i], "should be equal")
	}
}

func TestSet(t *testing.T) {
	t.Helper()
	m := Newf64(5)
	m.Set(2, 3, 10.0)
	assert.Equal(t, 10.0, m.vals[13], "should be equal")
}

func TestSetCol(t *testing.T) {
	t.Helper()
	m := Newf64(3, 4)
	m.SetCol(-1, 3.0)
	n := m.Col(-1)
	for i := range n.vals {
		assert.Equal(t, 3.0, n.vals[i], "should be equal")
	}
	m.SetCol(-1, []float64{0.0, 0.0, 0.0})
	n = m.Col(-1)
	for i := range n.vals {
		assert.Equal(t, 0.0, n.vals[i], "should be equal")
	}
	m.SetCol(1, 3.0)
	n = m.Col(1)
	for i := range n.vals {
		assert.Equal(t, 3.0, n.vals[i], "should be equal")
	}
	m.SetCol(1, []float64{0.0, 0.0, 0.0})
	n = m.Col(1)
	for i := range n.vals {
		assert.Equal(t, 0.0, n.vals[i], "should be equal")
	}

	assert.Panics(t, func() { m.SetCol(-5, 2.0) }, "should panic")
	assert.Panics(t, func() { m.SetCol(5, 2.0) }, "should panic")
	assert.Panics(t, func() { m.SetCol(-1, []float64{0.0}) }, "should panic")
	assert.Panics(t, func() { m.SetCol(1, []float64{0.0}) }, "should panic")
	assert.Panics(t, func() { m.SetCol(-1, 1) }, "should panic")
	assert.Panics(t, func() { m.SetCol(1, 1) }, "should panic")
}

func TestSetRow(t *testing.T) {
	t.Helper()
	m := Newf64(3, 4)
	m.SetRow(-1, 3.0)
	n := m.Row(-1)
	for i := range n.vals {
		assert.Equal(t, 3.0, n.vals[i], "should be equal")
	}
	m.SetRow(-1, []float64{0.0, 0.0, 0.0, 0.0})
	n = m.Row(-1)
	for i := range n.vals {
		assert.Equal(t, 0.0, n.vals[i], "should be equal")
	}
	m.SetRow(1, 3.0)
	n = m.Row(1)
	for i := range n.vals {
		assert.Equal(t, 3.0, n.vals[i], "should be equal")
	}
	m.SetRow(1, []float64{0.0, 0.0, 0.0, 0.0})
	n = m.Row(1)
	for i := range n.vals {
		assert.Equal(t, 0.0, n.vals[i], "should be equal")
	}

	assert.Panics(t, func() { m.SetRow(-5, 2.0) }, "should panic")
	assert.Panics(t, func() { m.SetRow(5, 2.0) }, "should panic")
	assert.Panics(t, func() { m.SetRow(-1, []float64{0.0}) }, "should panic")
	assert.Panics(t, func() { m.SetRow(1, []float64{0.0}) }, "should panic")
	assert.Panics(t, func() { m.SetRow(-1, 1) }, "should panic")
	assert.Panics(t, func() { m.SetRow(1, 1) }, "should panic")
}

func TestCol(t *testing.T) {
	t.Helper()
	row := 3
	col := 4
	m := Newf64(row, col)
	for i := range m.vals {
		m.vals[i] = float64(i)
	}
	for i := 0; i < col; i++ {
		got := m.Col(i)
		for j := 0; j < row; j++ {
			assert.Equal(t, m.vals[j*m.c+i], got.vals[j], "should be equal")
		}
	}
	for i := col; i < 0; i-- {
		got := m.Col(-i)
		for j := 0; j < row; j++ {
			assert.Equal(t, m.vals[j*m.c+(row-i)], got.vals[j], "should be equal")
		}
	}
}

func BenchmarkCol(b *testing.B) {
	m := Newf64(1721, 311)
	for i := range m.vals {
		m.vals[i] = float64(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = m.Col(211)
	}
}

func TestRow(t *testing.T) {
	t.Helper()
	row := 3
	col := 4
	m := Newf64(row, col)
	for i := range m.vals {
		m.vals[i] = float64(i)
	}
	idx := 0
	for i := 0; i < row; i++ {
		got := m.Row(i)
		for j := 0; j < col; j++ {
			assert.Equal(t, m.vals[idx], got.vals[j], "should be equal")
			idx++
		}
	}
}

func BenchmarkRow(b *testing.B) {
	m := Newf64(1721, 311)
	for i := range m.vals {
		m.vals[i] = float64(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = m.Row(211)
	}
}

func TestMin(t *testing.T) {
	t.Helper()
	m := Newf64(3, 4)
	m.Set(2, 1, -100.0)
	_, minVal := m.Min()
	assert.Equal(t, -100.0, minVal, "should be equal")
	idx, minVal := m.Min(0, 2)
	assert.Equal(t, -100.0, minVal, "should be equal")
	assert.Equal(t, 1, idx, "should be equal")
	idx, minVal = m.Min(1, 1)
	assert.Equal(t, -100.0, minVal, "should be equal")
	assert.Equal(t, 2, idx, "should be equal")
}

func TestMax(t *testing.T) {
	t.Helper()
	m := Newf64(3, 4)
	m.Set(2, 1, 100.0)
	_, maxVal := m.Max()
	assert.Equal(t, 100.0, maxVal, "should be equal")
	idx, maxVal := m.Max(0, 2)
	assert.Equal(t, 100.0, maxVal, "should be equal")
	assert.Equal(t, 1, idx, "should be equal")
	idx, maxVal = m.Max(1, 1)
	assert.Equal(t, 100.0, maxVal, "should be equal")
	assert.Equal(t, 2, idx, "should be equal")
}

func TestEquals(t *testing.T) {
	t.Helper()
	m := Newf64(13, 12)
	if !m.Equals(m) {
		t.Errorf("m is not equal itself")
	}
}

func TestCopy(t *testing.T) {
	t.Helper()
	rows, cols := 17, 13
	m := Newf64(rows, cols)
	for i := range m.vals {
		m.vals[i] = float64(i)
	}
	n := m.Copy()
	for i := 0; i < rows*cols; i++ {
		assert.Equal(t, m.vals[i], n.vals[i], "should be equal")
	}
}

func TestT(t *testing.T) {
	t.Helper()
	m := Newf64(12, 3)
	for i := range m.vals {
		m.vals[i] = float64(i)
	}
	n := m.T()
	p := m.ToSlice()
	q := n.ToSlice()
	for i := 0; i < m.r; i++ {
		for j := 0; j < m.c; j++ {
			assert.Equal(t, p[i][j], q[j][i], "should be equal")
		}
	}
}

func BenchmarkT(b *testing.B) {
	m := Newf64(1000, 251)
	for i := range m.vals {
		m.vals[i] = float64(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = m.T()
	}
}

func TestAll(t *testing.T) {
	t.Helper()
	m := Newf64(100, 21)
	for i := range m.vals {
		m.vals[i] = float64(i + 1)
	}
	assert.True(t, m.All(Positivef64), "All should be > 0")
	isOne := func(i *float64) bool {
		return *i == 1.0
	}
	m.SetAll(1.0)
	assert.True(t, m.All(isOne), "All should be 1.0s")
}

func TestAny(t *testing.T) {
	t.Helper()
	m := Newf64(100, 21)
	for i := range m.vals {
		m.vals[i] = float64(i)
	}
	assert.False(t, m.Any(Negativef64), "should have no negatives")
	assert.True(t, m.Any(Positivef64), "should have positives")
}

func TestMul(t *testing.T) {
	t.Helper()
	rows, cols := 13, 90
	m := Newf64(rows, cols)
	for i := range m.vals {
		m.vals[i] = float64(i)
	}
	n := m.Copy()
	m.Mul(m)
	for i := 0; i < rows*cols; i++ {
		assert.Equal(t, n.vals[i]*n.vals[i], m.vals[i], "should be equal")
	}
}

func BenchmarkMul(b *testing.B) {
	n := Newf64(1000, 1000)
	for i := range n.vals {
		n.vals[i] = float64(i)
	}
	m := Newf64(1000, 1000)
	for i := range m.vals {
		m.vals[i] = float64(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = m.Mul(n)
	}
}

func TestAdd(t *testing.T) {
	t.Helper()
	rows, cols := 13, 90
	m := Newf64(rows, cols)
	for i := range m.vals {
		m.vals[i] = float64(i)
	}
	n := m.Copy()
	m.Add(m)
	for i := 0; i < rows*cols; i++ {
		assert.Equal(t, n.vals[i]+n.vals[i], m.vals[i], "should be equal")
	}
}

func TestSub(t *testing.T) {
	t.Helper()
	rows, cols := 13, 90
	m := Newf64(rows, cols)
	for i := range m.vals {
		m.vals[i] = float64(i)
	}
	m.Sub(m)
	for i := 0; i < rows*cols; i++ {
		assert.Equal(t, 0.0, m.vals[i], "should be equal")
	}
}

func TestDiv(t *testing.T) {
	t.Helper()
	rows, cols := 13, 90
	m := Newf64(rows, cols)
	for i := range m.vals {
		m.vals[i] = float64(i)
	}
	m.vals[0] = 1.0
	m.Div(m)
	for i := 0; i < rows*cols; i++ {
		assert.Equal(t, 1.0, m.vals[i], "should be equal")
	}
}

func TestSum(t *testing.T) {
	t.Helper()
	row := 12
	col := 17
	m := Newf64(row, col).SetAll(1.0)
	for i := 0; i < row; i++ {
		assert.Equal(t, float64(col), m.Sum(0, i), "should be equal")
	}
	for i := 0; i < col; i++ {
		assert.Equal(t, float64(row), m.Sum(1, i), "should be equal")
	}
}

func TestAvg(t *testing.T) {
	t.Helper()
	row := 12
	col := 17
	m := Newf64(row, col).SetAll(1.0)
	for i := 0; i < row; i++ {
		assert.Equal(t, 1.0, m.Avg(0, i), "should be equal")
	}
	for i := 0; i < col; i++ {
		assert.Equal(t, 1.0, m.Avg(1, i), "should be equal")
	}
}

func TestPrd(t *testing.T) {
	t.Helper()
	row := 12
	col := 17
	m := Newf64(row, col).SetAll(1.0)
	for i := 0; i < row; i++ {
		assert.Equal(t, 1.0, m.Prd(0, i), "should be equal")
	}
	for i := 0; i < col; i++ {
		assert.Equal(t, 1.0, m.Prd(1, i), "should be equal")
	}
}

func TestStd(t *testing.T) {
	t.Helper()
	row := 12
	col := 17
	m := Newf64(row, col).SetAll(1.0)
	for i := 0; i < row; i++ {
		assert.Equal(t, 0.0, m.Std(0, i), "should be equal")
	}
	for i := 0; i < col; i++ {
		assert.Equal(t, 0.0, m.Std(1, i), "should be equal")
	}
}

func TestDot(t *testing.T) {
	t.Helper()
	var (
		row = 10
		col = 4
	)
	m := Newf64(row, col)
	for i := range m.vals {
		m.vals[i] = float64(i)
	}
	n := Newf64(col, row)
	for i := range n.vals {
		n.vals[i] = float64(i)
	}
	o := m.Dot(n)
	assert.Equal(t, row, o.r, "should be equal")
	assert.Equal(t, row, o.c, "should be equal")
	p := Newf64(row, row)
	q := o.Dot(p)
	for i := 0; i < row*row; i++ {
		assert.Equal(t, 0.0, q.vals[i], "should be zero")
	}
}

func BenchmarkDot(b *testing.B) {
	row, col := 150, 130
	m := Newf64(row, col)
	for i := range m.vals {
		m.vals[i] = float64(i)
	}
	n := Newf64(col, row)
	for i := range n.vals {
		n.vals[i] = float64(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = m.Dot(n)
	}
}

func TestAppendCol(t *testing.T) {
	t.Helper()
	var (
		row = 10
		col = 4
	)
	m := Newf64(row, col)
	for i := range m.vals {
		m.vals[i] = float64(i)
	}
	v := make([]float64, row)
	m.AppendCol(v)
	assert.Equal(t, col+1, m.c, "should have one more column")
	m.AppendCol(v)
	assert.Equal(t, col+2, m.c, "should have two more columns")
	m.AppendCol(v)
	assert.Equal(t, col+3, m.c, "should have three more columns")
}

func TestAppendRow(t *testing.T) {
	t.Helper()
	var (
		row = 3
		col = 4
	)
	m := Newf64(row, col)
	for i := range m.vals {
		m.vals[i] = float64(i)
	}
	v := make([]float64, col)
	for i := range v {
		v[i] = float64(i * i * i)
	}
	m.AppendRow(v)
	assert.Equal(t, row+1, m.r, "should have one more row")
	m.AppendRow(v)
	assert.Equal(t, row+2, m.r, "should have two more rows")
	m.AppendRow(v)
	assert.Equal(t, row+3, m.r, "should have three more rows")
}

func TestConcat(t *testing.T) {
	t.Helper()
	var (
		row = 10
		col = 4
	)
	m := Newf64(row, col)
	for i := range m.vals {
		m.vals[i] = float64(i)
	}
	n := Newf64(row, row)
	for i := range n.vals {
		n.vals[i] = float64(i)
	}
	m.Concat(n)
	if m.c != row+col {
		t.Errorf("Expected number of cols to be %d, but got %d", row+col, m.c)
	}
	idx1 := 0
	idx2 := 0
	for i := 0; i < row; i++ {
		for j := 0; j < col+row; j++ {
			if j < col {
				assert.Equal(t, float64(idx1), m.vals[i*m.c+j], "should be equal")
				idx1++
				continue
			}
			assert.Equal(t, float64(idx2), m.vals[i*m.c+j], "should be equal")
			idx2++
		}
	}
}