/*
* creating mutable and immutable config
 */
package config

import "encoding/json"
import "fmt"
import "os"
import "reflect"
import "github.com/tannhauser9000/utils/lock"

// constant error
const ErrNotFound = Error("Target config not found")
const ErrNotSupport = Error("Target config type is currently not supported")

// structure for configuration items
type boolConf struct {
	value   bool
	mutable bool
}

type float32Conf struct {
	value   float32
	mutable bool
}

type float64Conf struct {
	value   float64
	mutable bool
}

type intConf struct {
	value   int
	mutable bool
}

type int8Conf struct {
	value   int8
	mutable bool
}

type int16Conf struct {
	value   int16
	mutable bool
}

type int32Conf struct {
	value   int32
	mutable bool
}

type int64Conf struct {
	value   int64
	mutable bool
}

type stringConf struct {
	value   string
	mutable bool
}

type uintConf struct {
	value   uint
	mutable bool
}

type uint8Conf struct {
	value   uint8
	mutable bool
}

type uint16Conf struct {
	value   uint16
	mutable bool
}

type uint32Conf struct {
	value   uint32
	mutable bool
}

type uint64Conf struct {
	value   uint64
	mutable bool
}

// actual configuration instance
type MutableConfSt struct {
	reload  string                  // a random string to check if we should reload the config
	prefix  string                  // prefix for environment variable
	sleep   int                     // update routine sleep second
	routine *routineSt              // routine structure for updating env
	lock    *lock.RWLock            // conf lock
	init    bool                    // is initialzed?
	b       map[string]*boolConf    // map for boolean value
	f32     map[string]*float32Conf // map for float64 value
	f64     map[string]*float64Conf // map for float64 value
	i       map[string]*intConf     // map for int value
	i8      map[string]*int8Conf    // map for int8 value
	i16     map[string]*int16Conf   // map for int16 value
	i32     map[string]*int32Conf   // map for int32 value
	i64     map[string]*int64Conf   // map for int64 value
	s       map[string]*stringConf  // map for string value
	ui      map[string]*uintConf    // map for uint value
	ui8     map[string]*uint8Conf   // map for uint8 value
	ui16    map[string]*uint16Conf  // map for uint16 value
	ui32    map[string]*uint32Conf  // map for uint32 value
	ui64    map[string]*uint64Conf  // map for uint64 value
}

// initiate configuration
func InitJSONMutableConf(path string, conf interface{}, mute *map[string]bool) (*MutableConfSt, error) {
	fp, err := os.Open(path)
	defer fp.Close()
	if err != nil {
		return nil, err
	}
	decoder := json.NewDecoder(fp)
	err = decoder.Decode(conf)
	if err != nil {
		return nil, err
	}
	v := reflect.ValueOf(conf)
	v = reflect.Indirect(v)
	t := v.Type()
	assign := false
	st := &MutableConfSt{
		b:    make(map[string]*boolConf),    // map for boolean value
		f32:  make(map[string]*float32Conf), // map for float64 value
		f64:  make(map[string]*float64Conf), // map for float64 value
		i:    make(map[string]*intConf),     // map for int value
		i8:   make(map[string]*int8Conf),    // map for int8 value
		i16:  make(map[string]*int16Conf),   // map for int16 value
		i32:  make(map[string]*int32Conf),   // map for int32 value
		i64:  make(map[string]*int64Conf),   // map for int64 value
		s:    make(map[string]*stringConf),  // map for string value
		ui:   make(map[string]*uintConf),    // map for uint value
		ui8:  make(map[string]*uint8Conf),   // map for uint8 value
		ui16: make(map[string]*uint16Conf),  // map for uint16 value
		ui32: make(map[string]*uint32Conf),  // map for uint32 value
		ui64: make(map[string]*uint64Conf),  // map for uint64 value
	}
	(*st).lock, _ = lock.GetRWLock()
	mutable, ok := false, false
	(*st).lock.Lock()
	defer (*st).lock.Unlock()
	for i := 0; i < v.NumField(); i++ {
		assign = false
		mutable, ok = (*mute)[t.Field(i).Name]
		mutable = mutable && ok
		if !assign && v.Field(i).Type().Name() == "bool" {
			(*st).b[t.Field(i).Name] = &boolConf{
				value:   v.Field(i).Interface().(bool),
				mutable: mutable,
			}
			assign = true
		}
		if !assign && v.Field(i).Type().Name() == "float32" {
			(*st).f32[t.Field(i).Name] = &float32Conf{
				value:   v.Field(i).Interface().(float32),
				mutable: mutable,
			}
			assign = true
		}
		if !assign && v.Field(i).Type().Name() == "float64" {
			(*st).f64[t.Field(i).Name] = &float64Conf{
				value:   v.Field(i).Interface().(float64),
				mutable: mutable,
			}
			assign = true
		}
		if !assign && v.Field(i).Type().Name() == "int" {
			(*st).i[t.Field(i).Name] = &intConf{
				value:   v.Field(i).Interface().(int),
				mutable: mutable,
			}
			assign = true
		}
		if !assign && v.Field(i).Type().Name() == "int8" {
			(*st).i8[t.Field(i).Name] = &int8Conf{
				value:   v.Field(i).Interface().(int8),
				mutable: mutable,
			}
			assign = true
		}
		if !assign && v.Field(i).Type().Name() == "int16" {
			(*st).i16[t.Field(i).Name] = &int16Conf{
				value:   v.Field(i).Interface().(int16),
				mutable: mutable,
			}
			assign = true
		}
		if !assign && v.Field(i).Type().Name() == "int32" {
			(*st).i32[t.Field(i).Name] = &int32Conf{
				value:   v.Field(i).Interface().(int32),
				mutable: mutable,
			}
			assign = true
		}
		if !assign && v.Field(i).Type().Name() == "int64" {
			(*st).i64[t.Field(i).Name] = &int64Conf{
				value:   v.Field(i).Interface().(int64),
				mutable: mutable,
			}
			assign = true
		}
		if !assign && v.Field(i).Type().Name() == "string" {
			(*st).s[t.Field(i).Name] = &stringConf{
				value:   v.Field(i).Interface().(string),
				mutable: mutable,
			}
			assign = true
		}
		if !assign && v.Field(i).Type().Name() == "uint" {
			(*st).ui[t.Field(i).Name] = &uintConf{
				value:   v.Field(i).Interface().(uint),
				mutable: mutable,
			}
			assign = true
		}
		if !assign && v.Field(i).Type().Name() == "uint8" {
			(*st).ui8[t.Field(i).Name] = &uint8Conf{
				value:   v.Field(i).Interface().(uint8),
				mutable: mutable,
			}
			assign = true
		}
		if !assign && v.Field(i).Type().Name() == "uint16" {
			(*st).ui16[t.Field(i).Name] = &uint16Conf{
				value:   v.Field(i).Interface().(uint16),
				mutable: mutable,
			}
			assign = true
		}
		if !assign && v.Field(i).Type().Name() == "uint32" {
			(*st).ui32[t.Field(i).Name] = &uint32Conf{
				value:   v.Field(i).Interface().(uint32),
				mutable: mutable,
			}
			assign = true
		}
		if !assign && v.Field(i).Type().Name() == "uint64" {
			(*st).ui64[t.Field(i).Name] = &uint64Conf{
				value:   v.Field(i).Interface().(uint64),
				mutable: mutable,
			}
			assign = true
		}
		if !assign {
			return nil, ErrNotSupport
		}
	}
	(*st).init = true
	return st, nil
}

// is conf initialized?
func (m *MutableConfSt) Init() bool {
	return (*m).init
}

// get configuration item
func (m *MutableConfSt) GetBool(name string) *bool {
	if !(*m).init {
		return nil
	}
	(*m).lock.RLock()
	defer (*m).lock.RUnlock()
	this, ok := (*m).b[name]
	if !ok {
		return nil
	}
	value := (*this).value
	return &value
}

func (m *MutableConfSt) GetFloat32(name string) *float32 {
	if !(*m).init {
		return nil
	}
	(*m).lock.RLock()
	defer (*m).lock.RUnlock()
	this, ok := (*m).f32[name]
	if !ok {
		return nil
	}
	value := (*this).value
	return &value
}

func (m *MutableConfSt) GetFloat64(name string) *float64 {
	if !(*m).init {
		return nil
	}
	(*m).lock.RLock()
	defer (*m).lock.RUnlock()
	this, ok := (*m).f64[name]
	if !ok {
		return nil
	}
	value := (*this).value
	return &value
}

func (m *MutableConfSt) GetInt(name string) *int {
	if !(*m).init {
		return nil
	}
	(*m).lock.RLock()
	defer (*m).lock.RUnlock()
	this, ok := (*m).i[name]
	if !ok {
		return nil
	}
	value := (*this).value
	return &value
}

func (m *MutableConfSt) GetInt8(name string) *int8 {
	if !(*m).init {
		return nil
	}
	(*m).lock.RLock()
	defer (*m).lock.RUnlock()
	this, ok := (*m).i8[name]
	if !ok {
		return nil
	}
	value := (*this).value
	return &value
}

func (m *MutableConfSt) GetInt16(name string) *int16 {
	if !(*m).init {
		return nil
	}
	(*m).lock.RLock()
	defer (*m).lock.RUnlock()
	this, ok := (*m).i16[name]
	if !ok {
		return nil
	}
	value := (*this).value
	return &value
}

func (m *MutableConfSt) GetInt32(name string) *int32 {
	if !(*m).init {
		return nil
	}
	(*m).lock.RLock()
	defer (*m).lock.RUnlock()
	this, ok := (*m).i32[name]
	if !ok {
		return nil
	}
	value := (*this).value
	return &value
}

func (m *MutableConfSt) GetInt64(name string) *int64 {
	if !(*m).init {
		return nil
	}
	(*m).lock.RLock()
	defer (*m).lock.RUnlock()
	this, ok := (*m).i64[name]
	if !ok {
		return nil
	}
	value := (*this).value
	return &value
}

func (m *MutableConfSt) GetString(name string) *string {
	if !(*m).init {
		return nil
	}
	(*m).lock.RLock()
	defer (*m).lock.RUnlock()
	this, ok := (*m).s[name]
	if !ok {
		return nil
	}
	value := (*this).value
	return &value
}

func (m *MutableConfSt) GetUint(name string) *uint {
	if !(*m).init {
		return nil
	}
	(*m).lock.RLock()
	defer (*m).lock.RUnlock()
	this, ok := (*m).ui[name]
	if !ok {
		return nil
	}
	value := (*this).value
	return &value
}

func (m *MutableConfSt) GetUint8(name string) *uint8 {
	if !(*m).init {
		return nil
	}
	(*m).lock.RLock()
	defer (*m).lock.RUnlock()
	this, ok := (*m).ui8[name]
	if !ok {
		return nil
	}
	value := (*this).value
	return &value
}

func (m *MutableConfSt) GetUint16(name string) *uint16 {
	if !(*m).init {
		return nil
	}
	(*m).lock.RLock()
	defer (*m).lock.RUnlock()
	this, ok := (*m).ui16[name]
	if !ok {
		return nil
	}
	value := (*this).value
	return &value
}

func (m *MutableConfSt) GetUint32(name string) *uint32 {
	if !(*m).init {
		return nil
	}
	(*m).lock.RLock()
	defer (*m).lock.RUnlock()
	this, ok := (*m).ui32[name]
	if !ok {
		return nil
	}
	value := (*this).value
	return &value
}

func (m *MutableConfSt) GetUint64(name string) *uint64 {
	if !(*m).init {
		return nil
	}
	(*m).lock.RLock()
	defer (*m).lock.RUnlock()
	this, ok := (*m).ui64[name]
	if !ok {
		return nil
	}
	value := (*this).value
	return &value
}

// set mutable configurations, if immutable or not found, return false
func (m *MutableConfSt) SetBool(name string, value bool) bool {
	if !(*m).init {
		return false
	}
	this, ok := (*m).b[name]
	if !ok {
		return false
	}
	if (*this).mutable {
		(*this).value = value
	}
	return (*this).mutable
}

func (m *MutableConfSt) SetFloat32(name string, value float32) bool {
	if !(*m).init {
		return false
	}
	this, ok := (*m).f32[name]
	if !ok {
		return false
	}
	if (*this).mutable {
		(*this).value = value
	}
	return (*this).mutable
}

func (m *MutableConfSt) SetFloat64(name string, value float64) bool {
	if !(*m).init {
		return false
	}
	this, ok := (*m).f64[name]
	if !ok {
		return false
	}
	if (*this).mutable {
		(*this).value = value
	}
	return (*this).mutable
}

func (m *MutableConfSt) SetInt(name string, value int) bool {
	if !(*m).init {
		return false
	}
	this, ok := (*m).i[name]
	if !ok {
		return false
	}
	if (*this).mutable {
		(*this).value = value
	}
	return (*this).mutable
}

func (m *MutableConfSt) SetInt8(name string, value int8) bool {
	if !(*m).init {
		return false
	}
	this, ok := (*m).i8[name]
	if !ok {
		return false
	}
	if (*this).mutable {
		(*this).value = value
	}
	return (*this).mutable
}

func (m *MutableConfSt) SetInt16(name string, value int16) bool {
	if !(*m).init {
		return false
	}
	this, ok := (*m).i16[name]
	if !ok {
		return false
	}
	if (*this).mutable {
		(*this).value = value
	}
	return (*this).mutable
}

func (m *MutableConfSt) SetInt32(name string, value int32) bool {
	if !(*m).init {
		return false
	}
	this, ok := (*m).i32[name]
	if !ok {
		return false
	}
	if (*this).mutable {
		(*this).value = value
	}
	return (*this).mutable
}

func (m *MutableConfSt) SetInt64(name string, value int64) bool {
	if !(*m).init {
		return false
	}
	this, ok := (*m).i64[name]
	if !ok {
		return false
	}
	if (*this).mutable {
		(*this).value = value
	}
	return (*this).mutable
}

func (m *MutableConfSt) SetString(name string, value string) bool {
	if !(*m).init {
		return false
	}
	this, ok := (*m).s[name]
	if !ok {
		return false
	}
	if (*this).mutable {
		(*this).value = value
	}
	return (*this).mutable
}

func (m *MutableConfSt) SetUint(name string, value uint) bool {
	if !(*m).init {
		return false
	}
	this, ok := (*m).ui[name]
	if !ok {
		return false
	}
	if (*this).mutable {
		(*this).value = value
	}
	return (*this).mutable
}

func (m *MutableConfSt) SetUint8(name string, value uint8) bool {
	if !(*m).init {
		return false
	}
	this, ok := (*m).ui8[name]
	if !ok {
		return false
	}
	if (*this).mutable {
		(*this).value = value
	}
	return (*this).mutable
}

func (m *MutableConfSt) SetUint16(name string, value uint16) bool {
	if !(*m).init {
		return false
	}
	this, ok := (*m).ui16[name]
	if !ok {
		return false
	}
	if (*this).mutable {
		(*this).value = value
	}
	return (*this).mutable
}

func (m *MutableConfSt) SetUint32(name string, value uint32) bool {
	if !(*m).init {
		return false
	}
	this, ok := (*m).ui32[name]
	if !ok {
		return false
	}
	if (*this).mutable {
		(*this).value = value
	}
	return (*this).mutable
}

func (m *MutableConfSt) SetUint64(name string, value uint64) bool {
	if !(*m).init {
		return false
	}
	this, ok := (*m).ui64[name]
	if !ok {
		return false
	}
	if (*this).mutable {
		(*this).value = value
	}
	return (*this).mutable
}

// print configuration
func (m *MutableConfSt) Print() {
	if !(*m).init {
		fmt.Printf("{\"error\": \"MutableConfSt is not yet initialized!\"}\n")
		return
	}
	c := fmt.Sprintf("{\n")
	for k, v := range (*m).b {
		c = fmt.Sprintf("%s  \"%s\": {\n", c, k)
		c = fmt.Sprintf("%s    \"value\": %v,\n", c, v.value)
		c = fmt.Sprintf("%s    \"mutable\": %v\n", c, v.mutable)
		c = fmt.Sprintf("%s  },\n", c)
	}
	for k, v := range (*m).f32 {
		c = fmt.Sprintf("%s  \"%s\": {\n", c, k)
		c = fmt.Sprintf("%s    \"value\": %v,\n", c, v.value)
		c = fmt.Sprintf("%s    \"mutable\": %v\n", c, v.mutable)
		c = fmt.Sprintf("%s  },\n", c)
	}
	for k, v := range (*m).f64 {
		c = fmt.Sprintf("%s  \"%s\": {\n", c, k)
		c = fmt.Sprintf("%s    \"value\": %v,\n", c, v.value)
		c = fmt.Sprintf("%s    \"mutable\": %v\n", c, v.mutable)
		c = fmt.Sprintf("%s  },\n", c)
	}
	for k, v := range (*m).i {
		c = fmt.Sprintf("%s  \"%s\": {\n", c, k)
		c = fmt.Sprintf("%s    \"value\": %v,\n", c, v.value)
		c = fmt.Sprintf("%s    \"mutable\": %v\n", c, v.mutable)
		c = fmt.Sprintf("%s  },\n", c)
	}
	for k, v := range (*m).i8 {
		c = fmt.Sprintf("%s  \"%s\": {\n", c, k)
		c = fmt.Sprintf("%s    \"value\": %v,\n", c, v.value)
		c = fmt.Sprintf("%s    \"mutable\": %v\n", c, v.mutable)
		c = fmt.Sprintf("%s  },\n", c)
	}
	for k, v := range (*m).i16 {
		c = fmt.Sprintf("%s  \"%s\": {\n", c, k)
		c = fmt.Sprintf("%s    \"value\": %v,\n", c, v.value)
		c = fmt.Sprintf("%s    \"mutable\": %v\n", c, v.mutable)
		c = fmt.Sprintf("%s  },\n", c)
	}
	for k, v := range (*m).i32 {
		c = fmt.Sprintf("%s  \"%s\": {\n", c, k)
		c = fmt.Sprintf("%s    \"value\": %v,\n", c, v.value)
		c = fmt.Sprintf("%s    \"mutable\": %v\n", c, v.mutable)
		c = fmt.Sprintf("%s  },\n", c)
	}
	for k, v := range (*m).i64 {
		c = fmt.Sprintf("%s  \"%s\": {\n", c, k)
		c = fmt.Sprintf("%s    \"value\": %v,\n", c, v.value)
		c = fmt.Sprintf("%s    \"mutable\": %v\n", c, v.mutable)
		c = fmt.Sprintf("%s  },\n", c)
	}
	for k, v := range (*m).s {
		c = fmt.Sprintf("%s  \"%s\": {\n", c, k)
		c = fmt.Sprintf("%s    \"value\": \"%v\",\n", c, v.value)
		c = fmt.Sprintf("%s    \"mutable\": %v\n", c, v.mutable)
		c = fmt.Sprintf("%s  },\n", c)
	}
	for k, v := range (*m).ui {
		c = fmt.Sprintf("%s  \"%s\": {\n", c, k)
		c = fmt.Sprintf("%s    \"value\": %v,\n", c, v.value)
		c = fmt.Sprintf("%s    \"mutable\": %v\n", c, v.mutable)
		c = fmt.Sprintf("%s  },\n", c)
	}
	for k, v := range (*m).ui8 {
		c = fmt.Sprintf("%s  \"%s\": {\n", c, k)
		c = fmt.Sprintf("%s    \"value\": %v,\n", c, v.value)
		c = fmt.Sprintf("%s    \"mutable\": %v\n", c, v.mutable)
		c = fmt.Sprintf("%s  },\n", c)
	}
	for k, v := range (*m).ui16 {
		c = fmt.Sprintf("%s  \"%s\": {\n", c, k)
		c = fmt.Sprintf("%s    \"value\": %v,\n", c, v.value)
		c = fmt.Sprintf("%s    \"mutable\": %v\n", c, v.mutable)
		c = fmt.Sprintf("%s  },\n", c)
	}
	for k, v := range (*m).ui32 {
		c = fmt.Sprintf("%s  \"%s\": {\n", c, k)
		c = fmt.Sprintf("%s    \"value\": %v,\n", c, v.value)
		c = fmt.Sprintf("%s    \"mutable\": %v\n", c, v.mutable)
		c = fmt.Sprintf("%s  },\n", c)
	}
	for k, v := range (*m).ui64 {
		c = fmt.Sprintf("%s  \"%s\": {\n", c, k)
		c = fmt.Sprintf("%s    \"value\": %v,\n", c, v.value)
		c = fmt.Sprintf("%s    \"mutable\": %v\n", c, v.mutable)
		c = fmt.Sprintf("%s  },\n", c)
	}
	if len(c) > 2 {
		c = c[:len(c)-2]
	}
	c = fmt.Sprintf("%s\n}\n", c)
	fmt.Printf("%s", c)
}
