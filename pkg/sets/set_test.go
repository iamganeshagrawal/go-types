package sets

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewSet(t *testing.T) {
	t.Run("Integer_Set", func(t *testing.T) {
		t.Parallel()
		set := NewSet(1, 2, 3)
		require.Equal(t, 3, set.Size())
		require.False(t, set.IsEmpty())
	})
	t.Run("Integer_Set_With_Duplicate_Values", func(t *testing.T) {
		t.Parallel()
		set := NewSet(1, 2, 3, 2, 3, 4)
		require.Equal(t, 4, set.Size())
		require.False(t, set.IsEmpty())
	})
}

func TestNewEmptySet(t *testing.T) {
	set := NewEmptySet[int]()
	require.True(t, set.IsEmpty())
	require.Equal(t, 0, set.Size())
}

func TestSet(t *testing.T) {
	t.Run("Method_Clear", func(t *testing.T) {
		t.Parallel()
		set := NewSet(1, 2, 3)
		require.False(t, set.IsEmpty())
		set.Clear()
		require.True(t, set.IsEmpty())
	})
	t.Run("Method_Size", func(t *testing.T) {
		t.Parallel()
		require.Equal(t, 1, NewSet(1).Size())
		require.Equal(t, 2, NewSet(1, 2).Size())
		require.Equal(t, 2, NewSet(1, 2, 2).Size())
	})
	t.Run("Method_AddIfNotExist", func(t *testing.T) {
		t.Parallel()
		set := NewEmptySet[int]()
		require.True(t, set.AddIfNotExist(1))
		require.False(t, set.AddIfNotExist(1))
		require.True(t, set.AddIfNotExist(2))
		require.Equal(t, 2, set.Size())
	})
	t.Run("Method_Add", func(t *testing.T) {
		t.Parallel()
		set := NewEmptySet[int]()
		set.Add(1, 2, 3, 5, 6, 6)
		require.Equal(t, 5, set.Size())
		set.Add()
		require.Equal(t, 5, set.Size())
	})
	t.Run("Method_Has", func(t *testing.T) {
		t.Parallel()
		set := NewEmptySet[int]()
		require.False(t, set.Has(1))
		set.AddIfNotExist(1)
		require.True(t, set.Has(1))
	})
	t.Run("Method_Contains", func(t *testing.T) {
		t.Parallel()
		set := NewSet(1, 2, 3)
		require.True(t, set.Contains(1, 2, 3))
		require.False(t, set.Contains(1, 2, 4))
	})
	t.Run("Method_Remove", func(t *testing.T) {
		t.Parallel()
		set := NewSet(1, 2, 3)
		require.True(t, set.Has(1))
		set.Remove(1)
		require.False(t, set.Has(1))
	})
	t.Run("Method_Clone", func(t *testing.T) {
		t.Parallel()
		set := NewSet(1, 2, 3)
		setCopy := set.Clone()
		require.Equal(t, set, setCopy)
	})
	t.Run("Method_Union", func(t *testing.T) {
		t.Parallel()
		setA := NewSet(1, 2, 3)
		setB := NewSet(3, 4, 5)
		resultAUB := setA.Union(setB)
		resultBUA := setB.Union(setA)
		require.Equal(t, resultAUB, resultBUA)
		require.Equal(t, 5, resultBUA.Size())
	})
	t.Run("Method_Union_NilPtr", func(t *testing.T) {
		t.Parallel()
		setA := NewSet(1, 2, 3)
		require.Panics(t, func() { setA.Union(nil) })
	})
	t.Run("Method_Intersect", func(t *testing.T) {
		t.Parallel()
		setA := NewSet(1, 2, 3)
		setB := NewSet(3, 4, 5)
		resultAIB := setA.Intersect(setB)
		resultBIA := setB.Intersect(setA)
		require.Equal(t, resultAIB, resultBIA)
		require.Equal(t, 1, resultBIA.Size())
	})
	t.Run("Method_Intersect_NilPtr", func(t *testing.T) {
		t.Parallel()
		setA := NewSet(1, 2, 3)
		require.Panics(t, func() { setA.Intersect(nil) })
	})
	t.Run("Method_Intersect_Optimization", func(t *testing.T) {
		t.Parallel()
		setA := NewSet(1, 2, 3, 4)
		setB := NewSet(3, 4, 5)
		result := setA.Intersect(setB)
		require.Equal(t, 2, result.Size())
	})
	t.Run("Method_Difference", func(t *testing.T) {
		t.Parallel()
		setA := NewSet(1, 2, 3)
		setB := NewSet(3, 4, 5)
		expected := NewSet(1, 2)
		actual := setA.Difference(setB)
		require.Equal(t, expected, actual)
		require.NotEqual(t, setA.Difference(setB), setB.Difference(setA))
		// edge: difference with self reqturn empty set
		require.Equal(t, NewEmptySet[int](), setA.Difference(setA))
	})
	t.Run("Method_Difference_NilPtr", func(t *testing.T) {
		t.Parallel()
		setA := NewSet(1, 2, 3)
		require.Panics(t, func() { setA.Difference(nil) })
	})
	t.Run("Method_Values", func(t *testing.T) {
		t.Parallel()
		set := NewSet(1, 2, 3, 3)
		actual := set.Values()
		require.True(t, len(actual) == 3)
	})
}
