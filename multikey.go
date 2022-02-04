package multikey

type MultiKeyMap interface {
	// Set
	Set(keys []int, val interface{})
	// Get
	Get(keys []int) interface{}
	// Delete
	// Foreach
}

type mapInstance struct {
	vmap    map[uint64]interface{}
	kmap    map[interface{}]uint64
	counter uint64
}

func New() MultiKeyMap {
	m := &mapInstance{}
	m.vmap = make(map[uint64]interface{})
	m.kmap = make(map[interface{}]uint64)
	m.counter = 1

	return m
}

func (m *mapInstance) Set(keys []int, val interface{}) {
	var counter uint64

	for _, key := range keys {
		if cnt, ok := m.kmap[key]; !ok {
			m.kmap[key] = m.counter
			counter |= m.counter
			m.counter = m.counter << 1
		} else {
			counter |= cnt
		}
	}

	m.vmap[counter] = val
}

func (m *mapInstance) Get(keys []int) interface{} {
	var counter uint64

	for _, key := range keys {
		if cnt, ok := m.kmap[key]; !ok {
			m.kmap[key] = m.counter
			counter |= m.counter
			m.counter = m.counter << 1
		} else {
			counter |= cnt
		}
	}

	return m.vmap[counter]
}
