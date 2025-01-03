package random_test

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/roman-kartash/gocommons/random"
)

const letterBytes = random.ASCIILetters + "-"

type TestStringData struct {
	Len int
}

func GetTestStringData(cnt int, maxlength int) []TestStringData {
	data := make([]TestStringData, 0, cnt)

	for range cnt {
		data = append(data, TestStringData{Len: random.Int(1, maxlength)})
	}

	return data
}

func TestStringFrom(t *testing.T) {
	t.Parallel()

	data := []TestStringData{
		{Len: 0},
		{Len: 5},
		{Len: 10},
		{Len: 100},
		{Len: 1000},
	}

	for _, d := range data {
		res := random.StringFromBytes(d.Len, letterBytes)
		require.Len(t, res, d.Len)
	}
}

func TestStringFromPanic(t *testing.T) {
	t.Parallel()

	data := []TestStringData{
		{Len: -1},
		{Len: -5},
		{Len: -10},
		{Len: -100},
		{Len: -1000},
	}

	for _, d := range data {
		require.Panics(t, func() {
			res := random.StringFromBytes(d.Len, letterBytes)
			runtime.KeepAlive(res)
		})
	}
}

func TestStringFromUnique(t *testing.T) {
	t.Parallel()

	if testing.Short() {
		t.Log("Skip because short mode")
		return
	}

	data := GetTestStringData(100_000, 1_000)
	genStrs := make(map[string]int, len(data))

	for _, d := range data {
		res := random.StringFromBytes(d.Len, letterBytes)
		require.Len(t, res, d.Len)

		genStrs[res]++
	}

	notUnique := map[string]int{}

	for res, cnt := range genStrs {
		if cnt > 1 {
			notUnique[res] = cnt
		}
	}

	require.Less(t, float64(len(notUnique))/float64(len(genStrs)), .001)
}

func TestStringFromUniqueFixedSize(t *testing.T) {
	t.Parallel()

	if testing.Short() {
		t.Log("Skip because short mode")
		return
	}

	strCnt := 1_000_000
	strLen := 100
	genStr := make(map[string]int)

	for range strCnt {
		res := random.StringFromBytes(strLen, letterBytes)
		genStr[res]++
	}

	notUnique := map[string]int{}

	for res, cnt := range genStr {
		if cnt > 1 {
			notUnique[res] = cnt
		}
	}

	require.Less(t, float64(len(notUnique))/float64(len(genStr)), .001)
}

var globalString string

func BenchmarkStringFromFixedSize16(b *testing.B) {
	var local string

	size := 16

	for range b.N {
		local = random.StringFromBytes(size, letterBytes)
	}

	globalString = local
}

func BenchmarkStringFromFixedSize1000(b *testing.B) {
	var local string

	size := 1_000

	for range b.N {
		local = random.StringFromBytes(size, letterBytes)
	}

	globalString = local
}
