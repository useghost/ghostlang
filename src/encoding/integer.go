// You can edit this code!
// Click here and start typing.
package encoding

import (
	"encoding/binary"
)
func WriteI8LE(n uint8) []byte {
	return []byte{
		byte(n),
	};
}

func ReadI8LE(b []byte) uint8 {
	return uint8(b[0]);
} 


func WriteI16LE(n uint16) []byte {
	buf := make([]byte, 2)
	binary.LittleEndian.PutUint16(buf, n);
	return buf;
}

func ReadI16LE(b []byte) uint16 {
	return binary.LittleEndian.Uint16(b);
} 

func WriteI32LE(n uint32) []byte {
	buf := make([]byte, 4)
	binary.LittleEndian.PutUint32(buf, n);
	return buf;
}

func ReadI32LE(b []byte) uint32 {
	return binary.LittleEndian.Uint32(b);
}


func WriteI64LE(n uint64) []byte {
	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, n);
	return buf;
}

func ReadI64LE(b []byte) uint64 {
	return binary.LittleEndian.Uint64(b);
}


func WriteI8BE(n uint8) []byte {
	return []byte{
		byte(n),
	};
}

func ReadI8BE(b []byte) uint8 {
	return uint8(b[0]);
} 


func WriteI16BE(n uint16) []byte {
	buf := make([]byte, 2)
	binary.BigEndian.PutUint16(buf, n);
	return buf;
}

func ReadI16BE(b []byte) uint16 {
	return binary.BigEndian.Uint16(b);
} 

func WriteI32BE(n uint32) []byte {
	buf := make([]byte, 4)
	binary.BigEndian.PutUint32(buf, n);
	return buf;
}

func ReadI32BE(b []byte) uint32 {
	return binary.BigEndian.Uint32(b);
}


func WriteI64BE(n uint64) []byte {
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, n);
	return buf;
}

func ReadI64BE(b []byte) uint64 {
	return binary.BigEndian.Uint64(b);
}
