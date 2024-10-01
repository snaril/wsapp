// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: waMsgTransport/WAMsgTransport.proto

package waMsgTransport

import (
	reflect "reflect"
	sync "sync"

	waCommon "github.com/snaril/wsapp/proto/waCommon"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"

	_ "embed"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type MessageTransport_Protocol_Ancillary_BackupDirective_ActionType int32

const (
	MessageTransport_Protocol_Ancillary_BackupDirective_NOOP              MessageTransport_Protocol_Ancillary_BackupDirective_ActionType = 0
	MessageTransport_Protocol_Ancillary_BackupDirective_UPSERT            MessageTransport_Protocol_Ancillary_BackupDirective_ActionType = 1
	MessageTransport_Protocol_Ancillary_BackupDirective_DELETE            MessageTransport_Protocol_Ancillary_BackupDirective_ActionType = 2
	MessageTransport_Protocol_Ancillary_BackupDirective_UPSERT_AND_DELETE MessageTransport_Protocol_Ancillary_BackupDirective_ActionType = 3
)

// Enum value maps for MessageTransport_Protocol_Ancillary_BackupDirective_ActionType.
var (
	MessageTransport_Protocol_Ancillary_BackupDirective_ActionType_name = map[int32]string{
		0: "NOOP",
		1: "UPSERT",
		2: "DELETE",
		3: "UPSERT_AND_DELETE",
	}
	MessageTransport_Protocol_Ancillary_BackupDirective_ActionType_value = map[string]int32{
		"NOOP":              0,
		"UPSERT":            1,
		"DELETE":            2,
		"UPSERT_AND_DELETE": 3,
	}
)

func (x MessageTransport_Protocol_Ancillary_BackupDirective_ActionType) Enum() *MessageTransport_Protocol_Ancillary_BackupDirective_ActionType {
	p := new(MessageTransport_Protocol_Ancillary_BackupDirective_ActionType)
	*p = x
	return p
}

func (x MessageTransport_Protocol_Ancillary_BackupDirective_ActionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MessageTransport_Protocol_Ancillary_BackupDirective_ActionType) Descriptor() protoreflect.EnumDescriptor {
	return file_waMsgTransport_WAMsgTransport_proto_enumTypes[0].Descriptor()
}

func (MessageTransport_Protocol_Ancillary_BackupDirective_ActionType) Type() protoreflect.EnumType {
	return &file_waMsgTransport_WAMsgTransport_proto_enumTypes[0]
}

func (x MessageTransport_Protocol_Ancillary_BackupDirective_ActionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *MessageTransport_Protocol_Ancillary_BackupDirective_ActionType) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = MessageTransport_Protocol_Ancillary_BackupDirective_ActionType(num)
	return nil
}

// Deprecated: Use MessageTransport_Protocol_Ancillary_BackupDirective_ActionType.Descriptor instead.
func (MessageTransport_Protocol_Ancillary_BackupDirective_ActionType) EnumDescriptor() ([]byte, []int) {
	return file_waMsgTransport_WAMsgTransport_proto_rawDescGZIP(), []int{0, 1, 0, 0, 0}
}

type MessageTransport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payload  *MessageTransport_Payload  `protobuf:"bytes,1,opt,name=payload" json:"payload,omitempty"`
	Protocol *MessageTransport_Protocol `protobuf:"bytes,2,opt,name=protocol" json:"protocol,omitempty"`
}

func (x *MessageTransport) Reset() {
	*x = MessageTransport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_waMsgTransport_WAMsgTransport_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageTransport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageTransport) ProtoMessage() {}

func (x *MessageTransport) ProtoReflect() protoreflect.Message {
	mi := &file_waMsgTransport_WAMsgTransport_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageTransport.ProtoReflect.Descriptor instead.
func (*MessageTransport) Descriptor() ([]byte, []int) {
	return file_waMsgTransport_WAMsgTransport_proto_rawDescGZIP(), []int{0}
}

func (x *MessageTransport) GetPayload() *MessageTransport_Payload {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *MessageTransport) GetProtocol() *MessageTransport_Protocol {
	if x != nil {
		return x.Protocol
	}
	return nil
}

type DeviceListMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SenderKeyHash      []byte  `protobuf:"bytes,1,opt,name=senderKeyHash" json:"senderKeyHash,omitempty"`
	SenderTimestamp    *uint64 `protobuf:"varint,2,opt,name=senderTimestamp" json:"senderTimestamp,omitempty"`
	RecipientKeyHash   []byte  `protobuf:"bytes,8,opt,name=recipientKeyHash" json:"recipientKeyHash,omitempty"`
	RecipientTimestamp *uint64 `protobuf:"varint,9,opt,name=recipientTimestamp" json:"recipientTimestamp,omitempty"`
}

func (x *DeviceListMetadata) Reset() {
	*x = DeviceListMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_waMsgTransport_WAMsgTransport_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceListMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceListMetadata) ProtoMessage() {}

func (x *DeviceListMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_waMsgTransport_WAMsgTransport_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceListMetadata.ProtoReflect.Descriptor instead.
func (*DeviceListMetadata) Descriptor() ([]byte, []int) {
	return file_waMsgTransport_WAMsgTransport_proto_rawDescGZIP(), []int{1}
}

func (x *DeviceListMetadata) GetSenderKeyHash() []byte {
	if x != nil {
		return x.SenderKeyHash
	}
	return nil
}

func (x *DeviceListMetadata) GetSenderTimestamp() uint64 {
	if x != nil && x.SenderTimestamp != nil {
		return *x.SenderTimestamp
	}
	return 0
}

func (x *DeviceListMetadata) GetRecipientKeyHash() []byte {
	if x != nil {
		return x.RecipientKeyHash
	}
	return nil
}

func (x *DeviceListMetadata) GetRecipientTimestamp() uint64 {
	if x != nil && x.RecipientTimestamp != nil {
		return *x.RecipientTimestamp
	}
	return 0
}

type MessageTransport_Payload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApplicationPayload *waCommon.SubProtocol         `protobuf:"bytes,1,opt,name=applicationPayload" json:"applicationPayload,omitempty"`
	FutureProof        *waCommon.FutureProofBehavior `protobuf:"varint,3,opt,name=futureProof,enum=WACommon.FutureProofBehavior" json:"futureProof,omitempty"`
}

func (x *MessageTransport_Payload) Reset() {
	*x = MessageTransport_Payload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_waMsgTransport_WAMsgTransport_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageTransport_Payload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageTransport_Payload) ProtoMessage() {}

func (x *MessageTransport_Payload) ProtoReflect() protoreflect.Message {
	mi := &file_waMsgTransport_WAMsgTransport_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageTransport_Payload.ProtoReflect.Descriptor instead.
func (*MessageTransport_Payload) Descriptor() ([]byte, []int) {
	return file_waMsgTransport_WAMsgTransport_proto_rawDescGZIP(), []int{0, 0}
}

func (x *MessageTransport_Payload) GetApplicationPayload() *waCommon.SubProtocol {
	if x != nil {
		return x.ApplicationPayload
	}
	return nil
}

func (x *MessageTransport_Payload) GetFutureProof() waCommon.FutureProofBehavior {
	if x != nil && x.FutureProof != nil {
		return *x.FutureProof
	}
	return waCommon.FutureProofBehavior(0)
}

type MessageTransport_Protocol struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Integral  *MessageTransport_Protocol_Integral  `protobuf:"bytes,1,opt,name=integral" json:"integral,omitempty"`
	Ancillary *MessageTransport_Protocol_Ancillary `protobuf:"bytes,2,opt,name=ancillary" json:"ancillary,omitempty"`
}

func (x *MessageTransport_Protocol) Reset() {
	*x = MessageTransport_Protocol{}
	if protoimpl.UnsafeEnabled {
		mi := &file_waMsgTransport_WAMsgTransport_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageTransport_Protocol) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageTransport_Protocol) ProtoMessage() {}

func (x *MessageTransport_Protocol) ProtoReflect() protoreflect.Message {
	mi := &file_waMsgTransport_WAMsgTransport_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageTransport_Protocol.ProtoReflect.Descriptor instead.
func (*MessageTransport_Protocol) Descriptor() ([]byte, []int) {
	return file_waMsgTransport_WAMsgTransport_proto_rawDescGZIP(), []int{0, 1}
}

func (x *MessageTransport_Protocol) GetIntegral() *MessageTransport_Protocol_Integral {
	if x != nil {
		return x.Integral
	}
	return nil
}

func (x *MessageTransport_Protocol) GetAncillary() *MessageTransport_Protocol_Ancillary {
	if x != nil {
		return x.Ancillary
	}
	return nil
}

type MessageTransport_Protocol_Ancillary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Skdm               *MessageTransport_Protocol_Ancillary_SenderKeyDistributionMessage `protobuf:"bytes,2,opt,name=skdm" json:"skdm,omitempty"`
	DeviceListMetadata *DeviceListMetadata                                               `protobuf:"bytes,3,opt,name=deviceListMetadata" json:"deviceListMetadata,omitempty"`
	Icdc               *MessageTransport_Protocol_Ancillary_ICDCParticipantDevices       `protobuf:"bytes,4,opt,name=icdc" json:"icdc,omitempty"`
	BackupDirective    *MessageTransport_Protocol_Ancillary_BackupDirective              `protobuf:"bytes,5,opt,name=backupDirective" json:"backupDirective,omitempty"`
}

func (x *MessageTransport_Protocol_Ancillary) Reset() {
	*x = MessageTransport_Protocol_Ancillary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_waMsgTransport_WAMsgTransport_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageTransport_Protocol_Ancillary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageTransport_Protocol_Ancillary) ProtoMessage() {}

func (x *MessageTransport_Protocol_Ancillary) ProtoReflect() protoreflect.Message {
	mi := &file_waMsgTransport_WAMsgTransport_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageTransport_Protocol_Ancillary.ProtoReflect.Descriptor instead.
func (*MessageTransport_Protocol_Ancillary) Descriptor() ([]byte, []int) {
	return file_waMsgTransport_WAMsgTransport_proto_rawDescGZIP(), []int{0, 1, 0}
}

func (x *MessageTransport_Protocol_Ancillary) GetSkdm() *MessageTransport_Protocol_Ancillary_SenderKeyDistributionMessage {
	if x != nil {
		return x.Skdm
	}
	return nil
}

func (x *MessageTransport_Protocol_Ancillary) GetDeviceListMetadata() *DeviceListMetadata {
	if x != nil {
		return x.DeviceListMetadata
	}
	return nil
}

func (x *MessageTransport_Protocol_Ancillary) GetIcdc() *MessageTransport_Protocol_Ancillary_ICDCParticipantDevices {
	if x != nil {
		return x.Icdc
	}
	return nil
}

func (x *MessageTransport_Protocol_Ancillary) GetBackupDirective() *MessageTransport_Protocol_Ancillary_BackupDirective {
	if x != nil {
		return x.BackupDirective
	}
	return nil
}

type MessageTransport_Protocol_Integral struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Padding []byte                                                `protobuf:"bytes,1,opt,name=padding" json:"padding,omitempty"`
	DSM     *MessageTransport_Protocol_Integral_DeviceSentMessage `protobuf:"bytes,2,opt,name=DSM" json:"DSM,omitempty"`
}

func (x *MessageTransport_Protocol_Integral) Reset() {
	*x = MessageTransport_Protocol_Integral{}
	if protoimpl.UnsafeEnabled {
		mi := &file_waMsgTransport_WAMsgTransport_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageTransport_Protocol_Integral) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageTransport_Protocol_Integral) ProtoMessage() {}

func (x *MessageTransport_Protocol_Integral) ProtoReflect() protoreflect.Message {
	mi := &file_waMsgTransport_WAMsgTransport_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageTransport_Protocol_Integral.ProtoReflect.Descriptor instead.
func (*MessageTransport_Protocol_Integral) Descriptor() ([]byte, []int) {
	return file_waMsgTransport_WAMsgTransport_proto_rawDescGZIP(), []int{0, 1, 1}
}

func (x *MessageTransport_Protocol_Integral) GetPadding() []byte {
	if x != nil {
		return x.Padding
	}
	return nil
}

func (x *MessageTransport_Protocol_Integral) GetDSM() *MessageTransport_Protocol_Integral_DeviceSentMessage {
	if x != nil {
		return x.DSM
	}
	return nil
}

type MessageTransport_Protocol_Ancillary_BackupDirective struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageID       *string                                                         `protobuf:"bytes,1,opt,name=messageID" json:"messageID,omitempty"`
	ActionType      *MessageTransport_Protocol_Ancillary_BackupDirective_ActionType `protobuf:"varint,2,opt,name=actionType,enum=WAMsgTransport.MessageTransport_Protocol_Ancillary_BackupDirective_ActionType" json:"actionType,omitempty"`
	SupplementalKey *string                                                         `protobuf:"bytes,3,opt,name=supplementalKey" json:"supplementalKey,omitempty"`
}

func (x *MessageTransport_Protocol_Ancillary_BackupDirective) Reset() {
	*x = MessageTransport_Protocol_Ancillary_BackupDirective{}
	if protoimpl.UnsafeEnabled {
		mi := &file_waMsgTransport_WAMsgTransport_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageTransport_Protocol_Ancillary_BackupDirective) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageTransport_Protocol_Ancillary_BackupDirective) ProtoMessage() {}

func (x *MessageTransport_Protocol_Ancillary_BackupDirective) ProtoReflect() protoreflect.Message {
	mi := &file_waMsgTransport_WAMsgTransport_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageTransport_Protocol_Ancillary_BackupDirective.ProtoReflect.Descriptor instead.
func (*MessageTransport_Protocol_Ancillary_BackupDirective) Descriptor() ([]byte, []int) {
	return file_waMsgTransport_WAMsgTransport_proto_rawDescGZIP(), []int{0, 1, 0, 0}
}

func (x *MessageTransport_Protocol_Ancillary_BackupDirective) GetMessageID() string {
	if x != nil && x.MessageID != nil {
		return *x.MessageID
	}
	return ""
}

func (x *MessageTransport_Protocol_Ancillary_BackupDirective) GetActionType() MessageTransport_Protocol_Ancillary_BackupDirective_ActionType {
	if x != nil && x.ActionType != nil {
		return *x.ActionType
	}
	return MessageTransport_Protocol_Ancillary_BackupDirective_NOOP
}

func (x *MessageTransport_Protocol_Ancillary_BackupDirective) GetSupplementalKey() string {
	if x != nil && x.SupplementalKey != nil {
		return *x.SupplementalKey
	}
	return ""
}

type MessageTransport_Protocol_Ancillary_ICDCParticipantDevices struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SenderIdentity      *MessageTransport_Protocol_Ancillary_ICDCParticipantDevices_ICDCIdentityListDescription   `protobuf:"bytes,1,opt,name=senderIdentity" json:"senderIdentity,omitempty"`
	RecipientIdentities []*MessageTransport_Protocol_Ancillary_ICDCParticipantDevices_ICDCIdentityListDescription `protobuf:"bytes,2,rep,name=recipientIdentities" json:"recipientIdentities,omitempty"`
	RecipientUserJIDs   []string                                                                                  `protobuf:"bytes,3,rep,name=recipientUserJIDs" json:"recipientUserJIDs,omitempty"`
}

func (x *MessageTransport_Protocol_Ancillary_ICDCParticipantDevices) Reset() {
	*x = MessageTransport_Protocol_Ancillary_ICDCParticipantDevices{}
	if protoimpl.UnsafeEnabled {
		mi := &file_waMsgTransport_WAMsgTransport_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageTransport_Protocol_Ancillary_ICDCParticipantDevices) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageTransport_Protocol_Ancillary_ICDCParticipantDevices) ProtoMessage() {}

func (x *MessageTransport_Protocol_Ancillary_ICDCParticipantDevices) ProtoReflect() protoreflect.Message {
	mi := &file_waMsgTransport_WAMsgTransport_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageTransport_Protocol_Ancillary_ICDCParticipantDevices.ProtoReflect.Descriptor instead.
func (*MessageTransport_Protocol_Ancillary_ICDCParticipantDevices) Descriptor() ([]byte, []int) {
	return file_waMsgTransport_WAMsgTransport_proto_rawDescGZIP(), []int{0, 1, 0, 1}
}

func (x *MessageTransport_Protocol_Ancillary_ICDCParticipantDevices) GetSenderIdentity() *MessageTransport_Protocol_Ancillary_ICDCParticipantDevices_ICDCIdentityListDescription {
	if x != nil {
		return x.SenderIdentity
	}
	return nil
}

func (x *MessageTransport_Protocol_Ancillary_ICDCParticipantDevices) GetRecipientIdentities() []*MessageTransport_Protocol_Ancillary_ICDCParticipantDevices_ICDCIdentityListDescription {
	if x != nil {
		return x.RecipientIdentities
	}
	return nil
}

func (x *MessageTransport_Protocol_Ancillary_ICDCParticipantDevices) GetRecipientUserJIDs() []string {
	if x != nil {
		return x.RecipientUserJIDs
	}
	return nil
}

type MessageTransport_Protocol_Ancillary_SenderKeyDistributionMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupID                             *string `protobuf:"bytes,1,opt,name=groupID" json:"groupID,omitempty"`
	AxolotlSenderKeyDistributionMessage []byte  `protobuf:"bytes,2,opt,name=axolotlSenderKeyDistributionMessage" json:"axolotlSenderKeyDistributionMessage,omitempty"`
}

func (x *MessageTransport_Protocol_Ancillary_SenderKeyDistributionMessage) Reset() {
	*x = MessageTransport_Protocol_Ancillary_SenderKeyDistributionMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_waMsgTransport_WAMsgTransport_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageTransport_Protocol_Ancillary_SenderKeyDistributionMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageTransport_Protocol_Ancillary_SenderKeyDistributionMessage) ProtoMessage() {}

func (x *MessageTransport_Protocol_Ancillary_SenderKeyDistributionMessage) ProtoReflect() protoreflect.Message {
	mi := &file_waMsgTransport_WAMsgTransport_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageTransport_Protocol_Ancillary_SenderKeyDistributionMessage.ProtoReflect.Descriptor instead.
func (*MessageTransport_Protocol_Ancillary_SenderKeyDistributionMessage) Descriptor() ([]byte, []int) {
	return file_waMsgTransport_WAMsgTransport_proto_rawDescGZIP(), []int{0, 1, 0, 2}
}

func (x *MessageTransport_Protocol_Ancillary_SenderKeyDistributionMessage) GetGroupID() string {
	if x != nil && x.GroupID != nil {
		return *x.GroupID
	}
	return ""
}

func (x *MessageTransport_Protocol_Ancillary_SenderKeyDistributionMessage) GetAxolotlSenderKeyDistributionMessage() []byte {
	if x != nil {
		return x.AxolotlSenderKeyDistributionMessage
	}
	return nil
}

type MessageTransport_Protocol_Ancillary_ICDCParticipantDevices_ICDCIdentityListDescription struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Seq              *int32   `protobuf:"varint,1,opt,name=seq" json:"seq,omitempty"`
	SigningDevice    []byte   `protobuf:"bytes,2,opt,name=signingDevice" json:"signingDevice,omitempty"`
	UnknownDevices   [][]byte `protobuf:"bytes,3,rep,name=unknownDevices" json:"unknownDevices,omitempty"`
	UnknownDeviceIDs []int32  `protobuf:"varint,4,rep,name=unknownDeviceIDs" json:"unknownDeviceIDs,omitempty"`
}

func (x *MessageTransport_Protocol_Ancillary_ICDCParticipantDevices_ICDCIdentityListDescription) Reset() {
	*x = MessageTransport_Protocol_Ancillary_ICDCParticipantDevices_ICDCIdentityListDescription{}
	if protoimpl.UnsafeEnabled {
		mi := &file_waMsgTransport_WAMsgTransport_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageTransport_Protocol_Ancillary_ICDCParticipantDevices_ICDCIdentityListDescription) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageTransport_Protocol_Ancillary_ICDCParticipantDevices_ICDCIdentityListDescription) ProtoMessage() {
}

func (x *MessageTransport_Protocol_Ancillary_ICDCParticipantDevices_ICDCIdentityListDescription) ProtoReflect() protoreflect.Message {
	mi := &file_waMsgTransport_WAMsgTransport_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageTransport_Protocol_Ancillary_ICDCParticipantDevices_ICDCIdentityListDescription.ProtoReflect.Descriptor instead.
func (*MessageTransport_Protocol_Ancillary_ICDCParticipantDevices_ICDCIdentityListDescription) Descriptor() ([]byte, []int) {
	return file_waMsgTransport_WAMsgTransport_proto_rawDescGZIP(), []int{0, 1, 0, 1, 0}
}

func (x *MessageTransport_Protocol_Ancillary_ICDCParticipantDevices_ICDCIdentityListDescription) GetSeq() int32 {
	if x != nil && x.Seq != nil {
		return *x.Seq
	}
	return 0
}

func (x *MessageTransport_Protocol_Ancillary_ICDCParticipantDevices_ICDCIdentityListDescription) GetSigningDevice() []byte {
	if x != nil {
		return x.SigningDevice
	}
	return nil
}

func (x *MessageTransport_Protocol_Ancillary_ICDCParticipantDevices_ICDCIdentityListDescription) GetUnknownDevices() [][]byte {
	if x != nil {
		return x.UnknownDevices
	}
	return nil
}

func (x *MessageTransport_Protocol_Ancillary_ICDCParticipantDevices_ICDCIdentityListDescription) GetUnknownDeviceIDs() []int32 {
	if x != nil {
		return x.UnknownDeviceIDs
	}
	return nil
}

type MessageTransport_Protocol_Integral_DeviceSentMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DestinationJID *string `protobuf:"bytes,1,opt,name=destinationJID" json:"destinationJID,omitempty"`
	Phash          *string `protobuf:"bytes,2,opt,name=phash" json:"phash,omitempty"`
}

func (x *MessageTransport_Protocol_Integral_DeviceSentMessage) Reset() {
	*x = MessageTransport_Protocol_Integral_DeviceSentMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_waMsgTransport_WAMsgTransport_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageTransport_Protocol_Integral_DeviceSentMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageTransport_Protocol_Integral_DeviceSentMessage) ProtoMessage() {}

func (x *MessageTransport_Protocol_Integral_DeviceSentMessage) ProtoReflect() protoreflect.Message {
	mi := &file_waMsgTransport_WAMsgTransport_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageTransport_Protocol_Integral_DeviceSentMessage.ProtoReflect.Descriptor instead.
func (*MessageTransport_Protocol_Integral_DeviceSentMessage) Descriptor() ([]byte, []int) {
	return file_waMsgTransport_WAMsgTransport_proto_rawDescGZIP(), []int{0, 1, 1, 0}
}

func (x *MessageTransport_Protocol_Integral_DeviceSentMessage) GetDestinationJID() string {
	if x != nil && x.DestinationJID != nil {
		return *x.DestinationJID
	}
	return ""
}

func (x *MessageTransport_Protocol_Integral_DeviceSentMessage) GetPhash() string {
	if x != nil && x.Phash != nil {
		return *x.Phash
	}
	return ""
}

var File_waMsgTransport_WAMsgTransport_proto protoreflect.FileDescriptor

//go:embed WAMsgTransport.pb.raw
var file_waMsgTransport_WAMsgTransport_proto_rawDesc []byte

var (
	file_waMsgTransport_WAMsgTransport_proto_rawDescOnce sync.Once
	file_waMsgTransport_WAMsgTransport_proto_rawDescData = file_waMsgTransport_WAMsgTransport_proto_rawDesc
)

func file_waMsgTransport_WAMsgTransport_proto_rawDescGZIP() []byte {
	file_waMsgTransport_WAMsgTransport_proto_rawDescOnce.Do(func() {
		file_waMsgTransport_WAMsgTransport_proto_rawDescData = protoimpl.X.CompressGZIP(file_waMsgTransport_WAMsgTransport_proto_rawDescData)
	})
	return file_waMsgTransport_WAMsgTransport_proto_rawDescData
}

var file_waMsgTransport_WAMsgTransport_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_waMsgTransport_WAMsgTransport_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_waMsgTransport_WAMsgTransport_proto_goTypes = []any{
	(MessageTransport_Protocol_Ancillary_BackupDirective_ActionType)(0), // 0: WAMsgTransport.MessageTransport.Protocol.Ancillary.BackupDirective.ActionType
	(*MessageTransport)(nil),                                                                       // 1: WAMsgTransport.MessageTransport
	(*DeviceListMetadata)(nil),                                                                     // 2: WAMsgTransport.DeviceListMetadata
	(*MessageTransport_Payload)(nil),                                                               // 3: WAMsgTransport.MessageTransport.Payload
	(*MessageTransport_Protocol)(nil),                                                              // 4: WAMsgTransport.MessageTransport.Protocol
	(*MessageTransport_Protocol_Ancillary)(nil),                                                    // 5: WAMsgTransport.MessageTransport.Protocol.Ancillary
	(*MessageTransport_Protocol_Integral)(nil),                                                     // 6: WAMsgTransport.MessageTransport.Protocol.Integral
	(*MessageTransport_Protocol_Ancillary_BackupDirective)(nil),                                    // 7: WAMsgTransport.MessageTransport.Protocol.Ancillary.BackupDirective
	(*MessageTransport_Protocol_Ancillary_ICDCParticipantDevices)(nil),                             // 8: WAMsgTransport.MessageTransport.Protocol.Ancillary.ICDCParticipantDevices
	(*MessageTransport_Protocol_Ancillary_SenderKeyDistributionMessage)(nil),                       // 9: WAMsgTransport.MessageTransport.Protocol.Ancillary.SenderKeyDistributionMessage
	(*MessageTransport_Protocol_Ancillary_ICDCParticipantDevices_ICDCIdentityListDescription)(nil), // 10: WAMsgTransport.MessageTransport.Protocol.Ancillary.ICDCParticipantDevices.ICDCIdentityListDescription
	(*MessageTransport_Protocol_Integral_DeviceSentMessage)(nil),                                   // 11: WAMsgTransport.MessageTransport.Protocol.Integral.DeviceSentMessage
	(*waCommon.SubProtocol)(nil),                                                                   // 12: WACommon.SubProtocol
	(waCommon.FutureProofBehavior)(0),                                                              // 13: WACommon.FutureProofBehavior
}
var file_waMsgTransport_WAMsgTransport_proto_depIdxs = []int32{
	3,  // 0: WAMsgTransport.MessageTransport.payload:type_name -> WAMsgTransport.MessageTransport.Payload
	4,  // 1: WAMsgTransport.MessageTransport.protocol:type_name -> WAMsgTransport.MessageTransport.Protocol
	12, // 2: WAMsgTransport.MessageTransport.Payload.applicationPayload:type_name -> WACommon.SubProtocol
	13, // 3: WAMsgTransport.MessageTransport.Payload.futureProof:type_name -> WACommon.FutureProofBehavior
	6,  // 4: WAMsgTransport.MessageTransport.Protocol.integral:type_name -> WAMsgTransport.MessageTransport.Protocol.Integral
	5,  // 5: WAMsgTransport.MessageTransport.Protocol.ancillary:type_name -> WAMsgTransport.MessageTransport.Protocol.Ancillary
	9,  // 6: WAMsgTransport.MessageTransport.Protocol.Ancillary.skdm:type_name -> WAMsgTransport.MessageTransport.Protocol.Ancillary.SenderKeyDistributionMessage
	2,  // 7: WAMsgTransport.MessageTransport.Protocol.Ancillary.deviceListMetadata:type_name -> WAMsgTransport.DeviceListMetadata
	8,  // 8: WAMsgTransport.MessageTransport.Protocol.Ancillary.icdc:type_name -> WAMsgTransport.MessageTransport.Protocol.Ancillary.ICDCParticipantDevices
	7,  // 9: WAMsgTransport.MessageTransport.Protocol.Ancillary.backupDirective:type_name -> WAMsgTransport.MessageTransport.Protocol.Ancillary.BackupDirective
	11, // 10: WAMsgTransport.MessageTransport.Protocol.Integral.DSM:type_name -> WAMsgTransport.MessageTransport.Protocol.Integral.DeviceSentMessage
	0,  // 11: WAMsgTransport.MessageTransport.Protocol.Ancillary.BackupDirective.actionType:type_name -> WAMsgTransport.MessageTransport.Protocol.Ancillary.BackupDirective.ActionType
	10, // 12: WAMsgTransport.MessageTransport.Protocol.Ancillary.ICDCParticipantDevices.senderIdentity:type_name -> WAMsgTransport.MessageTransport.Protocol.Ancillary.ICDCParticipantDevices.ICDCIdentityListDescription
	10, // 13: WAMsgTransport.MessageTransport.Protocol.Ancillary.ICDCParticipantDevices.recipientIdentities:type_name -> WAMsgTransport.MessageTransport.Protocol.Ancillary.ICDCParticipantDevices.ICDCIdentityListDescription
	14, // [14:14] is the sub-list for method output_type
	14, // [14:14] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_waMsgTransport_WAMsgTransport_proto_init() }
func file_waMsgTransport_WAMsgTransport_proto_init() {
	if File_waMsgTransport_WAMsgTransport_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_waMsgTransport_WAMsgTransport_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*MessageTransport); i {
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
		file_waMsgTransport_WAMsgTransport_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*DeviceListMetadata); i {
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
		file_waMsgTransport_WAMsgTransport_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*MessageTransport_Payload); i {
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
		file_waMsgTransport_WAMsgTransport_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*MessageTransport_Protocol); i {
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
		file_waMsgTransport_WAMsgTransport_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*MessageTransport_Protocol_Ancillary); i {
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
		file_waMsgTransport_WAMsgTransport_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*MessageTransport_Protocol_Integral); i {
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
		file_waMsgTransport_WAMsgTransport_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*MessageTransport_Protocol_Ancillary_BackupDirective); i {
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
		file_waMsgTransport_WAMsgTransport_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*MessageTransport_Protocol_Ancillary_ICDCParticipantDevices); i {
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
		file_waMsgTransport_WAMsgTransport_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*MessageTransport_Protocol_Ancillary_SenderKeyDistributionMessage); i {
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
		file_waMsgTransport_WAMsgTransport_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*MessageTransport_Protocol_Ancillary_ICDCParticipantDevices_ICDCIdentityListDescription); i {
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
		file_waMsgTransport_WAMsgTransport_proto_msgTypes[10].Exporter = func(v any, i int) any {
			switch v := v.(*MessageTransport_Protocol_Integral_DeviceSentMessage); i {
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
			RawDescriptor: file_waMsgTransport_WAMsgTransport_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_waMsgTransport_WAMsgTransport_proto_goTypes,
		DependencyIndexes: file_waMsgTransport_WAMsgTransport_proto_depIdxs,
		EnumInfos:         file_waMsgTransport_WAMsgTransport_proto_enumTypes,
		MessageInfos:      file_waMsgTransport_WAMsgTransport_proto_msgTypes,
	}.Build()
	File_waMsgTransport_WAMsgTransport_proto = out.File
	file_waMsgTransport_WAMsgTransport_proto_rawDesc = nil
	file_waMsgTransport_WAMsgTransport_proto_goTypes = nil
	file_waMsgTransport_WAMsgTransport_proto_depIdxs = nil
}
