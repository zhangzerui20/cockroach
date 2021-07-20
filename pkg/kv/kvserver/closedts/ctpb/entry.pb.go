// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: kv/kvserver/closedts/ctpb/entry.proto

package ctpb

import (
	fmt "fmt"
	github_com_cockroachdb_cockroach_pkg_roachpb "github.com/cockroachdb/cockroach/pkg/roachpb"
	hlc "github.com/cockroachdb/cockroach/pkg/util/hlc"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_sortkeys "github.com/gogo/protobuf/sortkeys"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// An Entry is a closed timestamp update. It consists of a closed timestamp
// (i.e. a timestamp at or below which the origin node guarantees no more new
// writes are going to be permitted), an associated epoch in which the origin
// node promises it was live (for the closed timestamp), a map of minimum lease
// applied indexes (which have to be caught up to before being allowed to use
// the closed timestamp) as well as an indicator of whether this update supplies
// a full initial state or an increment to be merged into a previous state. In
// practice, the first Entry received for each epoch is full, while the remainder
// are incremental. An incremental update represents the implicit promise that
// the state accumulated since the last full Entry is the true full state.
type Entry struct {
	Epoch           Epoch                                                        `protobuf:"varint,1,opt,name=epoch,proto3,casttype=Epoch" json:"epoch,omitempty"`
	ClosedTimestamp hlc.Timestamp                                                `protobuf:"bytes,2,opt,name=closed_timestamp,json=closedTimestamp,proto3" json:"closed_timestamp"`
	MLAI            map[github_com_cockroachdb_cockroach_pkg_roachpb.RangeID]LAI `protobuf:"bytes,3,rep,name=mlai,proto3,castkey=github.com/cockroachdb/cockroach/pkg/roachpb.RangeID,castvalue=LAI" json:"mlai,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	// Full is true if the emitter promises that any future write to any range
	// mentioned in this Entry will be reflected in a subsequent Entry before any
	// stale follower reads are possible. For example, if range 1 is assigned an
	// MLAI of 12 in this Entry and isn't mentioned in the five subsequent
	// entries, the recipient may behave as if the MLAI of 12 were repeated across
	// all of these entries.
	//
	// In practice, a Full message is received when a stream of Entries is first
	// established (or the Epoch changes), and all other updates are incremental
	// (i.e. not Full).
	Full bool `protobuf:"varint,4,opt,name=full,proto3" json:"full,omitempty"`
}

func (m *Entry) Reset()      { *m = Entry{} }
func (*Entry) ProtoMessage() {}
func (*Entry) Descriptor() ([]byte, []int) {
	return fileDescriptor_db146746651382e6, []int{0}
}
func (m *Entry) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Entry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *Entry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Entry.Merge(m, src)
}
func (m *Entry) XXX_Size() int {
	return m.Size()
}
func (m *Entry) XXX_DiscardUnknown() {
	xxx_messageInfo_Entry.DiscardUnknown(m)
}

var xxx_messageInfo_Entry proto.InternalMessageInfo

// Reactions flow in the direction opposite to Entries and request for ranges to
// be included in the next Entry. Under rare circumstances, ranges may be omitted
// from closed timestamp updates, and so serving follower reads from them would
// fail. The Reaction mechanism serves to explicitly request the missing information
// when that happens.
type Reaction struct {
	Requested []github_com_cockroachdb_cockroach_pkg_roachpb.RangeID `protobuf:"varint,1,rep,packed,name=Requested,proto3,casttype=github.com/cockroachdb/cockroach/pkg/roachpb.RangeID" json:"Requested,omitempty"`
}

func (m *Reaction) Reset()      { *m = Reaction{} }
func (*Reaction) ProtoMessage() {}
func (*Reaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_db146746651382e6, []int{1}
}
func (m *Reaction) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Reaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *Reaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Reaction.Merge(m, src)
}
func (m *Reaction) XXX_Size() int {
	return m.Size()
}
func (m *Reaction) XXX_DiscardUnknown() {
	xxx_messageInfo_Reaction.DiscardUnknown(m)
}

var xxx_messageInfo_Reaction proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Entry)(nil), "cockroach.kv.kvserver.ctupdate.Entry")
	proto.RegisterMapType((map[github_com_cockroachdb_cockroach_pkg_roachpb.RangeID]LAI)(nil), "cockroach.kv.kvserver.ctupdate.Entry.MlaiEntry")
	proto.RegisterType((*Reaction)(nil), "cockroach.kv.kvserver.ctupdate.Reaction")
}

func init() {
	proto.RegisterFile("kv/kvserver/closedts/ctpb/entry.proto", fileDescriptor_db146746651382e6)
}

var fileDescriptor_db146746651382e6 = []byte{
	// 442 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x3f, 0x6f, 0xd4, 0x30,
	0x18, 0xc6, 0xe3, 0x4b, 0x82, 0x7a, 0xee, 0x40, 0x15, 0x75, 0x88, 0x4e, 0xe0, 0x44, 0x95, 0x40,
	0x99, 0x6c, 0xa9, 0x20, 0x51, 0x75, 0x6b, 0x44, 0x85, 0x4e, 0x6a, 0x19, 0x2c, 0xc4, 0xc0, 0x82,
	0x1c, 0x9f, 0x49, 0xa2, 0xf8, 0xe2, 0x90, 0x38, 0x91, 0xba, 0x32, 0x22, 0x06, 0x46, 0x46, 0x3e,
	0xce, 0x8d, 0x1d, 0x2b, 0x86, 0x2b, 0xe4, 0xbe, 0x45, 0x27, 0x14, 0x87, 0xde, 0x6d, 0x0c, 0xdd,
	0x7e, 0x79, 0xff, 0x3d, 0x4f, 0xde, 0xd7, 0xf0, 0x59, 0xd1, 0x91, 0xa2, 0x6b, 0x44, 0xdd, 0x89,
	0x9a, 0x70, 0xa9, 0x1a, 0xb1, 0xd0, 0x0d, 0xe1, 0xba, 0x4a, 0x88, 0x28, 0x75, 0x7d, 0x85, 0xab,
	0x5a, 0x69, 0xe5, 0x21, 0xae, 0x78, 0x51, 0x2b, 0xc6, 0x33, 0x5c, 0x74, 0xf8, 0xbe, 0x01, 0x73,
	0xdd, 0x56, 0x0b, 0xa6, 0xc5, 0xec, 0x30, 0x55, 0xa9, 0x32, 0xa5, 0x64, 0xa0, 0xb1, 0x6b, 0xf6,
	0x24, 0x55, 0x2a, 0x95, 0x82, 0xb0, 0x2a, 0x27, 0xac, 0x2c, 0x95, 0x66, 0x3a, 0x57, 0x65, 0xf3,
	0x2f, 0xeb, 0xb7, 0x3a, 0x97, 0x24, 0x93, 0x9c, 0xe8, 0x7c, 0x29, 0x1a, 0xcd, 0x96, 0xd5, 0x98,
	0x39, 0xfa, 0x35, 0x81, 0xee, 0xf9, 0xa0, 0xee, 0x05, 0xd0, 0x15, 0x95, 0xe2, 0x99, 0x0f, 0x42,
	0x10, 0xd9, 0xf1, 0xf4, 0x6e, 0x1d, 0xb8, 0xe7, 0x43, 0x80, 0x8e, 0x71, 0xef, 0x2d, 0x3c, 0x18,
	0x5d, 0x7f, 0xdc, 0x0e, 0xf1, 0x27, 0x21, 0x88, 0xf6, 0x8f, 0x9f, 0xe2, 0x9d, 0xe7, 0x41, 0x09,
	0x67, 0x92, 0xe3, 0x77, 0xf7, 0x45, 0xb1, 0xb3, 0x5a, 0x07, 0x16, 0x7d, 0x3c, 0x36, 0x6f, 0xc3,
	0xde, 0x37, 0x00, 0x9d, 0xa5, 0x64, 0xb9, 0x6f, 0x87, 0x76, 0xb4, 0x7f, 0x4c, 0xf0, 0xff, 0x7f,
	0x1c, 0x1b, 0x9b, 0xf8, 0x52, 0xb2, 0xdc, 0x50, 0xfc, 0xa6, 0x5f, 0x07, 0xce, 0xe5, 0xc5, 0xd9,
	0xfc, 0xcb, 0x6d, 0xf0, 0x32, 0xcd, 0x75, 0xd6, 0x26, 0x98, 0xab, 0x25, 0xd9, 0x8e, 0x59, 0x24,
	0x3b, 0x26, 0x55, 0x91, 0x12, 0x43, 0x55, 0x82, 0x29, 0x2b, 0x53, 0x31, 0x7f, 0xfd, 0xf5, 0x36,
	0xb0, 0x2f, 0xce, 0xe6, 0xd4, 0xb8, 0xf0, 0x3c, 0xe8, 0x7c, 0x6a, 0xa5, 0xf4, 0x9d, 0x10, 0x44,
	0x7b, 0xd4, 0xf0, 0xec, 0x15, 0x9c, 0x6e, 0xf5, 0xbc, 0x03, 0x68, 0x17, 0xe2, 0xca, 0xac, 0xc7,
	0xa5, 0x03, 0x7a, 0x87, 0xd0, 0xed, 0x98, 0x6c, 0x85, 0x59, 0x83, 0x4d, 0xc7, 0x8f, 0xd3, 0xc9,
	0x09, 0x38, 0x75, 0x7e, 0xfc, 0x0c, 0xac, 0xa3, 0x0c, 0xee, 0x51, 0xc1, 0xf8, 0x70, 0x09, 0xef,
	0x3d, 0x9c, 0x52, 0xf1, 0xb9, 0x15, 0x8d, 0x16, 0x0b, 0x1f, 0x84, 0x76, 0xe4, 0xc6, 0x27, 0x77,
	0xeb, 0x87, 0x19, 0xa7, 0xbb, 0x51, 0xa3, 0x52, 0xfc, 0x7c, 0xf5, 0x07, 0x59, 0xab, 0x1e, 0x81,
	0xeb, 0x1e, 0x81, 0x9b, 0x1e, 0x81, 0xdf, 0x3d, 0x02, 0xdf, 0x37, 0xc8, 0xba, 0xde, 0x20, 0xeb,
	0x66, 0x83, 0xac, 0x0f, 0xce, 0xf0, 0xd0, 0x92, 0x47, 0xe6, 0xea, 0x2f, 0xfe, 0x06, 0x00, 0x00,
	0xff, 0xff, 0x39, 0x68, 0xb2, 0x7f, 0x8c, 0x02, 0x00, 0x00,
}

func (m *Entry) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Entry) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Entry) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Full {
		i--
		if m.Full {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if len(m.MLAI) > 0 {
		keysForMLAI := make([]int32, 0, len(m.MLAI))
		for k := range m.MLAI {
			keysForMLAI = append(keysForMLAI, int32(k))
		}
		github_com_gogo_protobuf_sortkeys.Int32s(keysForMLAI)
		for iNdEx := len(keysForMLAI) - 1; iNdEx >= 0; iNdEx-- {
			v := m.MLAI[github_com_cockroachdb_cockroach_pkg_roachpb.RangeID(keysForMLAI[iNdEx])]
			baseI := i
			i = encodeVarintEntry(dAtA, i, uint64(v))
			i--
			dAtA[i] = 0x10
			i = encodeVarintEntry(dAtA, i, uint64(keysForMLAI[iNdEx]))
			i--
			dAtA[i] = 0x8
			i = encodeVarintEntry(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x1a
		}
	}
	{
		size, err := m.ClosedTimestamp.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintEntry(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.Epoch != 0 {
		i = encodeVarintEntry(dAtA, i, uint64(m.Epoch))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Reaction) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Reaction) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Reaction) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Requested) > 0 {
		dAtA3 := make([]byte, len(m.Requested)*10)
		var j2 int
		for _, num1 := range m.Requested {
			num := uint64(num1)
			for num >= 1<<7 {
				dAtA3[j2] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j2++
			}
			dAtA3[j2] = uint8(num)
			j2++
		}
		i -= j2
		copy(dAtA[i:], dAtA3[:j2])
		i = encodeVarintEntry(dAtA, i, uint64(j2))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintEntry(dAtA []byte, offset int, v uint64) int {
	offset -= sovEntry(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Entry) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Epoch != 0 {
		n += 1 + sovEntry(uint64(m.Epoch))
	}
	l = m.ClosedTimestamp.Size()
	n += 1 + l + sovEntry(uint64(l))
	if len(m.MLAI) > 0 {
		for k, v := range m.MLAI {
			_ = k
			_ = v
			mapEntrySize := 1 + sovEntry(uint64(k)) + 1 + sovEntry(uint64(v))
			n += mapEntrySize + 1 + sovEntry(uint64(mapEntrySize))
		}
	}
	if m.Full {
		n += 2
	}
	return n
}

func (m *Reaction) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Requested) > 0 {
		l = 0
		for _, e := range m.Requested {
			l += sovEntry(uint64(e))
		}
		n += 1 + sovEntry(uint64(l)) + l
	}
	return n
}

func sovEntry(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEntry(x uint64) (n int) {
	return sovEntry(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Entry) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEntry
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Entry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Entry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Epoch", wireType)
			}
			m.Epoch = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEntry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Epoch |= Epoch(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClosedTimestamp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEntry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEntry
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEntry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ClosedTimestamp.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MLAI", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEntry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEntry
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEntry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.MLAI == nil {
				m.MLAI = make(map[github_com_cockroachdb_cockroach_pkg_roachpb.RangeID]LAI)
			}
			var mapkey int32
			var mapvalue int64
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowEntry
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					wire |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				fieldNum := int32(wire >> 3)
				if fieldNum == 1 {
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowEntry
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapkey |= int32(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
				} else if fieldNum == 2 {
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowEntry
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapvalue |= int64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipEntry(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthEntry
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.MLAI[github_com_cockroachdb_cockroach_pkg_roachpb.RangeID(mapkey)] = ((LAI)(mapvalue))
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Full", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEntry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Full = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipEntry(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEntry
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Reaction) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEntry
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Reaction: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Reaction: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType == 0 {
				var v github_com_cockroachdb_cockroach_pkg_roachpb.RangeID
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowEntry
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= github_com_cockroachdb_cockroach_pkg_roachpb.RangeID(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Requested = append(m.Requested, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowEntry
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthEntry
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthEntry
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.Requested) == 0 {
					m.Requested = make([]github_com_cockroachdb_cockroach_pkg_roachpb.RangeID, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v github_com_cockroachdb_cockroach_pkg_roachpb.RangeID
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowEntry
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= github_com_cockroachdb_cockroach_pkg_roachpb.RangeID(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Requested = append(m.Requested, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Requested", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEntry(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEntry
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipEntry(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEntry
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowEntry
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowEntry
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthEntry
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEntry
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEntry
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEntry        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEntry          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEntry = fmt.Errorf("proto: unexpected end of group")
)
