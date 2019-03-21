/*
* testing for mutable config
 */
package config

import "errors"
import "fmt"
import "os"
import "testing"

const dummyConf = "{\"booli\": true, \"boolm\": true, \"float32i\": 1, \"float32m\": 2, \"float64i\": 1, \"float64m\": 2, \"inti\": 1, \"intm\": 2, \"int8i\": 1, \"int8m\": 2, \"int16i\": 1, \"int16m\": 2, \"int32i\": 1, \"int32m\": 2, \"int64i\": 1, \"int64m\": 2, \"stringi\": \"immutable\", \"stringm\": \"mutable\", \"uinti\": 1, \"uintm\": 2, \"uint8i\": 1, \"uint8m\": 2, \"uint16i\": 1, \"uint16m\": 2, \"uint32i\": 1, \"uint32m\": 2, \"uint64i\": 1, \"uint64m\": 2}"
const dummyPath = "dummy.json"
const dummyLock = "dummy.lock"
const dummyTry = 3

var created bool
var m *MutableConfSt

type dummyConfSt struct {
	Booli    bool    `json:"booli"`
	Boolm    bool    `json:"boolm"`
	Float32i float32 `json:"float32i"`
	Float32m float32 `json:"float32m"`
	Float64i float64 `json:"float64i"`
	Float64m float64 `json:"float64m"`
	Inti     int     `json:"inti"`
	Intm     int     `json:"intm"`
	Int8i    int8    `json:"int8i"`
	Int8m    int8    `json:"int8m"`
	Int16i   int16   `json:"int16i"`
	Int16m   int16   `json:"int16m"`
	Int32i   int32   `json:"int32i"`
	Int32m   int32   `json:"int32m"`
	Int64i   int64   `json:"int64i"`
	Int64m   int64   `json:"int64m"`
	Stringi  string  `json:"stringi"`
	Stringm  string  `json:"stringm"`
	Uinti    uint    `json:"uinti"`
	Uintm    uint    `json:"uintm"`
	Uint8i   uint8   `json:"uint8i"`
	Uint8m   uint8   `json:"uint8m"`
	Uint16i  uint16  `json:"uint16i"`
	Uint16m  uint16  `json:"uint16m"`
	Uint32i  uint32  `json:"uint32i"`
	Uint32m  uint32  `json:"uint32m"`
	Uint64i  uint64  `json:"uint64i"`
	Uint64m  uint64  `json:"uint64m"`
}

func TestInit(t *testing.T) {
	fmt.Printf("Testing Init()...\n")
	err := initConf()
	if err != nil {
		t.Error(err.Error())
		return
	}
	// test Init()
	if m.Init() != true {
		t.Error("(m *MutableConfSt) Init() should return true, got ", m.Init())
	}
	return
}

// test boolean
func TestBool(t *testing.T) {
	fmt.Printf("Testing boolean...\n")
	err := initConf()
	if err != nil {
		t.Error(err.Error())
		return
	}
	// immutable
	// test GetBool()
	b := m.GetBool("Booli")
	if b == nil {
		t.Error("(m *MutableConfSt) GetBool() should not return nil on exist item, got ", b)
		return
	}
	if !*b {
		t.Error("(m *MutableConfSt) GetBool() should return true, got ", *b)
		return
	}
	//test SetBool()
	ok := m.SetBool("Booli", false)
	if ok {
		t.Error("(m *MutableConfSt) SetBool() should return false on immutable item, got ", ok)
		return
	}
	// test GetBool()
	b = m.GetBool("Booli")
	if b == nil {
		t.Error("(m *MutableConfSt) GetBool() should not return nil on exist item, got ", b)
		return
	}
	if !*b {
		t.Error("(m *MutableConfSt) GetBool() should return true on immutable item, got ", *b)
		return
	}
	// test Get() on boolean item
	iface := m.Get("Booli")
	if iface == nil {
		t.Error("(m *MutableConfSt) Get() should not be nil on exist item: ", iface)
		return
	}
	b, ok = iface.(*bool)
	if !ok {
		t.Error("(m *MutableConfSt) Get() should be able to be casted back to *bool: ", ok)
		return
	}
	if !*b {
		t.Error("(m *MutableConfSt) Get() should return true on immutable item: ", *b)
		return
	}
	// test Set() on boolean
	ok = m.Set("Booli", "false")
	if ok {
		t.Error("(m *MutableConfSt) Set() should return on immutable item, got ", ok)
		return
	}
	// test Get() on boolean item after Set()
	iface = m.Get("Booli")
	if iface == nil {
		t.Error("(m *MutableConfSt) Get() should not be nil on exist item: ", iface)
		return
	}
	b, ok = iface.(*bool)
	if !ok {
		t.Error("(m *MutableConfSt) Get() should be able to be casted back to *bool: ", ok)
		return
	}
	if !*b {
		t.Error("(m *MutableConfSt) Get() should return true on immutable item: ", *b)
		return
	}

	// mutable
	// test GetBool()
	b = m.GetBool("Boolm")
	if b == nil {
		t.Error("(m *MutableConfSt) GetBool() should not return nil on exist item, got ", b)
		return
	}
	if !*b {
		t.Error("(m *MutableConfSt) GetBool() should return true, got ", *b)
		return
	}
	//test SetBool()
	ok = m.SetBool("Boolm", false)
	if !ok {
		t.Error("(m *MutableConfSt) SetBool() should return true on mutable item, got ", ok)
		return
	}
	// test GetBool()
	b = m.GetBool("Boolm")
	if b == nil {
		t.Error("(m *MutableConfSt) GetBool() should not return nil on exist item, got ", b)
		return
	}
	if *b {
		t.Error("(m *MutableConfSt) GetBool() should return on mutable, got ", *b)
		return
	}
	// test Get() on boolean item
	iface = m.Get("Boolm")
	if iface == nil {
		t.Error("(m *MutableConfSt) Get() should not be nil on exist item: ", iface)
		return
	}
	b, ok = iface.(*bool)
	if !ok {
		t.Error("(m *MutableConfSt) Get() should be able to be casted back to *bool: ", ok)
		return
	}
	if *b {
		t.Error("(m *MutableConfSt) Get() should return on mutable item: ", *b)
		return
	}
	// test Set() on boolean
	ok = m.Set("Boolm", "true")
	if !ok {
		t.Error("(m *MutableConfSt) Set() should return true on mutable item, got ", ok)
		return
	}
	// test Get() on boolean item after Set()
	iface = m.Get("Boolm")
	if iface == nil {
		t.Error("(m *MutableConfSt) Get() should not be nil on exist item: ", iface)
		return
	}
	b, ok = iface.(*bool)
	if !ok {
		t.Error("(m *MutableConfSt) Get() should be able to be casted back to *bool: ", ok)
		return
	}
	if !*b {
		t.Error("(m *MutableConfSt) Get() should return true on mutable item: ", *b)
		return
	}

	// non exist
	// test GetBool() on non-exist item
	b = m.GetBool("Booln")
	if b != nil {
		t.Error("(m *MutableConfSt) GetBool() should return nil on non-exist item, got ", *b)
		return
	}
	// test SetBool() on non-exist item
	ok = m.SetBool("Booln", false)
	if ok {
		t.Error("(m *MutableConfSt) SetBool() should return on non-exist item, got ", ok)
		return
	}
	// test Get() on non-exist item
	iface = m.Get("Booln")
	if iface != nil {
		t.Error("(m *MutableConfSt) Get() should return nil on non-exist item, got ", iface)
		return
	}
	ok = m.Set("Booln", "false")
	if ok {
		t.Error("(m *MutableConfSt) Set() should return on non-exist item, got ", ok)
		return
	}

	// invalid input
	// test Set() on invalid value
	ok = m.Set("Boolm", "test")
	if ok {
		t.Error("(m *MutableConfSt) Set() should return on invalid value, got ", ok)
		return
	}
	b = m.GetBool("Boolm")
	if b == nil {
		t.Error("(m *MutableConfSt) GetBool() should not return nil on exist item, got ", b)
		return
	}
	if !*b {
		t.Error("(m *MutableConfSt) GetBool() should return true after setting invalid value, got ", *b)
		return
	}
	return
}

// test float32
func TestFloat32(t *testing.T) {
	fmt.Printf("Testing float32...\n")
	err := initConf()
	if err != nil {
		t.Error(err.Error())
		return
	}
	// immutable
	// test GetFloat32()
	b := m.GetFloat32("Float32i")
	if b == nil {
		t.Error("(m *MutableConfSt) GetFloat32() should not return nil on exist item, got ", b)
		return
	}
	if *b != 1 {
		t.Error("(m *MutableConfSt) GetFloat32() should return 1, got ", *b)
		return
	}
	//test SetFloat32()
	ok := m.SetFloat32("Float32i", 3)
	if ok {
		t.Error("(m *MutableConfSt) SetFloat32() should return on immutable item, got ", ok)
		return
	}
	// test GetFloat32()
	b = m.GetFloat32("Float32i")
	if b == nil {
		t.Error("(m *MutableConfSt) GetFloat32() should not return nil on exist item, got ", b)
		return
	}
	if *b != 1 {
		t.Error("(m *MutableConfSt) GetFloat32() should return 1 on immutable item, got ", *b)
		return
	}
	// test Get() on float32 item
	iface := m.Get("Float32i")
	if iface == nil {
		t.Error("(m *MutableConfSt) Get() should not be nil on exist item: ", iface)
		return
	}
	b, ok = iface.(*float32)
	if !ok {
		t.Error("(m *MutableConfSt) Get() should be able to be casted back to *float32: ", ok)
		return
	}
	if *b != 1 {
		t.Error("(m *MutableConfSt) Get() should return 1 on immutable item: ", *b)
		return
	}
	// test Set() on float32
	ok = m.Set("Float32i", "3")
	if ok {
		t.Error("(m *MutableConfSt) Set() should return on immutable item, got ", ok)
		return
	}
	// test Get() on float32 item after Set()
	iface = m.Get("Float32i")
	if iface == nil {
		t.Error("(m *MutableConfSt) Get() should not be nil on exist item: ", iface)
		return
	}
	b, ok = iface.(*float32)
	if !ok {
		t.Error("(m *MutableConfSt) Get() should be able to be casted back to *float32: ", ok)
		return
	}
	if *b != 1 {
		t.Error("(m *MutableConfSt) Get() should return 1 on immutable item: ", *b)
		return
	}

	// mutable
	// test GetFloat32()
	b = m.GetFloat32("Float32m")
	if b == nil {
		t.Error("(m *MutableConfSt) GetFloat32() should not return nil on exist item, got ", b)
		return
	}
	if *b != 2 {
		t.Error("(m *MutableConfSt) GetFloat32() should return 2, got ", *b)
		return
	}
	//test SetFloat32()
	ok = m.SetFloat32("Float32m", 4)
	if !ok {
		t.Error("(m *MutableConfSt) SetFloat32() should return on mutable item, got ", ok)
		return
	}
	// test GetFloat32()
	b = m.GetFloat32("Float32m")
	if b == nil {
		t.Error("(m *MutableConfSt) GetFloat32() should not return nil on exist item, got ", b)
		return
	}
	if *b != 4 {
		t.Error("(m *MutableConfSt) GetFloat32() should return 4 on mutable, got ", *b)
		return
	}
	// test Get() on float32 item
	iface = m.Get("Float32m")
	if iface == nil {
		t.Error("(m *MutableConfSt) Get() should not be nil on exist item: ", iface)
		return
	}
	b, ok = iface.(*float32)
	if !ok {
		t.Error("(m *MutableConfSt) Get() should be able to be casted back to *float32: ", ok)
		return
	}
	if *b != 4 {
		t.Error("(m *MutableConfSt) Get() should return 4 on mutable item: ", *b)
		return
	}
	// test Set() on float32
	ok = m.Set("Float32m", "2")
	if !ok {
		t.Error("(m *MutableConfSt) Set() should return on mutable item, got ", ok)
		return
	}
	// test Get() on float32 item after Set()
	iface = m.Get("Float32m")
	if iface == nil {
		t.Error("(m *MutableConfSt) Get() should not be nil on exist item: ", iface)
		return
	}
	b, ok = iface.(*float32)
	if !ok {
		t.Error("(m *MutableConfSt) Get() should be able to be casted back to *float32: ", ok)
		return
	}
	if *b != 2 {
		t.Error("(m *MutableConfSt) Get() should return 2 on mutable item: ", *b)
		return
	}

	// non exist
	// test GetFloat32() on non-exist item
	b = m.GetFloat32("Float32n")
	if b != nil {
		t.Error("(m *MutableConfSt) GetFloat32() should return nil on non-exist item, got ", *b)
		return
	}
	// test SetFloat32() on non-exist item
	ok = m.SetFloat32("Float32n", 9)
	if ok {
		t.Error("(m *MutableConfSt) SetFloat32() should return on non-exist item, got ", ok)
		return
	}
	// test Get() on non-exist item
	iface = m.Get("Float32n")
	if iface != nil {
		t.Error("(m *MutableConfSt) Get() should return nil on non-exist item, got ", iface)
		return
	}
	ok = m.Set("Float32n", "9")
	if ok {
		t.Error("(m *MutableConfSt) Set() should return on non-exist item, got ", ok)
		return
	}

	// invalid input
	// test Set() on invalid value
	ok = m.Set("Float32m", "test")
	if ok {
		t.Error("(m *MutableConfSt) Set() should return on invalid value, got ", ok)
		return
	}
	b = m.GetFloat32("Float32m")
	if b == nil {
		t.Error("(m *MutableConfSt) GetFloat32() should not return nil on exist item, got ", b)
		return
	}
	if *b != 2 {
		t.Error("(m *MutableConfSt) GetFloat32() should return 2 after setting invalid value, got ", *b)
		return
	}
	return
}

// test float64
func TestFloat64(t *testing.T) {
	fmt.Printf("Testing float64...\n")
	err := initConf()
	if err != nil {
		t.Error(err.Error())
		return
	}
	// immutable
	// test GetFloat64()
	b := m.GetFloat64("Float64i")
	if b == nil {
		t.Error("(m *MutableConfSt) GetFloat64() should not return nil on exist item, got ", b)
		return
	}
	if *b != 1 {
		t.Error("(m *MutableConfSt) GetFloat64() should return 1, got ", *b)
		return
	}
	//test SetFloat64()
	ok := m.SetFloat64("Float64i", 3)
	if ok {
		t.Error("(m *MutableConfSt) SetFloat64() should return on immutable item, got ", ok)
		return
	}
	// test GetFloat64()
	b = m.GetFloat64("Float64i")
	if b == nil {
		t.Error("(m *MutableConfSt) GetFloat64() should not return nil on exist item, got ", b)
		return
	}
	if *b != 1 {
		t.Error("(m *MutableConfSt) GetFloat64() should return 1 on immutable item, got ", *b)
		return
	}
	// test Get() on float64 item
	iface := m.Get("Float64i")
	if iface == nil {
		t.Error("(m *MutableConfSt) Get() should not be nil on exist item: ", iface)
		return
	}
	b, ok = iface.(*float64)
	if !ok {
		t.Error("(m *MutableConfSt) Get() should be able to be casted back to *float64: ", ok)
		return
	}
	if *b != 1 {
		t.Error("(m *MutableConfSt) Get() should return 1 on immutable item: ", *b)
		return
	}
	// test Set() on float64
	ok = m.Set("Float64i", "3")
	if ok {
		t.Error("(m *MutableConfSt) Set() should return on immutable item, got ", ok)
		return
	}
	// test Get() on float64 item after Set()
	iface = m.Get("Float64i")
	if iface == nil {
		t.Error("(m *MutableConfSt) Get() should not be nil on exist item: ", iface)
		return
	}
	b, ok = iface.(*float64)
	if !ok {
		t.Error("(m *MutableConfSt) Get() should be able to be casted back to *float64: ", ok)
		return
	}
	if *b != 1 {
		t.Error("(m *MutableConfSt) Get() should return 1 on immutable item: ", *b)
		return
	}

	// mutable
	// test GetFloat64()
	b = m.GetFloat64("Float64m")
	if b == nil {
		t.Error("(m *MutableConfSt) GetFloat64() should not return nil on exist item, got ", b)
		return
	}
	if *b != 2 {
		t.Error("(m *MutableConfSt) GetFloat64() should return 2, got ", *b)
		return
	}
	//test SetFloat64()
	ok = m.SetFloat64("Float64m", 4)
	if !ok {
		t.Error("(m *MutableConfSt) SetFloat64() should return on mutable item, got ", ok)
		return
	}
	// test GetFloat64()
	b = m.GetFloat64("Float64m")
	if b == nil {
		t.Error("(m *MutableConfSt) GetFloat64() should not return nil on exist item, got ", b)
		return
	}
	if *b != 4 {
		t.Error("(m *MutableConfSt) GetFloat64() should return 4 on mutable, got ", *b)
		return
	}
	// test Get() on float64 item
	iface = m.Get("Float64m")
	if iface == nil {
		t.Error("(m *MutableConfSt) Get() should not be nil on exist item: ", iface)
		return
	}
	b, ok = iface.(*float64)
	if !ok {
		t.Error("(m *MutableConfSt) Get() should be able to be casted back to *float64: ", ok)
		return
	}
	if *b != 4 {
		t.Error("(m *MutableConfSt) Get() should return 4 on mutable item: ", *b)
		return
	}
	// test Set() on float64
	ok = m.Set("Float64m", "2")
	if !ok {
		t.Error("(m *MutableConfSt) Set() should return on mutable item, got ", ok)
		return
	}
	// test Get() on float64 item after Set()
	iface = m.Get("Float64m")
	if iface == nil {
		t.Error("(m *MutableConfSt) Get() should not be nil on exist item: ", iface)
		return
	}
	b, ok = iface.(*float64)
	if !ok {
		t.Error("(m *MutableConfSt) Get() should be able to be casted back to *float64: ", ok)
		return
	}
	if *b != 2 {
		t.Error("(m *MutableConfSt) Get() should return 2 on mutable item: ", *b)
		return
	}

	// non exist
	// test GetFloat64() on non-exist item
	b = m.GetFloat64("Float64n")
	if b != nil {
		t.Error("(m *MutableConfSt) GetFloat64() should return nil on non-exist item, got ", *b)
		return
	}
	// test SetFloat64() on non-exist item
	ok = m.SetFloat64("Float64n", 9)
	if ok {
		t.Error("(m *MutableConfSt) SetFloat64() should return on non-exist item, got ", ok)
		return
	}
	// test Get() on non-exist item
	iface = m.Get("Float64n")
	if iface != nil {
		t.Error("(m *MutableConfSt) Get() should return nil on non-exist item, got ", iface)
		return
	}
	ok = m.Set("Float64n", "9")
	if ok {
		t.Error("(m *MutableConfSt) Set() should return on non-exist item, got ", ok)
		return
	}

	// invalid input
	// test Set() on invalid value
	ok = m.Set("Float64m", "test")
	if ok {
		t.Error("(m *MutableConfSt) Set() should return on invalid value, got ", ok)
		return
	}
	b = m.GetFloat64("Float64m")
	if b == nil {
		t.Error("(m *MutableConfSt) GetFloat64() should not return nil on exist item, got ", b)
		return
	}
	if *b != 2 {
		t.Error("(m *MutableConfSt) GetFloat64() should return 2 after setting invalid value, got ", *b)
		return
	}
	return
}

// test int
func TestInt(t *testing.T) {
	fmt.Printf("Testing int...\n")
	err := initConf()
	if err != nil {
		t.Error(err.Error())
		return
	}
	// immutable
	// test GetInt()
	b := m.GetInt("Inti")
	if b == nil {
		t.Error("(m *MutableConfSt) GetInt() should not return nil on exist item, got ", b)
		return
	}
	if *b != 1 {
		t.Error("(m *MutableConfSt) GetInt() should return 1, got ", *b)
		return
	}
	//test SetInt()
	ok := m.SetInt("Inti", 3)
	if ok {
		t.Error("(m *MutableConfSt) SetInt() should return on immutable item, got ", ok)
		return
	}
	// test GetInt()
	b = m.GetInt("Inti")
	if b == nil {
		t.Error("(m *MutableConfSt) GetInt() should not return nil on exist item, got ", b)
		return
	}
	if *b != 1 {
		t.Error("(m *MutableConfSt) GetInt() should return 1 on immutable item, got ", *b)
		return
	}
	// test Get() on int item
	iface := m.Get("Inti")
	if iface == nil {
		t.Error("(m *MutableConfSt) Get() should not be nil on exist item: ", iface)
		return
	}
	b, ok = iface.(*int)
	if !ok {
		t.Error("(m *MutableConfSt) Get() should be able to be casted back to *int: ", ok)
		return
	}
	if *b != 1 {
		t.Error("(m *MutableConfSt) Get() should return 1 on immutable item: ", *b)
		return
	}
	// test Set() on int
	ok = m.Set("Inti", "3")
	if ok {
		t.Error("(m *MutableConfSt) Set() should return on immutable item, got ", ok)
		return
	}
	// test Get() on int item after Set()
	iface = m.Get("Inti")
	if iface == nil {
		t.Error("(m *MutableConfSt) Get() should not be nil on exist item: ", iface)
		return
	}
	b, ok = iface.(*int)
	if !ok {
		t.Error("(m *MutableConfSt) Get() should be able to be casted back to *int: ", ok)
		return
	}
	if *b != 1 {
		t.Error("(m *MutableConfSt) Get() should return 1 on immutable item: ", *b)
		return
	}

	// mutable
	// test GetInt()
	b = m.GetInt("Intm")
	if b == nil {
		t.Error("(m *MutableConfSt) GetInt() should not return nil on exist item, got ", b)
		return
	}
	if *b != 2 {
		t.Error("(m *MutableConfSt) GetInt() should return 2, got ", *b)
		return
	}
	//test SetInt()
	ok = m.SetInt("Intm", 4)
	if !ok {
		t.Error("(m *MutableConfSt) SetInt() should return on mutable item, got ", ok)
		return
	}
	// test GetInt()
	b = m.GetInt("Intm")
	if b == nil {
		t.Error("(m *MutableConfSt) GetInt() should not return nil on exist item, got ", b)
		return
	}
	if *b != 4 {
		t.Error("(m *MutableConfSt) GetInt() should return 4 on mutable, got ", *b)
		return
	}
	// test Get() on int item
	iface = m.Get("Intm")
	if iface == nil {
		t.Error("(m *MutableConfSt) Get() should not be nil on exist item: ", iface)
		return
	}
	b, ok = iface.(*int)
	if !ok {
		t.Error("(m *MutableConfSt) Get() should be able to be casted back to *int: ", ok)
		return
	}
	if *b != 4 {
		t.Error("(m *MutableConfSt) Get() should return 4 on mutable item: ", *b)
		return
	}
	// test Set() on int
	ok = m.Set("Intm", "2")
	if !ok {
		t.Error("(m *MutableConfSt) Set() should return on mutable item, got ", ok)
		return
	}
	// test Get() on int item after Set()
	iface = m.Get("Intm")
	if iface == nil {
		t.Error("(m *MutableConfSt) Get() should not be nil on exist item: ", iface)
		return
	}
	b, ok = iface.(*int)
	if !ok {
		t.Error("(m *MutableConfSt) Get() should be able to be casted back to *int: ", ok)
		return
	}
	if *b != 2 {
		t.Error("(m *MutableConfSt) Get() should return 2 on mutable item: ", *b)
		return
	}

	// non exist
	// test GetInt() on non-exist item
	b = m.GetInt("Intn")
	if b != nil {
		t.Error("(m *MutableConfSt) GetInt() should return nil on non-exist item, got ", *b)
		return
	}
	// test SetInt() on non-exist item
	ok = m.SetInt("Intn", 9)
	if ok {
		t.Error("(m *MutableConfSt) SetInt() should return on non-exist item, got ", ok)
		return
	}
	// test Get() on non-exist item
	iface = m.Get("Intn")
	if iface != nil {
		t.Error("(m *MutableConfSt) Get() should return nil on non-exist item, got ", iface)
		return
	}
	ok = m.Set("Intn", "9")
	if ok {
		t.Error("(m *MutableConfSt) Set() should return on non-exist item, got ", ok)
		return
	}

	// invalid input
	// test Set() on invalid value
	ok = m.Set("Intm", "test")
	if ok {
		t.Error("(m *MutableConfSt) Set() should return on invalid value, got ", ok)
		return
	}
	b = m.GetInt("Intm")
	if b == nil {
		t.Error("(m *MutableConfSt) GetInt() should not return nil on exist item, got ", b)
		return
	}
	if *b != 2 {
		t.Error("(m *MutableConfSt) GetInt() should return 2 after setting invalid value, got ", *b)
		return
	}
	return
}

// test int8
func TestInt8(t *testing.T) {
	fmt.Printf("Testing int8...\n")
	err := initConf()
	if err != nil {
		t.Error(err.Error())
		return
	}
	// immutable
	// test GetInt8()
	b := m.GetInt8("Int8i")
	if b == nil {
		t.Error("(m *MutableConfSt) GetInt8() should not return nil on exist item, got ", b)
		return
	}
	if *b != 1 {
		t.Error("(m *MutableConfSt) GetInt8() should return 1, got ", *b)
		return
	}
	//test SetInt8()
	ok := m.SetInt8("Int8i", 3)
	if ok {
		t.Error("(m *MutableConfSt) SetInt8() should return on immutable item, got ", ok)
		return
	}
	// test GetInt8()
	b = m.GetInt8("Int8i")
	if b == nil {
		t.Error("(m *MutableConfSt) GetInt8() should not return nil on exist item, got ", b)
		return
	}
	if *b != 1 {
		t.Error("(m *MutableConfSt) GetInt8() should return 1 on immutable item, got ", *b)
		return
	}
	// test Get() on int8 item
	iface := m.Get("Int8i")
	if iface == nil {
		t.Error("(m *MutableConfSt) Get() should not be nil on exist item: ", iface)
		return
	}
	b, ok = iface.(*int8)
	if !ok {
		t.Error("(m *MutableConfSt) Get() should be able to be casted back to *int8: ", ok)
		return
	}
	if *b != 1 {
		t.Error("(m *MutableConfSt) Get() should return 1 on immutable item: ", *b)
		return
	}
	// test Set() on int8
	ok = m.Set("Int8i", "3")
	if ok {
		t.Error("(m *MutableConfSt) Set() should return on immutable item, got ", ok)
		return
	}
	// test Get() on int8 item after Set()
	iface = m.Get("Int8i")
	if iface == nil {
		t.Error("(m *MutableConfSt) Get() should not be nil on exist item: ", iface)
		return
	}
	b, ok = iface.(*int8)
	if !ok {
		t.Error("(m *MutableConfSt) Get() should be able to be casted back to *int8: ", ok)
		return
	}
	if *b != 1 {
		t.Error("(m *MutableConfSt) Get() should return 1 on immutable item: ", *b)
		return
	}

	// mutable
	// test GetInt8()
	b = m.GetInt8("Int8m")
	if b == nil {
		t.Error("(m *MutableConfSt) GetInt8() should not return nil on exist item, got ", b)
		return
	}
	if *b != 2 {
		t.Error("(m *MutableConfSt) GetInt8() should return 2, got ", *b)
		return
	}
	//test SetInt8()
	ok = m.SetInt8("Int8m", 4)
	if !ok {
		t.Error("(m *MutableConfSt) SetInt8() should return on mutable item, got ", ok)
		return
	}
	// test GetInt8()
	b = m.GetInt8("Int8m")
	if b == nil {
		t.Error("(m *MutableConfSt) GetInt8() should not return nil on exist item, got ", b)
		return
	}
	if *b != 4 {
		t.Error("(m *MutableConfSt) GetInt8() should return 4 on mutable, got ", *b)
		return
	}
	// test Get() on int8 item
	iface = m.Get("Int8m")
	if iface == nil {
		t.Error("(m *MutableConfSt) Get() should not be nil on exist item: ", iface)
		return
	}
	b, ok = iface.(*int8)
	if !ok {
		t.Error("(m *MutableConfSt) Get() should be able to be casted back to *int8: ", ok)
		return
	}
	if *b != 4 {
		t.Error("(m *MutableConfSt) Get() should return 4 on mutable item: ", *b)
		return
	}
	// test Set() on int8
	ok = m.Set("Int8m", "2")
	if !ok {
		t.Error("(m *MutableConfSt) Set() should return on mutable item, got ", ok)
		return
	}
	// test Get() on int8 item after Set()
	iface = m.Get("Int8m")
	if iface == nil {
		t.Error("(m *MutableConfSt) Get() should not be nil on exist item: ", iface)
		return
	}
	b, ok = iface.(*int8)
	if !ok {
		t.Error("(m *MutableConfSt) Get() should be able to be casted back to *int8: ", ok)
		return
	}
	if *b != 2 {
		t.Error("(m *MutableConfSt) Get() should return 2 on mutable item: ", *b)
		return
	}

	// non exist
	// test GetInt8() on non-exist item
	b = m.GetInt8("Int8n")
	if b != nil {
		t.Error("(m *MutableConfSt) GetInt8() should return nil on non-exist item, got ", *b)
		return
	}
	// test SetInt8() on non-exist item
	ok = m.SetInt8("Int8n", 9)
	if ok {
		t.Error("(m *MutableConfSt) SetInt8() should return on non-exist item, got ", ok)
		return
	}
	// test Get() on non-exist item
	iface = m.Get("Int8n")
	if iface != nil {
		t.Error("(m *MutableConfSt) Get() should return nil on non-exist item, got ", iface)
		return
	}
	ok = m.Set("Int8n", "9")
	if ok {
		t.Error("(m *MutableConfSt) Set() should return on non-exist item, got ", ok)
		return
	}

	// invalid input
	// test Set() on invalid value
	ok = m.Set("Int8m", "test")
	if ok {
		t.Error("(m *MutableConfSt) Set() should return on invalid value, got ", ok)
		return
	}
	b = m.GetInt8("Int8m")
	if b == nil {
		t.Error("(m *MutableConfSt) GetInt8() should not return nil on exist item, got ", b)
		return
	}
	if *b != 2 {
		t.Error("(m *MutableConfSt) GetInt8() should return 2 after setting invalid value, got ", *b)
		return
	}
	return
}

// test int16
func TestInt16(t *testing.T) {
	fmt.Printf("Testing int16...\n")
	err := initConf()
	if err != nil {
		t.Error(err.Error())
		return
	}
	// immutable
	// test GetInt16()
	b := m.GetInt16("Int16i")
	if b == nil {
		t.Error("(m *MutableConfSt) GetInt16() should not return nil on exist item, got ", b)
		return
	}
	if *b != 1 {
		t.Error("(m *MutableConfSt) GetInt16() should return 1, got ", *b)
		return
	}
	//test SetInt16()
	ok := m.SetInt16("Int16i", 3)
	if ok {
		t.Error("(m *MutableConfSt) SetInt16() should return on immutable item, got ", ok)
		return
	}
	// test GetInt16()
	b = m.GetInt16("Int16i")
	if b == nil {
		t.Error("(m *MutableConfSt) GetInt16() should not return nil on exist item, got ", b)
		return
	}
	if *b != 1 {
		t.Error("(m *MutableConfSt) GetInt16() should return 1 on immutable item, got ", *b)
		return
	}
	// test Get() on int16 item
	iface := m.Get("Int16i")
	if iface == nil {
		t.Error("(m *MutableConfSt) Get() should not be nil on exist item: ", iface)
		return
	}
	b, ok = iface.(*int16)
	if !ok {
		t.Error("(m *MutableConfSt) Get() should be able to be casted back to *int16: ", ok)
		return
	}
	if *b != 1 {
		t.Error("(m *MutableConfSt) Get() should return 1 on immutable item: ", *b)
		return
	}
	// test Set() on int16
	ok = m.Set("Int16i", "3")
	if ok {
		t.Error("(m *MutableConfSt) Set() should return on immutable item, got ", ok)
		return
	}
	// test Get() on int16 item after Set()
	iface = m.Get("Int16i")
	if iface == nil {
		t.Error("(m *MutableConfSt) Get() should not be nil on exist item: ", iface)
		return
	}
	b, ok = iface.(*int16)
	if !ok {
		t.Error("(m *MutableConfSt) Get() should be able to be casted back to *int16: ", ok)
		return
	}
	if *b != 1 {
		t.Error("(m *MutableConfSt) Get() should return 1 on immutable item: ", *b)
		return
	}

	// mutable
	// test GetInt16()
	b = m.GetInt16("Int16m")
	if b == nil {
		t.Error("(m *MutableConfSt) GetInt16() should not return nil on exist item, got ", b)
		return
	}
	if *b != 2 {
		t.Error("(m *MutableConfSt) GetInt16() should return 2, got ", *b)
		return
	}
	//test SetInt16()
	ok = m.SetInt16("Int16m", 4)
	if !ok {
		t.Error("(m *MutableConfSt) SetInt16() should return on mutable item, got ", ok)
		return
	}
	// test GetInt16()
	b = m.GetInt16("Int16m")
	if b == nil {
		t.Error("(m *MutableConfSt) GetInt16() should not return nil on exist item, got ", b)
		return
	}
	if *b != 4 {
		t.Error("(m *MutableConfSt) GetInt16() should return 4 on mutable, got ", *b)
		return
	}
	// test Get() on int16 item
	iface = m.Get("Int16m")
	if iface == nil {
		t.Error("(m *MutableConfSt) Get() should not be nil on exist item: ", iface)
		return
	}
	b, ok = iface.(*int16)
	if !ok {
		t.Error("(m *MutableConfSt) Get() should be able to be casted back to *int16: ", ok)
		return
	}
	if *b != 4 {
		t.Error("(m *MutableConfSt) Get() should return 4 on mutable item: ", *b)
		return
	}
	// test Set() on int16
	ok = m.Set("Int16m", "2")
	if !ok {
		t.Error("(m *MutableConfSt) Set() should return on mutable item, got ", ok)
		return
	}
	// test Get() on int16 item after Set()
	iface = m.Get("Int16m")
	if iface == nil {
		t.Error("(m *MutableConfSt) Get() should not be nil on exist item: ", iface)
		return
	}
	b, ok = iface.(*int16)
	if !ok {
		t.Error("(m *MutableConfSt) Get() should be able to be casted back to *int16: ", ok)
		return
	}
	if *b != 2 {
		t.Error("(m *MutableConfSt) Get() should return 2 on mutable item: ", *b)
		return
	}

	// non exist
	// test GetInt16() on non-exist item
	b = m.GetInt16("Int16n")
	if b != nil {
		t.Error("(m *MutableConfSt) GetInt16() should return nil on non-exist item, got ", *b)
		return
	}
	// test SetInt16() on non-exist item
	ok = m.SetInt16("Int16n", 9)
	if ok {
		t.Error("(m *MutableConfSt) SetInt16() should return on non-exist item, got ", ok)
		return
	}
	// test Get() on non-exist item
	iface = m.Get("Int16n")
	if iface != nil {
		t.Error("(m *MutableConfSt) Get() should return nil on non-exist item, got ", iface)
		return
	}
	ok = m.Set("Int16n", "9")
	if ok {
		t.Error("(m *MutableConfSt) Set() should return on non-exist item, got ", ok)
		return
	}

	// invalid input
	// test Set() on invalid value
	ok = m.Set("Int16m", "test")
	if ok {
		t.Error("(m *MutableConfSt) Set() should return on invalid value, got ", ok)
		return
	}
	b = m.GetInt16("Int16m")
	if b == nil {
		t.Error("(m *MutableConfSt) GetInt16() should not return nil on exist item, got ", b)
		return
	}
	if *b != 2 {
		t.Error("(m *MutableConfSt) GetInt16() should return 2 after setting invalid value, got ", *b)
		return
	}
	return
}

// test int32
func TestInt32(t *testing.T) {
	fmt.Printf("Testing int32...\n")
	err := initConf()
	if err != nil {
		t.Error(err.Error())
		return
	}
	// immutable
	// test GetInt32()
	b := m.GetInt32("Int32i")
	if b == nil {
		t.Error("(m *MutableConfSt) GetInt32() should not return nil on exist item, got ", b)
		return
	}
	if *b != 1 {
		t.Error("(m *MutableConfSt) GetInt32() should return 1, got ", *b)
		return
	}
	//test SetInt32()
	ok := m.SetInt32("Int32i", 3)
	if ok {
		t.Error("(m *MutableConfSt) SetInt32() should return on immutable item, got ", ok)
		return
	}
	// test GetInt32()
	b = m.GetInt32("Int32i")
	if b == nil {
		t.Error("(m *MutableConfSt) GetInt32() should not return nil on exist item, got ", b)
		return
	}
	if *b != 1 {
		t.Error("(m *MutableConfSt) GetInt32() should return 1 on immutable item, got ", *b)
		return
	}
	// test Get() on int32 item
	iface := m.Get("Int32i")
	if iface == nil {
		t.Error("(m *MutableConfSt) Get() should not be nil on exist item: ", iface)
		return
	}
	b, ok = iface.(*int32)
	if !ok {
		t.Error("(m *MutableConfSt) Get() should be able to be casted back to *int32: ", ok)
		return
	}
	if *b != 1 {
		t.Error("(m *MutableConfSt) Get() should return 1 on immutable item: ", *b)
		return
	}
	// test Set() on int32
	ok = m.Set("Int32i", "3")
	if ok {
		t.Error("(m *MutableConfSt) Set() should return on immutable item, got ", ok)
		return
	}
	// test Get() on int32 item after Set()
	iface = m.Get("Int32i")
	if iface == nil {
		t.Error("(m *MutableConfSt) Get() should not be nil on exist item: ", iface)
		return
	}
	b, ok = iface.(*int32)
	if !ok {
		t.Error("(m *MutableConfSt) Get() should be able to be casted back to *int32: ", ok)
		return
	}
	if *b != 1 {
		t.Error("(m *MutableConfSt) Get() should return 1 on immutable item: ", *b)
		return
	}

	// mutable
	// test GetInt32()
	b = m.GetInt32("Int32m")
	if b == nil {
		t.Error("(m *MutableConfSt) GetInt32() should not return nil on exist item, got ", b)
		return
	}
	if *b != 2 {
		t.Error("(m *MutableConfSt) GetInt32() should return 2, got ", *b)
		return
	}
	//test SetInt32()
	ok = m.SetInt32("Int32m", 4)
	if !ok {
		t.Error("(m *MutableConfSt) SetInt32() should return on mutable item, got ", ok)
		return
	}
	// test GetInt32()
	b = m.GetInt32("Int32m")
	if b == nil {
		t.Error("(m *MutableConfSt) GetInt32() should not return nil on exist item, got ", b)
		return
	}
	if *b != 4 {
		t.Error("(m *MutableConfSt) GetInt32() should return 4 on mutable, got ", *b)
		return
	}
	// test Get() on int32 item
	iface = m.Get("Int32m")
	if iface == nil {
		t.Error("(m *MutableConfSt) Get() should not be nil on exist item: ", iface)
		return
	}
	b, ok = iface.(*int32)
	if !ok {
		t.Error("(m *MutableConfSt) Get() should be able to be casted back to *int32: ", ok)
		return
	}
	if *b != 4 {
		t.Error("(m *MutableConfSt) Get() should return 4 on mutable item: ", *b)
		return
	}
	// test Set() on int32
	ok = m.Set("Int32m", "2")
	if !ok {
		t.Error("(m *MutableConfSt) Set() should return on mutable item, got ", ok)
		return
	}
	// test Get() on int32 item after Set()
	iface = m.Get("Int32m")
	if iface == nil {
		t.Error("(m *MutableConfSt) Get() should not be nil on exist item: ", iface)
		return
	}
	b, ok = iface.(*int32)
	if !ok {
		t.Error("(m *MutableConfSt) Get() should be able to be casted back to *int32: ", ok)
		return
	}
	if *b != 2 {
		t.Error("(m *MutableConfSt) Get() should return 2 on mutable item: ", *b)
		return
	}

	// non exist
	// test GetInt32() on non-exist item
	b = m.GetInt32("Int32n")
	if b != nil {
		t.Error("(m *MutableConfSt) GetInt32() should return nil on non-exist item, got ", *b)
		return
	}
	// test SetInt32() on non-exist item
	ok = m.SetInt32("Int32n", 9)
	if ok {
		t.Error("(m *MutableConfSt) SetInt32() should return on non-exist item, got ", ok)
		return
	}
	// test Get() on non-exist item
	iface = m.Get("Int32n")
	if iface != nil {
		t.Error("(m *MutableConfSt) Get() should return nil on non-exist item, got ", iface)
		return
	}
	ok = m.Set("Int32n", "9")
	if ok {
		t.Error("(m *MutableConfSt) Set() should return on non-exist item, got ", ok)
		return
	}

	// invalid input
	// test Set() on invalid value
	ok = m.Set("Int32m", "test")
	if ok {
		t.Error("(m *MutableConfSt) Set() should return on invalid value, got ", ok)
		return
	}
	b = m.GetInt32("Int32m")
	if b == nil {
		t.Error("(m *MutableConfSt) GetInt32() should not return nil on exist item, got ", b)
		return
	}
	if *b != 2 {
		t.Error("(m *MutableConfSt) GetInt32() should return 2 after setting invalid value, got ", *b)
		return
	}
	return
}

// test int64
func TestInt64(t *testing.T) {
	fmt.Printf("Testing int64...\n")
	err := initConf()
	if err != nil {
		t.Error(err.Error())
		return
	}
	// immutable
	// test GetInt64()
	b := m.GetInt64("Int64i")
	if b == nil {
		t.Error("(m *MutableConfSt) GetInt64() should not return nil on exist item, got ", b)
		return
	}
	if *b != 1 {
		t.Error("(m *MutableConfSt) GetInt64() should return 1, got ", *b)
		return
	}
	//test SetInt64()
	ok := m.SetInt64("Int64i", 3)
	if ok {
		t.Error("(m *MutableConfSt) SetInt64() should return on immutable item, got ", ok)
		return
	}
	// test GetInt64()
	b = m.GetInt64("Int64i")
	if b == nil {
		t.Error("(m *MutableConfSt) GetInt64() should not return nil on exist item, got ", b)
		return
	}
	if *b != 1 {
		t.Error("(m *MutableConfSt) GetInt64() should return 1 on immutable item, got ", *b)
		return
	}
	// test Get() on int64 item
	iface := m.Get("Int64i")
	if iface == nil {
		t.Error("(m *MutableConfSt) Get() should not be nil on exist item: ", iface)
		return
	}
	b, ok = iface.(*int64)
	if !ok {
		t.Error("(m *MutableConfSt) Get() should be able to be casted back to *int64: ", ok)
		return
	}
	if *b != 1 {
		t.Error("(m *MutableConfSt) Get() should return 1 on immutable item: ", *b)
		return
	}
	// test Set() on int64
	ok = m.Set("Int64i", "3")
	if ok {
		t.Error("(m *MutableConfSt) Set() should return on immutable item, got ", ok)
		return
	}
	// test Get() on int64 item after Set()
	iface = m.Get("Int64i")
	if iface == nil {
		t.Error("(m *MutableConfSt) Get() should not be nil on exist item: ", iface)
		return
	}
	b, ok = iface.(*int64)
	if !ok {
		t.Error("(m *MutableConfSt) Get() should be able to be casted back to *int64: ", ok)
		return
	}
	if *b != 1 {
		t.Error("(m *MutableConfSt) Get() should return 1 on immutable item: ", *b)
		return
	}

	// mutable
	// test GetInt64()
	b = m.GetInt64("Int64m")
	if b == nil {
		t.Error("(m *MutableConfSt) GetInt64() should not return nil on exist item, got ", b)
		return
	}
	if *b != 2 {
		t.Error("(m *MutableConfSt) GetInt64() should return 2, got ", *b)
		return
	}
	//test SetInt64()
	ok = m.SetInt64("Int64m", 4)
	if !ok {
		t.Error("(m *MutableConfSt) SetInt64() should return on mutable item, got ", ok)
		return
	}
	// test GetInt64()
	b = m.GetInt64("Int64m")
	if b == nil {
		t.Error("(m *MutableConfSt) GetInt64() should not return nil on exist item, got ", b)
		return
	}
	if *b != 4 {
		t.Error("(m *MutableConfSt) GetInt64() should return 4 on mutable, got ", *b)
		return
	}
	// test Get() on int64 item
	iface = m.Get("Int64m")
	if iface == nil {
		t.Error("(m *MutableConfSt) Get() should not be nil on exist item: ", iface)
		return
	}
	b, ok = iface.(*int64)
	if !ok {
		t.Error("(m *MutableConfSt) Get() should be able to be casted back to *int64: ", ok)
		return
	}
	if *b != 4 {
		t.Error("(m *MutableConfSt) Get() should return 4 on mutable item: ", *b)
		return
	}
	// test Set() on int64
	ok = m.Set("Int64m", "2")
	if !ok {
		t.Error("(m *MutableConfSt) Set() should return on mutable item, got ", ok)
		return
	}
	// test Get() on int64 item after Set()
	iface = m.Get("Int64m")
	if iface == nil {
		t.Error("(m *MutableConfSt) Get() should not be nil on exist item: ", iface)
		return
	}
	b, ok = iface.(*int64)
	if !ok {
		t.Error("(m *MutableConfSt) Get() should be able to be casted back to *int64: ", ok)
		return
	}
	if *b != 2 {
		t.Error("(m *MutableConfSt) Get() should return 2 on mutable item: ", *b)
		return
	}

	// non exist
	// test GetInt64() on non-exist item
	b = m.GetInt64("Int64n")
	if b != nil {
		t.Error("(m *MutableConfSt) GetInt64() should return nil on non-exist item, got ", *b)
		return
	}
	// test SetInt64() on non-exist item
	ok = m.SetInt64("Int64n", 9)
	if ok {
		t.Error("(m *MutableConfSt) SetInt64() should return on non-exist item, got ", ok)
		return
	}
	// test Get() on non-exist item
	iface = m.Get("Int64n")
	if iface != nil {
		t.Error("(m *MutableConfSt) Get() should return nil on non-exist item, got ", iface)
		return
	}
	ok = m.Set("Int64n", "9")
	if ok {
		t.Error("(m *MutableConfSt) Set() should return on non-exist item, got ", ok)
		return
	}

	// invalid input
	// test Set() on invalid value
	ok = m.Set("Int64m", "test")
	if ok {
		t.Error("(m *MutableConfSt) Set() should return on invalid value, got ", ok)
		return
	}
	b = m.GetInt64("Int64m")
	if b == nil {
		t.Error("(m *MutableConfSt) GetInt64() should not return nil on exist item, got ", b)
		return
	}
	if *b != 2 {
		t.Error("(m *MutableConfSt) GetInt64() should return 2 after setting invalid value, got ", *b)
		return
	}
	return
}

func TestString(t *testing.T) {
	fmt.Printf("Testing string...\n")
	err := initConf()
	if err != nil {
		t.Error(err.Error())
		return
	}
	// immutable
	// test GetString()
	b := m.GetString("Stringi")
	if b == nil {
		t.Error("(m *MutableConfSt) GetString() should not return nil on exist item, got ", b)
		return
	}
	if *b != "immutable" {
		t.Error("(m *MutableConfSt) GetString() should return immutable, got ", *b)
		return
	}
	//test SetString()
	ok := m.SetString("Stringi", "mutation")
	if ok {
		t.Error("(m *MutableConfSt) SetString() should return false on immutable item, got ", ok)
		return
	}
	// test GetString()
	b = m.GetString("Stringi")
	if b == nil {
		t.Error("(m *MutableConfSt) GetString() should not return nil on exist item, got ", b)
		return
	}
	if *b != "immutable" {
		t.Error("(m *MutableConfSt) GetString() should return immutable on immutable item, got ", *b)
		return
	}
	// test Get() on string item
	iface := m.Get("Stringi")
	if iface == nil {
		t.Error("(m *MutableConfSt) Get() should not be nil on exist item: ", iface)
		return
	}
	b, ok = iface.(*string)
	if !ok {
		t.Error("(m *MutableConfSt) Get() should be able to be casted back to *string: ", ok)
		return
	}
	if *b != "immutable" {
		t.Error("(m *MutableConfSt) Get() should return immutable on immutable item: ", *b)
		return
	}
	// test Set() on string
	ok = m.Set("Stringi", "mutation")
	if ok {
		t.Error("(m *MutableConfSt) Set() should return false on immutable item, got ", ok)
		return
	}
	// test Get() on string item after Set()
	iface = m.Get("Stringi")
	if iface == nil {
		t.Error("(m *MutableConfSt) Get() should not be nil on exist item: ", iface)
		return
	}
	b, ok = iface.(*string)
	if !ok {
		t.Error("(m *MutableConfSt) Get() should be able to be casted back to *string: ", ok)
		return
	}
	if *b != "immutable" {
		t.Error("(m *MutableConfSt) Get() should return immutable on immutable item: ", *b)
		return
	}

	// mutable
	// test GetString()
	b = m.GetString("Stringm")
	if b == nil {
		t.Error("(m *MutableConfSt) GetString() should not return nil on exist item, got ", b)
		return
	}
	if *b != "mutable" {
		t.Error("(m *MutableConfSt) GetString() should return mutable, got ", *b)
		return
	}
	//test SetString()
	ok = m.SetString("Stringm", "mutation")
	if !ok {
		t.Error("(m *MutableConfSt) SetString() should return true on mutable item, got ", ok)
		return
	}
	// test GetString()
	b = m.GetString("Stringm")
	if b == nil {
		t.Error("(m *MutableConfSt) GetString() should not return nil on exist item, got ", b)
		return
	}
	if *b != "mutation" {
		t.Error("(m *MutableConfSt) GetString() should return mutation on mutable, got ", *b)
		return
	}
	// test Get() on string item
	iface = m.Get("Stringm")
	if iface == nil {
		t.Error("(m *MutableConfSt) Get() should not be nil on exist item: ", iface)
		return
	}
	b, ok = iface.(*string)
	if !ok {
		t.Error("(m *MutableConfSt) Get() should be able to be casted back to *string: ", ok)
		return
	}
	if *b != "mutation" {
		t.Error("(m *MutableConfSt) Get() should return 4 on mutable item: ", *b)
		return
	}
	// test Set() on string
	ok = m.Set("Stringm", "mutable")
	if !ok {
		t.Error("(m *MutableConfSt) Set() should return on mutable item, got ", ok)
		return
	}
	// test Get() on string item after Set()
	iface = m.Get("Stringm")
	if iface == nil {
		t.Error("(m *MutableConfSt) Get() should not be nil on exist item: ", iface)
		return
	}
	b, ok = iface.(*string)
	if !ok {
		t.Error("(m *MutableConfSt) Get() should be able to be casted back to *string: ", ok)
		return
	}
	if *b != "mutable" {
		t.Error("(m *MutableConfSt) Get() should return mutable on mutable item: ", *b)
		return
	}

	// non exist
	// test GetString() on non-exist item
	b = m.GetString("Stringn")
	if b != nil {
		t.Error("(m *MutableConfSt) GetString() should return nil on non-exist item, got ", *b)
		return
	}
	// test SetString() on non-exist item
	ok = m.SetString("Stringn", "nonexist")
	if ok {
		t.Error("(m *MutableConfSt) SetString() should return on non-exist item, got ", ok)
		return
	}
	// test Get() on non-exist item
	iface = m.Get("Stringn")
	if iface != nil {
		t.Error("(m *MutableConfSt) Get() should return nil on non-exist item, got ", iface)
		return
	}
	ok = m.Set("Stringn", "nonexist")
	if ok {
		t.Error("(m *MutableConfSt) Set() should return on non-exist item, got ", ok)
		return
	}
	return
}

func TestUint(t *testing.T) {
	fmt.Printf("Testing uint...\n")
	err := initConf()
	if err != nil {
		t.Error(err.Error())
		return
	}
	// immutable
	// test GetUint()
	b := m.GetUint("Uinti")
	if b == nil {
		t.Error("(m *MutableConfSt) GetUint() should not return nil on exist item, got ", b)
		return
	}
	if *b != 1 {
		t.Error("(m *MutableConfSt) GetUint() should return 1, got ", *b)
		return
	}
	//test SetUint()
	ok := m.SetUint("Uinti", 3)
	if ok {
		t.Error("(m *MutableConfSt) SetUint() should return on immutable item, got ", ok)
		return
	}
	// test GetUint()
	b = m.GetUint("Uinti")
	if b == nil {
		t.Error("(m *MutableConfSt) GetUint() should not return nil on exist item, got ", b)
		return
	}
	if *b != 1 {
		t.Error("(m *MutableConfSt) GetUint() should return 1 on immutable item, got ", *b)
		return
	}
	// test Get() on uint item
	iface := m.Get("Uinti")
	if iface == nil {
		t.Error("(m *MutableConfSt) Get() should not be nil on exist item: ", iface)
		return
	}
	b, ok = iface.(*uint)
	if !ok {
		t.Error("(m *MutableConfSt) Get() should be able to be casted back to *uint: ", ok)
		return
	}
	if *b != 1 {
		t.Error("(m *MutableConfSt) Get() should return 1 on immutable item: ", *b)
		return
	}
	// test Set() on uint
	ok = m.Set("Uinti", "3")
	if ok {
		t.Error("(m *MutableConfSt) Set() should return on immutable item, got ", ok)
		return
	}
	// test Get() on uint item after Set()
	iface = m.Get("Uinti")
	if iface == nil {
		t.Error("(m *MutableConfSt) Get() should not be nil on exist item: ", iface)
		return
	}
	b, ok = iface.(*uint)
	if !ok {
		t.Error("(m *MutableConfSt) Get() should be able to be casted back to *uint: ", ok)
		return
	}
	if *b != 1 {
		t.Error("(m *MutableConfSt) Get() should return 1 on immutable item: ", *b)
		return
	}

	// mutable
	// test GetUint()
	b = m.GetUint("Uintm")
	if b == nil {
		t.Error("(m *MutableConfSt) GetUint() should not return nil on exist item, got ", b)
		return
	}
	if *b != 2 {
		t.Error("(m *MutableConfSt) GetUint() should return 2, got ", *b)
		return
	}
	//test SetUint()
	ok = m.SetUint("Uintm", 4)
	if !ok {
		t.Error("(m *MutableConfSt) SetUint() should return on mutable item, got ", ok)
		return
	}
	// test GetUint()
	b = m.GetUint("Uintm")
	if b == nil {
		t.Error("(m *MutableConfSt) GetUint() should not return nil on exist item, got ", b)
		return
	}
	if *b != 4 {
		t.Error("(m *MutableConfSt) GetUint() should return 4 on mutable, got ", *b)
		return
	}
	// test Get() on uint item
	iface = m.Get("Uintm")
	if iface == nil {
		t.Error("(m *MutableConfSt) Get() should not be nil on exist item: ", iface)
		return
	}
	b, ok = iface.(*uint)
	if !ok {
		t.Error("(m *MutableConfSt) Get() should be able to be casted back to *uint: ", ok)
		return
	}
	if *b != 4 {
		t.Error("(m *MutableConfSt) Get() should return 4 on mutable item: ", *b)
		return
	}
	// test Set() on uint
	ok = m.Set("Uintm", "2")
	if !ok {
		t.Error("(m *MutableConfSt) Set() should return on mutable item, got ", ok)
		return
	}
	// test Get() on uint item after Set()
	iface = m.Get("Uintm")
	if iface == nil {
		t.Error("(m *MutableConfSt) Get() should not be nil on exist item: ", iface)
		return
	}
	b, ok = iface.(*uint)
	if !ok {
		t.Error("(m *MutableConfSt) Get() should be able to be casted back to *uint: ", ok)
		return
	}
	if *b != 2 {
		t.Error("(m *MutableConfSt) Get() should return 2 on mutable item: ", *b)
		return
	}

	// non exist
	// test GetUint() on non-exist item
	b = m.GetUint("Uintn")
	if b != nil {
		t.Error("(m *MutableConfSt) GetUint() should return nil on non-exist item, got ", *b)
		return
	}
	// test SetUint() on non-exist item
	ok = m.SetUint("Uintn", 9)
	if ok {
		t.Error("(m *MutableConfSt) SetUint() should return on non-exist item, got ", ok)
		return
	}
	// test Get() on non-exist item
	iface = m.Get("Uintn")
	if iface != nil {
		t.Error("(m *MutableConfSt) Get() should return nil on non-exist item, got ", iface)
		return
	}
	ok = m.Set("Uintn", "9")
	if ok {
		t.Error("(m *MutableConfSt) Set() should return on non-exist item, got ", ok)
		return
	}

	// invalid input
	// test Set() on invalid value
	ok = m.Set("Uintm", "test")
	if ok {
		t.Error("(m *MutableConfSt) Set() should return on invalid value, got ", ok)
		return
	}
	b = m.GetUint("Uintm")
	if b == nil {
		t.Error("(m *MutableConfSt) GetUint() should not return nil on exist item, got ", b)
		return
	}
	if *b != 2 {
		t.Error("(m *MutableConfSt) GetUint() should return 2 after setting invalid value, got ", *b)
		return
	}
	return
}

// test uint8
func TestUint8(t *testing.T) {
	fmt.Printf("Testing uint8...\n")
	err := initConf()
	if err != nil {
		t.Error(err.Error())
		return
	}
	// immutable
	// test GetUint8()
	b := m.GetUint8("Uint8i")
	if b == nil {
		t.Error("(m *MutableConfSt) GetUint8() should not return nil on exist item, got ", b)
		return
	}
	if *b != 1 {
		t.Error("(m *MutableConfSt) GetUint8() should return 1, got ", *b)
		return
	}
	//test SetUint8()
	ok := m.SetUint8("Uint8i", 3)
	if ok {
		t.Error("(m *MutableConfSt) SetUint8() should return on immutable item, got ", ok)
		return
	}
	// test GetUint8()
	b = m.GetUint8("Uint8i")
	if b == nil {
		t.Error("(m *MutableConfSt) GetUint8() should not return nil on exist item, got ", b)
		return
	}
	if *b != 1 {
		t.Error("(m *MutableConfSt) GetUint8() should return 1 on immutable item, got ", *b)
		return
	}
	// test Get() on uint8 item
	iface := m.Get("Uint8i")
	if iface == nil {
		t.Error("(m *MutableConfSt) Get() should not be nil on exist item: ", iface)
		return
	}
	b, ok = iface.(*uint8)
	if !ok {
		t.Error("(m *MutableConfSt) Get() should be able to be casted back to *uint8: ", ok)
		return
	}
	if *b != 1 {
		t.Error("(m *MutableConfSt) Get() should return 1 on immutable item: ", *b)
		return
	}
	// test Set() on uint8
	ok = m.Set("Uint8i", "3")
	if ok {
		t.Error("(m *MutableConfSt) Set() should return on immutable item, got ", ok)
		return
	}
	// test Get() on uint8 item after Set()
	iface = m.Get("Uint8i")
	if iface == nil {
		t.Error("(m *MutableConfSt) Get() should not be nil on exist item: ", iface)
		return
	}
	b, ok = iface.(*uint8)
	if !ok {
		t.Error("(m *MutableConfSt) Get() should be able to be casted back to *uint8: ", ok)
		return
	}
	if *b != 1 {
		t.Error("(m *MutableConfSt) Get() should return 1 on immutable item: ", *b)
		return
	}

	// mutable
	// test GetUint8()
	b = m.GetUint8("Uint8m")
	if b == nil {
		t.Error("(m *MutableConfSt) GetUint8() should not return nil on exist item, got ", b)
		return
	}
	if *b != 2 {
		t.Error("(m *MutableConfSt) GetUint8() should return 2, got ", *b)
		return
	}
	//test SetUint8()
	ok = m.SetUint8("Uint8m", 4)
	if !ok {
		t.Error("(m *MutableConfSt) SetUint8() should return on mutable item, got ", ok)
		return
	}
	// test GetUint8()
	b = m.GetUint8("Uint8m")
	if b == nil {
		t.Error("(m *MutableConfSt) GetUint8() should not return nil on exist item, got ", b)
		return
	}
	if *b != 4 {
		t.Error("(m *MutableConfSt) GetUint8() should return 4 on mutable, got ", *b)
		return
	}
	// test Get() on uint8 item
	iface = m.Get("Uint8m")
	if iface == nil {
		t.Error("(m *MutableConfSt) Get() should not be nil on exist item: ", iface)
		return
	}
	b, ok = iface.(*uint8)
	if !ok {
		t.Error("(m *MutableConfSt) Get() should be able to be casted back to *uint8: ", ok)
		return
	}
	if *b != 4 {
		t.Error("(m *MutableConfSt) Get() should return 4 on mutable item: ", *b)
		return
	}
	// test Set() on uint8
	ok = m.Set("Uint8m", "2")
	if !ok {
		t.Error("(m *MutableConfSt) Set() should return on mutable item, got ", ok)
		return
	}
	// test Get() on uint8 item after Set()
	iface = m.Get("Uint8m")
	if iface == nil {
		t.Error("(m *MutableConfSt) Get() should not be nil on exist item: ", iface)
		return
	}
	b, ok = iface.(*uint8)
	if !ok {
		t.Error("(m *MutableConfSt) Get() should be able to be casted back to *uint8: ", ok)
		return
	}
	if *b != 2 {
		t.Error("(m *MutableConfSt) Get() should return 2 on mutable item: ", *b)
		return
	}

	// non exist
	// test GetUint8() on non-exist item
	b = m.GetUint8("Uint8n")
	if b != nil {
		t.Error("(m *MutableConfSt) GetUint8() should return nil on non-exist item, got ", *b)
		return
	}
	// test SetUint8() on non-exist item
	ok = m.SetUint8("Uint8n", 9)
	if ok {
		t.Error("(m *MutableConfSt) SetUint8() should return on non-exist item, got ", ok)
		return
	}
	// test Get() on non-exist item
	iface = m.Get("Uint8n")
	if iface != nil {
		t.Error("(m *MutableConfSt) Get() should return nil on non-exist item, got ", iface)
		return
	}
	ok = m.Set("Uint8n", "9")
	if ok {
		t.Error("(m *MutableConfSt) Set() should return on non-exist item, got ", ok)
		return
	}

	// invalid input
	// test Set() on invalid value
	ok = m.Set("Uint8m", "test")
	if ok {
		t.Error("(m *MutableConfSt) Set() should return on invalid value, got ", ok)
		return
	}
	b = m.GetUint8("Uint8m")
	if b == nil {
		t.Error("(m *MutableConfSt) GetUint8() should not return nil on exist item, got ", b)
		return
	}
	if *b != 2 {
		t.Error("(m *MutableConfSt) GetUint8() should return 2 after setting invalid value, got ", *b)
		return
	}
	return
}

// test uint16
func TestUint16(t *testing.T) {
	fmt.Printf("Testing uint16...\n")
	err := initConf()
	if err != nil {
		t.Error(err.Error())
		return
	}
	// immutable
	// test GetUint16()
	b := m.GetUint16("Uint16i")
	if b == nil {
		t.Error("(m *MutableConfSt) GetUint16() should not return nil on exist item, got ", b)
		return
	}
	if *b != 1 {
		t.Error("(m *MutableConfSt) GetUint16() should return 1, got ", *b)
		return
	}
	//test SetUint16()
	ok := m.SetUint16("Uint16i", 3)
	if ok {
		t.Error("(m *MutableConfSt) SetUint16() should return on immutable item, got ", ok)
		return
	}
	// test GetUint16()
	b = m.GetUint16("Uint16i")
	if b == nil {
		t.Error("(m *MutableConfSt) GetUint16() should not return nil on exist item, got ", b)
		return
	}
	if *b != 1 {
		t.Error("(m *MutableConfSt) GetUint16() should return 1 on immutable item, got ", *b)
		return
	}
	// test Get() on uint16 item
	iface := m.Get("Uint16i")
	if iface == nil {
		t.Error("(m *MutableConfSt) Get() should not be nil on exist item: ", iface)
		return
	}
	b, ok = iface.(*uint16)
	if !ok {
		t.Error("(m *MutableConfSt) Get() should be able to be casted back to *uint16: ", ok)
		return
	}
	if *b != 1 {
		t.Error("(m *MutableConfSt) Get() should return 1 on immutable item: ", *b)
		return
	}
	// test Set() on uint16
	ok = m.Set("Uint16i", "3")
	if ok {
		t.Error("(m *MutableConfSt) Set() should return on immutable item, got ", ok)
		return
	}
	// test Get() on uint16 item after Set()
	iface = m.Get("Uint16i")
	if iface == nil {
		t.Error("(m *MutableConfSt) Get() should not be nil on exist item: ", iface)
		return
	}
	b, ok = iface.(*uint16)
	if !ok {
		t.Error("(m *MutableConfSt) Get() should be able to be casted back to *uint16: ", ok)
		return
	}
	if *b != 1 {
		t.Error("(m *MutableConfSt) Get() should return 1 on immutable item: ", *b)
		return
	}

	// mutable
	// test GetUint16()
	b = m.GetUint16("Uint16m")
	if b == nil {
		t.Error("(m *MutableConfSt) GetUint16() should not return nil on exist item, got ", b)
		return
	}
	if *b != 2 {
		t.Error("(m *MutableConfSt) GetUint16() should return 2, got ", *b)
		return
	}
	//test SetUint16()
	ok = m.SetUint16("Uint16m", 4)
	if !ok {
		t.Error("(m *MutableConfSt) SetUint16() should return on mutable item, got ", ok)
		return
	}
	// test GetUint16()
	b = m.GetUint16("Uint16m")
	if b == nil {
		t.Error("(m *MutableConfSt) GetUint16() should not return nil on exist item, got ", b)
		return
	}
	if *b != 4 {
		t.Error("(m *MutableConfSt) GetUint16() should return 4 on mutable, got ", *b)
		return
	}
	// test Get() on uint16 item
	iface = m.Get("Uint16m")
	if iface == nil {
		t.Error("(m *MutableConfSt) Get() should not be nil on exist item: ", iface)
		return
	}
	b, ok = iface.(*uint16)
	if !ok {
		t.Error("(m *MutableConfSt) Get() should be able to be casted back to *uint16: ", ok)
		return
	}
	if *b != 4 {
		t.Error("(m *MutableConfSt) Get() should return 4 on mutable item: ", *b)
		return
	}
	// test Set() on uint16
	ok = m.Set("Uint16m", "2")
	if !ok {
		t.Error("(m *MutableConfSt) Set() should return on mutable item, got ", ok)
		return
	}
	// test Get() on uint16 item after Set()
	iface = m.Get("Uint16m")
	if iface == nil {
		t.Error("(m *MutableConfSt) Get() should not be nil on exist item: ", iface)
		return
	}
	b, ok = iface.(*uint16)
	if !ok {
		t.Error("(m *MutableConfSt) Get() should be able to be casted back to *uint16: ", ok)
		return
	}
	if *b != 2 {
		t.Error("(m *MutableConfSt) Get() should return 2 on mutable item: ", *b)
		return
	}

	// non exist
	// test GetUint16() on non-exist item
	b = m.GetUint16("Uint16n")
	if b != nil {
		t.Error("(m *MutableConfSt) GetUint16() should return nil on non-exist item, got ", *b)
		return
	}
	// test SetUint16() on non-exist item
	ok = m.SetUint16("Uint16n", 9)
	if ok {
		t.Error("(m *MutableConfSt) SetUint16() should return on non-exist item, got ", ok)
		return
	}
	// test Get() on non-exist item
	iface = m.Get("Uint16n")
	if iface != nil {
		t.Error("(m *MutableConfSt) Get() should return nil on non-exist item, got ", iface)
		return
	}
	ok = m.Set("Uint16n", "9")
	if ok {
		t.Error("(m *MutableConfSt) Set() should return on non-exist item, got ", ok)
		return
	}

	// invalid input
	// test Set() on invalid value
	ok = m.Set("Uint16m", "test")
	if ok {
		t.Error("(m *MutableConfSt) Set() should return on invalid value, got ", ok)
		return
	}
	b = m.GetUint16("Uint16m")
	if b == nil {
		t.Error("(m *MutableConfSt) GetUint16() should not return nil on exist item, got ", b)
		return
	}
	if *b != 2 {
		t.Error("(m *MutableConfSt) GetUint16() should return 2 after setting invalid value, got ", *b)
		return
	}
	return
}

// test uint32
func TestUint32(t *testing.T) {
	fmt.Printf("Testing uint32...\n")
	err := initConf()
	if err != nil {
		t.Error(err.Error())
		return
	}
	// immutable
	// test GetUint32()
	b := m.GetUint32("Uint32i")
	if b == nil {
		t.Error("(m *MutableConfSt) GetUint32() should not return nil on exist item, got ", b)
		return
	}
	if *b != 1 {
		t.Error("(m *MutableConfSt) GetUint32() should return 1, got ", *b)
		return
	}
	//test SetUint32()
	ok := m.SetUint32("Uint32i", 3)
	if ok {
		t.Error("(m *MutableConfSt) SetUint32() should return on immutable item, got ", ok)
		return
	}
	// test GetUint32()
	b = m.GetUint32("Uint32i")
	if b == nil {
		t.Error("(m *MutableConfSt) GetUint32() should not return nil on exist item, got ", b)
		return
	}
	if *b != 1 {
		t.Error("(m *MutableConfSt) GetUint32() should return 1 on immutable item, got ", *b)
		return
	}
	// test Get() on uint32 item
	iface := m.Get("Uint32i")
	if iface == nil {
		t.Error("(m *MutableConfSt) Get() should not be nil on exist item: ", iface)
		return
	}
	b, ok = iface.(*uint32)
	if !ok {
		t.Error("(m *MutableConfSt) Get() should be able to be casted back to *uint32: ", ok)
		return
	}
	if *b != 1 {
		t.Error("(m *MutableConfSt) Get() should return 1 on immutable item: ", *b)
		return
	}
	// test Set() on uint32
	ok = m.Set("Uint32i", "3")
	if ok {
		t.Error("(m *MutableConfSt) Set() should return on immutable item, got ", ok)
		return
	}
	// test Get() on uint32 item after Set()
	iface = m.Get("Uint32i")
	if iface == nil {
		t.Error("(m *MutableConfSt) Get() should not be nil on exist item: ", iface)
		return
	}
	b, ok = iface.(*uint32)
	if !ok {
		t.Error("(m *MutableConfSt) Get() should be able to be casted back to *uint32: ", ok)
		return
	}
	if *b != 1 {
		t.Error("(m *MutableConfSt) Get() should return 1 on immutable item: ", *b)
		return
	}

	// mutable
	// test GetUint32()
	b = m.GetUint32("Uint32m")
	if b == nil {
		t.Error("(m *MutableConfSt) GetUint32() should not return nil on exist item, got ", b)
		return
	}
	if *b != 2 {
		t.Error("(m *MutableConfSt) GetUint32() should return 2, got ", *b)
		return
	}
	//test SetUint32()
	ok = m.SetUint32("Uint32m", 4)
	if !ok {
		t.Error("(m *MutableConfSt) SetUint32() should return on mutable item, got ", ok)
		return
	}
	// test GetUint32()
	b = m.GetUint32("Uint32m")
	if b == nil {
		t.Error("(m *MutableConfSt) GetUint32() should not return nil on exist item, got ", b)
		return
	}
	if *b != 4 {
		t.Error("(m *MutableConfSt) GetUint32() should return 4 on mutable, got ", *b)
		return
	}
	// test Get() on uint32 item
	iface = m.Get("Uint32m")
	if iface == nil {
		t.Error("(m *MutableConfSt) Get() should not be nil on exist item: ", iface)
		return
	}
	b, ok = iface.(*uint32)
	if !ok {
		t.Error("(m *MutableConfSt) Get() should be able to be casted back to *uint32: ", ok)
		return
	}
	if *b != 4 {
		t.Error("(m *MutableConfSt) Get() should return 4 on mutable item: ", *b)
		return
	}
	// test Set() on uint32
	ok = m.Set("Uint32m", "2")
	if !ok {
		t.Error("(m *MutableConfSt) Set() should return on mutable item, got ", ok)
		return
	}
	// test Get() on uint32 item after Set()
	iface = m.Get("Uint32m")
	if iface == nil {
		t.Error("(m *MutableConfSt) Get() should not be nil on exist item: ", iface)
		return
	}
	b, ok = iface.(*uint32)
	if !ok {
		t.Error("(m *MutableConfSt) Get() should be able to be casted back to *uint32: ", ok)
		return
	}
	if *b != 2 {
		t.Error("(m *MutableConfSt) Get() should return 2 on mutable item: ", *b)
		return
	}

	// non exist
	// test GetUint32() on non-exist item
	b = m.GetUint32("Uint32n")
	if b != nil {
		t.Error("(m *MutableConfSt) GetUint32() should return nil on non-exist item, got ", *b)
		return
	}
	// test SetUint32() on non-exist item
	ok = m.SetUint32("Uint32n", 9)
	if ok {
		t.Error("(m *MutableConfSt) SetUint32() should return on non-exist item, got ", ok)
		return
	}
	// test Get() on non-exist item
	iface = m.Get("Uint32n")
	if iface != nil {
		t.Error("(m *MutableConfSt) Get() should return nil on non-exist item, got ", iface)
		return
	}
	ok = m.Set("Uint32n", "9")
	if ok {
		t.Error("(m *MutableConfSt) Set() should return on non-exist item, got ", ok)
		return
	}

	// invalid input
	// test Set() on invalid value
	ok = m.Set("Uint32m", "test")
	if ok {
		t.Error("(m *MutableConfSt) Set() should return on invalid value, got ", ok)
		return
	}
	b = m.GetUint32("Uint32m")
	if b == nil {
		t.Error("(m *MutableConfSt) GetUint32() should not return nil on exist item, got ", b)
		return
	}
	if *b != 2 {
		t.Error("(m *MutableConfSt) GetUint32() should return 2 after setting invalid value, got ", *b)
		return
	}
	return
}

// test uint64
func TestUint64(t *testing.T) {
	fmt.Printf("Testing uint64...\n")
	err := initConf()
	if err != nil {
		t.Error(err.Error())
		return
	}
	// immutable
	// test GetUint64()
	b := m.GetUint64("Uint64i")
	if b == nil {
		t.Error("(m *MutableConfSt) GetUint64() should not return nil on exist item, got ", b)
		return
	}
	if *b != 1 {
		t.Error("(m *MutableConfSt) GetUint64() should return 1, got ", *b)
		return
	}
	//test SetUint64()
	ok := m.SetUint64("Uint64i", 3)
	if ok {
		t.Error("(m *MutableConfSt) SetUint64() should return on immutable item, got ", ok)
		return
	}
	// test GetUint64()
	b = m.GetUint64("Uint64i")
	if b == nil {
		t.Error("(m *MutableConfSt) GetUint64() should not return nil on exist item, got ", b)
		return
	}
	if *b != 1 {
		t.Error("(m *MutableConfSt) GetUint64() should return 1 on immutable item, got ", *b)
		return
	}
	// test Get() on uint64 item
	iface := m.Get("Uint64i")
	if iface == nil {
		t.Error("(m *MutableConfSt) Get() should not be nil on exist item: ", iface)
		return
	}
	b, ok = iface.(*uint64)
	if !ok {
		t.Error("(m *MutableConfSt) Get() should be able to be casted back to *uint64: ", ok)
		return
	}
	if *b != 1 {
		t.Error("(m *MutableConfSt) Get() should return 1 on immutable item: ", *b)
		return
	}
	// test Set() on uint64
	ok = m.Set("Uint64i", "3")
	if ok {
		t.Error("(m *MutableConfSt) Set() should return on immutable item, got ", ok)
		return
	}
	// test Get() on uint64 item after Set()
	iface = m.Get("Uint64i")
	if iface == nil {
		t.Error("(m *MutableConfSt) Get() should not be nil on exist item: ", iface)
		return
	}
	b, ok = iface.(*uint64)
	if !ok {
		t.Error("(m *MutableConfSt) Get() should be able to be casted back to *uint64: ", ok)
		return
	}
	if *b != 1 {
		t.Error("(m *MutableConfSt) Get() should return 1 on immutable item: ", *b)
		return
	}

	// mutable
	// test GetUint64()
	b = m.GetUint64("Uint64m")
	if b == nil {
		t.Error("(m *MutableConfSt) GetUint64() should not return nil on exist item, got ", b)
		return
	}
	if *b != 2 {
		t.Error("(m *MutableConfSt) GetUint64() should return 2, got ", *b)
		return
	}
	//test SetUint64()
	ok = m.SetUint64("Uint64m", 4)
	if !ok {
		t.Error("(m *MutableConfSt) SetUint64() should return on mutable item, got ", ok)
		return
	}
	// test GetUint64()
	b = m.GetUint64("Uint64m")
	if b == nil {
		t.Error("(m *MutableConfSt) GetUint64() should not return nil on exist item, got ", b)
		return
	}
	if *b != 4 {
		t.Error("(m *MutableConfSt) GetUint64() should return 4 on mutable, got ", *b)
		return
	}
	// test Get() on uint64 item
	iface = m.Get("Uint64m")
	if iface == nil {
		t.Error("(m *MutableConfSt) Get() should not be nil on exist item: ", iface)
		return
	}
	b, ok = iface.(*uint64)
	if !ok {
		t.Error("(m *MutableConfSt) Get() should be able to be casted back to *uint64: ", ok)
		return
	}
	if *b != 4 {
		t.Error("(m *MutableConfSt) Get() should return 4 on mutable item: ", *b)
		return
	}
	// test Set() on uint64
	ok = m.Set("Uint64m", "2")
	if !ok {
		t.Error("(m *MutableConfSt) Set() should return on mutable item, got ", ok)
		return
	}
	// test Get() on uint64 item after Set()
	iface = m.Get("Uint64m")
	if iface == nil {
		t.Error("(m *MutableConfSt) Get() should not be nil on exist item: ", iface)
		return
	}
	b, ok = iface.(*uint64)
	if !ok {
		t.Error("(m *MutableConfSt) Get() should be able to be casted back to *uint64: ", ok)
		return
	}
	if *b != 2 {
		t.Error("(m *MutableConfSt) Get() should return 2 on mutable item: ", *b)
		return
	}

	// non exist
	// test GetUint64() on non-exist item
	b = m.GetUint64("Uint64n")
	if b != nil {
		t.Error("(m *MutableConfSt) GetUint64() should return nil on non-exist item, got ", *b)
		return
	}
	// test SetUint64() on non-exist item
	ok = m.SetUint64("Uint64n", 9)
	if ok {
		t.Error("(m *MutableConfSt) SetUint64() should return on non-exist item, got ", ok)
		return
	}
	// test Get() on non-exist item
	iface = m.Get("Uint64n")
	if iface != nil {
		t.Error("(m *MutableConfSt) Get() should return nil on non-exist item, got ", iface)
		return
	}
	ok = m.Set("Uint64n", "9")
	if ok {
		t.Error("(m *MutableConfSt) Set() should return on non-exist item, got ", ok)
		return
	}

	// invalid input
	// test Set() on invalid value
	ok = m.Set("Uint64m", "test")
	if ok {
		t.Error("(m *MutableConfSt) Set() should return on invalid value, got ", ok)
		return
	}
	b = m.GetUint64("Uint64m")
	if b == nil {
		t.Error("(m *MutableConfSt) GetUint64() should not return nil on exist item, got ", b)
		return
	}
	if *b != 2 {
		t.Error("(m *MutableConfSt) GetUint64() should return 2 after setting invalid value, got ", *b)
		return
	}
	return
}

// prepare a dummy config file
func createDummyConf() error {
	created = false
	fp, err := os.OpenFile(dummyPath, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0644)
	defer fp.Close()
	if err != nil {
		return err
	}
	created = true
	fmt.Fprintf(fp, "%s", dummyConf)
	return nil
}

func getLock() bool {
	pass := false
	for i := 0; i < dummyTry; i++ {
		_, err := os.Stat(dummyLock)
		if err != nil {
			pass = true
			var fp *os.File
			fp, err = os.Create(dummyLock)
			defer fp.Close()
			break
		}
		if err != nil {
			fmt.Printf("failed to create file lock: %v\n", err)
			return false
		}
	}
	if !pass {
		fmt.Printf("failed to get file lock, please try later\n")
		return false
	}
	return true
}

// release file lock
func releaseLock() {
	os.Remove(dummyLock)
}

// clean up dummy config file
func cleanup() {
	os.Remove(dummyPath)
}

// initialize conf if not initialized
func initConf() error {
	if m == nil || !m.Init() {
		lock := getLock()
		if !lock {
			return errors.New("Failed to get lock")
		}
		err := createDummyConf()
		if err != nil {
			return errors.New(fmt.Sprintf("Failed to create dummy conf: %v", err))
		}
		defer func() {
			cleanup()
			releaseLock()
		}()
		c := &dummyConfSt{}
		mutable := make(map[string]bool)
		mutable["Boolm"] = true
		mutable["Float32m"] = true
		mutable["Float64m"] = true
		mutable["Intm"] = true
		mutable["Int8m"] = true
		mutable["Int16m"] = true
		mutable["Int32m"] = true
		mutable["Int64m"] = true
		mutable["Stringm"] = true
		mutable["Uintm"] = true
		mutable["Uint8m"] = true
		mutable["Uint16m"] = true
		mutable["Uint32m"] = true
		mutable["Uint64m"] = true

		// test initilzation
		m, err = InitJSONMutableConf(dummyPath, c, &mutable)
		if err != nil {
			return errors.New(fmt.Sprintf("Failed to initialize conf: %v", err))
		}
	}
	return nil
}
