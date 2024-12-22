// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v5.29.1
// source: cinema.proto

package cinema

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

type Status int32

const (
	Status_SUCCESS Status = 0
	Status_FAILURE Status = 1
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "SUCCESS",
		1: "FAILURE",
	}
	Status_value = map[string]int32{
		"SUCCESS": 0,
		"FAILURE": 1,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_cinema_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_cinema_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_cinema_proto_rawDescGZIP(), []int{0}
}

type CinemaConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Rows    int32  `protobuf:"varint,2,opt,name=rows,proto3" json:"rows,omitempty"`
	Columns int32  `protobuf:"varint,3,opt,name=columns,proto3" json:"columns,omitempty"`
}

func (x *CinemaConfig) Reset() {
	*x = CinemaConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cinema_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CinemaConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CinemaConfig) ProtoMessage() {}

func (x *CinemaConfig) ProtoReflect() protoreflect.Message {
	mi := &file_cinema_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CinemaConfig.ProtoReflect.Descriptor instead.
func (*CinemaConfig) Descriptor() ([]byte, []int) {
	return file_cinema_proto_rawDescGZIP(), []int{0}
}

func (x *CinemaConfig) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CinemaConfig) GetRows() int32 {
	if x != nil {
		return x.Rows
	}
	return 0
}

func (x *CinemaConfig) GetColumns() int32 {
	if x != nil {
		return x.Columns
	}
	return 0
}

type QueryReservedSeatsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CinemaId int32 `protobuf:"varint,1,opt,name=cinema_id,json=cinemaId,proto3" json:"cinema_id,omitempty"`
}

func (x *QueryReservedSeatsRequest) Reset() {
	*x = QueryReservedSeatsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cinema_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryReservedSeatsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryReservedSeatsRequest) ProtoMessage() {}

func (x *QueryReservedSeatsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cinema_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryReservedSeatsRequest.ProtoReflect.Descriptor instead.
func (*QueryReservedSeatsRequest) Descriptor() ([]byte, []int) {
	return file_cinema_proto_rawDescGZIP(), []int{1}
}

func (x *QueryReservedSeatsRequest) GetCinemaId() int32 {
	if x != nil {
		return x.CinemaId
	}
	return 0
}

type QueryReservedSeatsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Seats       []*Seat `protobuf:"bytes,1,rep,name=seats,proto3" json:"seats,omitempty"`
	MaxRows     int32   `protobuf:"varint,2,opt,name=max_rows,json=maxRows,proto3" json:"max_rows,omitempty"`
	MaxColumns  int32   `protobuf:"varint,3,opt,name=max_columns,json=maxColumns,proto3" json:"max_columns,omitempty"`
	MinDistance int32   `protobuf:"varint,4,opt,name=min_distance,json=minDistance,proto3" json:"min_distance,omitempty"`
}

func (x *QueryReservedSeatsResponse) Reset() {
	*x = QueryReservedSeatsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cinema_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryReservedSeatsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryReservedSeatsResponse) ProtoMessage() {}

func (x *QueryReservedSeatsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cinema_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryReservedSeatsResponse.ProtoReflect.Descriptor instead.
func (*QueryReservedSeatsResponse) Descriptor() ([]byte, []int) {
	return file_cinema_proto_rawDescGZIP(), []int{2}
}

func (x *QueryReservedSeatsResponse) GetSeats() []*Seat {
	if x != nil {
		return x.Seats
	}
	return nil
}

func (x *QueryReservedSeatsResponse) GetMaxRows() int32 {
	if x != nil {
		return x.MaxRows
	}
	return 0
}

func (x *QueryReservedSeatsResponse) GetMaxColumns() int32 {
	if x != nil {
		return x.MaxColumns
	}
	return 0
}

func (x *QueryReservedSeatsResponse) GetMinDistance() int32 {
	if x != nil {
		return x.MinDistance
	}
	return 0
}

type Seat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Row    int32 `protobuf:"varint,1,opt,name=row,proto3" json:"row,omitempty"`
	Column int32 `protobuf:"varint,2,opt,name=column,proto3" json:"column,omitempty"`
}

func (x *Seat) Reset() {
	*x = Seat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cinema_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Seat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Seat) ProtoMessage() {}

func (x *Seat) ProtoReflect() protoreflect.Message {
	mi := &file_cinema_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Seat.ProtoReflect.Descriptor instead.
func (*Seat) Descriptor() ([]byte, []int) {
	return file_cinema_proto_rawDescGZIP(), []int{3}
}

func (x *Seat) GetRow() int32 {
	if x != nil {
		return x.Row
	}
	return 0
}

func (x *Seat) GetColumn() int32 {
	if x != nil {
		return x.Column
	}
	return 0
}

type SeatReservationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CinemaId int32   `protobuf:"varint,1,opt,name=cinema_id,json=cinemaId,proto3" json:"cinema_id,omitempty"`
	Seats    []*Seat `protobuf:"bytes,2,rep,name=seats,proto3" json:"seats,omitempty"`
}

func (x *SeatReservationRequest) Reset() {
	*x = SeatReservationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cinema_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SeatReservationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SeatReservationRequest) ProtoMessage() {}

func (x *SeatReservationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cinema_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SeatReservationRequest.ProtoReflect.Descriptor instead.
func (*SeatReservationRequest) Descriptor() ([]byte, []int) {
	return file_cinema_proto_rawDescGZIP(), []int{4}
}

func (x *SeatReservationRequest) GetCinemaId() int32 {
	if x != nil {
		return x.CinemaId
	}
	return 0
}

func (x *SeatReservationRequest) GetSeats() []*Seat {
	if x != nil {
		return x.Seats
	}
	return nil
}

type SeatCancellationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CinemaId int32   `protobuf:"varint,1,opt,name=cinema_id,json=cinemaId,proto3" json:"cinema_id,omitempty"`
	Seats    []*Seat `protobuf:"bytes,2,rep,name=seats,proto3" json:"seats,omitempty"`
}

func (x *SeatCancellationRequest) Reset() {
	*x = SeatCancellationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cinema_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SeatCancellationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SeatCancellationRequest) ProtoMessage() {}

func (x *SeatCancellationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cinema_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SeatCancellationRequest.ProtoReflect.Descriptor instead.
func (*SeatCancellationRequest) Descriptor() ([]byte, []int) {
	return file_cinema_proto_rawDescGZIP(), []int{5}
}

func (x *SeatCancellationRequest) GetCinemaId() int32 {
	if x != nil {
		return x.CinemaId
	}
	return 0
}

func (x *SeatCancellationRequest) GetSeats() []*Seat {
	if x != nil {
		return x.Seats
	}
	return nil
}

type StatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  Status `protobuf:"varint,1,opt,name=status,proto3,enum=cinema.Status" json:"status,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *StatusResponse) Reset() {
	*x = StatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cinema_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusResponse) ProtoMessage() {}

func (x *StatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cinema_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusResponse.ProtoReflect.Descriptor instead.
func (*StatusResponse) Descriptor() ([]byte, []int) {
	return file_cinema_proto_rawDescGZIP(), []int{6}
}

func (x *StatusResponse) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_SUCCESS
}

func (x *StatusResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_cinema_proto protoreflect.FileDescriptor

var file_cinema_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x69, 0x6e, 0x65, 0x6d, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x63, 0x69, 0x6e, 0x65, 0x6d, 0x61, 0x22, 0x50, 0x0a, 0x0c, 0x43, 0x69, 0x6e, 0x65, 0x6d, 0x61,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f,
	0x77, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x72, 0x6f, 0x77, 0x73, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x07, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x22, 0x38, 0x0a, 0x19, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x53, 0x65, 0x61, 0x74, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x69, 0x6e, 0x65, 0x6d, 0x61, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x63, 0x69, 0x6e, 0x65, 0x6d, 0x61,
	0x49, 0x64, 0x22, 0x9f, 0x01, 0x0a, 0x1a, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x64, 0x53, 0x65, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x22, 0x0a, 0x05, 0x73, 0x65, 0x61, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x63, 0x69, 0x6e, 0x65, 0x6d, 0x61, 0x2e, 0x53, 0x65, 0x61, 0x74, 0x52, 0x05,
	0x73, 0x65, 0x61, 0x74, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x61, 0x78, 0x5f, 0x72, 0x6f, 0x77,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6d, 0x61, 0x78, 0x52, 0x6f, 0x77, 0x73,
	0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x61, 0x78, 0x5f, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x6d, 0x61, 0x78, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e,
	0x73, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x69, 0x6e, 0x5f, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x6d, 0x69, 0x6e, 0x44, 0x69, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x22, 0x30, 0x0a, 0x04, 0x53, 0x65, 0x61, 0x74, 0x12, 0x10, 0x0a, 0x03,
	0x72, 0x6f, 0x77, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x72, 0x6f, 0x77, 0x12, 0x16,
	0x0a, 0x06, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x22, 0x59, 0x0a, 0x16, 0x53, 0x65, 0x61, 0x74, 0x52, 0x65,
	0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1b, 0x0a, 0x09, 0x63, 0x69, 0x6e, 0x65, 0x6d, 0x61, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x63, 0x69, 0x6e, 0x65, 0x6d, 0x61, 0x49, 0x64, 0x12, 0x22, 0x0a,
	0x05, 0x73, 0x65, 0x61, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63,
	0x69, 0x6e, 0x65, 0x6d, 0x61, 0x2e, 0x53, 0x65, 0x61, 0x74, 0x52, 0x05, 0x73, 0x65, 0x61, 0x74,
	0x73, 0x22, 0x5a, 0x0a, 0x17, 0x53, 0x65, 0x61, 0x74, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09,
	0x63, 0x69, 0x6e, 0x65, 0x6d, 0x61, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x63, 0x69, 0x6e, 0x65, 0x6d, 0x61, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x05, 0x73, 0x65, 0x61,
	0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x69, 0x6e, 0x65, 0x6d,
	0x61, 0x2e, 0x53, 0x65, 0x61, 0x74, 0x52, 0x05, 0x73, 0x65, 0x61, 0x74, 0x73, 0x22, 0x52, 0x0a,
	0x0e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x26, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x0e, 0x2e, 0x63, 0x69, 0x6e, 0x65, 0x6d, 0x61, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2a, 0x22, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x53,
	0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x46, 0x41, 0x49, 0x4c,
	0x55, 0x52, 0x45, 0x10, 0x01, 0x32, 0xbd, 0x02, 0x0a, 0x0d, 0x43, 0x69, 0x6e, 0x65, 0x6d, 0x61,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3f, 0x0a, 0x0f, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x75, 0x72, 0x65, 0x43, 0x69, 0x6e, 0x65, 0x6d, 0x61, 0x12, 0x14, 0x2e, 0x63, 0x69, 0x6e,
	0x65, 0x6d, 0x61, 0x2e, 0x43, 0x69, 0x6e, 0x65, 0x6d, 0x61, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x1a, 0x16, 0x2e, 0x63, 0x69, 0x6e, 0x65, 0x6d, 0x61, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5b, 0x0a, 0x12, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x53, 0x65, 0x61, 0x74, 0x73, 0x12, 0x21,
	0x2e, 0x63, 0x69, 0x6e, 0x65, 0x6d, 0x61, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x64, 0x53, 0x65, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x22, 0x2e, 0x63, 0x69, 0x6e, 0x65, 0x6d, 0x61, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x53, 0x65, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x53, 0x65, 0x61, 0x74, 0x73, 0x12, 0x1e, 0x2e, 0x63, 0x69, 0x6e, 0x65, 0x6d, 0x61, 0x2e, 0x53,
	0x65, 0x61, 0x74, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x63, 0x69, 0x6e, 0x65, 0x6d, 0x61, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a,
	0x0b, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x53, 0x65, 0x61, 0x74, 0x73, 0x12, 0x1f, 0x2e, 0x63,
	0x69, 0x6e, 0x65, 0x6d, 0x61, 0x2e, 0x53, 0x65, 0x61, 0x74, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x63, 0x69, 0x6e, 0x65, 0x6d, 0x61, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x63, 0x69, 0x6e, 0x65, 0x6d,
	0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cinema_proto_rawDescOnce sync.Once
	file_cinema_proto_rawDescData = file_cinema_proto_rawDesc
)

func file_cinema_proto_rawDescGZIP() []byte {
	file_cinema_proto_rawDescOnce.Do(func() {
		file_cinema_proto_rawDescData = protoimpl.X.CompressGZIP(file_cinema_proto_rawDescData)
	})
	return file_cinema_proto_rawDescData
}

var file_cinema_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_cinema_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_cinema_proto_goTypes = []interface{}{
	(Status)(0),                        // 0: cinema.Status
	(*CinemaConfig)(nil),               // 1: cinema.CinemaConfig
	(*QueryReservedSeatsRequest)(nil),  // 2: cinema.QueryReservedSeatsRequest
	(*QueryReservedSeatsResponse)(nil), // 3: cinema.QueryReservedSeatsResponse
	(*Seat)(nil),                       // 4: cinema.Seat
	(*SeatReservationRequest)(nil),     // 5: cinema.SeatReservationRequest
	(*SeatCancellationRequest)(nil),    // 6: cinema.SeatCancellationRequest
	(*StatusResponse)(nil),             // 7: cinema.StatusResponse
}
var file_cinema_proto_depIdxs = []int32{
	4, // 0: cinema.QueryReservedSeatsResponse.seats:type_name -> cinema.Seat
	4, // 1: cinema.SeatReservationRequest.seats:type_name -> cinema.Seat
	4, // 2: cinema.SeatCancellationRequest.seats:type_name -> cinema.Seat
	0, // 3: cinema.StatusResponse.status:type_name -> cinema.Status
	1, // 4: cinema.CinemaService.ConfigureCinema:input_type -> cinema.CinemaConfig
	2, // 5: cinema.CinemaService.QueryReservedSeats:input_type -> cinema.QueryReservedSeatsRequest
	5, // 6: cinema.CinemaService.ReserveSeats:input_type -> cinema.SeatReservationRequest
	6, // 7: cinema.CinemaService.CancelSeats:input_type -> cinema.SeatCancellationRequest
	7, // 8: cinema.CinemaService.ConfigureCinema:output_type -> cinema.StatusResponse
	3, // 9: cinema.CinemaService.QueryReservedSeats:output_type -> cinema.QueryReservedSeatsResponse
	7, // 10: cinema.CinemaService.ReserveSeats:output_type -> cinema.StatusResponse
	7, // 11: cinema.CinemaService.CancelSeats:output_type -> cinema.StatusResponse
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_cinema_proto_init() }
func file_cinema_proto_init() {
	if File_cinema_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cinema_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CinemaConfig); i {
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
		file_cinema_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryReservedSeatsRequest); i {
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
		file_cinema_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryReservedSeatsResponse); i {
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
		file_cinema_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Seat); i {
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
		file_cinema_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SeatReservationRequest); i {
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
		file_cinema_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SeatCancellationRequest); i {
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
		file_cinema_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusResponse); i {
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
			RawDescriptor: file_cinema_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cinema_proto_goTypes,
		DependencyIndexes: file_cinema_proto_depIdxs,
		EnumInfos:         file_cinema_proto_enumTypes,
		MessageInfos:      file_cinema_proto_msgTypes,
	}.Build()
	File_cinema_proto = out.File
	file_cinema_proto_rawDesc = nil
	file_cinema_proto_goTypes = nil
	file_cinema_proto_depIdxs = nil
}