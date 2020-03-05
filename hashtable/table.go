package hashtable

const maxBucketSize = 3

type bucket struct {
	length uint
	first  *node
}

type node struct {
	key   string
	value interface{}
	next  *node
}

type Table struct {
	buckets  []bucket
	hashFunc func(s string) uint64
}

func New(size uint64, hashFunc func(s string) uint64) *Table {
	return &Table{
		buckets:  make([]bucket, size),
		hashFunc: hashFunc,
	}
}

func (t *Table) ind(key string) uint64 {
	return t.hashFunc(key) % uint64(len(t.buckets))
}

func (t *Table) Add(key string, value interface{}) {
	ind := t.ind(key)
	b := t.buckets[ind]

	if b.length == maxBucketSize {
		t.grow()
	}
	t.buckets[ind].first = &node{
		key:   key,
		value: value,
		next:  b.first,
	}
	t.buckets[ind].length++
}

func (t *Table) Get(key string) (value interface{}, found bool) {
	b := t.buckets[t.ind(key)]
	for n := b.first; n != nil; n = n.next {
		if n.key == key {
			return n.value, true
		}
	}
	return nil, false
}

func (t *Table) Remove(key string) (value interface{}, found bool) {
	ind := t.ind(key)
	prev := t.buckets[ind].first
	if prev == nil {
		return nil, false
	}
	if prev.key == key {
		t.buckets[ind].first = prev.next
		return prev.value, true
	}
	for n := prev.next; n != nil; prev, n = n, n.next {
		if n.key == key {
			v := n.value
			prev.next = n.next
			return v, true
		}
	}
	return nil, false
}

func (t *Table) grow() {
	bb := t.buckets
	t.buckets = make([]bucket, len(t.buckets)*2+1)
	for _, b := range bb {
		if b.first == nil {
			continue
		}
		for n := b.first; n.next != nil; n = n.next {
			t.Add(n.key, n.value)
		}
	}
}
