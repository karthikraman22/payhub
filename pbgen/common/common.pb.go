// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.11.2
// source: proto/common.proto

package common

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/golang/protobuf/ptypes/timestamp"
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

type Id struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrgId *OrgId `protobuf:"bytes,1,opt,name=orgId,proto3" json:"orgId,omitempty"`
}

func (x *Id) Reset() {
	*x = Id{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Id) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Id) ProtoMessage() {}

func (x *Id) ProtoReflect() protoreflect.Message {
	mi := &file_proto_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Id.ProtoReflect.Descriptor instead.
func (*Id) Descriptor() ([]byte, []int) {
	return file_proto_common_proto_rawDescGZIP(), []int{0}
}

func (x *Id) GetOrgId() *OrgId {
	if x != nil {
		return x.OrgId
	}
	return nil
}

type OrgId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Othr *Othr `protobuf:"bytes,1,opt,name=othr,proto3" json:"othr,omitempty"`
}

func (x *OrgId) Reset() {
	*x = OrgId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrgId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrgId) ProtoMessage() {}

func (x *OrgId) ProtoReflect() protoreflect.Message {
	mi := &file_proto_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrgId.ProtoReflect.Descriptor instead.
func (*OrgId) Descriptor() ([]byte, []int) {
	return file_proto_common_proto_rawDescGZIP(), []int{1}
}

func (x *OrgId) GetOthr() *Othr {
	if x != nil {
		return x.Othr
	}
	return nil
}

type Othr struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Othr) Reset() {
	*x = Othr{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Othr) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Othr) ProtoMessage() {}

func (x *Othr) ProtoReflect() protoreflect.Message {
	mi := &file_proto_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Othr.ProtoReflect.Descriptor instead.
func (*Othr) Descriptor() ([]byte, []int) {
	return file_proto_common_proto_rawDescGZIP(), []int{2}
}

func (x *Othr) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type InitiatingParty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id *Id `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *InitiatingParty) Reset() {
	*x = InitiatingParty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitiatingParty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitiatingParty) ProtoMessage() {}

func (x *InitiatingParty) ProtoReflect() protoreflect.Message {
	mi := &file_proto_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitiatingParty.ProtoReflect.Descriptor instead.
func (*InitiatingParty) Descriptor() ([]byte, []int) {
	return file_proto_common_proto_rawDescGZIP(), []int{3}
}

func (x *InitiatingParty) GetId() *Id {
	if x != nil {
		return x.Id
	}
	return nil
}

type Money struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CurrencyCode string  `protobuf:"bytes,1,opt,name=currency_code,json=currencyCode,proto3" json:"currency_code,omitempty"`
	Value        float64 `protobuf:"fixed64,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Money) Reset() {
	*x = Money{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_common_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Money) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Money) ProtoMessage() {}

func (x *Money) ProtoReflect() protoreflect.Message {
	mi := &file_proto_common_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Money.ProtoReflect.Descriptor instead.
func (*Money) Descriptor() ([]byte, []int) {
	return file_proto_common_proto_rawDescGZIP(), []int{4}
}

func (x *Money) GetCurrencyCode() string {
	if x != nil {
		return x.CurrencyCode
	}
	return ""
}

func (x *Money) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

type PaymentIdentification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EndToEndId    string `protobuf:"bytes,1,opt,name=end_to_end_id,json=endToEndId,proto3" json:"end_to_end_id,omitempty"`
	TransactionId string `protobuf:"bytes,2,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`
	Uetr          string `protobuf:"bytes,3,opt,name=uetr,proto3" json:"uetr,omitempty"`
}

func (x *PaymentIdentification) Reset() {
	*x = PaymentIdentification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_common_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaymentIdentification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentIdentification) ProtoMessage() {}

func (x *PaymentIdentification) ProtoReflect() protoreflect.Message {
	mi := &file_proto_common_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentIdentification.ProtoReflect.Descriptor instead.
func (*PaymentIdentification) Descriptor() ([]byte, []int) {
	return file_proto_common_proto_rawDescGZIP(), []int{5}
}

func (x *PaymentIdentification) GetEndToEndId() string {
	if x != nil {
		return x.EndToEndId
	}
	return ""
}

func (x *PaymentIdentification) GetTransactionId() string {
	if x != nil {
		return x.TransactionId
	}
	return ""
}

func (x *PaymentIdentification) GetUetr() string {
	if x != nil {
		return x.Uetr
	}
	return ""
}

type PaymentTypeInformation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceLevel    string `protobuf:"bytes,1,opt,name=service_level,json=serviceLevel,proto3" json:"service_level,omitempty"`
	LocalInstrument string `protobuf:"bytes,2,opt,name=local_instrument,json=localInstrument,proto3" json:"local_instrument,omitempty"`
	PaymentChannel  string `protobuf:"bytes,3,opt,name=payment_channel,json=paymentChannel,proto3" json:"payment_channel,omitempty"`
}

func (x *PaymentTypeInformation) Reset() {
	*x = PaymentTypeInformation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_common_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaymentTypeInformation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentTypeInformation) ProtoMessage() {}

func (x *PaymentTypeInformation) ProtoReflect() protoreflect.Message {
	mi := &file_proto_common_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentTypeInformation.ProtoReflect.Descriptor instead.
func (*PaymentTypeInformation) Descriptor() ([]byte, []int) {
	return file_proto_common_proto_rawDescGZIP(), []int{6}
}

func (x *PaymentTypeInformation) GetServiceLevel() string {
	if x != nil {
		return x.ServiceLevel
	}
	return ""
}

func (x *PaymentTypeInformation) GetLocalInstrument() string {
	if x != nil {
		return x.LocalInstrument
	}
	return ""
}

func (x *PaymentTypeInformation) GetPaymentChannel() string {
	if x != nil {
		return x.PaymentChannel
	}
	return ""
}

var File_proto_common_proto protoreflect.FileDescriptor

var file_proto_common_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x29, 0x0a, 0x02, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x05,
	0x6f, 0x72, 0x67, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4f, 0x72, 0x67, 0x49, 0x64, 0x52, 0x05, 0x6f, 0x72, 0x67, 0x49,
	0x64, 0x22, 0x29, 0x0a, 0x05, 0x4f, 0x72, 0x67, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x04, 0x6f, 0x74,
	0x68, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x4f, 0x74, 0x68, 0x72, 0x52, 0x04, 0x6f, 0x74, 0x68, 0x72, 0x22, 0x16, 0x0a, 0x04,
	0x4f, 0x74, 0x68, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x2d, 0x0a, 0x0f, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x74, 0x69,
	0x6e, 0x67, 0x50, 0x61, 0x72, 0x74, 0x79, 0x12, 0x1a, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x64, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x42, 0x0a, 0x05, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x12, 0x23, 0x0a, 0x0d,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x94, 0x01, 0x0a, 0x15, 0x50, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x2b, 0x0a, 0x0d, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x6f, 0x5f, 0x65, 0x6e, 0x64, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0xb0,
	0x01, 0x01, 0x52, 0x0a, 0x65, 0x6e, 0x64, 0x54, 0x6f, 0x45, 0x6e, 0x64, 0x49, 0x64, 0x12, 0x30,
	0x0a, 0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x20, 0x01, 0x28,
	0x24, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x12, 0x1c, 0x0a, 0x04, 0x75, 0x65, 0x74, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08,
	0xfa, 0x42, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x04, 0x75, 0x65, 0x74, 0x72, 0x22, 0x91,
	0x01, 0x0a, 0x16, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x29,
	0x0a, 0x10, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x49,
	0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x70, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x42, 0x27, 0x5a, 0x25, 0x61, 0x63, 0x68, 0x75, 0x61, 0x6c, 0x61, 0x2e, 0x69, 0x6e,
	0x2f, 0x70, 0x61, 0x79, 0x68, 0x75, 0x62, 0x2f, 0x70, 0x62, 0x67, 0x65, 0x6e, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto_common_proto_rawDescOnce sync.Once
	file_proto_common_proto_rawDescData = file_proto_common_proto_rawDesc
)

func file_proto_common_proto_rawDescGZIP() []byte {
	file_proto_common_proto_rawDescOnce.Do(func() {
		file_proto_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_common_proto_rawDescData)
	})
	return file_proto_common_proto_rawDescData
}

var file_proto_common_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_common_proto_goTypes = []interface{}{
	(*Id)(nil),                     // 0: common.Id
	(*OrgId)(nil),                  // 1: common.OrgId
	(*Othr)(nil),                   // 2: common.Othr
	(*InitiatingParty)(nil),        // 3: common.InitiatingParty
	(*Money)(nil),                  // 4: common.Money
	(*PaymentIdentification)(nil),  // 5: common.PaymentIdentification
	(*PaymentTypeInformation)(nil), // 6: common.PaymentTypeInformation
}
var file_proto_common_proto_depIdxs = []int32{
	1, // 0: common.Id.orgId:type_name -> common.OrgId
	2, // 1: common.OrgId.othr:type_name -> common.Othr
	0, // 2: common.InitiatingParty.id:type_name -> common.Id
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_common_proto_init() }
func file_proto_common_proto_init() {
	if File_proto_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Id); i {
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
		file_proto_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrgId); i {
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
		file_proto_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Othr); i {
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
		file_proto_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitiatingParty); i {
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
		file_proto_common_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Money); i {
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
		file_proto_common_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaymentIdentification); i {
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
		file_proto_common_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaymentTypeInformation); i {
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
			RawDescriptor: file_proto_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_common_proto_goTypes,
		DependencyIndexes: file_proto_common_proto_depIdxs,
		MessageInfos:      file_proto_common_proto_msgTypes,
	}.Build()
	File_proto_common_proto = out.File
	file_proto_common_proto_rawDesc = nil
	file_proto_common_proto_goTypes = nil
	file_proto_common_proto_depIdxs = nil
}