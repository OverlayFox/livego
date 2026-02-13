package pool

type Pool struct {
	pos int
	buf []byte
}

const maxpoolsize = 500 * 1024

func (pool *Pool) Get(size int) []byte {
	if size > maxpoolsize {
		return make([]byte, size)
	}

	if pool.buf == nil || maxpoolsize-pool.pos < size {
		pool.pos = 0
		pool.buf = make([]byte, maxpoolsize)
	}

	b := pool.buf[pool.pos : pool.pos+size]
	pool.pos += size
	return b
}

func NewPool() *Pool {
	return &Pool{
		buf: nil, // delay allocation until first Get call
	}
}
