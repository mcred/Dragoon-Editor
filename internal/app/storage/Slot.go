package storage

import (
	"encoding/binary"
)

type Slot struct {
	Head int
	Data []byte
}

func (s *Slot) GetValueByAttribute(a Attribute) int {
	b := s.Data[a.Location:(a.Location + a.Bytes)]
	if a.Reversed {
		switch a.Bytes {
		case 2:
			return int(binary.LittleEndian.Uint16(b))
		case 4:
			return int(binary.LittleEndian.Uint32(b))
		default:
			return int(b[0])
		}
	} else {
		switch a.Bytes {
		case 2:
			return int(binary.BigEndian.Uint16(b))
		case 4:
			return int(binary.LittleEndian.Uint32(b))
		default:
			return int(b[0])
		}
	}
}

func (s *Slot) SetValueAtLocation(a Attribute, val int) {
	b := make([]byte, a.Bytes)
	if a.Reversed {
		switch a.Bytes {
		case 2:
			binary.LittleEndian.PutUint16(b, uint16(val))
		case 4:
			binary.LittleEndian.PutUint32(b, uint32(val))
		default:
			b[0] = byte(val)
		}
	} else {
		switch a.Bytes {
		case 2:
			binary.BigEndian.PutUint16(b, uint16(val))
		case 4:
			binary.BigEndian.PutUint32(b, uint32(val))
		default:
			b[0] = byte(val)
		}
	}
	for i := 0; i < a.Bytes; i++ {
		s.Data[a.Location+i] = b[i]
	}
}
