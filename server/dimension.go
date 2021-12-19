package server

import (
	"bytes"
	"github.com/Tnze/go-mc/data/packetid"
	pk "github.com/Tnze/go-mc/net/packet"
	"github.com/Tnze/go-mc/save"
	"io"
	"math/bits"
	"sync"
)

type Dimension interface {
	Info() DimInfo
	PlayerJoin(p *Player)
	PlayerQuit(p *Player)
}

type DimInfo struct {
	Name       string
	HashedSeed uint64
}

type ChunkPos struct{ X, Z int }
type Chunk struct {
	sync.Mutex
	sections   []Section
	HeightMaps *save.BitStorage
}

type Section struct {
	blockCount int16
	States     *save.PaletteContainer
	Biomes     *save.PaletteContainer
}

func (s *Section) WriteTo(w io.Writer) (int64, error) {
	return pk.Tuple{
		pk.Short(s.blockCount),
		s.States,
		s.Biomes,
	}.WriteTo(w)
}

func (s *Section) ReadFrom(r io.Reader) (int64, error) {
	return pk.Tuple{
		pk.Short(s.blockCount),
		s.States,
		s.Biomes,
	}.ReadFrom(r)
}

func EmptyChunk(secs int) *Chunk {
	sections := make([]Section, secs)
	for i := range sections {
		sections[i] = Section{
			blockCount: 0,
			States:     save.NewStatesPaletteContainer(nil, 16*16*16),
			Biomes:     save.NewBiomesPaletteContainer(nil, 4*4*4),
		}
	}
	return &Chunk{
		sections:   sections,
		HeightMaps: save.NewBitStorage(bits.Len(uint(secs)*16), 16*16, nil),
	}
}

func (c *Chunk) WriteTo(w io.Writer) (int64, error) {
	data, err := c.Data()
	if err != nil {
		return 0, err
	}
	return pk.Tuple{
		// Heightmaps
		pk.NBT(struct {
			MotionBlocking []uint64 `nbt:"MOTION_BLOCKING"`
		}{c.HeightMaps.Raw()}),
		pk.ByteArray(data),
		pk.VarInt(0), // TODO: Block Entity
	}.WriteTo(w)
}

func (c *Chunk) Data() ([]byte, error) {
	var buff bytes.Buffer
	for _, section := range c.sections {
		_, err := section.WriteTo(&buff)
		if err != nil {
			return nil, err
		}
	}
	return buff.Bytes(), nil
}

type lightData struct {
	SkyLightMask   pk.BitSet
	BlockLightMask pk.BitSet
	SkyLight       []pk.ByteArray
	BlockLight     []pk.ByteArray
}

func bitSetRev(set pk.BitSet) pk.BitSet {
	rev := make(pk.BitSet, len(set))
	for i := range rev {
		rev[i] = ^set[i]
	}
	return rev
}

func (l *lightData) WriteTo(w io.Writer) (int64, error) {
	return pk.Tuple{
		pk.Boolean(true), // Trust Edges
		l.SkyLightMask,
		l.BlockLightMask,
		bitSetRev(l.SkyLightMask),
		bitSetRev(l.BlockLightMask),
		pk.Array(l.SkyLight),
		pk.Array(l.BlockLight),
	}.WriteTo(w)
}

type SimpleDim struct {
	numOfSection int
	Columns      map[ChunkPos]*Chunk
}

func NewSimpleDim(secs int) *SimpleDim {
	return &SimpleDim{
		numOfSection: secs,
		Columns:      make(map[ChunkPos]*Chunk),
	}
}

func (s *SimpleDim) LoadChunk(pos ChunkPos, c *Chunk) {
	s.Columns[pos] = c
}

func (s *SimpleDim) Info() DimInfo {
	return DimInfo{
		Name:       "minecraft:overworld",
		HashedSeed: 0,
	}
}

func (s *SimpleDim) PlayerJoin(p *Player) {
	for pos, column := range s.Columns {
		column.Lock()
		packet := pk.Marshal(
			packetid.ClientboundLevelChunkWithLight,
			pk.Int(pos.X), pk.Int(pos.Z),
			column,
			&lightData{
				SkyLightMask:   make(pk.BitSet, (16*16*16-1)>>6+1),
				BlockLightMask: make(pk.BitSet, (16*16*16-1)>>6+1),
				SkyLight:       []pk.ByteArray{},
				BlockLight:     []pk.ByteArray{},
			},
		)
		column.Unlock()

		err := p.WritePacket(Packet757(packet))
		if err != nil {
			return
		}
	}

	err := p.WritePacket(Packet757(pk.Marshal(
		packetid.ClientboundPlayerPosition,
		pk.Double(0), pk.Double(0), pk.Double(0),
		pk.Float(0), pk.Float(0),
		pk.Byte(0),
		pk.VarInt(0),
		pk.Boolean(true),
	)))
	if err != nil {
		return
	}
}

func (s *SimpleDim) PlayerQuit(p *Player) {

}