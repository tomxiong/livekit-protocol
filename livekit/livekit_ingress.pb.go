// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: livekit_ingress.proto

package livekit

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type IngressInput int32

const (
	IngressInput_RTMP_INPUT IngressInput = 0
)

// Enum value maps for IngressInput.
var (
	IngressInput_name = map[int32]string{
		0: "RTMP_INPUT",
	}
	IngressInput_value = map[string]int32{
		"RTMP_INPUT": 0,
	}
)

func (x IngressInput) Enum() *IngressInput {
	p := new(IngressInput)
	*p = x
	return p
}

func (x IngressInput) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (IngressInput) Descriptor() protoreflect.EnumDescriptor {
	return file_livekit_ingress_proto_enumTypes[0].Descriptor()
}

func (IngressInput) Type() protoreflect.EnumType {
	return &file_livekit_ingress_proto_enumTypes[0]
}

func (x IngressInput) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use IngressInput.Descriptor instead.
func (IngressInput) EnumDescriptor() ([]byte, []int) {
	return file_livekit_ingress_proto_rawDescGZIP(), []int{0}
}

type IngressInfo_State int32

const (
	IngressInfo_ENDPOINT_WAITING IngressInfo_State = 0
	IngressInfo_ENDPOINT_ACTIVE  IngressInfo_State = 1
	IngressInfo_ENDPOINT_ENDED   IngressInfo_State = 2
)

// Enum value maps for IngressInfo_State.
var (
	IngressInfo_State_name = map[int32]string{
		0: "ENDPOINT_WAITING",
		1: "ENDPOINT_ACTIVE",
		2: "ENDPOINT_ENDED",
	}
	IngressInfo_State_value = map[string]int32{
		"ENDPOINT_WAITING": 0,
		"ENDPOINT_ACTIVE":  1,
		"ENDPOINT_ENDED":   2,
	}
)

func (x IngressInfo_State) Enum() *IngressInfo_State {
	p := new(IngressInfo_State)
	*p = x
	return p
}

func (x IngressInfo_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (IngressInfo_State) Descriptor() protoreflect.EnumDescriptor {
	return file_livekit_ingress_proto_enumTypes[1].Descriptor()
}

func (IngressInfo_State) Type() protoreflect.EnumType {
	return &file_livekit_ingress_proto_enumTypes[1]
}

func (x IngressInfo_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use IngressInfo_State.Descriptor instead.
func (IngressInfo_State) EnumDescriptor() ([]byte, []int) {
	return file_livekit_ingress_proto_rawDescGZIP(), []int{0, 0}
}

type IngressInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                  string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	InputType           IngressInput      `protobuf:"varint,2,opt,name=input_type,json=inputType,proto3,enum=livekit.IngressInput" json:"input_type,omitempty"`
	State               IngressInfo_State `protobuf:"varint,3,opt,name=state,proto3,enum=livekit.IngressInfo_State" json:"state,omitempty"`
	Room                string            `protobuf:"bytes,4,opt,name=room,proto3" json:"room,omitempty"`
	ParticipantIdentity string            `protobuf:"bytes,5,opt,name=participant_identity,json=participantIdentity,proto3" json:"participant_identity,omitempty"`
	ParticipantName     string            `protobuf:"bytes,6,opt,name=participant_name,json=participantName,proto3" json:"participant_name,omitempty"`
	Transcode           bool              `protobuf:"varint,7,opt,name=transcode,proto3" json:"transcode,omitempty"`
	// for RTMP input, it'll be a rtmp:// URL
	// for FILE input, it'll be a http:// URL
	// for SRT input, it'll be a srt:// URL
	Url    string       `protobuf:"bytes,8,opt,name=url,proto3" json:"url,omitempty"`
	Tracks []*TrackInfo `protobuf:"bytes,9,rep,name=tracks,proto3" json:"tracks,omitempty"`
}

func (x *IngressInfo) Reset() {
	*x = IngressInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_livekit_ingress_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IngressInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IngressInfo) ProtoMessage() {}

func (x *IngressInfo) ProtoReflect() protoreflect.Message {
	mi := &file_livekit_ingress_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IngressInfo.ProtoReflect.Descriptor instead.
func (*IngressInfo) Descriptor() ([]byte, []int) {
	return file_livekit_ingress_proto_rawDescGZIP(), []int{0}
}

func (x *IngressInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *IngressInfo) GetInputType() IngressInput {
	if x != nil {
		return x.InputType
	}
	return IngressInput_RTMP_INPUT
}

func (x *IngressInfo) GetState() IngressInfo_State {
	if x != nil {
		return x.State
	}
	return IngressInfo_ENDPOINT_WAITING
}

func (x *IngressInfo) GetRoom() string {
	if x != nil {
		return x.Room
	}
	return ""
}

func (x *IngressInfo) GetParticipantIdentity() string {
	if x != nil {
		return x.ParticipantIdentity
	}
	return ""
}

func (x *IngressInfo) GetParticipantName() string {
	if x != nil {
		return x.ParticipantName
	}
	return ""
}

func (x *IngressInfo) GetTranscode() bool {
	if x != nil {
		return x.Transcode
	}
	return false
}

func (x *IngressInfo) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *IngressInfo) GetTracks() []*TrackInfo {
	if x != nil {
		return x.Tracks
	}
	return nil
}

type CreateIngressRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InputType IngressInput `protobuf:"varint,1,opt,name=input_type,json=inputType,proto3,enum=livekit.IngressInput" json:"input_type,omitempty"`
	// room to publish to
	RoomName string `protobuf:"bytes,2,opt,name=room_name,json=roomName,proto3" json:"room_name,omitempty"`
	// publish as participant
	ParticipantIdentity string `protobuf:"bytes,3,opt,name=participant_identity,json=participantIdentity,proto3" json:"participant_identity,omitempty"`
	// name of publishing participant (used for display only)
	ParticipantName string               `protobuf:"bytes,4,opt,name=participant_name,json=participantName,proto3" json:"participant_name,omitempty"`
	Audio           *IngressAudioOptions `protobuf:"bytes,5,opt,name=audio,proto3" json:"audio,omitempty"`
	Video           *IngressVideoOptions `protobuf:"bytes,6,opt,name=video,proto3" json:"video,omitempty"`
	// true to disable transcoding and publish input codecs
	// when this is enabled, and input codecs aren't compatible with WebRTC, we'll
	// throw an error.
	DisableTranscode bool `protobuf:"varint,7,opt,name=disable_transcode,json=disableTranscode,proto3" json:"disable_transcode,omitempty"`
}

func (x *CreateIngressRequest) Reset() {
	*x = CreateIngressRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_livekit_ingress_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateIngressRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateIngressRequest) ProtoMessage() {}

func (x *CreateIngressRequest) ProtoReflect() protoreflect.Message {
	mi := &file_livekit_ingress_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateIngressRequest.ProtoReflect.Descriptor instead.
func (*CreateIngressRequest) Descriptor() ([]byte, []int) {
	return file_livekit_ingress_proto_rawDescGZIP(), []int{1}
}

func (x *CreateIngressRequest) GetInputType() IngressInput {
	if x != nil {
		return x.InputType
	}
	return IngressInput_RTMP_INPUT
}

func (x *CreateIngressRequest) GetRoomName() string {
	if x != nil {
		return x.RoomName
	}
	return ""
}

func (x *CreateIngressRequest) GetParticipantIdentity() string {
	if x != nil {
		return x.ParticipantIdentity
	}
	return ""
}

func (x *CreateIngressRequest) GetParticipantName() string {
	if x != nil {
		return x.ParticipantName
	}
	return ""
}

func (x *CreateIngressRequest) GetAudio() *IngressAudioOptions {
	if x != nil {
		return x.Audio
	}
	return nil
}

func (x *CreateIngressRequest) GetVideo() *IngressVideoOptions {
	if x != nil {
		return x.Video
	}
	return nil
}

func (x *CreateIngressRequest) GetDisableTranscode() bool {
	if x != nil {
		return x.DisableTranscode
	}
	return false
}

type IngressAudioOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string      `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Source TrackSource `protobuf:"varint,2,opt,name=source,proto3,enum=livekit.TrackSource" json:"source,omitempty"`
	// desired mime_type to publish to room
	MimeType string `protobuf:"bytes,3,opt,name=mime_type,json=mimeType,proto3" json:"mime_type,omitempty"`
	Bitrate  uint32 `protobuf:"varint,4,opt,name=bitrate,proto3" json:"bitrate,omitempty"`
	Dtx      bool   `protobuf:"varint,5,opt,name=dtx,proto3" json:"dtx,omitempty"`
	Channels uint32 `protobuf:"varint,6,opt,name=channels,proto3" json:"channels,omitempty"`
}

func (x *IngressAudioOptions) Reset() {
	*x = IngressAudioOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_livekit_ingress_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IngressAudioOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IngressAudioOptions) ProtoMessage() {}

func (x *IngressAudioOptions) ProtoReflect() protoreflect.Message {
	mi := &file_livekit_ingress_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IngressAudioOptions.ProtoReflect.Descriptor instead.
func (*IngressAudioOptions) Descriptor() ([]byte, []int) {
	return file_livekit_ingress_proto_rawDescGZIP(), []int{2}
}

func (x *IngressAudioOptions) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *IngressAudioOptions) GetSource() TrackSource {
	if x != nil {
		return x.Source
	}
	return TrackSource_UNKNOWN
}

func (x *IngressAudioOptions) GetMimeType() string {
	if x != nil {
		return x.MimeType
	}
	return ""
}

func (x *IngressAudioOptions) GetBitrate() uint32 {
	if x != nil {
		return x.Bitrate
	}
	return 0
}

func (x *IngressAudioOptions) GetDtx() bool {
	if x != nil {
		return x.Dtx
	}
	return false
}

func (x *IngressAudioOptions) GetChannels() uint32 {
	if x != nil {
		return x.Channels
	}
	return 0
}

type IngressVideoOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string      `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Source TrackSource `protobuf:"varint,2,opt,name=source,proto3,enum=livekit.TrackSource" json:"source,omitempty"`
	// desired mime_type to publish to room
	MimeType string `protobuf:"bytes,3,opt,name=mime_type,json=mimeType,proto3" json:"mime_type,omitempty"`
	// simulcast layers to publish, when empty, it'll pick default simulcast
	// layers at 1/2 and 1/4 of the dimensions
	Layers []*VideoLayer `protobuf:"bytes,4,rep,name=layers,proto3" json:"layers,omitempty"`
}

func (x *IngressVideoOptions) Reset() {
	*x = IngressVideoOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_livekit_ingress_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IngressVideoOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IngressVideoOptions) ProtoMessage() {}

func (x *IngressVideoOptions) ProtoReflect() protoreflect.Message {
	mi := &file_livekit_ingress_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IngressVideoOptions.ProtoReflect.Descriptor instead.
func (*IngressVideoOptions) Descriptor() ([]byte, []int) {
	return file_livekit_ingress_proto_rawDescGZIP(), []int{3}
}

func (x *IngressVideoOptions) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *IngressVideoOptions) GetSource() TrackSource {
	if x != nil {
		return x.Source
	}
	return TrackSource_UNKNOWN
}

func (x *IngressVideoOptions) GetMimeType() string {
	if x != nil {
		return x.MimeType
	}
	return ""
}

func (x *IngressVideoOptions) GetLayers() []*VideoLayer {
	if x != nil {
		return x.Layers
	}
	return nil
}

type ListIngressRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// when blank, lists all ingress endpoints
	Room string `protobuf:"bytes,1,opt,name=room,proto3" json:"room,omitempty"`
}

func (x *ListIngressRequest) Reset() {
	*x = ListIngressRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_livekit_ingress_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListIngressRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListIngressRequest) ProtoMessage() {}

func (x *ListIngressRequest) ProtoReflect() protoreflect.Message {
	mi := &file_livekit_ingress_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListIngressRequest.ProtoReflect.Descriptor instead.
func (*ListIngressRequest) Descriptor() ([]byte, []int) {
	return file_livekit_ingress_proto_rawDescGZIP(), []int{4}
}

func (x *ListIngressRequest) GetRoom() string {
	if x != nil {
		return x.Room
	}
	return ""
}

type ListIngressResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*IngressInfo `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ListIngressResponse) Reset() {
	*x = ListIngressResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_livekit_ingress_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListIngressResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListIngressResponse) ProtoMessage() {}

func (x *ListIngressResponse) ProtoReflect() protoreflect.Message {
	mi := &file_livekit_ingress_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListIngressResponse.ProtoReflect.Descriptor instead.
func (*ListIngressResponse) Descriptor() ([]byte, []int) {
	return file_livekit_ingress_proto_rawDescGZIP(), []int{5}
}

func (x *ListIngressResponse) GetItems() []*IngressInfo {
	if x != nil {
		return x.Items
	}
	return nil
}

type DeleteIngressRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteIngressRequest) Reset() {
	*x = DeleteIngressRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_livekit_ingress_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteIngressRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteIngressRequest) ProtoMessage() {}

func (x *DeleteIngressRequest) ProtoReflect() protoreflect.Message {
	mi := &file_livekit_ingress_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteIngressRequest.ProtoReflect.Descriptor instead.
func (*DeleteIngressRequest) Descriptor() ([]byte, []int) {
	return file_livekit_ingress_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteIngressRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_livekit_ingress_proto protoreflect.FileDescriptor

var file_livekit_ingress_proto_rawDesc = []byte{
	0x0a, 0x15, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x5f, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x6c,
	0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x9b, 0x03, 0x0a, 0x0b, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x34, 0x0a, 0x0a, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69,
	0x74, 0x2e, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x52, 0x09,
	0x69, 0x6e, 0x70, 0x75, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x30, 0x0a, 0x05, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b,
	0x69, 0x74, 0x2e, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x72,
	0x6f, 0x6f, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x12,
	0x31, 0x0a, 0x14, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x70,
	0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x12, 0x29, 0x0a, 0x10, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e,
	0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x70, 0x61,
	0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x09, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x72, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x2a, 0x0a,
	0x06, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x06, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x73, 0x22, 0x46, 0x0a, 0x05, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x4e, 0x44, 0x50, 0x4f, 0x49, 0x4e, 0x54, 0x5f, 0x57,
	0x41, 0x49, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x4e, 0x44, 0x50,
	0x4f, 0x49, 0x4e, 0x54, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x01, 0x12, 0x12, 0x0a,
	0x0e, 0x45, 0x4e, 0x44, 0x50, 0x4f, 0x49, 0x4e, 0x54, 0x5f, 0x45, 0x4e, 0x44, 0x45, 0x44, 0x10,
	0x02, 0x22, 0xdc, 0x02, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x67, 0x72,
	0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x0a, 0x69, 0x6e,
	0x70, 0x75, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15,
	0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x52, 0x09, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x6f, 0x6f, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x31, 0x0a,
	0x14, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x70, 0x61, 0x72,
	0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x12, 0x29, 0x0a, 0x10, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x70, 0x61, 0x72, 0x74,
	0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x32, 0x0a, 0x05, 0x61,
	0x75, 0x64, 0x69, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6c, 0x69, 0x76,
	0x65, 0x6b, 0x69, 0x74, 0x2e, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x41, 0x75, 0x64, 0x69,
	0x6f, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x05, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x12,
	0x32, 0x0a, 0x05, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73,
	0x56, 0x69, 0x64, 0x65, 0x6f, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x05, 0x76, 0x69,
	0x64, 0x65, 0x6f, 0x12, 0x2b, 0x0a, 0x11, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10,
	0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x6f, 0x64, 0x65,
	0x22, 0xbc, 0x01, 0x0a, 0x13, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x41, 0x75, 0x64, 0x69,
	0x6f, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x06,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x6c,
	0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x53, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x69,
	0x6d, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d,
	0x69, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x69, 0x74, 0x72, 0x61,
	0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x62, 0x69, 0x74, 0x72, 0x61, 0x74,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x74, 0x78, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03,
	0x64, 0x74, 0x78, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x22,
	0xa1, 0x01, 0x0a, 0x13, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x56, 0x69, 0x64, 0x65, 0x6f,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x06, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x6c, 0x69,
	0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x53, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x69, 0x6d,
	0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x69,
	0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2b, 0x0a, 0x06, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74,
	0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x06, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x73, 0x22, 0x28, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x67, 0x72, 0x65,
	0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6f,
	0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x22, 0x41, 0x0a,
	0x13, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x49, 0x6e,
	0x67, 0x72, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x22, 0x26, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x73,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x2a, 0x1e, 0x0a, 0x0c, 0x49, 0x6e, 0x67, 0x72,
	0x65, 0x73, 0x73, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x0e, 0x0a, 0x0a, 0x52, 0x54, 0x4d, 0x50,
	0x5f, 0x49, 0x4e, 0x50, 0x55, 0x54, 0x10, 0x00, 0x32, 0xe1, 0x01, 0x0a, 0x07, 0x49, 0x6e, 0x67,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x44, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6e,
	0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1d, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x49,
	0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x48, 0x0a, 0x0b, 0x4c, 0x69,
	0x73, 0x74, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1b, 0x2e, 0x6c, 0x69, 0x76, 0x65,
	0x6b, 0x69, 0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6e,
	0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1d, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x26, 0x5a, 0x24,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x6f, 0x6d, 0x78, 0x69,
	0x6f, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x6c, 0x69, 0x76,
	0x65, 0x6b, 0x69, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_livekit_ingress_proto_rawDescOnce sync.Once
	file_livekit_ingress_proto_rawDescData = file_livekit_ingress_proto_rawDesc
)

func file_livekit_ingress_proto_rawDescGZIP() []byte {
	file_livekit_ingress_proto_rawDescOnce.Do(func() {
		file_livekit_ingress_proto_rawDescData = protoimpl.X.CompressGZIP(file_livekit_ingress_proto_rawDescData)
	})
	return file_livekit_ingress_proto_rawDescData
}

var file_livekit_ingress_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_livekit_ingress_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_livekit_ingress_proto_goTypes = []interface{}{
	(IngressInput)(0),            // 0: livekit.IngressInput
	(IngressInfo_State)(0),       // 1: livekit.IngressInfo.State
	(*IngressInfo)(nil),          // 2: livekit.IngressInfo
	(*CreateIngressRequest)(nil), // 3: livekit.CreateIngressRequest
	(*IngressAudioOptions)(nil),  // 4: livekit.IngressAudioOptions
	(*IngressVideoOptions)(nil),  // 5: livekit.IngressVideoOptions
	(*ListIngressRequest)(nil),   // 6: livekit.ListIngressRequest
	(*ListIngressResponse)(nil),  // 7: livekit.ListIngressResponse
	(*DeleteIngressRequest)(nil), // 8: livekit.DeleteIngressRequest
	(*TrackInfo)(nil),            // 9: livekit.TrackInfo
	(TrackSource)(0),             // 10: livekit.TrackSource
	(*VideoLayer)(nil),           // 11: livekit.VideoLayer
	(*emptypb.Empty)(nil),        // 12: google.protobuf.Empty
}
var file_livekit_ingress_proto_depIdxs = []int32{
	0,  // 0: livekit.IngressInfo.input_type:type_name -> livekit.IngressInput
	1,  // 1: livekit.IngressInfo.state:type_name -> livekit.IngressInfo.State
	9,  // 2: livekit.IngressInfo.tracks:type_name -> livekit.TrackInfo
	0,  // 3: livekit.CreateIngressRequest.input_type:type_name -> livekit.IngressInput
	4,  // 4: livekit.CreateIngressRequest.audio:type_name -> livekit.IngressAudioOptions
	5,  // 5: livekit.CreateIngressRequest.video:type_name -> livekit.IngressVideoOptions
	10, // 6: livekit.IngressAudioOptions.source:type_name -> livekit.TrackSource
	10, // 7: livekit.IngressVideoOptions.source:type_name -> livekit.TrackSource
	11, // 8: livekit.IngressVideoOptions.layers:type_name -> livekit.VideoLayer
	2,  // 9: livekit.ListIngressResponse.items:type_name -> livekit.IngressInfo
	3,  // 10: livekit.Ingress.CreateIngress:input_type -> livekit.CreateIngressRequest
	6,  // 11: livekit.Ingress.ListIngress:input_type -> livekit.ListIngressRequest
	8,  // 12: livekit.Ingress.DeleteIngress:input_type -> livekit.DeleteIngressRequest
	2,  // 13: livekit.Ingress.CreateIngress:output_type -> livekit.IngressInfo
	7,  // 14: livekit.Ingress.ListIngress:output_type -> livekit.ListIngressResponse
	12, // 15: livekit.Ingress.DeleteIngress:output_type -> google.protobuf.Empty
	13, // [13:16] is the sub-list for method output_type
	10, // [10:13] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_livekit_ingress_proto_init() }
func file_livekit_ingress_proto_init() {
	if File_livekit_ingress_proto != nil {
		return
	}
	file_livekit_models_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_livekit_ingress_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IngressInfo); i {
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
		file_livekit_ingress_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateIngressRequest); i {
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
		file_livekit_ingress_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IngressAudioOptions); i {
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
		file_livekit_ingress_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IngressVideoOptions); i {
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
		file_livekit_ingress_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListIngressRequest); i {
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
		file_livekit_ingress_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListIngressResponse); i {
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
		file_livekit_ingress_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteIngressRequest); i {
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
			RawDescriptor: file_livekit_ingress_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_livekit_ingress_proto_goTypes,
		DependencyIndexes: file_livekit_ingress_proto_depIdxs,
		EnumInfos:         file_livekit_ingress_proto_enumTypes,
		MessageInfos:      file_livekit_ingress_proto_msgTypes,
	}.Build()
	File_livekit_ingress_proto = out.File
	file_livekit_ingress_proto_rawDesc = nil
	file_livekit_ingress_proto_goTypes = nil
	file_livekit_ingress_proto_depIdxs = nil
}
