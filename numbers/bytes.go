package numbers;

import "encoding/binary";

func B2Uint16L(output *uint16, input []byte) (error) {
  if len(input) < 2 {
    return ErrInputBufferLength;
  }
  var i int
  var tmp uint16
  *output = uint16(0)
  for i = 0; i < 2; i++{
    tmp = uint16(input[1 - i]);
    *output = *output << 8
    *output = *output + tmp
  }
  return nil;
}

// tannhauser, transform an uint16 into a byte[] of length 2
func Ui16ToBL(output []byte, input uint16) (error) {
  if len(output) < 2 {
    return ErrOutputBufferLength;
  }
  binary.LittleEndian.PutUint16(output, input)
  return nil;
}

// tannhauser, transform a byte[] of length 4 into a uint32
func B2Uint32L(output *uint32,input []byte) (error) {
  if len(input) < 4 {
    return ErrInputBufferLength;
  }
  var i int
  var tmp uint32
  *output = uint32(0)
  for i = 0; i < 4; i++{
    tmp = uint32(input[3 - i])
    *output = *output << 8
    *output = *output + tmp
  }
  return nil;
}

// tannhauser, transform an uint32 into a byte[] of length 4
func Ui32ToBL(output []byte, input uint32) (error) {
  if len(output) < 4 {
    return ErrOutputBufferLength;
  }
  binary.LittleEndian.PutUint32(output, input)
  return nil;
}

// tannhauser, transform a byte[] of length 8 into a uint64
func B2Uint64L(output *uint64, input []byte) (error) {
  if len(input) < 8 {
    return ErrInputBufferLength;
  }
  var i int
  var tmp uint64
  *output = uint64(0)
  for i = 0; i < 8; i++{
    tmp = uint64(input[7 - i])
    *output = *output << 8
    *output = *output + tmp
  }
  return nil;
}

// tannhauser, transform an uint64 into a byte[] of length 8
func Ui64ToBL(output []byte, input uint64) (error) {
  if len(output) < 8 {
    return ErrOutputBufferLength;
  }
  binary.LittleEndian.PutUint64(output, input)
  return nil;
}

func B2Uint16B(output *uint16, input []byte) (error) {
  if len(input) < 2 {
    return ErrInputBufferLength;
  }
  var i int
  var tmp uint16
  *output = uint16(0)
  for i = 0; i < 2; i++{
    tmp = uint16(input[1 - i]);
    *output = *output << 8
    *output = *output + tmp
  }
  return nil;
}

// tannhauser, transform an uint16 into a byte[] of length 2
func Ui16ToBB(output []byte, input uint16) (error) {
  if len(output) < 2 {
    return ErrOutputBufferLength;
  }
  binary.BigEndian.PutUint16(output, input)
  return nil;
}

// tannhauser, transform a byte[] of length 4 into a uint32
func B2Uint32B(output *uint32,input []byte) (error) {
  if len(input) < 4 {
    return ErrInputBufferLength;
  }
  var i int
  var tmp uint32
  *output = uint32(0)
  for i = 0; i < 4; i++{
    tmp = uint32(input[3 - i])
    *output = *output << 8
    *output = *output + tmp
  }
  return nil;
}

// tannhauser, transform an uint32 into a byte[] of length 4
func Ui32ToBB(output []byte, input uint32) (error) {
  if len(output) < 4 {
    return ErrOutputBufferLength;
  }
  binary.BigEndian.PutUint32(output, input)
  return nil;
}

// tannhauser, transform a byte[] of length 8 into a uint64
func B2Uint64B(output *uint64, input []byte) (error) {
  if len(input) < 8 {
    return ErrInputBufferLength;
  }
  var i int
  var tmp uint64
  *output = uint64(0)
  for i = 0; i < 8; i++{
    tmp = uint64(input[7 - i])
    *output = *output << 8
    *output = *output + tmp
  }
  return nil;
}

// tannhauser, transform an uint64 into a byte[] of length 8
func Ui64ToBB(output []byte, input uint64) (error) {
  if len(output) < 8 {
    return ErrOutputBufferLength;
  }
  binary.BigEndian.PutUint64(output, input)
  return nil;
}

