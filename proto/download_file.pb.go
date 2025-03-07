// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v6.30.0
// source: proto/download_file.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type FileRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ProjectId     string                 `protobuf:"bytes,1,opt,name=ProjectId,proto3" json:"ProjectId,omitempty"`
	AppName       string                 `protobuf:"bytes,2,opt,name=AppName,proto3" json:"AppName,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FileRequest) Reset() {
	*x = FileRequest{}
	mi := &file_proto_download_file_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileRequest) ProtoMessage() {}

func (x *FileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_download_file_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileRequest.ProtoReflect.Descriptor instead.
func (*FileRequest) Descriptor() ([]byte, []int) {
	return file_proto_download_file_proto_rawDescGZIP(), []int{0}
}

func (x *FileRequest) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *FileRequest) GetAppName() string {
	if x != nil {
		return x.AppName
	}
	return ""
}

type FileResponse struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	FilesDownloaded []string               `protobuf:"bytes,1,rep,name=files_downloaded,json=filesDownloaded,proto3" json:"files_downloaded,omitempty"`
	Error           string                 `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *FileResponse) Reset() {
	*x = FileResponse{}
	mi := &file_proto_download_file_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileResponse) ProtoMessage() {}

func (x *FileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_download_file_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileResponse.ProtoReflect.Descriptor instead.
func (*FileResponse) Descriptor() ([]byte, []int) {
	return file_proto_download_file_proto_rawDescGZIP(), []int{1}
}

func (x *FileResponse) GetFilesDownloaded() []string {
	if x != nil {
		return x.FilesDownloaded
	}
	return nil
}

func (x *FileResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_proto_download_file_proto protoreflect.FileDescriptor

var file_proto_download_file_proto_rawDesc = string([]byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64,
	0x5f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x64, 0x6f, 0x77,
	0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x22, 0x45, 0x0a, 0x0b, 0x46, 0x69, 0x6c,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x50, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x70, 0x70, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x70, 0x70, 0x4e, 0x61, 0x6d, 0x65,
	0x22, 0x4f, 0x0a, 0x0c, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x29, 0x0a, 0x10, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x5f, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f,
	0x61, 0x64, 0x65, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0f, 0x66, 0x69, 0x6c, 0x65,
	0x73, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x32, 0x5c, 0x0a, 0x13, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61,
	0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x45, 0x0a, 0x0c, 0x44, 0x6f, 0x77, 0x6e,
	0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x19, 0x2e, 0x64, 0x6f, 0x77, 0x6e, 0x6c,
	0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69,
	0x6c, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x08, 0x5a, 0x06, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
})

var (
	file_proto_download_file_proto_rawDescOnce sync.Once
	file_proto_download_file_proto_rawDescData []byte
)

func file_proto_download_file_proto_rawDescGZIP() []byte {
	file_proto_download_file_proto_rawDescOnce.Do(func() {
		file_proto_download_file_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_download_file_proto_rawDesc), len(file_proto_download_file_proto_rawDesc)))
	})
	return file_proto_download_file_proto_rawDescData
}

var file_proto_download_file_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_download_file_proto_goTypes = []any{
	(*FileRequest)(nil),  // 0: downloadFile.FileRequest
	(*FileResponse)(nil), // 1: downloadFile.FileResponse
}
var file_proto_download_file_proto_depIdxs = []int32{
	0, // 0: downloadFile.FileDownloadService.DownloadFile:input_type -> downloadFile.FileRequest
	1, // 1: downloadFile.FileDownloadService.DownloadFile:output_type -> downloadFile.FileResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_download_file_proto_init() }
func file_proto_download_file_proto_init() {
	if File_proto_download_file_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_download_file_proto_rawDesc), len(file_proto_download_file_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_download_file_proto_goTypes,
		DependencyIndexes: file_proto_download_file_proto_depIdxs,
		MessageInfos:      file_proto_download_file_proto_msgTypes,
	}.Build()
	File_proto_download_file_proto = out.File
	file_proto_download_file_proto_goTypes = nil
	file_proto_download_file_proto_depIdxs = nil
}
