/*
* author: tannhauser ruan
* email: tannhauser.sphinx@gmail.com
* summary: testing for index pool
 */
package pool

import "fmt"
import "sync"
import "testing"
import "time"

// constant
const poolSize = 3              // testing pool size
const loopCount = 10000         // testing loop count for pool access
const timeout = 1 * time.Second // timeout for pool w/ timeout
const concurrency = 100         // concurrent go routine for pool access

// global variable
var errAry []error

func (p *IndexPool) goGetIndex(slot int, w *sync.WaitGroup) {
	index := 0
	err := error(nil)
	defer func(wg *sync.WaitGroup) {
		wg.Done()
		return
	}(w)
	for i := 0; i < loopCount; i++ {
		select {
		case index = <-(*p).index:
			err = nil
		case <-time.After(timeout):
			err = ErrTimeout
		}
		if err == nil {
			p.Free(index)
		} else {
			errAry[slot] = err
		}
	}
	return
}

func (p *TimeoutIndexPool) goGetTimeoutIndex(slot int, w *sync.WaitGroup) {
	index := 0
	err := error(nil)
	defer func(wg *sync.WaitGroup) {
		wg.Done()
		return
	}(w)
	for i := 0; i < loopCount; i++ {
		index, err = p.Get()
		if err == nil {
			p.Free(index)
		} else {
			errAry[slot] = err
		}
	}
	return
}

func TestPool(t *testing.T) {
	fmt.Printf("Testing Index Pool w/o Timeout:\n")
	errAry = make([]error, concurrency)
	defer func() {
		errAry = nil
	}()
	wg := sync.WaitGroup{}
	p := &IndexPool{}
	p.Init(poolSize)
	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go p.goGetIndex(i, &wg)
	}
	wg.Wait()
	for i := 0; i < concurrency; i++ {
		if errAry[i] != nil {
			t.Error(errAry[i])
			return
		}
	}
	return
}

func TestTimoutPool(t *testing.T) {
	fmt.Printf("Testing Index Pool w/ Timeout:\n")
	errAry = make([]error, concurrency)
	defer func() {
		errAry = nil
	}()
	wg := sync.WaitGroup{}
	p := &TimeoutIndexPool{}
	p.Init(poolSize, timeout)
	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go p.goGetTimeoutIndex(i, &wg)
	}
	wg.Wait()
	for i := 0; i < concurrency; i++ {
		if errAry[i] != nil {
			t.Error(errAry[i])
			return
		}
	}
	err := error(nil)
	for i := 0; i < poolSize; i++ {
		_, err = p.Get()
		if err != nil {
			t.Error(err)
		}
	}
	_, err = p.Get()
	if err != ErrTimeout {
		t.Error(ErrUnknown)
	}
	return
}
