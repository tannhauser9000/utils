package lock

/*
* author: tannhauser ruan
* email: tannhauser.sphin@gmail.com
* summary: lock utility from golang channel
 */

import "time"

import "github.com/tannhauser9000/utils/pool"

// Error is for contant error
type Error string

// Lock is for regular lock
type Lock struct {
	lock       *pool.IndexPool
	unlockable bool
}

// RWLock is for regular read-write lock, read-preferring implementation
type RWLock struct {
	rLock  *Lock
	wLock  *Lock
	reader int
}

// TimeoutLock is for lock with timeout, avoid indefinitely waiting
type TimeoutLock struct {
	lock       *pool.TimeoutIndexPool
	unlockable bool
}

// TimeoutRWLock is for read-write lock with timeout, avoid indefinitely waiting, read-preferring implementation
type TimeoutRWLock struct {
	rLock  *TimeoutLock
	wLock  *TimeoutLock
	reader int
}

// constant error for lock acquirement

// ErrTimeout is for lock acquire timeout
const ErrTimeout = Error("Lock timeout")

// ErrUnknown is for unknown situation
const ErrUnknown = Error("Unknown situation when acquiring lock")

// ErrDuplicatedUnlock is for duplicated unlock
const ErrDuplicatedUnlock = Error("Duplicated unlock operation")

// Error is for constant error
func (e Error) Error() string {
	return string(e)
}

// GetLock get a regular lock
func GetLock() *Lock {
	l := &Lock{}
	(*l).lock = &pool.IndexPool{}
	(*l).lock.Init(1)
	(*l).unlockable = false
	return l
}

// Lock lock a regular lock
func (l *Lock) Lock() {
	(*l).lock.Get()
	(*l).unlockable = true
}

// Unlock unlock a regular lock
func (l *Lock) Unlock() error {
	if !(*l).unlockable {
		return ErrDuplicatedUnlock
	}
	(*l).lock.Free(0)
	(*l).unlockable = false
	return nil
}

// GetRWLock get a regular read-write lock
func GetRWLock() *RWLock {
	rwLock := &RWLock{}
	(*rwLock).rLock = GetLock()
	(*rwLock).wLock = GetLock()
	(*rwLock).reader = 0
	return rwLock
}

// RLock read lock a read-write lock
func (rwLock *RWLock) RLock() {
	(*rwLock).rLock.Lock()
	(*rwLock).reader++
	if (*rwLock).reader == 1 {
		(*rwLock).wLock.Lock()
	}
	(*rwLock).rLock.Unlock()
}

// RUnlock read unlock a read-write lock
func (rwLock *RWLock) RUnlock() {
	(*rwLock).rLock.Lock()
	(*rwLock).reader--
	if (*rwLock).reader == 0 {
		(*rwLock).wLock.Unlock()
	}
	(*rwLock).rLock.Unlock()
}

// Lock write lock a read-write lock
func (rwLock *RWLock) Lock() {
	(*rwLock).wLock.Lock()
}

// Unlock write unlock a read-write lock
func (rwLock *RWLock) Unlock() {
	(*rwLock).wLock.Unlock()
}

// GetTimeoutLock get a timeout lock
func GetTimeoutLock(timeout time.Duration) *TimeoutLock {
	l := &TimeoutLock{}
	(*l).lock = &pool.TimeoutIndexPool{}
	(*l).lock.Init(1, timeout)
	(*l).unlockable = false
	return l
}

// Lock lock a timeout lock
func (l *TimeoutLock) Lock() error {
	_, err := (*l).lock.Get()
	if err != nil && err != pool.ErrTimeout {
		err = ErrUnknown
	}
	if err != nil && err == pool.ErrTimeout {
		err = ErrTimeout
	}
	if err == nil {
		(*l).unlockable = true
	}
	return err
}

// Unlock unlock a timeout lock
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

// GetTimeoutRWLock get a timeout read-write lock
func GetTimeoutRWLock(timeout time.Duration) *TimeoutRWLock {
	rwLock := &TimeoutRWLock{}
	(*rwLock).rLock = GetTimeoutLock(timeout)
	(*rwLock).wLock = GetTimeoutLock(timeout)
	(*rwLock).reader = 0
	return rwLock
}

// RLock read lock a read-write lock w/ timeout
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

// RUnlock read unlock a read-write lock w/ timeout
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

// Lock write lock a read-write lock w/ timeout
func (rwLock *TimeoutRWLock) Lock() error {
	return (*rwLock).wLock.Lock()
}

// Unlock write unlock a read-write lock w/ timeout
func (rwLock *TimeoutRWLock) Unlock() error {
	return (*rwLock).wLock.Unlock()
}
