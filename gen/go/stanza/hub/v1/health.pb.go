// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: stanza/hub/v1/health.proto

package hubv1

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Called by SDK to determine whether a Guard is overloaded at a given Feature's priority level. Used so that customer code can make good decisions about fail-fast or graceful degradation as high up the stack as possible. SDK may cache the result for up to 10 seconds.
type QueryGuardHealthRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Only tags which are used for quota management should be included here - i.e. the list of quota_tags returned by the GetGuardConfig endpoint for this Guard. If tags are in use only one quota token will be issued at a time.
	Selector *GuardFeatureSelector `protobuf:"bytes,1,opt,name=selector,proto3" json:"selector,omitempty"` // Required: GuardName, featureName, environment
	// Used to boost priority - SDK can increase or decrease priority of request, relative to normal feature priority. For instance, a customer may wish to boost the priority of paid user traffic over free tier. Priority boosts may also be negative - for example, one might deprioritise bot traffic.
	PriorityBoost *int32 `protobuf:"varint,4,opt,name=priority_boost,json=priorityBoost,proto3,oneof" json:"priority_boost,omitempty"`
}

func (x *QueryGuardHealthRequest) Reset() {
	*x = QueryGuardHealthRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stanza_hub_v1_health_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryGuardHealthRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryGuardHealthRequest) ProtoMessage() {}

func (x *QueryGuardHealthRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stanza_hub_v1_health_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryGuardHealthRequest.ProtoReflect.Descriptor instead.
func (*QueryGuardHealthRequest) Descriptor() ([]byte, []int) {
	return file_stanza_hub_v1_health_proto_rawDescGZIP(), []int{0}
}

func (x *QueryGuardHealthRequest) GetSelector() *GuardFeatureSelector {
	if x != nil {
		return x.Selector
	}
	return nil
}

func (x *QueryGuardHealthRequest) GetPriorityBoost() int32 {
	if x != nil && x.PriorityBoost != nil {
		return *x.PriorityBoost
	}
	return 0
}

type QueryGuardHealthResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Health Health `protobuf:"varint,1,opt,name=health,proto3,enum=stanza.hub.v1.Health" json:"health,omitempty"`
}

func (x *QueryGuardHealthResponse) Reset() {
	*x = QueryGuardHealthResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stanza_hub_v1_health_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryGuardHealthResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryGuardHealthResponse) ProtoMessage() {}

func (x *QueryGuardHealthResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stanza_hub_v1_health_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryGuardHealthResponse.ProtoReflect.Descriptor instead.
func (*QueryGuardHealthResponse) Descriptor() ([]byte, []int) {
	return file_stanza_hub_v1_health_proto_rawDescGZIP(), []int{1}
}

func (x *QueryGuardHealthResponse) GetHealth() Health {
	if x != nil {
		return x.Health
	}
	return Health_HEALTH_UNSPECIFIED
}

var File_stanza_hub_v1_health_proto protoreflect.FileDescriptor

var file_stanza_hub_v1_health_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x73, 0x74, 0x61, 0x6e, 0x7a, 0x61, 0x2f, 0x68, 0x75, 0x62, 0x2f, 0x76, 0x31, 0x2f,
	0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x73, 0x74,
	0x61, 0x6e, 0x7a, 0x61, 0x2e, 0x68, 0x75, 0x62, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x73, 0x74, 0x61, 0x6e, 0x7a,
	0x61, 0x2f, 0x68, 0x75, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x99, 0x01, 0x0a, 0x17, 0x51, 0x75, 0x65, 0x72, 0x79, 0x47,
	0x75, 0x61, 0x72, 0x64, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x3f, 0x0a, 0x08, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x73, 0x74, 0x61, 0x6e, 0x7a, 0x61, 0x2e, 0x68, 0x75, 0x62,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x75, 0x61, 0x72, 0x64, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x08, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x12, 0x2a, 0x0a, 0x0e, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x62,
	0x6f, 0x6f, 0x73, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x0d, 0x70, 0x72,
	0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x42, 0x6f, 0x6f, 0x73, 0x74, 0x88, 0x01, 0x01, 0x42, 0x11,
	0x0a, 0x0f, 0x5f, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x62, 0x6f, 0x6f, 0x73,
	0x74, 0x22, 0x49, 0x0a, 0x18, 0x51, 0x75, 0x65, 0x72, 0x79, 0x47, 0x75, 0x61, 0x72, 0x64, 0x48,
	0x65, 0x61, 0x6c, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a,
	0x06, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e,
	0x73, 0x74, 0x61, 0x6e, 0x7a, 0x61, 0x2e, 0x68, 0x75, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x52, 0x06, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x32, 0xc6, 0x02, 0x0a,
	0x0d, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xb4,
	0x02, 0x0a, 0x10, 0x51, 0x75, 0x65, 0x72, 0x79, 0x47, 0x75, 0x61, 0x72, 0x64, 0x48, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x12, 0x26, 0x2e, 0x73, 0x74, 0x61, 0x6e, 0x7a, 0x61, 0x2e, 0x68, 0x75, 0x62,
	0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x47, 0x75, 0x61, 0x72, 0x64, 0x48, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x73, 0x74,
	0x61, 0x6e, 0x7a, 0x61, 0x2e, 0x68, 0x75, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x47, 0x75, 0x61, 0x72, 0x64, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0xce, 0x01, 0x92, 0x41, 0xaf, 0x01, 0x12, 0x10, 0x47, 0x65, 0x74,
	0x20, 0x47, 0x75, 0x61, 0x72, 0x64, 0x20, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x1a, 0x61, 0x55,
	0x73, 0x65, 0x64, 0x20, 0x62, 0x79, 0x20, 0x53, 0x44, 0x4b, 0x20, 0x74, 0x6f, 0x20, 0x61, 0x6c,
	0x6c, 0x6f, 0x77, 0x20, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x72, 0x73, 0x20, 0x74,
	0x6f, 0x20, 0x6d, 0x61, 0x6b, 0x65, 0x20, 0x64, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x20, 0x61, 0x62, 0x6f, 0x75, 0x74, 0x20, 0x67, 0x72, 0x61, 0x63, 0x65, 0x66, 0x75, 0x6c, 0x20,
	0x64, 0x65, 0x67, 0x72, 0x61, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x6f, 0x66, 0x20, 0x62,
	0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x20, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x4a, 0x38, 0x0a, 0x03, 0x32, 0x30, 0x30, 0x12, 0x31, 0x0a, 0x02, 0x4f, 0x4b, 0x12, 0x2b, 0x0a,
	0x29, 0x1a, 0x27, 0x2e, 0x73, 0x74, 0x61, 0x6e, 0x7a, 0x61, 0x2e, 0x68, 0x75, 0x62, 0x2e, 0x76,
	0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x47, 0x75, 0x61, 0x72, 0x64, 0x48, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15,
	0x3a, 0x01, 0x2a, 0x22, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2f,
	0x67, 0x75, 0x61, 0x72, 0x64, 0x42, 0xaf, 0x01, 0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x74,
	0x61, 0x6e, 0x7a, 0x61, 0x2e, 0x68, 0x75, 0x62, 0x2e, 0x76, 0x31, 0x42, 0x0b, 0x48, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x74, 0x61, 0x6e, 0x7a, 0x61, 0x53, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x73, 0x2f, 0x68, 0x75, 0x62, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f,
	0x73, 0x74, 0x61, 0x6e, 0x7a, 0x61, 0x2f, 0x68, 0x75, 0x62, 0x2f, 0x76, 0x31, 0x3b, 0x68, 0x75,
	0x62, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x53, 0x48, 0x58, 0xaa, 0x02, 0x0d, 0x53, 0x74, 0x61, 0x6e,
	0x7a, 0x61, 0x2e, 0x48, 0x75, 0x62, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0d, 0x53, 0x74, 0x61, 0x6e,
	0x7a, 0x61, 0x5c, 0x48, 0x75, 0x62, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x19, 0x53, 0x74, 0x61, 0x6e,
	0x7a, 0x61, 0x5c, 0x48, 0x75, 0x62, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0f, 0x53, 0x74, 0x61, 0x6e, 0x7a, 0x61, 0x3a, 0x3a,
	0x48, 0x75, 0x62, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_stanza_hub_v1_health_proto_rawDescOnce sync.Once
	file_stanza_hub_v1_health_proto_rawDescData = file_stanza_hub_v1_health_proto_rawDesc
)

func file_stanza_hub_v1_health_proto_rawDescGZIP() []byte {
	file_stanza_hub_v1_health_proto_rawDescOnce.Do(func() {
		file_stanza_hub_v1_health_proto_rawDescData = protoimpl.X.CompressGZIP(file_stanza_hub_v1_health_proto_rawDescData)
	})
	return file_stanza_hub_v1_health_proto_rawDescData
}

var file_stanza_hub_v1_health_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_stanza_hub_v1_health_proto_goTypes = []interface{}{
	(*QueryGuardHealthRequest)(nil),  // 0: stanza.hub.v1.QueryGuardHealthRequest
	(*QueryGuardHealthResponse)(nil), // 1: stanza.hub.v1.QueryGuardHealthResponse
	(*GuardFeatureSelector)(nil),     // 2: stanza.hub.v1.GuardFeatureSelector
	(Health)(0),                      // 3: stanza.hub.v1.Health
}
var file_stanza_hub_v1_health_proto_depIdxs = []int32{
	2, // 0: stanza.hub.v1.QueryGuardHealthRequest.selector:type_name -> stanza.hub.v1.GuardFeatureSelector
	3, // 1: stanza.hub.v1.QueryGuardHealthResponse.health:type_name -> stanza.hub.v1.Health
	0, // 2: stanza.hub.v1.HealthService.QueryGuardHealth:input_type -> stanza.hub.v1.QueryGuardHealthRequest
	1, // 3: stanza.hub.v1.HealthService.QueryGuardHealth:output_type -> stanza.hub.v1.QueryGuardHealthResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_stanza_hub_v1_health_proto_init() }
func file_stanza_hub_v1_health_proto_init() {
	if File_stanza_hub_v1_health_proto != nil {
		return
	}
	file_stanza_hub_v1_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_stanza_hub_v1_health_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryGuardHealthRequest); i {
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
		file_stanza_hub_v1_health_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryGuardHealthResponse); i {
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
	file_stanza_hub_v1_health_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_stanza_hub_v1_health_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_stanza_hub_v1_health_proto_goTypes,
		DependencyIndexes: file_stanza_hub_v1_health_proto_depIdxs,
		MessageInfos:      file_stanza_hub_v1_health_proto_msgTypes,
	}.Build()
	File_stanza_hub_v1_health_proto = out.File
	file_stanza_hub_v1_health_proto_rawDesc = nil
	file_stanza_hub_v1_health_proto_goTypes = nil
	file_stanza_hub_v1_health_proto_depIdxs = nil
}