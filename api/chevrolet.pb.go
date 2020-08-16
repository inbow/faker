// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.3
// source: api/chevrolet.proto

package api

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ChevroletPushResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SlotCount           int64                        `protobuf:"varint,1,opt,name=slot_count,json=slotCount,proto3" json:"slot_count,omitempty"`
	ImpressionKey       string                       `protobuf:"bytes,2,opt,name=impression_key,json=impressionKey,proto3" json:"impression_key,omitempty"`
	ImpressionServedUrl string                       `protobuf:"bytes,3,opt,name=impression_served_url,json=impressionServedUrl,proto3" json:"impression_served_url,omitempty"`
	Articles            []*ChevroletPushItemResponse `protobuf:"bytes,4,rep,name=articles,proto3" json:"articles,omitempty"`
}

func (x *ChevroletPushResponse) Reset() {
	*x = ChevroletPushResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_chevrolet_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChevroletPushResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChevroletPushResponse) ProtoMessage() {}

func (x *ChevroletPushResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_chevrolet_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChevroletPushResponse.ProtoReflect.Descriptor instead.
func (*ChevroletPushResponse) Descriptor() ([]byte, []int) {
	return file_api_chevrolet_proto_rawDescGZIP(), []int{0}
}

func (x *ChevroletPushResponse) GetSlotCount() int64 {
	if x != nil {
		return x.SlotCount
	}
	return 0
}

func (x *ChevroletPushResponse) GetImpressionKey() string {
	if x != nil {
		return x.ImpressionKey
	}
	return ""
}

func (x *ChevroletPushResponse) GetImpressionServedUrl() string {
	if x != nil {
		return x.ImpressionServedUrl
	}
	return ""
}

func (x *ChevroletPushResponse) GetArticles() []*ChevroletPushItemResponse {
	if x != nil {
		return x.Articles
	}
	return nil
}

type ChevroletPushItemResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Slot        int64   `protobuf:"varint,1,opt,name=slot,proto3" json:"slot,omitempty"`
	Url         string  `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	Title       string  `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Image       string  `protobuf:"bytes,4,opt,name=image,proto3" json:"image,omitempty"`
	Sponsored   bool    `protobuf:"varint,5,opt,name=sponsored,proto3" json:"sponsored,omitempty"`
	SponsoredBy string  `protobuf:"bytes,6,opt,name=sponsored_by,json=sponsoredBy,proto3" json:"sponsored_by,omitempty"`
	Cpc         float64 `protobuf:"fixed64,7,opt,name=cpc,proto3" json:"cpc,omitempty"`
	Score       float64 `protobuf:"fixed64,8,opt,name=score,proto3" json:"score,omitempty"`
}

func (x *ChevroletPushItemResponse) Reset() {
	*x = ChevroletPushItemResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_chevrolet_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChevroletPushItemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChevroletPushItemResponse) ProtoMessage() {}

func (x *ChevroletPushItemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_chevrolet_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChevroletPushItemResponse.ProtoReflect.Descriptor instead.
func (*ChevroletPushItemResponse) Descriptor() ([]byte, []int) {
	return file_api_chevrolet_proto_rawDescGZIP(), []int{1}
}

func (x *ChevroletPushItemResponse) GetSlot() int64 {
	if x != nil {
		return x.Slot
	}
	return 0
}

func (x *ChevroletPushItemResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *ChevroletPushItemResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ChevroletPushItemResponse) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *ChevroletPushItemResponse) GetSponsored() bool {
	if x != nil {
		return x.Sponsored
	}
	return false
}

func (x *ChevroletPushItemResponse) GetSponsoredBy() string {
	if x != nil {
		return x.SponsoredBy
	}
	return ""
}

func (x *ChevroletPushItemResponse) GetCpc() float64 {
	if x != nil {
		return x.Cpc
	}
	return 0
}

func (x *ChevroletPushItemResponse) GetScore() float64 {
	if x != nil {
		return x.Score
	}
	return 0
}

var File_api_chevrolet_proto protoreflect.FileDescriptor

var file_api_chevrolet_proto_rawDesc = []byte{
	0x0a, 0x13, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x68, 0x65, 0x76, 0x72, 0x6f, 0x6c, 0x65, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc9, 0x01, 0x0a, 0x15, 0x43, 0x68, 0x65, 0x76, 0x72, 0x6f,
	0x6c, 0x65, 0x74, 0x50, 0x75, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x73, 0x6c, 0x6f, 0x74, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x6c, 0x6f, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x25,
	0x0a, 0x0e, 0x69, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x6b, 0x65, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x69, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x12, 0x32, 0x0a, 0x15, 0x69, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x69, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x64, 0x55, 0x72, 0x6c, 0x12, 0x36, 0x0a, 0x08, 0x61, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x43, 0x68,
	0x65, 0x76, 0x72, 0x6f, 0x6c, 0x65, 0x74, 0x50, 0x75, 0x73, 0x68, 0x49, 0x74, 0x65, 0x6d, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65,
	0x73, 0x22, 0xd6, 0x01, 0x0a, 0x19, 0x43, 0x68, 0x65, 0x76, 0x72, 0x6f, 0x6c, 0x65, 0x74, 0x50,
	0x75, 0x73, 0x68, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x73, 0x6c, 0x6f, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73,
	0x6c, 0x6f, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x6f, 0x72, 0x65, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x6f, 0x72, 0x65, 0x64, 0x12,
	0x21, 0x0a, 0x0c, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x6f, 0x72, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x6f, 0x72, 0x65, 0x64,
	0x42, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x70, 0x63, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x03, 0x63, 0x70, 0x63, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x42, 0x1e, 0x5a, 0x1c, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x78, 0x79, 0x64, 0x2d, 0x69, 0x6f,
	0x2f, 0x66, 0x61, 0x6b, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_api_chevrolet_proto_rawDescOnce sync.Once
	file_api_chevrolet_proto_rawDescData = file_api_chevrolet_proto_rawDesc
)

func file_api_chevrolet_proto_rawDescGZIP() []byte {
	file_api_chevrolet_proto_rawDescOnce.Do(func() {
		file_api_chevrolet_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_chevrolet_proto_rawDescData)
	})
	return file_api_chevrolet_proto_rawDescData
}

var file_api_chevrolet_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_chevrolet_proto_goTypes = []interface{}{
	(*ChevroletPushResponse)(nil),     // 0: ChevroletPushResponse
	(*ChevroletPushItemResponse)(nil), // 1: ChevroletPushItemResponse
}
var file_api_chevrolet_proto_depIdxs = []int32{
	1, // 0: ChevroletPushResponse.articles:type_name -> ChevroletPushItemResponse
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_chevrolet_proto_init() }
func file_api_chevrolet_proto_init() {
	if File_api_chevrolet_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_chevrolet_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChevroletPushResponse); i {
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
		file_api_chevrolet_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChevroletPushItemResponse); i {
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
			RawDescriptor: file_api_chevrolet_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_chevrolet_proto_goTypes,
		DependencyIndexes: file_api_chevrolet_proto_depIdxs,
		MessageInfos:      file_api_chevrolet_proto_msgTypes,
	}.Build()
	File_api_chevrolet_proto = out.File
	file_api_chevrolet_proto_rawDesc = nil
	file_api_chevrolet_proto_goTypes = nil
	file_api_chevrolet_proto_depIdxs = nil
}