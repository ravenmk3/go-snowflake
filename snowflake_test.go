package snowflake

import (
	"testing"
)

func TestSnowflake(t *testing.T) {
	g, err := NewGenerator(1)
	if err != nil {
		t.Error(err)
	}
	m := map[int64]int64{}
	for i := 0; i < 1000000; i++ {
		id, err := g.NextId()
		if err != nil {
			t.Error(err)
		}
		_, ok := m[id]
		if ok {
			t.Errorf("Duplicate id: %d", id)
			return
		}
		m[id] = id
	}
}
