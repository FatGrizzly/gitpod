// Copyright (c) 2020 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License-AGPL.txt in the project root for license information.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: imagespec.proto

package api

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// ImageSpec configures the image one wishes to pull from a registry facade
type ImageSpec struct {
	// base_ref points to an image in another registry
	BaseRef string `protobuf:"bytes,1,opt,name=base_ref,json=baseRef,proto3" json:"base_ref,omitempty"`
	// ide_ref points to an image denoting the IDE to use
	IdeRef string `protobuf:"bytes,2,opt,name=ide_ref,json=ideRef,proto3" json:"ide_ref,omitempty"`
	// content_layer describe the last few layers which provide the workspace's content
	ContentLayer         []*ContentLayer `protobuf:"bytes,3,rep,name=content_layer,json=contentLayer,proto3" json:"content_layer,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ImageSpec) Reset()         { *m = ImageSpec{} }
func (m *ImageSpec) String() string { return proto.CompactTextString(m) }
func (*ImageSpec) ProtoMessage()    {}
func (*ImageSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_d570058ee0092420, []int{0}
}

func (m *ImageSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImageSpec.Unmarshal(m, b)
}
func (m *ImageSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImageSpec.Marshal(b, m, deterministic)
}
func (m *ImageSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImageSpec.Merge(m, src)
}
func (m *ImageSpec) XXX_Size() int {
	return xxx_messageInfo_ImageSpec.Size(m)
}
func (m *ImageSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_ImageSpec.DiscardUnknown(m)
}

var xxx_messageInfo_ImageSpec proto.InternalMessageInfo

func (m *ImageSpec) GetBaseRef() string {
	if m != nil {
		return m.BaseRef
	}
	return ""
}

func (m *ImageSpec) GetIdeRef() string {
	if m != nil {
		return m.IdeRef
	}
	return ""
}

func (m *ImageSpec) GetContentLayer() []*ContentLayer {
	if m != nil {
		return m.ContentLayer
	}
	return nil
}

// ContentLayer is a layer that provides a workspace's content
type ContentLayer struct {
	// Types that are valid to be assigned to Spec:
	//	*ContentLayer_Remote
	//	*ContentLayer_Direct
	Spec                 isContentLayer_Spec `protobuf_oneof:"spec"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ContentLayer) Reset()         { *m = ContentLayer{} }
func (m *ContentLayer) String() string { return proto.CompactTextString(m) }
func (*ContentLayer) ProtoMessage()    {}
func (*ContentLayer) Descriptor() ([]byte, []int) {
	return fileDescriptor_d570058ee0092420, []int{1}
}

func (m *ContentLayer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContentLayer.Unmarshal(m, b)
}
func (m *ContentLayer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContentLayer.Marshal(b, m, deterministic)
}
func (m *ContentLayer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContentLayer.Merge(m, src)
}
func (m *ContentLayer) XXX_Size() int {
	return xxx_messageInfo_ContentLayer.Size(m)
}
func (m *ContentLayer) XXX_DiscardUnknown() {
	xxx_messageInfo_ContentLayer.DiscardUnknown(m)
}

var xxx_messageInfo_ContentLayer proto.InternalMessageInfo

type isContentLayer_Spec interface {
	isContentLayer_Spec()
}

type ContentLayer_Remote struct {
	Remote *RemoteContentLayer `protobuf:"bytes,1,opt,name=remote,proto3,oneof"`
}

type ContentLayer_Direct struct {
	Direct *DirectContentLayer `protobuf:"bytes,2,opt,name=direct,proto3,oneof"`
}

func (*ContentLayer_Remote) isContentLayer_Spec() {}

func (*ContentLayer_Direct) isContentLayer_Spec() {}

func (m *ContentLayer) GetSpec() isContentLayer_Spec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *ContentLayer) GetRemote() *RemoteContentLayer {
	if x, ok := m.GetSpec().(*ContentLayer_Remote); ok {
		return x.Remote
	}
	return nil
}

func (m *ContentLayer) GetDirect() *DirectContentLayer {
	if x, ok := m.GetSpec().(*ContentLayer_Direct); ok {
		return x.Direct
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ContentLayer) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ContentLayer_Remote)(nil),
		(*ContentLayer_Direct)(nil),
	}
}

// RemoteContentLayer is a layer which can be downloaded from a remote URL.
// If the diff_id is empty or equals the digest the layer is expected to be uncompressed.
type RemoteContentLayer struct {
	// url points to the actual content location. This must be a valid HTTPS URL pointing
	// to a tar.gz file.
	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	// digest is the digest (content hash) of the file the URL points to.
	Digest string `protobuf:"bytes,2,opt,name=digest,proto3" json:"digest,omitempty"`
	// diff_id is the digest (content hash) of the uncompressed data the URL points to if the
	// the URL points to a compressed file. If the file is uncompressed to begin with this field
	// can either be empty or the same as digest.
	DiffId string `protobuf:"bytes,3,opt,name=diff_id,json=diffId,proto3" json:"diff_id,omitempty"`
	// media_type is the content type of the layer and is expected to be one of:
	//  application/vnd.oci.image.layer.v1.tar
	//  application/vnd.oci.image.layer.v1.tar+gzip
	//  application/vnd.oci.image.layer.v1.tar+zstd
	MediaType string `protobuf:"bytes,4,opt,name=media_type,json=mediaType,proto3" json:"media_type,omitempty"`
	// size is the size of the layer download in bytes
	Size                 int64    `protobuf:"varint,5,opt,name=size,proto3" json:"size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoteContentLayer) Reset()         { *m = RemoteContentLayer{} }
func (m *RemoteContentLayer) String() string { return proto.CompactTextString(m) }
func (*RemoteContentLayer) ProtoMessage()    {}
func (*RemoteContentLayer) Descriptor() ([]byte, []int) {
	return fileDescriptor_d570058ee0092420, []int{2}
}

func (m *RemoteContentLayer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoteContentLayer.Unmarshal(m, b)
}
func (m *RemoteContentLayer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoteContentLayer.Marshal(b, m, deterministic)
}
func (m *RemoteContentLayer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoteContentLayer.Merge(m, src)
}
func (m *RemoteContentLayer) XXX_Size() int {
	return xxx_messageInfo_RemoteContentLayer.Size(m)
}
func (m *RemoteContentLayer) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoteContentLayer.DiscardUnknown(m)
}

var xxx_messageInfo_RemoteContentLayer proto.InternalMessageInfo

func (m *RemoteContentLayer) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *RemoteContentLayer) GetDigest() string {
	if m != nil {
		return m.Digest
	}
	return ""
}

func (m *RemoteContentLayer) GetDiffId() string {
	if m != nil {
		return m.DiffId
	}
	return ""
}

func (m *RemoteContentLayer) GetMediaType() string {
	if m != nil {
		return m.MediaType
	}
	return ""
}

func (m *RemoteContentLayer) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

// DirectContentLayer is an uncompressed tar file which is directly added as layer
type DirectContentLayer struct {
	// the bytes of the uncompressed tar file which is served as layer
	Content              []byte   `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DirectContentLayer) Reset()         { *m = DirectContentLayer{} }
func (m *DirectContentLayer) String() string { return proto.CompactTextString(m) }
func (*DirectContentLayer) ProtoMessage()    {}
func (*DirectContentLayer) Descriptor() ([]byte, []int) {
	return fileDescriptor_d570058ee0092420, []int{3}
}

func (m *DirectContentLayer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DirectContentLayer.Unmarshal(m, b)
}
func (m *DirectContentLayer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DirectContentLayer.Marshal(b, m, deterministic)
}
func (m *DirectContentLayer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DirectContentLayer.Merge(m, src)
}
func (m *DirectContentLayer) XXX_Size() int {
	return xxx_messageInfo_DirectContentLayer.Size(m)
}
func (m *DirectContentLayer) XXX_DiscardUnknown() {
	xxx_messageInfo_DirectContentLayer.DiscardUnknown(m)
}

var xxx_messageInfo_DirectContentLayer proto.InternalMessageInfo

func (m *DirectContentLayer) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

func init() {
	proto.RegisterType((*ImageSpec)(nil), "registryfacade.ImageSpec")
	proto.RegisterType((*ContentLayer)(nil), "registryfacade.ContentLayer")
	proto.RegisterType((*RemoteContentLayer)(nil), "registryfacade.RemoteContentLayer")
	proto.RegisterType((*DirectContentLayer)(nil), "registryfacade.DirectContentLayer")
}

func init() {
	proto.RegisterFile("imagespec.proto", fileDescriptor_d570058ee0092420)
}

var fileDescriptor_d570058ee0092420 = []byte{
	// 309 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x3f, 0x4f, 0xc3, 0x30,
	0x10, 0xc5, 0x09, 0x69, 0x53, 0x7a, 0x2d, 0x7f, 0x74, 0x03, 0x18, 0x09, 0xa4, 0x2a, 0x53, 0xa7,
	0x0c, 0x65, 0x65, 0xa1, 0x30, 0x50, 0x89, 0xc9, 0x30, 0xb1, 0x44, 0xae, 0x7d, 0xae, 0x2c, 0xb5,
	0x8d, 0xe5, 0x98, 0x21, 0x8c, 0x8c, 0x8c, 0x7c, 0x62, 0x64, 0xa7, 0x95, 0x4a, 0xcb, 0x76, 0x2f,
	0xbf, 0x7b, 0xb9, 0xa7, 0x67, 0x38, 0x37, 0x2b, 0xb1, 0xa0, 0xda, 0x92, 0x2c, 0xac, 0xab, 0x7c,
	0x85, 0x67, 0x8e, 0x16, 0xa6, 0xf6, 0xae, 0xd1, 0x42, 0x0a, 0x45, 0xf9, 0x57, 0x02, 0xfd, 0x59,
	0xd8, 0x79, 0xb5, 0x24, 0xf1, 0x1a, 0x4e, 0xe6, 0xa2, 0xa6, 0xd2, 0x91, 0x66, 0xc9, 0x28, 0x19,
	0xf7, 0x79, 0x2f, 0x68, 0x4e, 0x1a, 0xaf, 0xa0, 0x67, 0x54, 0x4b, 0x8e, 0x23, 0xc9, 0x8c, 0x8a,
	0xe0, 0x01, 0x4e, 0x65, 0xb5, 0xf6, 0xb4, 0xf6, 0xe5, 0x52, 0x34, 0xe4, 0x58, 0x3a, 0x4a, 0xc7,
	0x83, 0xc9, 0x4d, 0xf1, 0xf7, 0x52, 0xf1, 0xd8, 0x2e, 0xbd, 0x84, 0x1d, 0x3e, 0x94, 0x3b, 0x2a,
	0xff, 0x49, 0x60, 0xb8, 0x8b, 0xf1, 0x1e, 0x32, 0x47, 0xab, 0xca, 0x53, 0x4c, 0x31, 0x98, 0xe4,
	0xfb, 0x3f, 0xe3, 0x91, 0xee, 0x7a, 0x9e, 0x8f, 0xf8, 0xc6, 0x13, 0xdc, 0xca, 0x38, 0x92, 0x3e,
	0x26, 0xfd, 0xc7, 0xfd, 0x14, 0xe9, 0xbe, 0xbb, 0xf5, 0x4c, 0x33, 0xe8, 0x84, 0xbe, 0xf2, 0xef,
	0x04, 0xf0, 0xf0, 0x0c, 0x5e, 0x40, 0xfa, 0xe1, 0x96, 0x9b, 0x76, 0xc2, 0x88, 0x97, 0xe1, 0xdc,
	0x82, 0x6a, 0xbf, 0x2d, 0xa6, 0x55, 0xa1, 0x31, 0x65, 0xb4, 0x2e, 0x8d, 0x62, 0xe9, 0x16, 0x68,
	0x3d, 0x53, 0x78, 0x0b, 0xb0, 0x22, 0x65, 0x44, 0xe9, 0x1b, 0x4b, 0xac, 0x13, 0x59, 0x3f, 0x7e,
	0x79, 0x6b, 0x2c, 0x21, 0x42, 0xa7, 0x36, 0x9f, 0xc4, 0xba, 0xa3, 0x64, 0x9c, 0xf2, 0x38, 0xe7,
	0x05, 0xe0, 0x61, 0x68, 0x64, 0xd0, 0xdb, 0xf4, 0x18, 0xf3, 0x0c, 0xf9, 0x56, 0x4e, 0xbb, 0xef,
	0xa9, 0xb0, 0x66, 0x9e, 0xc5, 0x47, 0xbf, 0xfb, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xd6, 0x96, 0xc7,
	0x37, 0x07, 0x02, 0x00, 0x00,
}
