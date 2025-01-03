package random_test

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/roman-kartash/gocommons/random"
)

var globalInt int

type IntPanicData struct {
	Min   int
	Max   int
	Panic bool
}

func TestInt_Panics(t *testing.T) {
	t.Parallel()

	data := []IntPanicData{
		// should panic
		{Min: 1, Max: 0, Panic: true},
		{Min: -1, Max: 0, Panic: true},
		{Min: 0, Max: -1, Panic: true},
		{Min: 2, Max: 1, Panic: true},

		// should not panic
		{Min: 1, Max: 2},
		{Min: 0, Max: 1},
		{Min: 0, Max: 2},
		{Min: 10, Max: 100},
	}

	for _, d := range data {
		if d.Panic {
			require.Panics(t, func() {
				i := random.Int(d.Min, d.Max)
				runtime.KeepAlive(i)
			})
		} else {
			require.NotPanics(t, func() {
				i := random.Int(d.Min, d.Max)
				runtime.KeepAlive(i)
			})
		}
	}
}

type IntData struct {
	Min int
	Max int
}

func TestInt(t *testing.T) {
	t.Parallel()

	data := []IntData{
		{Min: 0, Max: 0},
		{Min: 0, Max: 1},
		{Min: 0, Max: 2},
		{Min: 0, Max: 3},
		{Min: 5, Max: 150},
		{Min: 10, Max: 25_000},
	}

	for _, d := range data {
		i := random.Int(d.Min, d.Max)
		if d.Min == d.Max {
			require.Equal(t, d.Min, i)
		} else {
			require.GreaterOrEqual(t, i, d.Min)
			require.Less(t, i, d.Max)
		}
	}
}

func BenchmarkInt(b *testing.B) {
	var local int

	for range b.N {
		local = random.Int(0, 100)
	}

	globalInt = local
}
