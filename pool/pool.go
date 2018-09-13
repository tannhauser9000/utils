/*
* author: tannhauser ruan
* email: tannhauser.sphinx@gmail.com
* summary: golang index pool via golang channel, for resource pool or lock
 */
package pool

import "fmt"
import "time"

// for contant error
type Error string

// regular index pool
type IndexPool struct {
	size  int
	index chan int
}

// index pool with timeout
type TimeoutIndexPool struct {
	size    int
	index   chan int
	timeout time.Duration
}

// constant error for index pool
const ErrTimeout = Error("Index timeout.")
const ErrFreeTimeout = Error("Index free timeout.")
const ErrUnknown = Error("Unknown situation when accessing index pool")

// function for constant error
func (e Error) Error() string {
	return string(e)
}

// initialize a regular index pool
func (p *IndexPool) Init(size int) {
	(*p).index = make(chan int, size)
	for i := 0; i < size; i++ {
		(*p).index <- i
	}
	(*p).size = size
	return
}

// get an index from a regular index pool
func (p *IndexPool) Get() int {
	return <-(*p).index
}

// free an index to the regular index pool
func (p *IndexPool) Free(index int) {
	(*p).index <- index
	return
}

// initialize a index pool with timeout
func (p *TimeoutIndexPool) Init(size int, timeout time.Duration) {
	(*p).index = make(chan int, size)
	for i := 0; i < size; i++ {
		(*p).index <- i
	}
	(*p).size = size
	(*p).timeout = timeout
	return
}

// get an index from a index pool with timeout
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

// free an index back to index pool wiith timeout
func (p *TimeoutIndexPool) Free(index int) error {
	err := error(nil)
	start := time.Now()
	select {
	case (*p).index <- index:
		err = nil
	case <-time.After((*p).timeout):
		fmt.Printf("[Free] duration: %d\n", time.Since(start))
		err = ErrFreeTimeout
	}
	return err
}
