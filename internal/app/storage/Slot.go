package storage

import (
	"LODeditor/internal/app/common"
	"encoding/binary"
)

type Slot struct {
	Head int
	Data []byte
}

func (s *Slot) GetValueByAttribute(a common.Attribute) int {
	b := s.Data[a.Location:(a.Location+a.Bytes)]
	if a.Bytes == 1 {
		return int(b[0])
	} else {
		if a.Reversed {
			return int(binary.LittleEndian.Uint16(b))
		} else {
			return int(binary.BigEndian.Uint16(b))
		}
	}
}

func (s *Slot) SetValueAtLocation(a common.Attribute, val int) {
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
		s.Data[a.Location + i] = b[i]
	}
}