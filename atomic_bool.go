package atomicbool

import "sync/atomic"

// AtomicBool atomic bool
type AtomicBool interface {
	Get() bool
	Set(bool)
}

type atomicBool struct {
	value int32
}

// Get get
func (b *atomicBool) Get() bool {
	return i2b(atomic.LoadInt32(&b.value))
}

// Set set
func (b *atomicBool) Set(v bool) {
	atomic.StoreInt32(&b.value, b2i(v))
}

// Create create
func Create(v bool) AtomicBool {
	return &atomicBool{
		value: b2i(v),
	}
}

func b2i(v bool) int32 {
	if v {
		return 1
	}
	return 0
}

func i2b(i int32) bool {
	if i == 0 {
		return false
	}
	return true
}
