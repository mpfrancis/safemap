package safemap_test

import (
	"testing"

	"github.com/mpfrancis/convert"
	"github.com/mpfrancis/safemap"
	"github.com/stretchr/testify/assert"
)

func Test_Delete(t *testing.T) {
	m := safemap.New[string, *string]()

	m.Set("key", convert.Ptr("value"))
	assert.Equal(t, convert.Ptr("value"), m.GetOrZero("key"))
	assert.Equal(t, 1, m.Size())

	m.Delete("key")
	if _, ok := m.Get("key"); ok {
		t.Errorf("expected key to be deleted")
	}

	assert.Equal(t, 0, m.Size())
}

func Test_GetOrZero(t *testing.T) {
	m := safemap.New[string, *string]()

	var empty *string
	m.Set("key", convert.Ptr("value"))
	assert.Equal(t, convert.Ptr("value"), m.GetOrZero("key"))
	assert.Equal(t, empty, m.GetOrZero("nonexistent"))
	assert.Equal(t, 1, m.Size())

	m.Delete("key")
	assert.Equal(t, empty, m.GetOrZero("key"))
	assert.Equal(t, 0, m.Size())
}

func Test_Get(t *testing.T) {
	m := safemap.New[string, *string]()

	var empty *string
	m.Set("key", convert.Ptr("value"))
	v, ok := m.Get("key")
	assert.Equal(t, convert.Ptr("value"), v)
	assert.True(t, ok)

	v, ok = m.Get("nonexistent")
	assert.Equal(t, empty, v)
	assert.False(t, ok)

	m.Delete("key")
	v, ok = m.Get("key")
	assert.Equal(t, empty, v)
	assert.False(t, ok)
}

func Test_GetAndDelete(t *testing.T) {
	m := safemap.New[string, *string]()

	var empty *string
	m.Set("key", convert.Ptr("value"))
	v, ok := m.GetAndDelete("key")
	assert.Equal(t, convert.Ptr("value"), v)
	assert.True(t, ok)
	assert.Equal(t, 0, m.Size())

	v, ok = m.GetAndDelete("nonexistent")
	assert.Equal(t, empty, v)
	assert.False(t, ok)
	assert.Equal(t, 0, m.Size())

	m.Set("key", convert.Ptr("value"))
	v, ok = m.GetAndDelete("key")
	assert.Equal(t, convert.Ptr("value"), v)
	assert.True(t, ok)
	assert.Equal(t, 0, m.Size())

	v, ok = m.GetAndDelete("nonexistent")
	assert.Equal(t, empty, v)
	assert.False(t, ok)
	assert.Equal(t, 0, m.Size())
}

func Test_GetOrSet(t *testing.T) {
	m := safemap.New[string, *string]()

	v, ok := m.GetOrSet("key", convert.Ptr("value"))
	assert.Equal(t, convert.Ptr("value"), v)
	assert.False(t, ok)
	assert.Equal(t, 1, m.Size())

	v, ok = m.GetOrSet("key", convert.Ptr("new_value"))
	assert.Equal(t, convert.Ptr("value"), v)
	assert.True(t, ok)
	assert.Equal(t, 1, m.Size())

	m.Delete("key")
	assert.Equal(t, 0, m.Size())

	v, ok = m.GetOrSet("key", convert.Ptr("new_value"))
	assert.Equal(t, convert.Ptr("new_value"), v)
	assert.False(t, ok)
	assert.Equal(t, 1, m.Size())
}

func Test_Range(t *testing.T) {
	m := safemap.New[string, *string]()

	m.Set("key1", convert.Ptr("value1"))
	m.Set("key2", convert.Ptr("value2"))

	count := 0
	m.Range(func(key string, value *string) bool {
		count++
		return true
	})
	assert.Equal(t, 2, count)

	count = 0
	m.Range(func(key string, value *string) bool {
		count++
		return false
	})
	assert.Equal(t, 1, count)
}
