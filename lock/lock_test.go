/*
* author: tannhauser ruan
* email: tannhauser.sphinx@gmail.com
* summary: testing for lock
 */
package lock

import "sync"
import "testing"
import "time"

// constants
const quickTimeout = 10 * time.Millisecond
const timeout = 20 * time.Millisecond
const longTimeout = 30 * time.Millisecond

// global variables

func backgroundLockAcquire(l *Lock, e *error, w *sync.WaitGroup, start time.Time, t time.Duration) {
	defer func(wg *sync.WaitGroup) {
		(*wg).Done()
	}(w)
	l.Lock()
	duration := time.Since(start)
	if duration < t {
		(*e) = ErrUnknown
	}
	l.Unlock()
	return
}

// mode:
//    0: lock released after lock timeout
//    1: lock released before lock timeout
func backgroundTimeoutLockAcquire(l *TimeoutLock, e *error, w *sync.WaitGroup, start time.Time, t time.Duration, mode int) {
	defer func(wg *sync.WaitGroup) {
		(*wg).Done()
	}(w)
	(*e) = l.Lock()
	duration := time.Since(start)
	if (*e) == nil {
		l.Unlock()
	}
	if (mode == 0 && (*e) != ErrTimeout) || duration < t {
		(*e) = ErrUnknown
	}
	return
}

func backgroundWriteLockAcquire(l *RWLock, e *error, w *sync.WaitGroup, start time.Time, t time.Duration) {
	defer func(wg *sync.WaitGroup) {
		(*wg).Done()
	}(w)
	l.Lock()
	duration := time.Since(start)
	if duration < t {
		(*e) = ErrUnknown
	}
	l.Unlock()
	return
}

// mode:
//    0: lock released after lock timeout
//    1: lock released before lock timeout
func backgroundTimeoutWriteLockAcquire(l *TimeoutRWLock, e *error, w *sync.WaitGroup, start time.Time, t time.Duration, mode int) {
	defer func(wg *sync.WaitGroup) {
		(*wg).Done()
	}(w)
	(*e) = l.Lock()
	duration := time.Since(start)
	err := error(nil)
	if (*e) == nil {
		err = l.Unlock()
	}
	if (mode == 0 && (*e) != ErrTimeout) || duration < t {
		(*e) = ErrUnknown
	}
	if err != nil {
		(*e) = err
	}
	return
}

// mode:
//    0: locked by read lock
//    1: locked by write
func backgroundReadLockAcquire(l *RWLock, e *error, w *sync.WaitGroup, start time.Time, t time.Duration, mode int) {
	defer func(wg *sync.WaitGroup) {
		(*wg).Done()
	}(w)
	l.RLock()
	duration := time.Since(start)
	if mode == 0 && duration > t {
		(*e) = ErrUnknown
	}
	if mode == 1 && duration < t {
		(*e) = ErrUnknown
	}
	l.RUnlock()
	return
}

// mode:
//    0: locked by read lock
//    1: locked by write
// timeoutMode:
//    0: lock released after lock timeout
//    1: lock released before lock timeout
func backgroundTimeoutReadLockAcquire(l *TimeoutRWLock, e *error, w *sync.WaitGroup, start time.Time, t time.Duration, mode int, timeoutMode int) {
	defer func(wg *sync.WaitGroup) {
		(*wg).Done()
	}(w)
	(*e) = l.RLock()
	duration := time.Since(start)
	err := error(nil)
	if (*e) == nil {
		err = l.RUnlock()
	}
	if mode == 0 && (duration > t || (*e) != nil) {
		(*e) = ErrUnknown
	}
	if mode == 1 && ((timeoutMode == 0 && (*e) != ErrTimeout) || duration < t) {
		(*e) = ErrUnknown
	}
	if err != nil {
		(*e) = err
	}
	return
}

func TestLock(t *testing.T) {
	l := GetLock()
	wg := sync.WaitGroup{}
	err := error(nil)
	l.Lock()
	wg.Add(1)
	start := time.Now()
	go backgroundLockAcquire(l, &err, &wg, start, quickTimeout)
	time.Sleep(quickTimeout)
	e := l.Unlock()
	wg.Wait()
	if err != nil {
		t.Error(err)
	}
	if e != nil {
		t.Error(e)
	}
	e = l.Unlock()
	if e != ErrDuplicatedUnlock {
		t.Error(e)
	}
}

func TestRWLock(t *testing.T) {
	l := GetRWLock()
	wg := sync.WaitGroup{}
	err1 := error(nil)
	err2 := error(nil)
	l.RLock()
	wg.Add(2)
	start := time.Now()
	go backgroundWriteLockAcquire(l, &err1, &wg, start, quickTimeout)
	go backgroundReadLockAcquire(l, &err2, &wg, start, quickTimeout, 0)
	time.Sleep(quickTimeout)
	l.RUnlock()
	wg.Wait()
	if err1 != nil {
		t.Error(err1)
	}
	if err2 != nil {
		t.Error(err2)
	}
	err1, err2 = nil, nil
	l.Lock()
	wg.Add(2)
	start = time.Now()
	go backgroundWriteLockAcquire(l, &err1, &wg, start, quickTimeout)
	go backgroundReadLockAcquire(l, &err2, &wg, start, quickTimeout, 1)
	time.Sleep(quickTimeout)
	l.Unlock()
	wg.Wait()
	if err1 != nil {
		t.Error(err1)
	}
	if err2 != nil {
		t.Error(err2)
	}
}

func TestTimeoutLock(t *testing.T) {
	l := GetTimeoutLock(timeout)
	wg := sync.WaitGroup{}
	err := error(nil)

	// test lock released before lock timeout
	l.Lock()
	wg.Add(1)
	start := time.Now()
	go backgroundTimeoutLockAcquire(l, &err, &wg, start, quickTimeout, 1)
	time.Sleep(quickTimeout)
	e := l.Unlock()
	wg.Wait()
	if err != nil {
		t.Error(err)
	}
	if e != nil {
		t.Error(e)
	}

	// test lock released after lock timeout
	l.Lock()
	wg.Add(1)
	start = time.Now()
	go backgroundTimeoutLockAcquire(l, &err, &wg, start, timeout, 0)
	time.Sleep(longTimeout)
	e = l.Unlock()
	wg.Wait()
	if err != nil && err != ErrTimeout {
		t.Error(err)
	}
	if e != nil {
		t.Error(e)
	}
	e = l.Unlock()
	if e != ErrDuplicatedUnlock {
		t.Error(e)
	}
}

func TestTimeoutRWLock(t *testing.T) {
	l := GetTimeoutRWLock(timeout)
	wg := sync.WaitGroup{}
	err1 := error(nil)
	err2 := error(nil)

	// test lock released before lock timeout, locked by read lock
	l.RLock()
	wg.Add(2)
	start := time.Now()
	go backgroundTimeoutWriteLockAcquire(l, &err1, &wg, start, quickTimeout, 1)
	go backgroundTimeoutReadLockAcquire(l, &err2, &wg, start, quickTimeout, 0, 1)
	time.Sleep(quickTimeout)
	l.RUnlock()
	wg.Wait()
	if err1 != nil {
		t.Error(err1)
	}
	if err2 != nil {
		t.Error(err2)
	}
	err1, err2 = nil, nil

	// test lock released after lock timeout, locked by read lock
	l.RLock()
	wg.Add(2)
	start = time.Now()
	go backgroundTimeoutWriteLockAcquire(l, &err1, &wg, start, timeout, 0)
	go backgroundTimeoutReadLockAcquire(l, &err2, &wg, start, timeout, 0, 0)
	time.Sleep(longTimeout)
	l.RUnlock()
	wg.Wait()
	if err1 != nil && err1 != ErrTimeout {
		t.Error(err1)
	}
	if err2 != nil {
		t.Error(err2)
	}
	err1, err2 = nil, nil

	// test lock released before lock timeout, locked by write lock
	l.Lock()
	wg.Add(2)
	start = time.Now()
	go backgroundTimeoutWriteLockAcquire(l, &err1, &wg, start, quickTimeout, 1)
	go backgroundTimeoutReadLockAcquire(l, &err2, &wg, start, quickTimeout, 1, 1)
	time.Sleep(quickTimeout)
	l.Unlock()
	wg.Wait()
	if err1 != nil {
		t.Error(err1)
	}
	if err2 != nil {
		t.Error(err2)
	}
	err1, err2 = nil, nil

	// test lock released after lock timeout, locked by write lock
	l.Lock()
	wg.Add(2)
	start = time.Now()
	go backgroundTimeoutWriteLockAcquire(l, &err1, &wg, start, timeout, 0)
	go backgroundTimeoutReadLockAcquire(l, &err2, &wg, start, timeout, 1, 0)
	time.Sleep(longTimeout)
	l.Unlock()
	wg.Wait()
	if err1 != nil && err1 != ErrTimeout {
		t.Error(err1)
	}
	if err2 != nil && err2 != ErrTimeout {
		t.Error(err2)
	}
}
