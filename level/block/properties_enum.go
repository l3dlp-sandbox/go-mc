// Code generated by generator/properties/main.go; DO NOT EDIT.

package block

import (
	"errors"
	"strconv"
)

type AttachFace byte

const (
	AttachFaceFloor AttachFace = iota
	AttachFaceWall
	AttachFaceCeiling
)

var strAttachFace = [...]string{"floor", "wall", "ceiling"}

func (a AttachFace) String() string {
	if int(a) < len(strAttachFace) {
		return strAttachFace[a]
	}
	return "invalid AttachFace"
}

func (a AttachFace) MarshalText() (text []byte, err error) {
	if int(a) < len(strAttachFace) {
		return []byte(strAttachFace[a]), nil
	}
	return nil, errors.New("invalid AttachFace: " + strconv.Itoa(int(a)))
}

func (a *AttachFace) UnmarshalText(text []byte) error {
	switch str := string(text); str {
	case "floor":
		*a = AttachFaceFloor
	case "wall":
		*a = AttachFaceWall
	case "ceiling":
		*a = AttachFaceCeiling
	default:
		return errors.New("unknown AttachFace: " + str)
	}
	return nil
}

type BambooLeaves byte

const (
	BambooLeavesNone BambooLeaves = iota
	BambooLeavesSmall
	BambooLeavesLarge
)

var strBambooLeaves = [...]string{"none", "small", "large"}

func (b BambooLeaves) String() string {
	if int(b) < len(strBambooLeaves) {
		return strBambooLeaves[b]
	}
	return "invalid BambooLeaves"
}

func (b BambooLeaves) MarshalText() (text []byte, err error) {
	if int(b) < len(strBambooLeaves) {
		return []byte(strBambooLeaves[b]), nil
	}
	return nil, errors.New("invalid BambooLeaves: " + strconv.Itoa(int(b)))
}

func (b *BambooLeaves) UnmarshalText(text []byte) error {
	switch str := string(text); str {
	case "none":
		*b = BambooLeavesNone
	case "small":
		*b = BambooLeavesSmall
	case "large":
		*b = BambooLeavesLarge
	default:
		return errors.New("unknown BambooLeaves: " + str)
	}
	return nil
}

type BedPart byte

const (
	BedPartHead BedPart = iota
	BedPartFoot
)

var strBedPart = [...]string{"head", "foot"}

func (b BedPart) String() string {
	if int(b) < len(strBedPart) {
		return strBedPart[b]
	}
	return "invalid BedPart"
}

func (b BedPart) MarshalText() (text []byte, err error) {
	if int(b) < len(strBedPart) {
		return []byte(strBedPart[b]), nil
	}
	return nil, errors.New("invalid BedPart: " + strconv.Itoa(int(b)))
}

func (b *BedPart) UnmarshalText(text []byte) error {
	switch str := string(text); str {
	case "head":
		*b = BedPartHead
	case "foot":
		*b = BedPartFoot
	default:
		return errors.New("unknown BedPart: " + str)
	}
	return nil
}

type BellAttachType byte

const (
	BellAttachTypeFloor BellAttachType = iota
	BellAttachTypeCeiling
	BellAttachTypeSingleWall
	BellAttachTypeDoubleWall
)

var strBellAttachType = [...]string{"floor", "ceiling", "single_wall", "double_wall"}

func (b BellAttachType) String() string {
	if int(b) < len(strBellAttachType) {
		return strBellAttachType[b]
	}
	return "invalid BellAttachType"
}

func (b BellAttachType) MarshalText() (text []byte, err error) {
	if int(b) < len(strBellAttachType) {
		return []byte(strBellAttachType[b]), nil
	}
	return nil, errors.New("invalid BellAttachType: " + strconv.Itoa(int(b)))
}

func (b *BellAttachType) UnmarshalText(text []byte) error {
	switch str := string(text); str {
	case "floor":
		*b = BellAttachTypeFloor
	case "ceiling":
		*b = BellAttachTypeCeiling
	case "single_wall":
		*b = BellAttachTypeSingleWall
	case "double_wall":
		*b = BellAttachTypeDoubleWall
	default:
		return errors.New("unknown BellAttachType: " + str)
	}
	return nil
}

type ChestType byte

const (
	ChestTypeSingle ChestType = iota
	ChestTypeLeft
	ChestTypeRight
)

var strChestType = [...]string{"single", "left", "right"}

func (c ChestType) String() string {
	if int(c) < len(strChestType) {
		return strChestType[c]
	}
	return "invalid ChestType"
}

func (c ChestType) MarshalText() (text []byte, err error) {
	if int(c) < len(strChestType) {
		return []byte(strChestType[c]), nil
	}
	return nil, errors.New("invalid ChestType: " + strconv.Itoa(int(c)))
}

func (c *ChestType) UnmarshalText(text []byte) error {
	switch str := string(text); str {
	case "single":
		*c = ChestTypeSingle
	case "left":
		*c = ChestTypeLeft
	case "right":
		*c = ChestTypeRight
	default:
		return errors.New("unknown ChestType: " + str)
	}
	return nil
}

type ComparatorMode byte

const (
	ComparatorModeCompare ComparatorMode = iota
	ComparatorModeSubtract
)

var strComparatorMode = [...]string{"compare", "subtract"}

func (c ComparatorMode) String() string {
	if int(c) < len(strComparatorMode) {
		return strComparatorMode[c]
	}
	return "invalid ComparatorMode"
}

func (c ComparatorMode) MarshalText() (text []byte, err error) {
	if int(c) < len(strComparatorMode) {
		return []byte(strComparatorMode[c]), nil
	}
	return nil, errors.New("invalid ComparatorMode: " + strconv.Itoa(int(c)))
}

func (c *ComparatorMode) UnmarshalText(text []byte) error {
	switch str := string(text); str {
	case "compare":
		*c = ComparatorModeCompare
	case "subtract":
		*c = ComparatorModeSubtract
	default:
		return errors.New("unknown ComparatorMode: " + str)
	}
	return nil
}

type Direction byte

const (
	Down Direction = iota
	Up
	North
	South
	West
	East
)

var strDirection = [...]string{"down", "up", "north", "south", "west", "east"}

func (d Direction) String() string {
	if int(d) < len(strDirection) {
		return strDirection[d]
	}
	return "invalid Direction"
}

func (d Direction) MarshalText() (text []byte, err error) {
	if int(d) < len(strDirection) {
		return []byte(strDirection[d]), nil
	}
	return nil, errors.New("invalid Direction: " + strconv.Itoa(int(d)))
}

func (d *Direction) UnmarshalText(text []byte) error {
	switch str := string(text); str {
	case "down":
		*d = Down
	case "up":
		*d = Up
	case "north":
		*d = North
	case "south":
		*d = South
	case "west":
		*d = West
	case "east":
		*d = East
	default:
		return errors.New("unknown Direction: " + str)
	}
	return nil
}

type Axis byte

const (
	X Axis = iota
	Y
	Z
)

var strAxis = [...]string{"x", "y", "z"}

func (a Axis) String() string {
	if int(a) < len(strAxis) {
		return strAxis[a]
	}
	return "invalid Axis"
}

func (a Axis) MarshalText() (text []byte, err error) {
	if int(a) < len(strAxis) {
		return []byte(strAxis[a]), nil
	}
	return nil, errors.New("invalid Axis: " + strconv.Itoa(int(a)))
}

func (a *Axis) UnmarshalText(text []byte) error {
	switch str := string(text); str {
	case "x":
		*a = X
	case "y":
		*a = Y
	case "z":
		*a = Z
	default:
		return errors.New("unknown Axis: " + str)
	}
	return nil
}

type DoorHingeSide byte

const (
	DoorHingeSideLeft DoorHingeSide = iota
	DoorHingeSideRight
)

var strDoorHingeSide = [...]string{"left", "right"}

func (d DoorHingeSide) String() string {
	if int(d) < len(strDoorHingeSide) {
		return strDoorHingeSide[d]
	}
	return "invalid DoorHingeSide"
}

func (d DoorHingeSide) MarshalText() (text []byte, err error) {
	if int(d) < len(strDoorHingeSide) {
		return []byte(strDoorHingeSide[d]), nil
	}
	return nil, errors.New("invalid DoorHingeSide: " + strconv.Itoa(int(d)))
}

func (d *DoorHingeSide) UnmarshalText(text []byte) error {
	switch str := string(text); str {
	case "left":
		*d = DoorHingeSideLeft
	case "right":
		*d = DoorHingeSideRight
	default:
		return errors.New("unknown DoorHingeSide: " + str)
	}
	return nil
}

type DoubleBlockHalf byte

const (
	DoubleBlockHalfUpper DoubleBlockHalf = iota
	DoubleBlockHalfLower
)

var strDoubleBlockHalf = [...]string{"upper", "lower"}

func (d DoubleBlockHalf) String() string {
	if int(d) < len(strDoubleBlockHalf) {
		return strDoubleBlockHalf[d]
	}
	return "invalid DoubleBlockHalf"
}

func (d DoubleBlockHalf) MarshalText() (text []byte, err error) {
	if int(d) < len(strDoubleBlockHalf) {
		return []byte(strDoubleBlockHalf[d]), nil
	}
	return nil, errors.New("invalid DoubleBlockHalf: " + strconv.Itoa(int(d)))
}

func (d *DoubleBlockHalf) UnmarshalText(text []byte) error {
	switch str := string(text); str {
	case "upper":
		*d = DoubleBlockHalfUpper
	case "lower":
		*d = DoubleBlockHalfLower
	default:
		return errors.New("unknown DoubleBlockHalf: " + str)
	}
	return nil
}

type DripstoneThickness byte

const (
	DripstoneThicknessTipMerge DripstoneThickness = iota
	DripstoneThicknessTip
	DripstoneThicknessFrustum
	DripstoneThicknessMiddle
	DripstoneThicknessBase
)

var strDripstoneThickness = [...]string{"tip_merge", "tip", "frustum", "middle", "base"}

func (d DripstoneThickness) String() string {
	if int(d) < len(strDripstoneThickness) {
		return strDripstoneThickness[d]
	}
	return "invalid DripstoneThickness"
}

func (d DripstoneThickness) MarshalText() (text []byte, err error) {
	if int(d) < len(strDripstoneThickness) {
		return []byte(strDripstoneThickness[d]), nil
	}
	return nil, errors.New("invalid DripstoneThickness: " + strconv.Itoa(int(d)))
}

func (d *DripstoneThickness) UnmarshalText(text []byte) error {
	switch str := string(text); str {
	case "tip_merge":
		*d = DripstoneThicknessTipMerge
	case "tip":
		*d = DripstoneThicknessTip
	case "frustum":
		*d = DripstoneThicknessFrustum
	case "middle":
		*d = DripstoneThicknessMiddle
	case "base":
		*d = DripstoneThicknessBase
	default:
		return errors.New("unknown DripstoneThickness: " + str)
	}
	return nil
}

type Half byte

const (
	Top Half = iota
	Bottom
)

var strHalf = [...]string{"top", "bottom"}

func (h Half) String() string {
	if int(h) < len(strHalf) {
		return strHalf[h]
	}
	return "invalid Half"
}

func (h Half) MarshalText() (text []byte, err error) {
	if int(h) < len(strHalf) {
		return []byte(strHalf[h]), nil
	}
	return nil, errors.New("invalid Half: " + strconv.Itoa(int(h)))
}

func (h *Half) UnmarshalText(text []byte) error {
	switch str := string(text); str {
	case "top":
		*h = Top
	case "bottom":
		*h = Bottom
	default:
		return errors.New("unknown Half: " + str)
	}
	return nil
}

type NoteBlockInstrument byte

const (
	NoteBlockInstrumentHarp NoteBlockInstrument = iota
	NoteBlockInstrumentBasedrum
	NoteBlockInstrumentSnare
	NoteBlockInstrumentHat
	NoteBlockInstrumentBass
	NoteBlockInstrumentFlute
	NoteBlockInstrumentBell
	NoteBlockInstrumentGuitar
	NoteBlockInstrumentChime
	NoteBlockInstrumentXylophone
	NoteBlockInstrumentIronXylophone
	NoteBlockInstrumentCowBell
	NoteBlockInstrumentDidgeridoo
	NoteBlockInstrumentBit
	NoteBlockInstrumentBanjo
	NoteBlockInstrumentPling
	NoteBlockInstrumentZombie
	NoteBlockInstrumentSkeleton
	NoteBlockInstrumentCreeper
	NoteBlockInstrumentDragon
	NoteBlockInstrumentWitherSkeleton
	NoteBlockInstrumentPiglin
	NoteBlockInstrumentCustomHead
)

var strNoteBlockInstrument = [...]string{"harp", "basedrum", "snare", "hat", "bass", "flute", "bell", "guitar", "chime", "xylophone", "iron_xylophone", "cow_bell", "didgeridoo", "bit", "banjo", "pling", "zombie", "skeleton", "creeper", "dragon", "wither_skeleton", "piglin", "custom_head"}

func (n NoteBlockInstrument) String() string {
	if int(n) < len(strNoteBlockInstrument) {
		return strNoteBlockInstrument[n]
	}
	return "invalid NoteBlockInstrument"
}

func (n NoteBlockInstrument) MarshalText() (text []byte, err error) {
	if int(n) < len(strNoteBlockInstrument) {
		return []byte(strNoteBlockInstrument[n]), nil
	}
	return nil, errors.New("invalid NoteBlockInstrument: " + strconv.Itoa(int(n)))
}

func (n *NoteBlockInstrument) UnmarshalText(text []byte) error {
	switch str := string(text); str {
	case "harp":
		*n = NoteBlockInstrumentHarp
	case "basedrum":
		*n = NoteBlockInstrumentBasedrum
	case "snare":
		*n = NoteBlockInstrumentSnare
	case "hat":
		*n = NoteBlockInstrumentHat
	case "bass":
		*n = NoteBlockInstrumentBass
	case "flute":
		*n = NoteBlockInstrumentFlute
	case "bell":
		*n = NoteBlockInstrumentBell
	case "guitar":
		*n = NoteBlockInstrumentGuitar
	case "chime":
		*n = NoteBlockInstrumentChime
	case "xylophone":
		*n = NoteBlockInstrumentXylophone
	case "iron_xylophone":
		*n = NoteBlockInstrumentIronXylophone
	case "cow_bell":
		*n = NoteBlockInstrumentCowBell
	case "didgeridoo":
		*n = NoteBlockInstrumentDidgeridoo
	case "bit":
		*n = NoteBlockInstrumentBit
	case "banjo":
		*n = NoteBlockInstrumentBanjo
	case "pling":
		*n = NoteBlockInstrumentPling
	case "zombie":
		*n = NoteBlockInstrumentZombie
	case "skeleton":
		*n = NoteBlockInstrumentSkeleton
	case "creeper":
		*n = NoteBlockInstrumentCreeper
	case "dragon":
		*n = NoteBlockInstrumentDragon
	case "wither_skeleton":
		*n = NoteBlockInstrumentWitherSkeleton
	case "piglin":
		*n = NoteBlockInstrumentPiglin
	case "custom_head":
		*n = NoteBlockInstrumentCustomHead
	default:
		return errors.New("unknown NoteBlockInstrument: " + str)
	}
	return nil
}

type PistonType byte

const (
	PistonTypeNormal PistonType = iota
	PistonTypeSticky
)

var strPistonType = [...]string{"normal", "sticky"}

func (p PistonType) String() string {
	if int(p) < len(strPistonType) {
		return strPistonType[p]
	}
	return "invalid PistonType"
}

func (p PistonType) MarshalText() (text []byte, err error) {
	if int(p) < len(strPistonType) {
		return []byte(strPistonType[p]), nil
	}
	return nil, errors.New("invalid PistonType: " + strconv.Itoa(int(p)))
}

func (p *PistonType) UnmarshalText(text []byte) error {
	switch str := string(text); str {
	case "normal":
		*p = PistonTypeNormal
	case "sticky":
		*p = PistonTypeSticky
	default:
		return errors.New("unknown PistonType: " + str)
	}
	return nil
}

type RailShape byte

const (
	RailShapeNorthSouth RailShape = iota
	RailShapeEastWest
	RailShapeAscendingEast
	RailShapeAscendingWest
	RailShapeAscendingNorth
	RailShapeAscendingSouth
	RailShapeSouthEast
	RailShapeSouthWest
	RailShapeNorthWest
	RailShapeNorthEast
)

var strRailShape = [...]string{"north_south", "east_west", "ascending_east", "ascending_west", "ascending_north", "ascending_south", "south_east", "south_west", "north_west", "north_east"}

func (r RailShape) String() string {
	if int(r) < len(strRailShape) {
		return strRailShape[r]
	}
	return "invalid RailShape"
}

func (r RailShape) MarshalText() (text []byte, err error) {
	if int(r) < len(strRailShape) {
		return []byte(strRailShape[r]), nil
	}
	return nil, errors.New("invalid RailShape: " + strconv.Itoa(int(r)))
}

func (r *RailShape) UnmarshalText(text []byte) error {
	switch str := string(text); str {
	case "north_south":
		*r = RailShapeNorthSouth
	case "east_west":
		*r = RailShapeEastWest
	case "ascending_east":
		*r = RailShapeAscendingEast
	case "ascending_west":
		*r = RailShapeAscendingWest
	case "ascending_north":
		*r = RailShapeAscendingNorth
	case "ascending_south":
		*r = RailShapeAscendingSouth
	case "south_east":
		*r = RailShapeSouthEast
	case "south_west":
		*r = RailShapeSouthWest
	case "north_west":
		*r = RailShapeNorthWest
	case "north_east":
		*r = RailShapeNorthEast
	default:
		return errors.New("unknown RailShape: " + str)
	}
	return nil
}

type RedstoneSide byte

const (
	RedstoneSideUp RedstoneSide = iota
	RedstoneSideSide
	RedstoneSideNone
)

var strRedstoneSide = [...]string{"up", "side", "none"}

func (r RedstoneSide) String() string {
	if int(r) < len(strRedstoneSide) {
		return strRedstoneSide[r]
	}
	return "invalid RedstoneSide"
}

func (r RedstoneSide) MarshalText() (text []byte, err error) {
	if int(r) < len(strRedstoneSide) {
		return []byte(strRedstoneSide[r]), nil
	}
	return nil, errors.New("invalid RedstoneSide: " + strconv.Itoa(int(r)))
}

func (r *RedstoneSide) UnmarshalText(text []byte) error {
	switch str := string(text); str {
	case "up":
		*r = RedstoneSideUp
	case "side":
		*r = RedstoneSideSide
	case "none":
		*r = RedstoneSideNone
	default:
		return errors.New("unknown RedstoneSide: " + str)
	}
	return nil
}

type SculkSensorPhase byte

const (
	SculkSensorPhaseInactive SculkSensorPhase = iota
	SculkSensorPhaseActive
	SculkSensorPhaseCooldown
)

var strSculkSensorPhase = [...]string{"inactive", "active", "cooldown"}

func (s SculkSensorPhase) String() string {
	if int(s) < len(strSculkSensorPhase) {
		return strSculkSensorPhase[s]
	}
	return "invalid SculkSensorPhase"
}

func (s SculkSensorPhase) MarshalText() (text []byte, err error) {
	if int(s) < len(strSculkSensorPhase) {
		return []byte(strSculkSensorPhase[s]), nil
	}
	return nil, errors.New("invalid SculkSensorPhase: " + strconv.Itoa(int(s)))
}

func (s *SculkSensorPhase) UnmarshalText(text []byte) error {
	switch str := string(text); str {
	case "inactive":
		*s = SculkSensorPhaseInactive
	case "active":
		*s = SculkSensorPhaseActive
	case "cooldown":
		*s = SculkSensorPhaseCooldown
	default:
		return errors.New("unknown SculkSensorPhase: " + str)
	}
	return nil
}

type SlabType byte

const (
	SlabTypeTop SlabType = iota
	SlabTypeBottom
	SlabTypeDouble
)

var strSlabType = [...]string{"top", "bottom", "double"}

func (s SlabType) String() string {
	if int(s) < len(strSlabType) {
		return strSlabType[s]
	}
	return "invalid SlabType"
}

func (s SlabType) MarshalText() (text []byte, err error) {
	if int(s) < len(strSlabType) {
		return []byte(strSlabType[s]), nil
	}
	return nil, errors.New("invalid SlabType: " + strconv.Itoa(int(s)))
}

func (s *SlabType) UnmarshalText(text []byte) error {
	switch str := string(text); str {
	case "top":
		*s = SlabTypeTop
	case "bottom":
		*s = SlabTypeBottom
	case "double":
		*s = SlabTypeDouble
	default:
		return errors.New("unknown SlabType: " + str)
	}
	return nil
}

type StairsShape byte

const (
	StairsShapeStraight StairsShape = iota
	StairsShapeInnerLeft
	StairsShapeInnerRight
	StairsShapeOuterLeft
	StairsShapeOuterRight
)

var strStairsShape = [...]string{"straight", "inner_left", "inner_right", "outer_left", "outer_right"}

func (s StairsShape) String() string {
	if int(s) < len(strStairsShape) {
		return strStairsShape[s]
	}
	return "invalid StairsShape"
}

func (s StairsShape) MarshalText() (text []byte, err error) {
	if int(s) < len(strStairsShape) {
		return []byte(strStairsShape[s]), nil
	}
	return nil, errors.New("invalid StairsShape: " + strconv.Itoa(int(s)))
}

func (s *StairsShape) UnmarshalText(text []byte) error {
	switch str := string(text); str {
	case "straight":
		*s = StairsShapeStraight
	case "inner_left":
		*s = StairsShapeInnerLeft
	case "inner_right":
		*s = StairsShapeInnerRight
	case "outer_left":
		*s = StairsShapeOuterLeft
	case "outer_right":
		*s = StairsShapeOuterRight
	default:
		return errors.New("unknown StairsShape: " + str)
	}
	return nil
}

type StructureMode byte

const (
	StructureModeSave StructureMode = iota
	StructureModeLoad
	StructureModeCorner
	StructureModeData
)

var strStructureMode = [...]string{"save", "load", "corner", "data"}

func (s StructureMode) String() string {
	if int(s) < len(strStructureMode) {
		return strStructureMode[s]
	}
	return "invalid StructureMode"
}

func (s StructureMode) MarshalText() (text []byte, err error) {
	if int(s) < len(strStructureMode) {
		return []byte(strStructureMode[s]), nil
	}
	return nil, errors.New("invalid StructureMode: " + strconv.Itoa(int(s)))
}

func (s *StructureMode) UnmarshalText(text []byte) error {
	switch str := string(text); str {
	case "save":
		*s = StructureModeSave
	case "load":
		*s = StructureModeLoad
	case "corner":
		*s = StructureModeCorner
	case "data":
		*s = StructureModeData
	default:
		return errors.New("unknown StructureMode: " + str)
	}
	return nil
}

type Tilt byte

const (
	TiltNone Tilt = iota
	TiltUnstable
	TiltPartial
	TiltFull
)

var strTilt = [...]string{"none", "unstable", "partial", "full"}

func (t Tilt) String() string {
	if int(t) < len(strTilt) {
		return strTilt[t]
	}
	return "invalid Tilt"
}

func (t Tilt) MarshalText() (text []byte, err error) {
	if int(t) < len(strTilt) {
		return []byte(strTilt[t]), nil
	}
	return nil, errors.New("invalid Tilt: " + strconv.Itoa(int(t)))
}

func (t *Tilt) UnmarshalText(text []byte) error {
	switch str := string(text); str {
	case "none":
		*t = TiltNone
	case "unstable":
		*t = TiltUnstable
	case "partial":
		*t = TiltPartial
	case "full":
		*t = TiltFull
	default:
		return errors.New("unknown Tilt: " + str)
	}
	return nil
}

type WallSide byte

const (
	WallSideNone WallSide = iota
	WallSideLow
	WallSideTall
)

var strWallSide = [...]string{"none", "low", "tall"}

func (w WallSide) String() string {
	if int(w) < len(strWallSide) {
		return strWallSide[w]
	}
	return "invalid WallSide"
}

func (w WallSide) MarshalText() (text []byte, err error) {
	if int(w) < len(strWallSide) {
		return []byte(strWallSide[w]), nil
	}
	return nil, errors.New("invalid WallSide: " + strconv.Itoa(int(w)))
}

func (w *WallSide) UnmarshalText(text []byte) error {
	switch str := string(text); str {
	case "none":
		*w = WallSideNone
	case "low":
		*w = WallSideLow
	case "tall":
		*w = WallSideTall
	default:
		return errors.New("unknown WallSide: " + str)
	}
	return nil
}

type FrontAndTop byte

const (
	DownEast FrontAndTop = iota
	DownNorth
	DownSouth
	DownWest
	UpEast
	UpNorth
	UpSouth
	UpWest
	WestUp
	EastUp
	NorthUp
	SouthUp
)

var strFrontAndTop = [...]string{"down_east", "down_north", "down_south", "down_west", "up_east", "up_north", "up_south", "up_west", "west_up", "east_up", "north_up", "south_up"}

func (f FrontAndTop) String() string {
	if int(f) < len(strFrontAndTop) {
		return strFrontAndTop[f]
	}
	return "invalid FrontAndTop"
}

func (f FrontAndTop) MarshalText() (text []byte, err error) {
	if int(f) < len(strFrontAndTop) {
		return []byte(strFrontAndTop[f]), nil
	}
	return nil, errors.New("invalid FrontAndTop: " + strconv.Itoa(int(f)))
}

func (f *FrontAndTop) UnmarshalText(text []byte) error {
	switch str := string(text); str {
	case "down_east":
		*f = DownEast
	case "down_north":
		*f = DownNorth
	case "down_south":
		*f = DownSouth
	case "down_west":
		*f = DownWest
	case "up_east":
		*f = UpEast
	case "up_north":
		*f = UpNorth
	case "up_south":
		*f = UpSouth
	case "up_west":
		*f = UpWest
	case "west_up":
		*f = WestUp
	case "east_up":
		*f = EastUp
	case "north_up":
		*f = NorthUp
	case "south_up":
		*f = SouthUp
	default:
		return errors.New("unknown FrontAndTop: " + str)
	}
	return nil
}

type VaultState byte

const (
	VaultStateInactive VaultState = iota
	VaultStateActive
	VaultStateUnlocking
	VaultStateEjecting
)

var strVaultState = [...]string{"inactive", "active", "unlocking", "ejecting"}

func (v VaultState) String() string {
	if int(v) < len(strVaultState) {
		return strVaultState[v]
	}
	return "invalid VaultState"
}

func (v VaultState) MarshalText() (text []byte, err error) {
	if int(v) < len(strVaultState) {
		return []byte(strVaultState[v]), nil
	}
	return nil, errors.New("invalid VaultState: " + strconv.Itoa(int(v)))
}

func (v *VaultState) UnmarshalText(text []byte) error {
	switch str := string(text); str {
	case "inactive":
		*v = VaultStateInactive
	case "active":
		*v = VaultStateActive
	case "unlocking":
		*v = VaultStateUnlocking
	case "ejecting":
		*v = VaultStateEjecting
	default:
		return errors.New("unknown VaultState: " + str)
	}
	return nil
}

type TrialSpawnerState byte

const (
	TrialSpawnerStateInactive TrialSpawnerState = iota
	TrialSpawnerStateWaitingForPlayers
	TrialSpawnerStateActive
	TrialSpawnerStateWaitingForRewardEjection
	TrialSpawnerStateEjectingReward
	TrialSpawnerStateCooldown
)

var strTrialSpawnerState = [...]string{"inactive", "waiting_for_players", "active", "waiting_for_reward_ejection", "ejecting_reward", "cooldown"}

func (t TrialSpawnerState) String() string {
	if int(t) < len(strTrialSpawnerState) {
		return strTrialSpawnerState[t]
	}
	return "invalid TrialSpawnerState"
}

func (t TrialSpawnerState) MarshalText() (text []byte, err error) {
	if int(t) < len(strTrialSpawnerState) {
		return []byte(strTrialSpawnerState[t]), nil
	}
	return nil, errors.New("invalid TrialSpawnerState: " + strconv.Itoa(int(t)))
}

func (t *TrialSpawnerState) UnmarshalText(text []byte) error {
	switch str := string(text); str {
	case "inactive":
		*t = TrialSpawnerStateInactive
	case "waiting_for_players":
		*t = TrialSpawnerStateWaitingForPlayers
	case "active":
		*t = TrialSpawnerStateActive
	case "waiting_for_reward_ejection":
		*t = TrialSpawnerStateWaitingForRewardEjection
	case "ejecting_reward":
		*t = TrialSpawnerStateEjectingReward
	case "cooldown":
		*t = TrialSpawnerStateCooldown
	default:
		return errors.New("unknown TrialSpawnerState: " + str)
	}
	return nil
}
