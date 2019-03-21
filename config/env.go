/*
* creating config from environment variable
 */
package config

import "os"
import "strconv"

// initialzation
func (m *MutableConfSt) SetEnvConf(prefix string) {
	var err error
	var this string
	var found bool
	(*m).lock.Lock()
	defer (*m).lock.Unlock()
	(*m).prefix = prefix
	for k, v := range (*m).b {
		this = os.Getenv(prefix + k)
		var value bool
		found = false
		if this != "" {
			found = true
			value, err = strconv.ParseBool(this)
		}
		if found && err == nil {
			(*v).value = value
		}
	}
	for k, v := range (*m).f32 {
		this = os.Getenv(prefix + k)
		var value float64
		found = false
		if this != "" {
			found = true
			value, err = strconv.ParseFloat(this, 32)
		}
		if found && err == nil {
			(*v).value = float32(value)
		}
	}
	for k, v := range (*m).f64 {
		this = os.Getenv(prefix + k)
		var value float64
		found = false
		if this != "" {
			found = true
			value, err = strconv.ParseFloat(this, 64)
		}
		if found && err == nil {
			(*v).value = float64(value)
		}
	}
	for k, v := range (*m).i {
		this = os.Getenv(prefix + k)
		var value int64
		found = false
		if this != "" {
			found = true
			value, err = strconv.ParseInt(this, 10, 32)
		}
		if found && err == nil {
			(*v).value = int(value)
		}
	}
	for k, v := range (*m).i8 {
		this = os.Getenv(prefix + k)
		var value int64
		found = false
		if this != "" {
			found = true
			value, err = strconv.ParseInt(this, 10, 8)
		}
		if found && err == nil {
			(*v).value = int8(value)
		}
	}
	for k, v := range (*m).i16 {
		this = os.Getenv(prefix + k)
		var value int64
		found = false
		if this != "" {
			found = true
			value, err = strconv.ParseInt(this, 10, 16)
		}
		if found && err == nil {
			(*v).value = int16(value)
		}
	}
	for k, v := range (*m).i32 {
		this = os.Getenv(prefix + k)
		var value int64
		found = false
		if this != "" {
			found = true
			value, err = strconv.ParseInt(this, 10, 32)
		}
		if found && err == nil {
			(*v).value = int32(value)
		}
	}
	for k, v := range (*m).i64 {
		this = os.Getenv(prefix + k)
		var value int64
		found = false
		if this != "" {
			found = true
			value, err = strconv.ParseInt(this, 10, 64)
		}
		if found && err == nil {
			(*v).value = int64(value)
		}
	}
	for k, v := range (*m).s {
		this = os.Getenv(prefix + k)
		found = false
		if this != "" {
			found = true
			(*v).value = this
		}
	}
	for k, v := range (*m).ui {
		this = os.Getenv(prefix + k)
		var value uint64
		found = false
		if this != "" {
			found = true
			value, err = strconv.ParseUint(this, 10, 32)
		}
		if found && err == nil {
			(*v).value = uint(value)
		}
	}
	for k, v := range (*m).ui8 {
		this = os.Getenv(prefix + k)
		var value uint64
		found = false
		if this != "" {
			found = true
			value, err = strconv.ParseUint(this, 10, 8)
		}
		if found && err == nil {
			(*v).value = uint8(value)
		}
	}
	for k, v := range (*m).ui16 {
		this = os.Getenv(prefix + k)
		var value uint64
		found = false
		if this != "" {
			found = true
			value, err = strconv.ParseUint(this, 10, 16)
		}
		if found && err == nil {
			(*v).value = uint16(value)
		}
	}
	for k, v := range (*m).ui32 {
		this = os.Getenv(prefix + k)
		var value uint64
		found = false
		if this != "" {
			found = true
			value, err = strconv.ParseUint(this, 10, 32)
		}
		if found && err == nil {
			(*v).value = uint32(value)
		}
	}
	for k, v := range (*m).ui64 {
		this = os.Getenv(prefix + k)
		var value uint64
		found = false
		if this != "" {
			found = true
			value, err = strconv.ParseUint(this, 10, 64)
		}
		if found && err == nil {
			(*v).value = uint64(value)
		}
	}
}
