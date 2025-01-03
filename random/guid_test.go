package random_test

import (
	"runtime"
	"testing"

	"github.com/roman-kartash/gocommons/random"
	"github.com/stretchr/testify/require"
)

func TestGuidNotPanic(t *testing.T) {
	t.Parallel()

	require.NotPanics(t, func() {
		g := random.Guid()
		runtime.KeepAlive(g)
	})
}

func TestGuidUniq100000(t *testing.T) {
	t.Parallel()

	if testing.Short() {
		t.Log("Skip because short mode")
		return
	}

	m := map[string]struct{}{}

	for range 100_000 {
		g := random.Guid()
		_, ok := m[g]
		require.False(t, ok, g)

		m[g] = struct{}{}
	}
}

func BenchmarkGuid(b *testing.B) {
	var local string

	for range b.N {
		local = random.Guid()
	}

	globalString = local
}
