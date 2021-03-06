// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/containerd/containerd/api/services/diff/v1/diff.proto

package diff

import (
	context "context"
	fmt "fmt"
	types "github.com/containerd/containerd/api/types"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_sortkeys "github.com/gogo/protobuf/sortkeys"
	grpc "google.golang.org/grpc"
	io "io"
	math "math"
	reflect "reflect"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type ApplyRequest struct {
	// Diff is the descriptor of the diff to be extracted
	Diff                 *types.Descriptor `protobuf:"bytes,1,opt,name=diff,proto3" json:"diff,omitempty"`
	Mounts               []*types.Mount    `protobuf:"bytes,2,rep,name=mounts,proto3" json:"mounts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ApplyRequest) Reset()      { *m = ApplyRequest{} }
func (*ApplyRequest) ProtoMessage() {}
func (*ApplyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b36a99e6faaa935, []int{0}
}
func (m *ApplyRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ApplyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ApplyRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ApplyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApplyRequest.Merge(m, src)
}
func (m *ApplyRequest) XXX_Size() int {
	return m.Size()
}
func (m *ApplyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ApplyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ApplyRequest proto.InternalMessageInfo

type ApplyResponse struct {
	// Applied is the descriptor for the object which was applied.
	// If the input was a compressed blob then the result will be
	// the descriptor for the uncompressed blob.
	Applied              *types.Descriptor `protobuf:"bytes,1,opt,name=applied,proto3" json:"applied,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ApplyResponse) Reset()      { *m = ApplyResponse{} }
func (*ApplyResponse) ProtoMessage() {}
func (*ApplyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b36a99e6faaa935, []int{1}
}
func (m *ApplyResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ApplyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ApplyResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ApplyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApplyResponse.Merge(m, src)
}
func (m *ApplyResponse) XXX_Size() int {
	return m.Size()
}
func (m *ApplyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ApplyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ApplyResponse proto.InternalMessageInfo

type DiffRequest struct {
	// Left are the mounts which represent the older copy
	// in which is the base of the computed changes.
	Left []*types.Mount `protobuf:"bytes,1,rep,name=left,proto3" json:"left,omitempty"`
	// Right are the mounts which represents the newer copy
	// in which changes from the left were made into.
	Right []*types.Mount `protobuf:"bytes,2,rep,name=right,proto3" json:"right,omitempty"`
	// MediaType is the media type descriptor for the created diff
	// object
	MediaType string `protobuf:"bytes,3,opt,name=media_type,json=mediaType,proto3" json:"media_type,omitempty"`
	// Ref identifies the pre-commit content store object. This
	// reference can be used to get the status from the content store.
	Ref string `protobuf:"bytes,4,opt,name=ref,proto3" json:"ref,omitempty"`
	// Labels are the labels to apply to the generated content
	// on content store commit.
	Labels               map[string]string `protobuf:"bytes,5,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *DiffRequest) Reset()      { *m = DiffRequest{} }
func (*DiffRequest) ProtoMessage() {}
func (*DiffRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b36a99e6faaa935, []int{2}
}
func (m *DiffRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DiffRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DiffRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DiffRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DiffRequest.Merge(m, src)
}
func (m *DiffRequest) XXX_Size() int {
	return m.Size()
}
func (m *DiffRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DiffRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DiffRequest proto.InternalMessageInfo

type DiffResponse struct {
	// Diff is the descriptor of the diff which can be applied
	Diff                 *types.Descriptor `protobuf:"bytes,3,opt,name=diff,proto3" json:"diff,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *DiffResponse) Reset()      { *m = DiffResponse{} }
func (*DiffResponse) ProtoMessage() {}
func (*DiffResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b36a99e6faaa935, []int{3}
}
func (m *DiffResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DiffResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DiffResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DiffResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DiffResponse.Merge(m, src)
}
func (m *DiffResponse) XXX_Size() int {
	return m.Size()
}
func (m *DiffResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DiffResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DiffResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ApplyRequest)(nil), "containerd.services.diff.v1.ApplyRequest")
	proto.RegisterType((*ApplyResponse)(nil), "containerd.services.diff.v1.ApplyResponse")
	proto.RegisterType((*DiffRequest)(nil), "containerd.services.diff.v1.DiffRequest")
	proto.RegisterMapType((map[string]string)(nil), "containerd.services.diff.v1.DiffRequest.LabelsEntry")
	proto.RegisterType((*DiffResponse)(nil), "containerd.services.diff.v1.DiffResponse")
}

func init() {
	proto.RegisterFile("github.com/containerd/containerd/api/services/diff/v1/diff.proto", fileDescriptor_3b36a99e6faaa935)
}

var fileDescriptor_3b36a99e6faaa935 = []byte{
	// 457 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0x4f, 0x6f, 0xd3, 0x30,
	0x14, 0xaf, 0xfb, 0x0f, 0xf5, 0x75, 0x48, 0xc8, 0x9a, 0x44, 0x14, 0x20, 0xaa, 0x7a, 0xea, 0x40,
	0x38, 0xac, 0xa0, 0x09, 0xb6, 0xcb, 0x40, 0x43, 0x5c, 0xc6, 0x25, 0xda, 0x01, 0x81, 0x04, 0x4a,
	0x9b, 0x97, 0xce, 0x22, 0x8d, 0xbd, 0xd8, 0xad, 0x94, 0x1b, 0xdf, 0x85, 0x8f, 0xc2, 0x65, 0x47,
	0x8e, 0x1c, 0x69, 0x3f, 0x09, 0xb2, 0x93, 0x40, 0x24, 0xa4, 0x12, 0x76, 0xca, 0xcb, 0xf3, 0xef,
	0x9f, 0xfd, 0x6c, 0x38, 0x5d, 0x70, 0x7d, 0xb9, 0x9a, 0xb1, 0xb9, 0x58, 0xfa, 0x73, 0x91, 0xea,
	0x90, 0xa7, 0x98, 0x45, 0xf5, 0x32, 0x94, 0xdc, 0x57, 0x98, 0xad, 0xf9, 0x1c, 0x95, 0x1f, 0xf1,
	0x38, 0xf6, 0xd7, 0x87, 0xf6, 0xcb, 0x64, 0x26, 0xb4, 0xa0, 0xf7, 0xfe, 0x60, 0x59, 0x85, 0x63,
	0x76, 0x7d, 0x7d, 0xe8, 0xee, 0x2f, 0xc4, 0x42, 0x58, 0x9c, 0x6f, 0xaa, 0x82, 0xe2, 0x1e, 0x35,
	0x32, 0xd5, 0xb9, 0x44, 0xe5, 0x2f, 0xc5, 0x2a, 0xd5, 0x25, 0xef, 0xe4, 0x3f, 0x78, 0x11, 0xaa,
	0x79, 0xc6, 0xa5, 0x16, 0x59, 0x41, 0x1e, 0x5f, 0xc1, 0xde, 0x4b, 0x29, 0x93, 0x3c, 0xc0, 0xab,
	0x15, 0x2a, 0x4d, 0x9f, 0x40, 0xd7, 0xa4, 0x74, 0xc8, 0x88, 0x4c, 0x86, 0xd3, 0xfb, 0xac, 0xb6,
	0x0d, 0xab, 0xc0, 0xce, 0x7e, 0x2b, 0x04, 0x16, 0x49, 0x7d, 0xe8, 0xdb, 0x34, 0xca, 0x69, 0x8f,
	0x3a, 0x93, 0xe1, 0xf4, 0xee, 0xdf, 0x9c, 0xb7, 0x66, 0x3d, 0x28, 0x61, 0xe3, 0x37, 0x70, 0xbb,
	0xb4, 0x54, 0x52, 0xa4, 0x0a, 0xe9, 0x11, 0xdc, 0x0a, 0xa5, 0x4c, 0x38, 0x46, 0x8d, 0x6c, 0x2b,
	0xf0, 0xf8, 0x6b, 0x1b, 0x86, 0x67, 0x3c, 0x8e, 0xab, 0xec, 0x8f, 0xa0, 0x9b, 0x60, 0xac, 0x1d,
	0xb2, 0x3b, 0x87, 0x05, 0xd1, 0xc7, 0xd0, 0xcb, 0xf8, 0xe2, 0x52, 0xff, 0x2b, 0x75, 0x81, 0xa2,
	0x0f, 0x00, 0x96, 0x18, 0xf1, 0xf0, 0x93, 0x59, 0x73, 0x3a, 0x23, 0x32, 0x19, 0x04, 0x03, 0xdb,
	0xb9, 0xc8, 0x25, 0xd2, 0x3b, 0xd0, 0xc9, 0x30, 0x76, 0xba, 0xb6, 0x6f, 0x4a, 0x7a, 0x0e, 0xfd,
	0x24, 0x9c, 0x61, 0xa2, 0x9c, 0x9e, 0x35, 0x78, 0xc6, 0x76, 0xdc, 0x08, 0x56, 0xdb, 0x06, 0x3b,
	0xb7, 0xb4, 0xd7, 0xa9, 0xce, 0xf2, 0xa0, 0xd4, 0x70, 0x5f, 0xc0, 0xb0, 0xd6, 0x36, 0x76, 0x9f,
	0x31, 0xb7, 0xa7, 0x35, 0x08, 0x4c, 0x49, 0xf7, 0xa1, 0xb7, 0x0e, 0x93, 0x15, 0x3a, 0x6d, 0xdb,
	0x2b, 0x7e, 0x8e, 0xdb, 0xcf, 0xc9, 0xf8, 0x14, 0xf6, 0x0a, 0xf5, 0xf2, 0xb4, 0xab, 0x09, 0x77,
	0x9a, 0x4e, 0x78, 0xfa, 0x8d, 0x40, 0xd7, 0x48, 0xd0, 0x8f, 0xd0, 0xb3, 0x93, 0xa3, 0x07, 0x3b,
	0x37, 0x53, 0xbf, 0x50, 0xee, 0xc3, 0x26, 0xd0, 0x32, 0xda, 0x87, 0xd2, 0x67, 0xd2, 0xf4, 0xac,
	0xdc, 0x83, 0x06, 0xc8, 0x42, 0xfc, 0xd5, 0xc5, 0xf5, 0xc6, 0x6b, 0xfd, 0xd8, 0x78, 0xad, 0x2f,
	0x5b, 0x8f, 0x5c, 0x6f, 0x3d, 0xf2, 0x7d, 0xeb, 0x91, 0x9f, 0x5b, 0x8f, 0xbc, 0x3f, 0xbe, 0xd1,
	0x6b, 0x3f, 0x31, 0xdf, 0x77, 0xad, 0x59, 0xdf, 0x3e, 0xa4, 0xa7, 0xbf, 0x02, 0x00, 0x00, 0xff,
	0xff, 0x61, 0xd1, 0x6e, 0x9e, 0x34, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DiffClient is the client API for Diff service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DiffClient interface {
	// Apply applies the content associated with the provided digests onto
	// the provided mounts. Archive content will be extracted and
	// decompressed if necessary.
	Apply(ctx context.Context, in *ApplyRequest, opts ...grpc.CallOption) (*ApplyResponse, error)
	// Diff creates a diff between the given mounts and uploads the result
	// to the content store.
	Diff(ctx context.Context, in *DiffRequest, opts ...grpc.CallOption) (*DiffResponse, error)
}

type diffClient struct {
	cc *grpc.ClientConn
}

func NewDiffClient(cc *grpc.ClientConn) DiffClient {
	return &diffClient{cc}
}

func (c *diffClient) Apply(ctx context.Context, in *ApplyRequest, opts ...grpc.CallOption) (*ApplyResponse, error) {
	out := new(ApplyResponse)
	err := c.cc.Invoke(ctx, "/containerd.services.diff.v1.Diff/Apply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *diffClient) Diff(ctx context.Context, in *DiffRequest, opts ...grpc.CallOption) (*DiffResponse, error) {
	out := new(DiffResponse)
	err := c.cc.Invoke(ctx, "/containerd.services.diff.v1.Diff/Diff", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DiffServer is the server API for Diff service.
type DiffServer interface {
	// Apply applies the content associated with the provided digests onto
	// the provided mounts. Archive content will be extracted and
	// decompressed if necessary.
	Apply(context.Context, *ApplyRequest) (*ApplyResponse, error)
	// Diff creates a diff between the given mounts and uploads the result
	// to the content store.
	Diff(context.Context, *DiffRequest) (*DiffResponse, error)
}

func RegisterDiffServer(s *grpc.Server, srv DiffServer) {
	s.RegisterService(&_Diff_serviceDesc, srv)
}

func _Diff_Apply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiffServer).Apply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/containerd.services.diff.v1.Diff/Apply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiffServer).Apply(ctx, req.(*ApplyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Diff_Diff_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DiffRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiffServer).Diff(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/containerd.services.diff.v1.Diff/Diff",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiffServer).Diff(ctx, req.(*DiffRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Diff_serviceDesc = grpc.ServiceDesc{
	ServiceName: "containerd.services.diff.v1.Diff",
	HandlerType: (*DiffServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Apply",
			Handler:    _Diff_Apply_Handler,
		},
		{
			MethodName: "Diff",
			Handler:    _Diff_Diff_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/containerd/containerd/api/services/diff/v1/diff.proto",
}

func (m *ApplyRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ApplyRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Diff != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintDiff(dAtA, i, uint64(m.Diff.Size()))
		n1, err := m.Diff.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if len(m.Mounts) > 0 {
		for _, msg := range m.Mounts {
			dAtA[i] = 0x12
			i++
			i = encodeVarintDiff(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *ApplyResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ApplyResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Applied != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintDiff(dAtA, i, uint64(m.Applied.Size()))
		n2, err := m.Applied.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *DiffRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DiffRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Left) > 0 {
		for _, msg := range m.Left {
			dAtA[i] = 0xa
			i++
			i = encodeVarintDiff(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.Right) > 0 {
		for _, msg := range m.Right {
			dAtA[i] = 0x12
			i++
			i = encodeVarintDiff(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.MediaType) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintDiff(dAtA, i, uint64(len(m.MediaType)))
		i += copy(dAtA[i:], m.MediaType)
	}
	if len(m.Ref) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintDiff(dAtA, i, uint64(len(m.Ref)))
		i += copy(dAtA[i:], m.Ref)
	}
	if len(m.Labels) > 0 {
		for k, _ := range m.Labels {
			dAtA[i] = 0x2a
			i++
			v := m.Labels[k]
			mapSize := 1 + len(k) + sovDiff(uint64(len(k))) + 1 + len(v) + sovDiff(uint64(len(v)))
			i = encodeVarintDiff(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintDiff(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			dAtA[i] = 0x12
			i++
			i = encodeVarintDiff(dAtA, i, uint64(len(v)))
			i += copy(dAtA[i:], v)
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *DiffResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DiffResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Diff != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintDiff(dAtA, i, uint64(m.Diff.Size()))
		n3, err := m.Diff.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintDiff(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *ApplyRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Diff != nil {
		l = m.Diff.Size()
		n += 1 + l + sovDiff(uint64(l))
	}
	if len(m.Mounts) > 0 {
		for _, e := range m.Mounts {
			l = e.Size()
			n += 1 + l + sovDiff(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ApplyResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Applied != nil {
		l = m.Applied.Size()
		n += 1 + l + sovDiff(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *DiffRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Left) > 0 {
		for _, e := range m.Left {
			l = e.Size()
			n += 1 + l + sovDiff(uint64(l))
		}
	}
	if len(m.Right) > 0 {
		for _, e := range m.Right {
			l = e.Size()
			n += 1 + l + sovDiff(uint64(l))
		}
	}
	l = len(m.MediaType)
	if l > 0 {
		n += 1 + l + sovDiff(uint64(l))
	}
	l = len(m.Ref)
	if l > 0 {
		n += 1 + l + sovDiff(uint64(l))
	}
	if len(m.Labels) > 0 {
		for k, v := range m.Labels {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovDiff(uint64(len(k))) + 1 + len(v) + sovDiff(uint64(len(v)))
			n += mapEntrySize + 1 + sovDiff(uint64(mapEntrySize))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *DiffResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Diff != nil {
		l = m.Diff.Size()
		n += 1 + l + sovDiff(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovDiff(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozDiff(x uint64) (n int) {
	return sovDiff(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *ApplyRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ApplyRequest{`,
		`Diff:` + strings.Replace(fmt.Sprintf("%v", this.Diff), "Descriptor", "types.Descriptor", 1) + `,`,
		`Mounts:` + strings.Replace(fmt.Sprintf("%v", this.Mounts), "Mount", "types.Mount", 1) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ApplyResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ApplyResponse{`,
		`Applied:` + strings.Replace(fmt.Sprintf("%v", this.Applied), "Descriptor", "types.Descriptor", 1) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func (this *DiffRequest) String() string {
	if this == nil {
		return "nil"
	}
	keysForLabels := make([]string, 0, len(this.Labels))
	for k, _ := range this.Labels {
		keysForLabels = append(keysForLabels, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForLabels)
	mapStringForLabels := "map[string]string{"
	for _, k := range keysForLabels {
		mapStringForLabels += fmt.Sprintf("%v: %v,", k, this.Labels[k])
	}
	mapStringForLabels += "}"
	s := strings.Join([]string{`&DiffRequest{`,
		`Left:` + strings.Replace(fmt.Sprintf("%v", this.Left), "Mount", "types.Mount", 1) + `,`,
		`Right:` + strings.Replace(fmt.Sprintf("%v", this.Right), "Mount", "types.Mount", 1) + `,`,
		`MediaType:` + fmt.Sprintf("%v", this.MediaType) + `,`,
		`Ref:` + fmt.Sprintf("%v", this.Ref) + `,`,
		`Labels:` + mapStringForLabels + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func (this *DiffResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&DiffResponse{`,
		`Diff:` + strings.Replace(fmt.Sprintf("%v", this.Diff), "Descriptor", "types.Descriptor", 1) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringDiff(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *ApplyRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDiff
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
			return fmt.Errorf("proto: ApplyRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ApplyRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Diff", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDiff
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
				return ErrInvalidLengthDiff
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDiff
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Diff == nil {
				m.Diff = &types.Descriptor{}
			}
			if err := m.Diff.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Mounts", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDiff
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
				return ErrInvalidLengthDiff
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDiff
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Mounts = append(m.Mounts, &types.Mount{})
			if err := m.Mounts[len(m.Mounts)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDiff(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDiff
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthDiff
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ApplyResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDiff
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
			return fmt.Errorf("proto: ApplyResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ApplyResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Applied", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDiff
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
				return ErrInvalidLengthDiff
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDiff
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Applied == nil {
				m.Applied = &types.Descriptor{}
			}
			if err := m.Applied.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDiff(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDiff
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthDiff
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *DiffRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDiff
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
			return fmt.Errorf("proto: DiffRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DiffRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Left", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDiff
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
				return ErrInvalidLengthDiff
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDiff
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Left = append(m.Left, &types.Mount{})
			if err := m.Left[len(m.Left)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Right", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDiff
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
				return ErrInvalidLengthDiff
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDiff
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Right = append(m.Right, &types.Mount{})
			if err := m.Right[len(m.Right)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MediaType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDiff
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDiff
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDiff
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MediaType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ref", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDiff
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDiff
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDiff
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Ref = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Labels", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDiff
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
				return ErrInvalidLengthDiff
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDiff
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Labels == nil {
				m.Labels = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowDiff
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
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowDiff
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthDiff
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthDiff
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowDiff
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return ErrInvalidLengthDiff
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue < 0 {
						return ErrInvalidLengthDiff
					}
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipDiff(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthDiff
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Labels[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDiff(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDiff
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthDiff
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *DiffResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDiff
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
			return fmt.Errorf("proto: DiffResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DiffResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Diff", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDiff
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
				return ErrInvalidLengthDiff
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDiff
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Diff == nil {
				m.Diff = &types.Descriptor{}
			}
			if err := m.Diff.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDiff(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDiff
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthDiff
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipDiff(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDiff
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
					return 0, ErrIntOverflowDiff
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowDiff
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
				return 0, ErrInvalidLengthDiff
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthDiff
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowDiff
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipDiff(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthDiff
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthDiff = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDiff   = fmt.Errorf("proto: integer overflow")
)
