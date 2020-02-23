package mhfpacket

import (
	"github.com/Andoryuuta/Erupe/network"
	"github.com/Andoryuuta/byteframe"
)

// MsgMhfGetSeibattle represents the MSG_MHF_GET_SEIBATTLE
type MsgMhfGetSeibattle struct {
	AckHandle uint32
	Unk0      uint8
	Unk1      uint8
	Unk2      uint32 // Some shared ID with MSG_SYS_RECORD_LOG, world ID?
	Unk3      uint8
	Unk4      uint16
}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfGetSeibattle) Opcode() network.PacketID {
	return network.MSG_MHF_GET_SEIBATTLE
}

// Parse parses the packet from binary
func (m *MsgMhfGetSeibattle) Parse(bf *byteframe.ByteFrame) error {
	m.AckHandle = bf.ReadUint32()
	m.Unk0 = bf.ReadUint8()
	m.Unk1 = bf.ReadUint8()
	m.Unk2 = bf.ReadUint32()
	m.Unk3 = bf.ReadUint8()
	m.Unk4 = bf.ReadUint16()
	return nil
}

// Build builds a binary packet from the current data.
func (m *MsgMhfGetSeibattle) Build(bf *byteframe.ByteFrame) error {
	panic("Not implemented")
}
