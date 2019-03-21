/*
* provide a monotonic increasing count with channel lock, go to zero if overflow
 */
package monocnt

import "github.com/tannhauser9000/utils/lock"

// types
const type8 = "uint8"
const type16 = "uint16"
const type32 = "uint32"

// struct for monocnt
type CntSt struct {
	cnt      uint32     // actual counter
	cnt_type string     // counter type, uint8, uint16, uint32 available
	lock     *lock.Lock // lock for atomic access
}

// initialization
func InitCount(cntType string) *CntSt {
	this := &CntSt{
		cnt:      uint32(0),
		cnt_type: cntType,
	}
	(*this).lock, _ = lock.GetLock()
	return this
}

// get count
func (c *CntSt) Get() uint {
	(*c).lock.Lock()
	this := (*c).cnt
	(*c).cnt++
	(*c).lock.Unlock()
	if (*c).cnt_type == type8 {
		return uint(uint8(this))
	}
	if (*c).cnt_type == type16 {
		return uint(uint16(this))
	}
	if (*c).cnt_type == type32 {
		return uint(this)
	}
	return 0
}
