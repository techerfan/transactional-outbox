// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.14.0
// source: shipment/shipment.proto

// cd contract/protobuf && protoc --go_out=../../contract/goproto/shipment --go_opt=paths=source_relative --go-grpc_out=../../contract/goproto/shipment --go-grpc_opt=paths=source_relative shipment/shipment.proto

package shipment

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

type Shipment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	OrderId        uint32 `protobuf:"varint,2,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	OrderOutboxId  uint32 `protobuf:"varint,3,opt,name=order_outbox_id,json=orderOutboxId,proto3" json:"order_outbox_id,omitempty"`
	IdempotencyKey string `protobuf:"bytes,4,opt,name=idempotency_key,json=idempotencyKey,proto3" json:"idempotency_key,omitempty"`
}

func (x *Shipment) Reset() {
	*x = Shipment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shipment_shipment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Shipment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Shipment) ProtoMessage() {}

func (x *Shipment) ProtoReflect() protoreflect.Message {
	mi := &file_shipment_shipment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Shipment.ProtoReflect.Descriptor instead.
func (*Shipment) Descriptor() ([]byte, []int) {
	return file_shipment_shipment_proto_rawDescGZIP(), []int{0}
}

func (x *Shipment) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Shipment) GetOrderId() uint32 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *Shipment) GetOrderOutboxId() uint32 {
	if x != nil {
		return x.OrderOutboxId
	}
	return 0
}

func (x *Shipment) GetIdempotencyKey() string {
	if x != nil {
		return x.IdempotencyKey
	}
	return ""
}

type CommitShipmentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Shipment *Shipment `protobuf:"bytes,1,opt,name=shipment,proto3" json:"shipment,omitempty"`
}

func (x *CommitShipmentRequest) Reset() {
	*x = CommitShipmentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shipment_shipment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommitShipmentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommitShipmentRequest) ProtoMessage() {}

func (x *CommitShipmentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shipment_shipment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommitShipmentRequest.ProtoReflect.Descriptor instead.
func (*CommitShipmentRequest) Descriptor() ([]byte, []int) {
	return file_shipment_shipment_proto_rawDescGZIP(), []int{1}
}

func (x *CommitShipmentRequest) GetShipment() *Shipment {
	if x != nil {
		return x.Shipment
	}
	return nil
}

type CommitShipmentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *CommitShipmentResponse) Reset() {
	*x = CommitShipmentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shipment_shipment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommitShipmentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommitShipmentResponse) ProtoMessage() {}

func (x *CommitShipmentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shipment_shipment_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommitShipmentResponse.ProtoReflect.Descriptor instead.
func (*CommitShipmentResponse) Descriptor() ([]byte, []int) {
	return file_shipment_shipment_proto_rawDescGZIP(), []int{2}
}

func (x *CommitShipmentResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_shipment_shipment_proto protoreflect.FileDescriptor

var file_shipment_shipment_proto_rawDesc = []byte{
	0x0a, 0x17, 0x73, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x73, 0x68, 0x69, 0x70, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x68, 0x69, 0x70, 0x6d,
	0x65, 0x6e, 0x74, 0x22, 0x86, 0x01, 0x0a, 0x08, 0x53, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x5f, 0x6f, 0x75, 0x74, 0x62, 0x6f, 0x78, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x4f, 0x75, 0x74, 0x62, 0x6f,
	0x78, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x69, 0x64, 0x65, 0x6d, 0x70, 0x6f, 0x74, 0x65, 0x6e,
	0x63, 0x79, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x69, 0x64,
	0x65, 0x6d, 0x70, 0x6f, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x4b, 0x65, 0x79, 0x22, 0x47, 0x0a, 0x15,
	0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x53, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x08, 0x73, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x68, 0x69, 0x70, 0x6d, 0x65,
	0x6e, 0x74, 0x2e, 0x53, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x08, 0x73, 0x68, 0x69,
	0x70, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x32, 0x0a, 0x16, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x53,
	0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x68, 0x0a, 0x0f, 0x53, 0x68, 0x69,
	0x70, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x55, 0x0a, 0x0e,
	0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x53, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1f,
	0x2e, 0x73, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74,
	0x53, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x20, 0x2e, 0x73, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69,
	0x74, 0x53, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x1b, 0x5a, 0x19, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2f,
	0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_shipment_shipment_proto_rawDescOnce sync.Once
	file_shipment_shipment_proto_rawDescData = file_shipment_shipment_proto_rawDesc
)

func file_shipment_shipment_proto_rawDescGZIP() []byte {
	file_shipment_shipment_proto_rawDescOnce.Do(func() {
		file_shipment_shipment_proto_rawDescData = protoimpl.X.CompressGZIP(file_shipment_shipment_proto_rawDescData)
	})
	return file_shipment_shipment_proto_rawDescData
}

var file_shipment_shipment_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_shipment_shipment_proto_goTypes = []interface{}{
	(*Shipment)(nil),               // 0: shipment.Shipment
	(*CommitShipmentRequest)(nil),  // 1: shipment.CommitShipmentRequest
	(*CommitShipmentResponse)(nil), // 2: shipment.CommitShipmentResponse
}
var file_shipment_shipment_proto_depIdxs = []int32{
	0, // 0: shipment.CommitShipmentRequest.shipment:type_name -> shipment.Shipment
	1, // 1: shipment.ShipmentService.CommitShipment:input_type -> shipment.CommitShipmentRequest
	2, // 2: shipment.ShipmentService.CommitShipment:output_type -> shipment.CommitShipmentResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_shipment_shipment_proto_init() }
func file_shipment_shipment_proto_init() {
	if File_shipment_shipment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_shipment_shipment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Shipment); i {
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
		file_shipment_shipment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommitShipmentRequest); i {
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
		file_shipment_shipment_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommitShipmentResponse); i {
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
			RawDescriptor: file_shipment_shipment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_shipment_shipment_proto_goTypes,
		DependencyIndexes: file_shipment_shipment_proto_depIdxs,
		MessageInfos:      file_shipment_shipment_proto_msgTypes,
	}.Build()
	File_shipment_shipment_proto = out.File
	file_shipment_shipment_proto_rawDesc = nil
	file_shipment_shipment_proto_goTypes = nil
	file_shipment_shipment_proto_depIdxs = nil
}
