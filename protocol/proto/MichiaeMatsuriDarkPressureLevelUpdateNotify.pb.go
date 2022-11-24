// Sorapointa - A server software re-implementation for a certain anime game, and avoid sorapointa.
// Copyright (C) 2022  Sorapointa Team
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: MichiaeMatsuriDarkPressureLevelUpdateNotify.proto

package proto

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

// CmdId: 8825
// EnetChannelId: 0
// EnetIsReliable: true
type MichiaeMatsuriDarkPressureLevelUpdateNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DarkPressureLevel uint32 `protobuf:"varint,8,opt,name=dark_pressure_level,json=darkPressureLevel,proto3" json:"dark_pressure_level,omitempty"`
}

func (x *MichiaeMatsuriDarkPressureLevelUpdateNotify) Reset() {
	*x = MichiaeMatsuriDarkPressureLevelUpdateNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MichiaeMatsuriDarkPressureLevelUpdateNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MichiaeMatsuriDarkPressureLevelUpdateNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MichiaeMatsuriDarkPressureLevelUpdateNotify) ProtoMessage() {}

func (x *MichiaeMatsuriDarkPressureLevelUpdateNotify) ProtoReflect() protoreflect.Message {
	mi := &file_MichiaeMatsuriDarkPressureLevelUpdateNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MichiaeMatsuriDarkPressureLevelUpdateNotify.ProtoReflect.Descriptor instead.
func (*MichiaeMatsuriDarkPressureLevelUpdateNotify) Descriptor() ([]byte, []int) {
	return file_MichiaeMatsuriDarkPressureLevelUpdateNotify_proto_rawDescGZIP(), []int{0}
}

func (x *MichiaeMatsuriDarkPressureLevelUpdateNotify) GetDarkPressureLevel() uint32 {
	if x != nil {
		return x.DarkPressureLevel
	}
	return 0
}

var File_MichiaeMatsuriDarkPressureLevelUpdateNotify_proto protoreflect.FileDescriptor

var file_MichiaeMatsuriDarkPressureLevelUpdateNotify_proto_rawDesc = []byte{
	0x0a, 0x31, 0x4d, 0x69, 0x63, 0x68, 0x69, 0x61, 0x65, 0x4d, 0x61, 0x74, 0x73, 0x75, 0x72, 0x69,
	0x44, 0x61, 0x72, 0x6b, 0x50, 0x72, 0x65, 0x73, 0x73, 0x75, 0x72, 0x65, 0x4c, 0x65, 0x76, 0x65,
	0x6c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5d, 0x0a, 0x2b, 0x4d, 0x69,
	0x63, 0x68, 0x69, 0x61, 0x65, 0x4d, 0x61, 0x74, 0x73, 0x75, 0x72, 0x69, 0x44, 0x61, 0x72, 0x6b,
	0x50, 0x72, 0x65, 0x73, 0x73, 0x75, 0x72, 0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x2e, 0x0a, 0x13, 0x64, 0x61, 0x72,
	0x6b, 0x5f, 0x70, 0x72, 0x65, 0x73, 0x73, 0x75, 0x72, 0x65, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x11, 0x64, 0x61, 0x72, 0x6b, 0x50, 0x72, 0x65, 0x73,
	0x73, 0x75, 0x72, 0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_MichiaeMatsuriDarkPressureLevelUpdateNotify_proto_rawDescOnce sync.Once
	file_MichiaeMatsuriDarkPressureLevelUpdateNotify_proto_rawDescData = file_MichiaeMatsuriDarkPressureLevelUpdateNotify_proto_rawDesc
)

func file_MichiaeMatsuriDarkPressureLevelUpdateNotify_proto_rawDescGZIP() []byte {
	file_MichiaeMatsuriDarkPressureLevelUpdateNotify_proto_rawDescOnce.Do(func() {
		file_MichiaeMatsuriDarkPressureLevelUpdateNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_MichiaeMatsuriDarkPressureLevelUpdateNotify_proto_rawDescData)
	})
	return file_MichiaeMatsuriDarkPressureLevelUpdateNotify_proto_rawDescData
}

var file_MichiaeMatsuriDarkPressureLevelUpdateNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_MichiaeMatsuriDarkPressureLevelUpdateNotify_proto_goTypes = []interface{}{
	(*MichiaeMatsuriDarkPressureLevelUpdateNotify)(nil), // 0: proto.MichiaeMatsuriDarkPressureLevelUpdateNotify
}
var file_MichiaeMatsuriDarkPressureLevelUpdateNotify_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_MichiaeMatsuriDarkPressureLevelUpdateNotify_proto_init() }
func file_MichiaeMatsuriDarkPressureLevelUpdateNotify_proto_init() {
	if File_MichiaeMatsuriDarkPressureLevelUpdateNotify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_MichiaeMatsuriDarkPressureLevelUpdateNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MichiaeMatsuriDarkPressureLevelUpdateNotify); i {
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
			RawDescriptor: file_MichiaeMatsuriDarkPressureLevelUpdateNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_MichiaeMatsuriDarkPressureLevelUpdateNotify_proto_goTypes,
		DependencyIndexes: file_MichiaeMatsuriDarkPressureLevelUpdateNotify_proto_depIdxs,
		MessageInfos:      file_MichiaeMatsuriDarkPressureLevelUpdateNotify_proto_msgTypes,
	}.Build()
	File_MichiaeMatsuriDarkPressureLevelUpdateNotify_proto = out.File
	file_MichiaeMatsuriDarkPressureLevelUpdateNotify_proto_rawDesc = nil
	file_MichiaeMatsuriDarkPressureLevelUpdateNotify_proto_goTypes = nil
	file_MichiaeMatsuriDarkPressureLevelUpdateNotify_proto_depIdxs = nil
}