/*
* author: tannhauser ruan
* email: tannhauser.sphin@gmail.com
* summary: lock utility from golang channel
 */
package lock

import "time"

import "github.com/tannhauser9000/utils/pool"

// for contant error
type Error string

// regular lock
type Lock struct {
	lock       *pool.IndexPool
	unlockable bool
}

// regular read-write lock, read-preferring implementation
type RWLock struct {
	rLock  *Lock
	wLock  *Lock
	reader int
}

// lock with timeout, avoid indefinitely waiting
type TimeoutLock struct {
	lock       *pool.TimeoutIndexPool
	unlockable bool
}

// read-write lock with timeout, avoid indefinitely waiting, read-preferring implementation
type TimeoutRWLock struct {
	rLock  *TimeoutLock
	wLock  *TimeoutLock
	reader int
}

// constant error for lock acquirement
const ErrTimeout = Error("Lock timeout")
const ErrUnknown = Error("Unknown situation when acquiring lock")
const ErrDuplicatedUnlock = Error("Duplicated unlock operation")

// function for constant error
func (e Error) Error() string {
	return string(e)
}

// get a regular lock
func GetLock() *Lock {
	l := &Lock{}
	(*l).lock = &pool.IndexPool{}
	(*l).lock.Init(1)
	(*l).unlockable = false
	return l
}

// lock a regular lock
func (l *Lock) Lock() {
	(*l).lock.Get()
	(*l).unlockable = true
}

// unlock a regular lock
func (l *Lock) Unlock() error {
	if !(*l).unlockable {
		return ErrDuplicatedUnlock
	}
	(*l).lock.Free(0)
	(*l).unlockable = false
	return nil
}

// get a regular read-write lock
func GetRWLock() *RWLock {
	rwLock := &RWLock{}
	(*rwLock).rLock = GetLock()
	(*rwLock).wLock = GetLock()
	(*rwLock).reader = 0
	return rwLock
}

// read lock a read-write lock
func (rwLock *RWLock) RLock() {
	(*rwLock).rLock.Lock()
	(*rwLock).reader++
	if (*rwLock).reader == 1 {
		(*rwLock).wLock.Lock()
	}
	(*rwLock).rLock.Unlock()
}

// read unlock a read-write lock
func (rwLock *RWLock) RUnlock() {
	(*rwLock).rLock.Lock()
	(*rwLock).reader--
	if (*rwLock).reader == 0 {
		(*rwLock).wLock.Unlock()
	}
	(*rwLock).rLock.Unlock()
}

// write lock a read-write lock
func (rwLock *RWLock) Lock() {
	(*rwLock).wLock.Lock()
}

// write unlock a read-write lock
func (rwLock *RWLock) Unlock() {
	(*rwLock).wLock.Unlock()
}

// get a timeout lock
func GetTimeoutLock(timeout time.Duration) *TimeoutLock {
	l := &TimeoutLock{}
	(*l).lock = &pool.TimeoutIndexPool{}
	(*l).lock.Init(1, timeout)
	(*l).unlockable = false
	return l
}

// lock a timeout lock
func (l *TimeoutLock) Lock() error {
	_, err := (*l).lock.Get()
	if err != nil && err != pool.ErrTimeout {
		err = ErrUnknown
	}
	if err != nil {
		err = ErrTimeout
	} else {
		(*l).unlockable = true
	}
	return err
}

// unlock a timeout lock
func (l *TimeoutLock) Unlock() error {
	if !(*l).unlockable {
		return ErrDuplicatedUnlock
	}
	err := (*l).lock.Free(0)
	if err != nil && err != pool.ErrFreeTimeout {
		err = ErrUnknown
	}
	if err != nil {
		err = ErrTimeout
	} else {
		(*l).unlockable = false
	}
	return err
}

// get a timeout read-write lock
func GetTimeoutRWLock(timeout time.Duration) *TimeoutRWLock {
	rwLock := &TimeoutRWLock{}
	(*rwLock).rLock = GetTimeoutLock(timeout)
	(*rwLock).wLock = GetTimeoutLock(timeout)
	(*rwLock).reader = 0
	return rwLock
}

// read lock a read-write lock
func (rwLock *TimeoutRWLock) RLock() error {
	err := (*rwLock).rLock.Lock()
	if err != nil {
		return err
	}
	(*rwLock).reader++
	if (*rwLock).reader == 1 {
		err = (*rwLock).wLock.Lock()
	}
	if err != nil {
		return err
	}
	return (*rwLock).rLock.Unlock()
}

// read unlock a read-write lock
func (rwLock *TimeoutRWLock) RUnlock() error {
	err := (*rwLock).rLock.Lock()
	if err != nil {
		return err
	}
	(*rwLock).reader--
	if (*rwLock).reader == 0 {
		err = (*rwLock).wLock.Unlock()
	}
	if err != nil {
		return err
	}
	return (*rwLock).rLock.Unlock()
}

// write lock a read-write lock
func (rwLock *TimeoutRWLock) Lock() error {
	return (*rwLock).wLock.Lock()
}

// write unlock a read-write lock
func (rwLock *TimeoutRWLock) Unlock() error {
	return (*rwLock).wLock.Unlock()
}
