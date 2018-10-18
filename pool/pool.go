package pool

/*
* author: tannhauser ruan
* email: tannhauser.sphinx@gmail.com
* summary: golang index pool via golang channel, for resource pool or lock
 */

import "time"

// Error for contant error
type Error string

// IndexPool regular index pool
type IndexPool struct {
	size  int
	index chan int
}

// TimeoutIndexPool index pool with timeout
type TimeoutIndexPool struct {
	size    int
	index   chan int
	timeout time.Duration
}

// constant error for index pool

// ErrTimeout for get index timeout
const ErrTimeout = Error("Index timeout.")

// ErrFreeTimeout for free index timeout
const ErrFreeTimeout = Error("Index free timeout.")

// ErrUnknown for unknown error
const ErrUnknown = Error("Unknown situation when accessing index pool")

// Error for constant error
func (e Error) Error() string {
	return string(e)
}

// Init initialize a regular index pool
func (p *IndexPool) Init(size int) {
	(*p).index = make(chan int, size)
	for i := 0; i < size; i++ {
		(*p).index <- i
	}
	(*p).size = size
	return
}

// Get get an index from a regular index pool
func (p *IndexPool) Get() int {
	return <-(*p).index
}

// Free free an index to the regular index pool
func (p *IndexPool) Free(index int) {
	(*p).index <- index
	return
}

// Init initialize a index pool with timeout
func (p *TimeoutIndexPool) Init(size int, timeout time.Duration) {
	(*p).index = make(chan int, size)
	for i := 0; i < size; i++ {
		(*p).index <- i
	}
	(*p).size = size
	(*p).timeout = timeout
	return
}

// Get get an index from a index pool with timeout
func (p *TimeoutIndexPool) Get() (int, error) {
	index := 0
	err := error(nil)
	select {
	case index = <-(*p).index:
		err = nil
	case <-time.After((*p).timeout):
		err = ErrTimeout
	}
	return index, err
}

// Free free an index back to index pool wiith timeout
func (p *TimeoutIndexPool) Free(index int) error {
	err := error(nil)
	select {
	case (*p).index <- index:
		err = nil
	case <-time.After((*p).timeout):
		err = ErrFreeTimeout
	}
	return err
}
