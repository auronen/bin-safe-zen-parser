package parsers
// Code generated by kaitai-struct-compiler from a .ksy source file. DO NOT EDIT.

import (
	"github.com/kaitai-io/kaitai_struct_go_runtime/kaitai"
	"io"
)

/**
 * File format definition for binary safe ZEN world archive format for the ZenGin.
 */
type BsZen struct {
	Header  *BsZen_Header
	World   *BsZen_World
	OCWorld *BsZen_OCWorld
	_io     *kaitai.Stream
	_root   *BsZen
	_parent interface{}
}

func NewBsZen() *BsZen {
	return &BsZen{}
}

func (this *BsZen) Read(io *kaitai.Stream, parent interface{}, root *BsZen) (err error) {
	this._io = io
	this._parent = parent
	this._root = root

	tmp1 := NewBsZen_Header()
	err = tmp1.Read(this._io, this, this._root)
	if err != nil {
		return err
	}
	this.Header = tmp1
	tmp2 := NewBsZen_World()
	err = tmp2.Read(this._io, this, this._root)
	if err != nil {
		return err
	}
	this.World = tmp2
	tmp3 := NewBsZen_OCWorld()
	err = tmp3.Read(this._io, this, this._root)
	if err != nil {
		return err
	}
	this.OCWorld = tmp3
	return err
}

type BsZen_TypeColor struct {
	R       uint8
	G       uint8
	B       uint8
	A       uint8
	_io     *kaitai.Stream
	_root   *BsZen
	_parent *BsZen_Property
}

func NewBsZen_TypeColor() *BsZen_TypeColor {
	return &BsZen_TypeColor{}
}

func (this *BsZen_TypeColor) Read(io *kaitai.Stream, parent *BsZen_Property, root *BsZen) (err error) {
	this._io = io
	this._parent = parent
	this._root = root

	tmp4, err := this._io.ReadU1()
	if err != nil {
		return err
	}
	this.R = tmp4
	tmp5, err := this._io.ReadU1()
	if err != nil {
		return err
	}
	this.G = tmp5
	tmp6, err := this._io.ReadU1()
	if err != nil {
		return err
	}
	this.B = tmp6
	tmp7, err := this._io.ReadU1()
	if err != nil {
		return err
	}
	this.A = tmp7
	return err
}

type BsZen_TypeRaw struct {
	Size    uint16
	RawData []byte
	_io     *kaitai.Stream
	_root   *BsZen
	_parent *BsZen_Property
}

func NewBsZen_TypeRaw() *BsZen_TypeRaw {
	return &BsZen_TypeRaw{}
}

func (this *BsZen_TypeRaw) Read(io *kaitai.Stream, parent *BsZen_Property, root *BsZen) (err error) {
	this._io = io
	this._parent = parent
	this._root = root

	tmp8, err := this._io.ReadU2le()
	if err != nil {
		return err
	}
	this.Size = uint16(tmp8)
	tmp9, err := this._io.ReadBytes(int(this.Size))
	if err != nil {
		return err
	}
	tmp9 = tmp9
	this.RawData = tmp9
	return err
}

type BsZen_TypeFloat struct {
	Value   float32
	_io     *kaitai.Stream
	_root   *BsZen
	_parent *BsZen_Property
}

func NewBsZen_TypeFloat() *BsZen_TypeFloat {
	return &BsZen_TypeFloat{}
}

func (this *BsZen_TypeFloat) Read(io *kaitai.Stream, parent *BsZen_Property, root *BsZen) (err error) {
	this._io = io
	this._parent = parent
	this._root = root

	tmp10, err := this._io.ReadF4le()
	if err != nil {
		return err
	}
	this.Value = float32(tmp10)
	return err
}

type BsZen_OCWorld struct {
	ObjectId   *BsZen_Property
	Mab        *BsZen_Mab
	Properties []*BsZen_Property
	_io        *kaitai.Stream
	_root      *BsZen
	_parent    *BsZen
}

func NewBsZen_OCWorld() *BsZen_OCWorld {
	return &BsZen_OCWorld{}
}

func (this *BsZen_OCWorld) Read(io *kaitai.Stream, parent *BsZen, root *BsZen) (err error) {
	this._io = io
	this._parent = parent
	this._root = root

	tmp11 := NewBsZen_Property()
	err = tmp11.Read(this._io, this, this._root)
	if err != nil {
		return err
	}
	this.ObjectId = tmp11
	tmp12 := NewBsZen_Mab()
	err = tmp12.Read(this._io, this, this._root)
	if err != nil {
		return err
	}
	this.Mab = tmp12
	for i := 1; ; i++ {
		tmp13 := NewBsZen_Property()
		err = tmp13.Read(this._io, this, this._root)
		if err != nil {
			return err
		}
		_it := tmp13
		this.Properties = append(this.Properties, _it)
		tmp14, err := this._io.Pos()
		if err != nil {
			return err
		}
		if (int64(this._parent.World.HashTableOffset) - tmp14) == 0 {
			break
		}
	}
	return err
}

type BsZen_TypeRawFloat struct {
	Size         uint16
	RawFloatData []byte
	_io          *kaitai.Stream
	_root        *BsZen
	_parent      *BsZen_Property
}

func NewBsZen_TypeRawFloat() *BsZen_TypeRawFloat {
	return &BsZen_TypeRawFloat{}
}

func (this *BsZen_TypeRawFloat) Read(io *kaitai.Stream, parent *BsZen_Property, root *BsZen) (err error) {
	this._io = io
	this._parent = parent
	this._root = root

	tmp15, err := this._io.ReadU2le()
	if err != nil {
		return err
	}
	this.Size = uint16(tmp15)
	tmp16, err := this._io.ReadBytes(int(this.Size))
	if err != nil {
		return err
	}
	tmp16 = tmp16
	this.RawFloatData = tmp16
	return err
}

type BsZen_HashTableChunk struct {
	NameLength  uint16
	LinearValue uint16
	HashValue   uint32
	Str         string
	_io         *kaitai.Stream
	_root       *BsZen
	_parent     *BsZen_HashTableHeader
}

func NewBsZen_HashTableChunk() *BsZen_HashTableChunk {
	return &BsZen_HashTableChunk{}
}

func (this *BsZen_HashTableChunk) Read(io *kaitai.Stream, parent *BsZen_HashTableHeader, root *BsZen) (err error) {
	this._io = io
	this._parent = parent
	this._root = root

	tmp17, err := this._io.ReadU2le()
	if err != nil {
		return err
	}
	this.NameLength = uint16(tmp17)
	tmp18, err := this._io.ReadU2le()
	if err != nil {
		return err
	}
	this.LinearValue = uint16(tmp18)
	tmp19, err := this._io.ReadU4le()
	if err != nil {
		return err
	}
	this.HashValue = uint32(tmp19)
	tmp20, err := this._io.ReadBytes(int(this.NameLength))
	if err != nil {
		return err
	}
	tmp20 = tmp20
	this.Str = string(tmp20)
	return err
}

type BsZen_TypeBool struct {
	Value   uint32
	_io     *kaitai.Stream
	_root   *BsZen
	_parent *BsZen_Property
}

func NewBsZen_TypeBool() *BsZen_TypeBool {
	return &BsZen_TypeBool{}
}

func (this *BsZen_TypeBool) Read(io *kaitai.Stream, parent *BsZen_Property, root *BsZen) (err error) {
	this._io = io
	this._parent = parent
	this._root = root

	tmp21, err := this._io.ReadU4le()
	if err != nil {
		return err
	}
	this.Value = uint32(tmp21)
	return err
}

type BsZen_World struct {
	BsVersion       uint32
	ObjectCount     uint32
	HashTableOffset uint32
	_io             *kaitai.Stream
	_root           *BsZen
	_parent         *BsZen
	_f_ht           bool
	ht              *BsZen_HashTableHeader
}

func NewBsZen_World() *BsZen_World {
	return &BsZen_World{}
}

func (this *BsZen_World) Read(io *kaitai.Stream, parent *BsZen, root *BsZen) (err error) {
	this._io = io
	this._parent = parent
	this._root = root

	tmp22, err := this._io.ReadU4le()
	if err != nil {
		return err
	}
	this.BsVersion = uint32(tmp22)
	tmp23, err := this._io.ReadU4le()
	if err != nil {
		return err
	}
	this.ObjectCount = uint32(tmp23)
	tmp24, err := this._io.ReadU4le()
	if err != nil {
		return err
	}
	this.HashTableOffset = uint32(tmp24)
	return err
}
func (this *BsZen_World) Ht() (v *BsZen_HashTableHeader, err error) {
	if this._f_ht {
		return this.ht, nil
	}
	_pos, err := this._io.Pos()
	if err != nil {
		return nil, err
	}
	_, err = this._io.Seek(int64(this.HashTableOffset), io.SeekStart)
	if err != nil {
		return nil, err
	}
	tmp25 := NewBsZen_HashTableHeader()
	err = tmp25.Read(this._io, this, this._root)
	if err != nil {
		return nil, err
	}
	this.ht = tmp25
	_, err = this._io.Seek(_pos, io.SeekStart)
	if err != nil {
		return nil, err
	}
	this._f_ht = true
	this._f_ht = true
	return this.ht, nil
}

type BsZen_Bsp struct {
	BspMagic uint32
	Size     uint32
	Version  uint32
	Data     []byte
	_io      *kaitai.Stream
	_root    *BsZen
	_parent  *BsZen_Mab
}

func NewBsZen_Bsp() *BsZen_Bsp {
	return &BsZen_Bsp{}
}

func (this *BsZen_Bsp) Read(io *kaitai.Stream, parent *BsZen_Mab, root *BsZen) (err error) {
	this._io = io
	this._parent = parent
	this._root = root

	tmp26, err := this._io.ReadU4le()
	if err != nil {
		return err
	}
	this.BspMagic = uint32(tmp26)
	tmp27, err := this._io.ReadU4le()
	if err != nil {
		return err
	}
	this.Size = uint32(tmp27)
	tmp28, err := this._io.ReadU4le()
	if err != nil {
		return err
	}
	this.Version = uint32(tmp28)
	tmp29, err := this._io.ReadBytes(int((this.Size + 1)))
	if err != nil {
		return err
	}
	tmp29 = tmp29
	this.Data = tmp29
	return err
}

type BsZen_TypeEnum struct {
	Value   uint32
	_io     *kaitai.Stream
	_root   *BsZen
	_parent *BsZen_Property
}

func NewBsZen_TypeEnum() *BsZen_TypeEnum {
	return &BsZen_TypeEnum{}
}

func (this *BsZen_TypeEnum) Read(io *kaitai.Stream, parent *BsZen_Property, root *BsZen) (err error) {
	this._io = io
	this._parent = parent
	this._root = root

	tmp30, err := this._io.ReadU4le()
	if err != nil {
		return err
	}
	this.Value = uint32(tmp30)
	return err
}

type BsZen_Mab struct {
	MeshAndBsp *BsZen_Property
	Bsp        *BsZen_Bsp
	_io        *kaitai.Stream
	_root      *BsZen
	_parent    *BsZen_OCWorld
}

func NewBsZen_Mab() *BsZen_Mab {
	return &BsZen_Mab{}
}

func (this *BsZen_Mab) Read(io *kaitai.Stream, parent *BsZen_OCWorld, root *BsZen) (err error) {
	this._io = io
	this._parent = parent
	this._root = root

	tmp31 := NewBsZen_Property()
	err = tmp31.Read(this._io, this, this._root)
	if err != nil {
		return err
	}
	this.MeshAndBsp = tmp31
	tmp32 := NewBsZen_Bsp()
	err = tmp32.Read(this._io, this, this._root)
	if err != nil {
		return err
	}
	this.Bsp = tmp32
	return err
}

type BsZen_String struct {
	Str     string
	_io     *kaitai.Stream
	_root   *BsZen
	_parent *BsZen_Header
}

func NewBsZen_String() *BsZen_String {
	return &BsZen_String{}
}

func (this *BsZen_String) Read(io *kaitai.Stream, parent *BsZen_Header, root *BsZen) (err error) {
	this._io = io
	this._parent = parent
	this._root = root

	tmp33, err := this._io.ReadBytesTerm(10, false, true, true)
	if err != nil {
		return err
	}
	this.Str = string(tmp33)
	return err
}

type BsZen_TypeWord struct {
	Value   []byte
	_io     *kaitai.Stream
	_root   *BsZen
	_parent *BsZen_Property
}

func NewBsZen_TypeWord() *BsZen_TypeWord {
	return &BsZen_TypeWord{}
}

func (this *BsZen_TypeWord) Read(io *kaitai.Stream, parent *BsZen_Property, root *BsZen) (err error) {
	this._io = io
	this._parent = parent
	this._root = root

	tmp34, err := this._io.ReadBytes(int(2))
	if err != nil {
		return err
	}
	tmp34 = tmp34
	this.Value = tmp34
	return err
}

type BsZen_TypeInteger struct {
	Value   uint32
	_io     *kaitai.Stream
	_root   *BsZen
	_parent *BsZen_Property
}

func NewBsZen_TypeInteger() *BsZen_TypeInteger {
	return &BsZen_TypeInteger{}
}

func (this *BsZen_TypeInteger) Read(io *kaitai.Stream, parent *BsZen_Property, root *BsZen) (err error) {
	this._io = io
	this._parent = parent
	this._root = root

	tmp35, err := this._io.ReadU4le()
	if err != nil {
		return err
	}
	this.Value = uint32(tmp35)
	return err
}

type BsZen_TypeString struct {
	Size     uint16
	ObjectId string
	_io      *kaitai.Stream
	_root    *BsZen
	_parent  *BsZen_Property
}

func NewBsZen_TypeString() *BsZen_TypeString {
	return &BsZen_TypeString{}
}

func (this *BsZen_TypeString) Read(io *kaitai.Stream, parent *BsZen_Property, root *BsZen) (err error) {
	this._io = io
	this._parent = parent
	this._root = root

	tmp36, err := this._io.ReadU2le()
	if err != nil {
		return err
	}
	this.Size = uint16(tmp36)
	tmp37, err := this._io.ReadBytes(int(this.Size))
	if err != nil {
		return err
	}
	tmp37 = tmp37
	this.ObjectId = string(tmp37)
	return err
}

type BsZen_Header struct {
	Magic         *BsZen_String
	Version       *BsZen_String
	ArchiverClass *BsZen_String
	ArchiverType  *BsZen_String
	SaveGame      *BsZen_String
	Date          *BsZen_String
	User          *BsZen_String
	End           *BsZen_String
	_io           *kaitai.Stream
	_root         *BsZen
	_parent       *BsZen
}

func NewBsZen_Header() *BsZen_Header {
	return &BsZen_Header{}
}

func (this *BsZen_Header) Read(io *kaitai.Stream, parent *BsZen, root *BsZen) (err error) {
	this._io = io
	this._parent = parent
	this._root = root

	tmp38 := NewBsZen_String()
	err = tmp38.Read(this._io, this, this._root)
	if err != nil {
		return err
	}
	this.Magic = tmp38
	tmp39 := NewBsZen_String()
	err = tmp39.Read(this._io, this, this._root)
	if err != nil {
		return err
	}
	this.Version = tmp39
	tmp40 := NewBsZen_String()
	err = tmp40.Read(this._io, this, this._root)
	if err != nil {
		return err
	}
	this.ArchiverClass = tmp40
	tmp41 := NewBsZen_String()
	err = tmp41.Read(this._io, this, this._root)
	if err != nil {
		return err
	}
	this.ArchiverType = tmp41
	tmp42 := NewBsZen_String()
	err = tmp42.Read(this._io, this, this._root)
	if err != nil {
		return err
	}
	this.SaveGame = tmp42
	tmp43 := NewBsZen_String()
	err = tmp43.Read(this._io, this, this._root)
	if err != nil {
		return err
	}
	this.Date = tmp43
	tmp44 := NewBsZen_String()
	err = tmp44.Read(this._io, this, this._root)
	if err != nil {
		return err
	}
	this.User = tmp44
	tmp45 := NewBsZen_String()
	err = tmp45.Read(this._io, this, this._root)
	if err != nil {
		return err
	}
	this.End = tmp45
	return err
}

type BsZen_Property struct {
	Type     uint8
	PropBody interface{}
	_io      *kaitai.Stream
	_root    *BsZen
	_parent  interface{}
}

func NewBsZen_Property() *BsZen_Property {
	return &BsZen_Property{}
}

func (this *BsZen_Property) Read(io *kaitai.Stream, parent interface{}, root *BsZen) (err error) {
	this._io = io
	this._parent = parent
	this._root = root

	tmp46, err := this._io.ReadU1()
	if err != nil {
		return err
	}
	this.Type = tmp46
	switch this.Type {
	case 17:
		tmp47 := NewBsZen_TypeEnum()
		err = tmp47.Read(this._io, this, this._root)
		if err != nil {
			return err
		}
		this.PropBody = tmp47
	case 4:
		tmp48 := NewBsZen_TypeByte()
		err = tmp48.Read(this._io, this, this._root)
		if err != nil {
			return err
		}
		this.PropBody = tmp48
	case 6:
		tmp49 := NewBsZen_TypeBool()
		err = tmp49.Read(this._io, this, this._root)
		if err != nil {
			return err
		}
		this.PropBody = tmp49
	case 7:
		tmp50 := NewBsZen_TypeVec3()
		err = tmp50.Read(this._io, this, this._root)
		if err != nil {
			return err
		}
		this.PropBody = tmp50
	case 1:
		tmp51 := NewBsZen_TypeString()
		err = tmp51.Read(this._io, this, this._root)
		if err != nil {
			return err
		}
		this.PropBody = tmp51
	case 3:
		tmp52 := NewBsZen_TypeFloat()
		err = tmp52.Read(this._io, this, this._root)
		if err != nil {
			return err
		}
		this.PropBody = tmp52
	case 5:
		tmp53 := NewBsZen_TypeWord()
		err = tmp53.Read(this._io, this, this._root)
		if err != nil {
			return err
		}
		this.PropBody = tmp53
	case 8:
		tmp54 := NewBsZen_TypeColor()
		err = tmp54.Read(this._io, this, this._root)
		if err != nil {
			return err
		}
		this.PropBody = tmp54
	case 9:
		tmp55 := NewBsZen_TypeRaw()
		err = tmp55.Read(this._io, this, this._root)
		if err != nil {
			return err
		}
		this.PropBody = tmp55
	case 16:
		tmp56 := NewBsZen_TypeRawFloat()
		err = tmp56.Read(this._io, this, this._root)
		if err != nil {
			return err
		}
		this.PropBody = tmp56
	case 18:
		tmp57 := NewBsZen_TypeHash()
		err = tmp57.Read(this._io, this, this._root)
		if err != nil {
			return err
		}
		this.PropBody = tmp57
	case 2:
		tmp58 := NewBsZen_TypeInteger()
		err = tmp58.Read(this._io, this, this._root)
		if err != nil {
			return err
		}
		this.PropBody = tmp58
	}
	return err
}

type BsZen_TypeVec3 struct {
	X       float32
	Y       float32
	Z       float32
	_io     *kaitai.Stream
	_root   *BsZen
	_parent *BsZen_Property
}

func NewBsZen_TypeVec3() *BsZen_TypeVec3 {
	return &BsZen_TypeVec3{}
}

func (this *BsZen_TypeVec3) Read(io *kaitai.Stream, parent *BsZen_Property, root *BsZen) (err error) {
	this._io = io
	this._parent = parent
	this._root = root

	tmp59, err := this._io.ReadF4le()
	if err != nil {
		return err
	}
	this.X = float32(tmp59)
	tmp60, err := this._io.ReadF4le()
	if err != nil {
		return err
	}
	this.Y = float32(tmp60)
	tmp61, err := this._io.ReadF4le()
	if err != nil {
		return err
	}
	this.Z = float32(tmp61)
	return err
}

type BsZen_TypeHash struct {
	Value   uint32
	_io     *kaitai.Stream
	_root   *BsZen
	_parent *BsZen_Property
}

func NewBsZen_TypeHash() *BsZen_TypeHash {
	return &BsZen_TypeHash{}
}

func (this *BsZen_TypeHash) Read(io *kaitai.Stream, parent *BsZen_Property, root *BsZen) (err error) {
	this._io = io
	this._parent = parent
	this._root = root

	tmp62, err := this._io.ReadU4le()
	if err != nil {
		return err
	}
	this.Value = uint32(tmp62)
	return err
}

type BsZen_TypeByte struct {
	Value   []byte
	_io     *kaitai.Stream
	_root   *BsZen
	_parent *BsZen_Property
}

func NewBsZen_TypeByte() *BsZen_TypeByte {
	return &BsZen_TypeByte{}
}

func (this *BsZen_TypeByte) Read(io *kaitai.Stream, parent *BsZen_Property, root *BsZen) (err error) {
	this._io = io
	this._parent = parent
	this._root = root

	tmp63, err := this._io.ReadBytes(int(1))
	if err != nil {
		return err
	}
	tmp63 = tmp63
	this.Value = tmp63
	return err
}

type BsZen_HashTableHeader struct {
	ChunkCount uint32
	Chunks     []*BsZen_HashTableChunk
	_io        *kaitai.Stream
	_root      *BsZen
	_parent    *BsZen_World
}

func NewBsZen_HashTableHeader() *BsZen_HashTableHeader {
	return &BsZen_HashTableHeader{}
}

func (this *BsZen_HashTableHeader) Read(io *kaitai.Stream, parent *BsZen_World, root *BsZen) (err error) {
	this._io = io
	this._parent = parent
	this._root = root

	tmp64, err := this._io.ReadU4le()
	if err != nil {
		return err
	}
	this.ChunkCount = uint32(tmp64)
	for i := 0; i < int(this.ChunkCount); i++ {
		_ = i
		tmp65 := NewBsZen_HashTableChunk()
		err = tmp65.Read(this._io, this, this._root)
		if err != nil {
			return err
		}
		this.Chunks = append(this.Chunks, tmp65)
	}
	return err
}
