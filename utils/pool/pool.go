package pool

import "sync"

type Pool struct {
	pos int
	buf []byte

	mtx sync.Mutex
}

const maxpoolsize = 500 * 1024

func (pool *Pool) Get(size int) []byte {
	if size > maxpoolsize {
		return make([]byte, size)
	}

	pool.mtx.Lock()
	defer pool.mtx.Unlock()

	if pool.buf == nil || maxpoolsize-pool.pos < size {
		pool.pos = 0
		pool.buf = make([]byte, maxpoolsize)
	}

	b := make([]byte, size)
	copy(b, pool.buf[pool.pos:pool.pos+size])
	pool.pos += size
	return b
}

func (pool *Pool) Reset() {
	pool.mtx.Lock()
	defer pool.mtx.Unlock()
	pool.pos = 0
}

func NewPool() *Pool {
	return &Pool{
		buf: nil, // delay allocation until first Get call
	}
}
