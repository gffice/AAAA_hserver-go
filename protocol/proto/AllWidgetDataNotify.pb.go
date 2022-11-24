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
// source: AllWidgetDataNotify.proto

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

// CmdId: 4271
// EnetChannelId: 0
// EnetIsReliable: true
type AllWidgetDataNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BackgroundActiveWidgetList       []uint32                        `protobuf:"varint,11,rep,packed,name=background_active_widget_list,json=backgroundActiveWidgetList,proto3" json:"background_active_widget_list,omitempty"`
	LunchBoxData                     *LunchBoxData                   `protobuf:"bytes,1,opt,name=lunch_box_data,json=lunchBoxData,proto3" json:"lunch_box_data,omitempty"`
	CoolDownGroupDataList            []*WidgetCoolDownData           `protobuf:"bytes,13,rep,name=cool_down_group_data_list,json=coolDownGroupDataList,proto3" json:"cool_down_group_data_list,omitempty"`
	AnchorPointList                  []*AnchorPointData              `protobuf:"bytes,3,rep,name=anchor_point_list,json=anchorPointList,proto3" json:"anchor_point_list,omitempty"`
	SlotList                         []*WidgetSlotData               `protobuf:"bytes,6,rep,name=slot_list,json=slotList,proto3" json:"slot_list,omitempty"`
	NextAnchorPointUsableTime        uint32                          `protobuf:"varint,10,opt,name=next_anchor_point_usable_time,json=nextAnchorPointUsableTime,proto3" json:"next_anchor_point_usable_time,omitempty"`
	ClientCollectorDataList          []*ClientCollectorData          `protobuf:"bytes,4,rep,name=client_collector_data_list,json=clientCollectorDataList,proto3" json:"client_collector_data_list,omitempty"`
	OneofGatherPointDetectorDataList []*OneofGatherPointDetectorData `protobuf:"bytes,15,rep,name=oneof_gather_point_detector_data_list,json=oneofGatherPointDetectorDataList,proto3" json:"oneof_gather_point_detector_data_list,omitempty"`
	NormalCoolDownDataList           []*WidgetCoolDownData           `protobuf:"bytes,9,rep,name=normal_cool_down_data_list,json=normalCoolDownDataList,proto3" json:"normal_cool_down_data_list,omitempty"`
	SkyCrystalDetectorData           *SkyCrystalDetectorData         `protobuf:"bytes,12,opt,name=sky_crystal_detector_data,json=skyCrystalDetectorData,proto3" json:"sky_crystal_detector_data,omitempty"`
}

func (x *AllWidgetDataNotify) Reset() {
	*x = AllWidgetDataNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AllWidgetDataNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllWidgetDataNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllWidgetDataNotify) ProtoMessage() {}

func (x *AllWidgetDataNotify) ProtoReflect() protoreflect.Message {
	mi := &file_AllWidgetDataNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllWidgetDataNotify.ProtoReflect.Descriptor instead.
func (*AllWidgetDataNotify) Descriptor() ([]byte, []int) {
	return file_AllWidgetDataNotify_proto_rawDescGZIP(), []int{0}
}

func (x *AllWidgetDataNotify) GetBackgroundActiveWidgetList() []uint32 {
	if x != nil {
		return x.BackgroundActiveWidgetList
	}
	return nil
}

func (x *AllWidgetDataNotify) GetLunchBoxData() *LunchBoxData {
	if x != nil {
		return x.LunchBoxData
	}
	return nil
}

func (x *AllWidgetDataNotify) GetCoolDownGroupDataList() []*WidgetCoolDownData {
	if x != nil {
		return x.CoolDownGroupDataList
	}
	return nil
}

func (x *AllWidgetDataNotify) GetAnchorPointList() []*AnchorPointData {
	if x != nil {
		return x.AnchorPointList
	}
	return nil
}

func (x *AllWidgetDataNotify) GetSlotList() []*WidgetSlotData {
	if x != nil {
		return x.SlotList
	}
	return nil
}

func (x *AllWidgetDataNotify) GetNextAnchorPointUsableTime() uint32 {
	if x != nil {
		return x.NextAnchorPointUsableTime
	}
	return 0
}

func (x *AllWidgetDataNotify) GetClientCollectorDataList() []*ClientCollectorData {
	if x != nil {
		return x.ClientCollectorDataList
	}
	return nil
}

func (x *AllWidgetDataNotify) GetOneofGatherPointDetectorDataList() []*OneofGatherPointDetectorData {
	if x != nil {
		return x.OneofGatherPointDetectorDataList
	}
	return nil
}

func (x *AllWidgetDataNotify) GetNormalCoolDownDataList() []*WidgetCoolDownData {
	if x != nil {
		return x.NormalCoolDownDataList
	}
	return nil
}

func (x *AllWidgetDataNotify) GetSkyCrystalDetectorData() *SkyCrystalDetectorData {
	if x != nil {
		return x.SkyCrystalDetectorData
	}
	return nil
}

var File_AllWidgetDataNotify_proto protoreflect.FileDescriptor

var file_AllWidgetDataNotify_proto_rawDesc = []byte{
	0x0a, 0x19, 0x41, 0x6c, 0x6c, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x15, 0x41, 0x6e, 0x63, 0x68, 0x6f, 0x72, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x44,
	0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x4c, 0x75, 0x6e, 0x63, 0x68, 0x42, 0x6f, 0x78, 0x44, 0x61,
	0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x4f, 0x6e, 0x65, 0x6f, 0x66, 0x47,
	0x61, 0x74, 0x68, 0x65, 0x72, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x53, 0x6b,
	0x79, 0x43, 0x72, 0x79, 0x73, 0x74, 0x61, 0x6c, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x44, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x57, 0x69, 0x64, 0x67,
	0x65, 0x74, 0x43, 0x6f, 0x6f, 0x6c, 0x44, 0x6f, 0x77, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x53, 0x6c, 0x6f, 0x74,
	0x44, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa2, 0x06, 0x0a, 0x13, 0x41,
	0x6c, 0x6c, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x79, 0x12, 0x41, 0x0a, 0x1d, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64,
	0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x1a, 0x62, 0x61, 0x63, 0x6b, 0x67,
	0x72, 0x6f, 0x75, 0x6e, 0x64, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x57, 0x69, 0x64, 0x67, 0x65,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x39, 0x0a, 0x0e, 0x6c, 0x75, 0x6e, 0x63, 0x68, 0x5f, 0x62,
	0x6f, 0x78, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x75, 0x6e, 0x63, 0x68, 0x42, 0x6f, 0x78, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x0c, 0x6c, 0x75, 0x6e, 0x63, 0x68, 0x42, 0x6f, 0x78, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x53, 0x0a, 0x19, 0x63, 0x6f, 0x6f, 0x6c, 0x5f, 0x64, 0x6f, 0x77, 0x6e, 0x5f, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0d, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x57, 0x69, 0x64, 0x67,
	0x65, 0x74, 0x43, 0x6f, 0x6f, 0x6c, 0x44, 0x6f, 0x77, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x15,
	0x63, 0x6f, 0x6f, 0x6c, 0x44, 0x6f, 0x77, 0x6e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x44, 0x61, 0x74,
	0x61, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x42, 0x0a, 0x11, 0x61, 0x6e, 0x63, 0x68, 0x6f, 0x72, 0x5f,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x6e, 0x63, 0x68, 0x6f, 0x72, 0x50,
	0x6f, 0x69, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0f, 0x61, 0x6e, 0x63, 0x68, 0x6f, 0x72,
	0x50, 0x6f, 0x69, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x09, 0x73, 0x6c, 0x6f,
	0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x53, 0x6c, 0x6f, 0x74, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x08, 0x73, 0x6c, 0x6f, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x40, 0x0a,
	0x1d, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x61, 0x6e, 0x63, 0x68, 0x6f, 0x72, 0x5f, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x5f, 0x75, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x19, 0x6e, 0x65, 0x78, 0x74, 0x41, 0x6e, 0x63, 0x68, 0x6f, 0x72,
	0x50, 0x6f, 0x69, 0x6e, 0x74, 0x55, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x57, 0x0a, 0x1a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x17, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x44, 0x61, 0x74, 0x61, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x74, 0x0a, 0x25, 0x6f, 0x6e, 0x65, 0x6f,
	0x66, 0x5f, 0x67, 0x61, 0x74, 0x68, 0x65, 0x72, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x64,
	0x65, 0x74, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6c, 0x69, 0x73,
	0x74, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x4f, 0x6e, 0x65, 0x6f, 0x66, 0x47, 0x61, 0x74, 0x68, 0x65, 0x72, 0x50, 0x6f, 0x69, 0x6e, 0x74,
	0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x44, 0x61, 0x74, 0x61, 0x52, 0x20, 0x6f, 0x6e,
	0x65, 0x6f, 0x66, 0x47, 0x61, 0x74, 0x68, 0x65, 0x72, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x44, 0x65,
	0x74, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x44, 0x61, 0x74, 0x61, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x55,
	0x0a, 0x1a, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x6f, 0x6c, 0x5f, 0x64, 0x6f,
	0x77, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x09, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x57, 0x69, 0x64, 0x67, 0x65,
	0x74, 0x43, 0x6f, 0x6f, 0x6c, 0x44, 0x6f, 0x77, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x16, 0x6e,
	0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x43, 0x6f, 0x6f, 0x6c, 0x44, 0x6f, 0x77, 0x6e, 0x44, 0x61, 0x74,
	0x61, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x58, 0x0a, 0x19, 0x73, 0x6b, 0x79, 0x5f, 0x63, 0x72, 0x79,
	0x73, 0x74, 0x61, 0x6c, 0x5f, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x53, 0x6b, 0x79, 0x43, 0x72, 0x79, 0x73, 0x74, 0x61, 0x6c, 0x44, 0x65, 0x74, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x44, 0x61, 0x74, 0x61, 0x52, 0x16, 0x73, 0x6b, 0x79, 0x43, 0x72, 0x79, 0x73,
	0x74, 0x61, 0x6c, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x44, 0x61, 0x74, 0x61, 0x42,
	0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_AllWidgetDataNotify_proto_rawDescOnce sync.Once
	file_AllWidgetDataNotify_proto_rawDescData = file_AllWidgetDataNotify_proto_rawDesc
)

func file_AllWidgetDataNotify_proto_rawDescGZIP() []byte {
	file_AllWidgetDataNotify_proto_rawDescOnce.Do(func() {
		file_AllWidgetDataNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_AllWidgetDataNotify_proto_rawDescData)
	})
	return file_AllWidgetDataNotify_proto_rawDescData
}

var file_AllWidgetDataNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_AllWidgetDataNotify_proto_goTypes = []interface{}{
	(*AllWidgetDataNotify)(nil),          // 0: proto.AllWidgetDataNotify
	(*LunchBoxData)(nil),                 // 1: proto.LunchBoxData
	(*WidgetCoolDownData)(nil),           // 2: proto.WidgetCoolDownData
	(*AnchorPointData)(nil),              // 3: proto.AnchorPointData
	(*WidgetSlotData)(nil),               // 4: proto.WidgetSlotData
	(*ClientCollectorData)(nil),          // 5: proto.ClientCollectorData
	(*OneofGatherPointDetectorData)(nil), // 6: proto.OneofGatherPointDetectorData
	(*SkyCrystalDetectorData)(nil),       // 7: proto.SkyCrystalDetectorData
}
var file_AllWidgetDataNotify_proto_depIdxs = []int32{
	1, // 0: proto.AllWidgetDataNotify.lunch_box_data:type_name -> proto.LunchBoxData
	2, // 1: proto.AllWidgetDataNotify.cool_down_group_data_list:type_name -> proto.WidgetCoolDownData
	3, // 2: proto.AllWidgetDataNotify.anchor_point_list:type_name -> proto.AnchorPointData
	4, // 3: proto.AllWidgetDataNotify.slot_list:type_name -> proto.WidgetSlotData
	5, // 4: proto.AllWidgetDataNotify.client_collector_data_list:type_name -> proto.ClientCollectorData
	6, // 5: proto.AllWidgetDataNotify.oneof_gather_point_detector_data_list:type_name -> proto.OneofGatherPointDetectorData
	2, // 6: proto.AllWidgetDataNotify.normal_cool_down_data_list:type_name -> proto.WidgetCoolDownData
	7, // 7: proto.AllWidgetDataNotify.sky_crystal_detector_data:type_name -> proto.SkyCrystalDetectorData
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_AllWidgetDataNotify_proto_init() }
func file_AllWidgetDataNotify_proto_init() {
	if File_AllWidgetDataNotify_proto != nil {
		return
	}
	file_AnchorPointData_proto_init()
	file_ClientCollectorData_proto_init()
	file_LunchBoxData_proto_init()
	file_OneofGatherPointDetectorData_proto_init()
	file_SkyCrystalDetectorData_proto_init()
	file_WidgetCoolDownData_proto_init()
	file_WidgetSlotData_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_AllWidgetDataNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllWidgetDataNotify); i {
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
			RawDescriptor: file_AllWidgetDataNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_AllWidgetDataNotify_proto_goTypes,
		DependencyIndexes: file_AllWidgetDataNotify_proto_depIdxs,
		MessageInfos:      file_AllWidgetDataNotify_proto_msgTypes,
	}.Build()
	File_AllWidgetDataNotify_proto = out.File
	file_AllWidgetDataNotify_proto_rawDesc = nil
	file_AllWidgetDataNotify_proto_goTypes = nil
	file_AllWidgetDataNotify_proto_depIdxs = nil
}