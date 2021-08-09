// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: torrent.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type FilesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Magnet string `protobuf:"bytes,1,opt,name=magnet,proto3" json:"magnet,omitempty"`
}

func (x *FilesRequest) Reset() {
	*x = FilesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_torrent_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilesRequest) ProtoMessage() {}

func (x *FilesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_torrent_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilesRequest.ProtoReflect.Descriptor instead.
func (*FilesRequest) Descriptor() ([]byte, []int) {
	return file_torrent_proto_rawDescGZIP(), []int{0}
}

func (x *FilesRequest) GetMagnet() string {
	if x != nil {
		return x.Magnet
	}
	return ""
}

type FilesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Files []*File `protobuf:"bytes,1,rep,name=files,proto3" json:"files,omitempty"`
}

func (x *FilesResponse) Reset() {
	*x = FilesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_torrent_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilesResponse) ProtoMessage() {}

func (x *FilesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_torrent_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilesResponse.ProtoReflect.Descriptor instead.
func (*FilesResponse) Descriptor() ([]byte, []int) {
	return file_torrent_proto_rawDescGZIP(), []int{1}
}

func (x *FilesResponse) GetFiles() []*File {
	if x != nil {
		return x.Files
	}
	return nil
}

type File struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TorrentHash string `protobuf:"bytes,1,opt,name=torrentHash,proto3" json:"torrentHash,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Length      int64  `protobuf:"varint,3,opt,name=length,proto3" json:"length,omitempty"`
}

func (x *File) Reset() {
	*x = File{}
	if protoimpl.UnsafeEnabled {
		mi := &file_torrent_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *File) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*File) ProtoMessage() {}

func (x *File) ProtoReflect() protoreflect.Message {
	mi := &file_torrent_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use File.ProtoReflect.Descriptor instead.
func (*File) Descriptor() ([]byte, []int) {
	return file_torrent_proto_rawDescGZIP(), []int{2}
}

func (x *File) GetTorrentHash() string {
	if x != nil {
		return x.TorrentHash
	}
	return ""
}

func (x *File) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *File) GetLength() int64 {
	if x != nil {
		return x.Length
	}
	return 0
}

type ReadAtRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	File  *File `protobuf:"bytes,1,opt,name=file,proto3" json:"file,omitempty"`
	Index int32 `protobuf:"varint,2,opt,name=index,proto3" json:"index,omitempty"`
	Off   int64 `protobuf:"varint,3,opt,name=off,proto3" json:"off,omitempty"`
	Ln    int64 `protobuf:"varint,4,opt,name=ln,proto3" json:"ln,omitempty"`
}

func (x *ReadAtRequest) Reset() {
	*x = ReadAtRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_torrent_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadAtRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadAtRequest) ProtoMessage() {}

func (x *ReadAtRequest) ProtoReflect() protoreflect.Message {
	mi := &file_torrent_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadAtRequest.ProtoReflect.Descriptor instead.
func (*ReadAtRequest) Descriptor() ([]byte, []int) {
	return file_torrent_proto_rawDescGZIP(), []int{3}
}

func (x *ReadAtRequest) GetFile() *File {
	if x != nil {
		return x.File
	}
	return nil
}

func (x *ReadAtRequest) GetIndex() int32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *ReadAtRequest) GetOff() int64 {
	if x != nil {
		return x.Off
	}
	return 0
}

func (x *ReadAtRequest) GetLn() int64 {
	if x != nil {
		return x.Ln
	}
	return 0
}

type ReadAtResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Buffer []byte `protobuf:"bytes,1,opt,name=buffer,proto3" json:"buffer,omitempty"`
}

func (x *ReadAtResponse) Reset() {
	*x = ReadAtResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_torrent_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadAtResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadAtResponse) ProtoMessage() {}

func (x *ReadAtResponse) ProtoReflect() protoreflect.Message {
	mi := &file_torrent_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadAtResponse.ProtoReflect.Descriptor instead.
func (*ReadAtResponse) Descriptor() ([]byte, []int) {
	return file_torrent_proto_rawDescGZIP(), []int{4}
}

func (x *ReadAtResponse) GetBuffer() []byte {
	if x != nil {
		return x.Buffer
	}
	return nil
}

type IsMagnetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Magnet string `protobuf:"bytes,1,opt,name=magnet,proto3" json:"magnet,omitempty"`
}

func (x *IsMagnetRequest) Reset() {
	*x = IsMagnetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_torrent_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsMagnetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsMagnetRequest) ProtoMessage() {}

func (x *IsMagnetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_torrent_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsMagnetRequest.ProtoReflect.Descriptor instead.
func (*IsMagnetRequest) Descriptor() ([]byte, []int) {
	return file_torrent_proto_rawDescGZIP(), []int{5}
}

func (x *IsMagnetRequest) GetMagnet() string {
	if x != nil {
		return x.Magnet
	}
	return ""
}

type IsMagnetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Valid bool `protobuf:"varint,1,opt,name=valid,proto3" json:"valid,omitempty"`
}

func (x *IsMagnetResponse) Reset() {
	*x = IsMagnetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_torrent_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsMagnetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsMagnetResponse) ProtoMessage() {}

func (x *IsMagnetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_torrent_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsMagnetResponse.ProtoReflect.Descriptor instead.
func (*IsMagnetResponse) Descriptor() ([]byte, []int) {
	return file_torrent_proto_rawDescGZIP(), []int{6}
}

func (x *IsMagnetResponse) GetValid() bool {
	if x != nil {
		return x.Valid
	}
	return false
}

var File_torrent_proto protoreflect.FileDescriptor

var file_torrent_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x74, 0x6f, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x26, 0x0a, 0x0c, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x6d, 0x61, 0x67, 0x6e, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x6d, 0x61, 0x67, 0x6e, 0x65, 0x74, 0x22, 0x2c, 0x0a, 0x0d, 0x46, 0x69, 0x6c, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x05, 0x66, 0x69, 0x6c, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x05,
	0x66, 0x69, 0x6c, 0x65, 0x73, 0x22, 0x54, 0x0a, 0x04, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x74, 0x6f, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x48, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x74, 0x6f, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x48, 0x61, 0x73, 0x68, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x22, 0x62, 0x0a, 0x0d, 0x52,
	0x65, 0x61, 0x64, 0x41, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x04,
	0x66, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x46, 0x69, 0x6c,
	0x65, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x10, 0x0a,
	0x03, 0x6f, 0x66, 0x66, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6f, 0x66, 0x66, 0x12,
	0x0e, 0x0a, 0x02, 0x6c, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x6c, 0x6e, 0x22,
	0x28, 0x0a, 0x0e, 0x52, 0x65, 0x61, 0x64, 0x41, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x06, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x22, 0x29, 0x0a, 0x0f, 0x49, 0x73, 0x4d,
	0x61, 0x67, 0x6e, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x6d, 0x61, 0x67, 0x6e, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x61,
	0x67, 0x6e, 0x65, 0x74, 0x22, 0x28, 0x0a, 0x10, 0x49, 0x73, 0x4d, 0x61, 0x67, 0x6e, 0x65, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x32, 0x93,
	0x01, 0x0a, 0x07, 0x54, 0x6f, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x28, 0x0a, 0x05, 0x46, 0x69,
	0x6c, 0x65, 0x73, 0x12, 0x0d, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x2b, 0x0a, 0x06, 0x52, 0x65, 0x61, 0x64, 0x41, 0x74, 0x12, 0x0e,
	0x2e, 0x52, 0x65, 0x61, 0x64, 0x41, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f,
	0x2e, 0x52, 0x65, 0x61, 0x64, 0x41, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x31, 0x0a, 0x08, 0x49, 0x73, 0x4d, 0x61, 0x67, 0x6e, 0x65, 0x74, 0x12, 0x10, 0x2e,
	0x49, 0x73, 0x4d, 0x61, 0x67, 0x6e, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x11, 0x2e, 0x49, 0x73, 0x4d, 0x61, 0x67, 0x6e, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x20, 0x5a, 0x1e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x63, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6e, 0x2f, 0x74, 0x6f, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_torrent_proto_rawDescOnce sync.Once
	file_torrent_proto_rawDescData = file_torrent_proto_rawDesc
)

func file_torrent_proto_rawDescGZIP() []byte {
	file_torrent_proto_rawDescOnce.Do(func() {
		file_torrent_proto_rawDescData = protoimpl.X.CompressGZIP(file_torrent_proto_rawDescData)
	})
	return file_torrent_proto_rawDescData
}

var file_torrent_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_torrent_proto_goTypes = []interface{}{
	(*FilesRequest)(nil),     // 0: FilesRequest
	(*FilesResponse)(nil),    // 1: FilesResponse
	(*File)(nil),             // 2: File
	(*ReadAtRequest)(nil),    // 3: ReadAtRequest
	(*ReadAtResponse)(nil),   // 4: ReadAtResponse
	(*IsMagnetRequest)(nil),  // 5: IsMagnetRequest
	(*IsMagnetResponse)(nil), // 6: IsMagnetResponse
}
var file_torrent_proto_depIdxs = []int32{
	2, // 0: FilesResponse.files:type_name -> File
	2, // 1: ReadAtRequest.file:type_name -> File
	0, // 2: Torrent.Files:input_type -> FilesRequest
	3, // 3: Torrent.ReadAt:input_type -> ReadAtRequest
	5, // 4: Torrent.IsMagnet:input_type -> IsMagnetRequest
	1, // 5: Torrent.Files:output_type -> FilesResponse
	4, // 6: Torrent.ReadAt:output_type -> ReadAtResponse
	6, // 7: Torrent.IsMagnet:output_type -> IsMagnetResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_torrent_proto_init() }
func file_torrent_proto_init() {
	if File_torrent_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_torrent_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilesRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_torrent_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilesResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_torrent_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*File); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_torrent_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadAtRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_torrent_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadAtResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_torrent_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsMagnetRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_torrent_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsMagnetResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_torrent_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_torrent_proto_goTypes,
		DependencyIndexes: file_torrent_proto_depIdxs,
		MessageInfos:      file_torrent_proto_msgTypes,
	}.Build()
	File_torrent_proto = out.File
	file_torrent_proto_rawDesc = nil
	file_torrent_proto_goTypes = nil
	file_torrent_proto_depIdxs = nil
}