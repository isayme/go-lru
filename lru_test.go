package lru

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLRU(t *testing.T) {
	require := require.New(t)

	t.Run("normal", func(t *testing.T) {
		l := New(11)
		require.Equal(11, l.size)
		require.Len(l.cache, 0)
		require.Len(l.oldCache, 0)

		_, ok := l.Get("any")
		require.False(ok)
		_, ok = l.Peek("any")
		require.False(ok)
		require.False(l.Has("any"))

		l.Set("any", 123)

		value, ok := l.Get("any")
		require.True(ok)
		require.Equal(123, value.(int))
		value, ok = l.Peek("any")
		require.True(ok)
		require.Equal(123, value.(int))
		require.True(l.Has("any"))

		l.Set("remove", "v")
		require.True(l.Has("remove"))
		require.Len(l.cache, 2)
		l.Remove("remove")
		require.False(l.Has("remove"))
		require.Len(l.cache, 1)

		l.Clear()
		require.Len(l.cache, 0)
		require.Len(l.oldCache, 0)
	})

	t.Run("oldCache + Get", func(t *testing.T) {
		l := New(2)

		l.Set(1, 1)
		require.Len(l.cache, 1)
		require.Len(l.oldCache, 0)
		l.Set(2, 2)
		require.Len(l.cache, 0)
		require.Len(l.oldCache, 2)
		l.Set(3, 3)
		require.Len(l.cache, 1)
		require.Len(l.oldCache, 2)
		_, ok := l.cache[3]
		require.True(ok)

		l.Get(1)
		require.Len(l.cache, 0)
		require.Len(l.oldCache, 2)
		_, ok = l.oldCache[1]
		require.True(ok)
		_, ok = l.oldCache[3]
		require.True(ok)
	})

	t.Run("oldCache + Remove", func(t *testing.T) {
		l := New(2)

		l.Set(1, 1)
		require.Len(l.cache, 1)
		require.Len(l.oldCache, 0)
		l.Set(2, 2)
		require.Len(l.cache, 0)
		require.Len(l.oldCache, 2)
		l.Set(3, 3)
		require.Len(l.cache, 1)
		require.Len(l.oldCache, 2)
		_, ok := l.cache[3]
		require.True(ok)

		l.Remove(1)
		require.Len(l.cache, 1)
		require.Len(l.oldCache, 1)

		l.Remove(3)
		require.Len(l.cache, 0)
		require.Len(l.oldCache, 1)
	})
}
