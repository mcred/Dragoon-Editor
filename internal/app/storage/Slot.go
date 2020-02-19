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
		if a.Reversed == true {
			return int(binary.LittleEndian.Uint16(b))
		} else {
			return int(binary.BigEndian.Uint16(b))
		}
	}
}

func (s *Slot) SetValueAtLocation(a common.Attribute, val int) {
	s.Data[a.Location] = byte(val)
}