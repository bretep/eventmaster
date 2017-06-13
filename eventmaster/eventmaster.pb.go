// Code generated by protoc-gen-go.
// source: eventmaster.proto
// DO NOT EDIT!

/*
Package eventmaster is a generated protocol buffer package.

It is generated from these files:
	eventmaster.proto

It has these top-level messages:
	Event
	Query
	Topic
	TopicResult
	UpdateTopicRequest
	DeleteTopicRequest
	Dc
	DcResult
	UpdateDcRequest
	WriteResponse
	EmptyRequest
*/
package eventmaster

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Event struct {
	EventId       string   `protobuf:"bytes,1,opt,name=event_id,json=eventId" json:"event_id,omitempty"`
	ParentEventId string   `protobuf:"bytes,2,opt,name=parent_event_id,json=parentEventId" json:"parent_event_id,omitempty"`
	EventTime     int64    `protobuf:"varint,3,opt,name=event_time,json=eventTime" json:"event_time,omitempty"`
	Dc            string   `protobuf:"bytes,4,opt,name=dc" json:"dc,omitempty"`
	TopicName     string   `protobuf:"bytes,5,opt,name=topic_name,json=topicName" json:"topic_name,omitempty"`
	TagSet        []string `protobuf:"bytes,6,rep,name=tag_set,json=tagSet" json:"tag_set,omitempty"`
	Host          string   `protobuf:"bytes,7,opt,name=host" json:"host,omitempty"`
	TargetHostSet []string `protobuf:"bytes,8,rep,name=target_host_set,json=targetHostSet" json:"target_host_set,omitempty"`
	User          string   `protobuf:"bytes,9,opt,name=user" json:"user,omitempty"`
	Data          []byte   `protobuf:"bytes,10,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *Event) Reset()                    { *m = Event{} }
func (m *Event) String() string            { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()               {}
func (*Event) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Event) GetEventId() string {
	if m != nil {
		return m.EventId
	}
	return ""
}

func (m *Event) GetParentEventId() string {
	if m != nil {
		return m.ParentEventId
	}
	return ""
}

func (m *Event) GetEventTime() int64 {
	if m != nil {
		return m.EventTime
	}
	return 0
}

func (m *Event) GetDc() string {
	if m != nil {
		return m.Dc
	}
	return ""
}

func (m *Event) GetTopicName() string {
	if m != nil {
		return m.TopicName
	}
	return ""
}

func (m *Event) GetTagSet() []string {
	if m != nil {
		return m.TagSet
	}
	return nil
}

func (m *Event) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *Event) GetTargetHostSet() []string {
	if m != nil {
		return m.TargetHostSet
	}
	return nil
}

func (m *Event) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *Event) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type Query struct {
	Dc                []string `protobuf:"bytes,1,rep,name=dc" json:"dc,omitempty"`
	Host              []string `protobuf:"bytes,2,rep,name=host" json:"host,omitempty"`
	TargetHostSet     []string `protobuf:"bytes,3,rep,name=target_host_set,json=targetHostSet" json:"target_host_set,omitempty"`
	User              []string `protobuf:"bytes,4,rep,name=user" json:"user,omitempty"`
	TopicName         []string `protobuf:"bytes,5,rep,name=topic_name,json=topicName" json:"topic_name,omitempty"`
	TagSet            []string `protobuf:"bytes,6,rep,name=tag_set,json=tagSet" json:"tag_set,omitempty"`
	ParentEventId     []string `protobuf:"bytes,7,rep,name=parent_event_id,json=parentEventId" json:"parent_event_id,omitempty"`
	Data              string   `protobuf:"bytes,8,opt,name=data" json:"data,omitempty"`
	StartEventTime    int64    `protobuf:"varint,9,opt,name=start_event_time,json=startEventTime" json:"start_event_time,omitempty"`
	EndEventTime      int64    `protobuf:"varint,10,opt,name=end_event_time,json=endEventTime" json:"end_event_time,omitempty"`
	StartReceivedTime int64    `protobuf:"varint,11,opt,name=start_received_time,json=startReceivedTime" json:"start_received_time,omitempty"`
	EndReceivedTime   int64    `protobuf:"varint,12,opt,name=end_received_time,json=endReceivedTime" json:"end_received_time,omitempty"`
	SortField         []string `protobuf:"bytes,13,rep,name=sort_field,json=sortField" json:"sort_field,omitempty"`
	SortAscending     []bool   `protobuf:"varint,14,rep,packed,name=sort_ascending,json=sortAscending" json:"sort_ascending,omitempty"`
	Start             int32    `protobuf:"varint,15,opt,name=start" json:"start,omitempty"`
	Limit             int32    `protobuf:"varint,16,opt,name=limit" json:"limit,omitempty"`
}

func (m *Query) Reset()                    { *m = Query{} }
func (m *Query) String() string            { return proto.CompactTextString(m) }
func (*Query) ProtoMessage()               {}
func (*Query) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Query) GetDc() []string {
	if m != nil {
		return m.Dc
	}
	return nil
}

func (m *Query) GetHost() []string {
	if m != nil {
		return m.Host
	}
	return nil
}

func (m *Query) GetTargetHostSet() []string {
	if m != nil {
		return m.TargetHostSet
	}
	return nil
}

func (m *Query) GetUser() []string {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *Query) GetTopicName() []string {
	if m != nil {
		return m.TopicName
	}
	return nil
}

func (m *Query) GetTagSet() []string {
	if m != nil {
		return m.TagSet
	}
	return nil
}

func (m *Query) GetParentEventId() []string {
	if m != nil {
		return m.ParentEventId
	}
	return nil
}

func (m *Query) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *Query) GetStartEventTime() int64 {
	if m != nil {
		return m.StartEventTime
	}
	return 0
}

func (m *Query) GetEndEventTime() int64 {
	if m != nil {
		return m.EndEventTime
	}
	return 0
}

func (m *Query) GetStartReceivedTime() int64 {
	if m != nil {
		return m.StartReceivedTime
	}
	return 0
}

func (m *Query) GetEndReceivedTime() int64 {
	if m != nil {
		return m.EndReceivedTime
	}
	return 0
}

func (m *Query) GetSortField() []string {
	if m != nil {
		return m.SortField
	}
	return nil
}

func (m *Query) GetSortAscending() []bool {
	if m != nil {
		return m.SortAscending
	}
	return nil
}

func (m *Query) GetStart() int32 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *Query) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type Topic struct {
	TopicName  string `protobuf:"bytes,1,opt,name=topic_name,json=topicName" json:"topic_name,omitempty"`
	DataSchema []byte `protobuf:"bytes,2,opt,name=data_schema,json=dataSchema,proto3" json:"data_schema,omitempty"`
}

func (m *Topic) Reset()                    { *m = Topic{} }
func (m *Topic) String() string            { return proto.CompactTextString(m) }
func (*Topic) ProtoMessage()               {}
func (*Topic) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Topic) GetTopicName() string {
	if m != nil {
		return m.TopicName
	}
	return ""
}

func (m *Topic) GetDataSchema() []byte {
	if m != nil {
		return m.DataSchema
	}
	return nil
}

type TopicResult struct {
	Results []*Topic `protobuf:"bytes,1,rep,name=results" json:"results,omitempty"`
}

func (m *TopicResult) Reset()                    { *m = TopicResult{} }
func (m *TopicResult) String() string            { return proto.CompactTextString(m) }
func (*TopicResult) ProtoMessage()               {}
func (*TopicResult) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *TopicResult) GetResults() []*Topic {
	if m != nil {
		return m.Results
	}
	return nil
}

type UpdateTopicRequest struct {
	OldName    string `protobuf:"bytes,1,opt,name=old_name,json=oldName" json:"old_name,omitempty"`
	NewName    string `protobuf:"bytes,2,opt,name=new_name,json=newName" json:"new_name,omitempty"`
	DataSchema []byte `protobuf:"bytes,3,opt,name=data_schema,json=dataSchema,proto3" json:"data_schema,omitempty"`
}

func (m *UpdateTopicRequest) Reset()                    { *m = UpdateTopicRequest{} }
func (m *UpdateTopicRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateTopicRequest) ProtoMessage()               {}
func (*UpdateTopicRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *UpdateTopicRequest) GetOldName() string {
	if m != nil {
		return m.OldName
	}
	return ""
}

func (m *UpdateTopicRequest) GetNewName() string {
	if m != nil {
		return m.NewName
	}
	return ""
}

func (m *UpdateTopicRequest) GetDataSchema() []byte {
	if m != nil {
		return m.DataSchema
	}
	return nil
}

type DeleteTopicRequest struct {
	TopicName string `protobuf:"bytes,1,opt,name=topic_name,json=topicName" json:"topic_name,omitempty"`
}

func (m *DeleteTopicRequest) Reset()                    { *m = DeleteTopicRequest{} }
func (m *DeleteTopicRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteTopicRequest) ProtoMessage()               {}
func (*DeleteTopicRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *DeleteTopicRequest) GetTopicName() string {
	if m != nil {
		return m.TopicName
	}
	return ""
}

type Dc struct {
	Dc string `protobuf:"bytes,1,opt,name=dc" json:"dc,omitempty"`
}

func (m *Dc) Reset()                    { *m = Dc{} }
func (m *Dc) String() string            { return proto.CompactTextString(m) }
func (*Dc) ProtoMessage()               {}
func (*Dc) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Dc) GetDc() string {
	if m != nil {
		return m.Dc
	}
	return ""
}

type DcResult struct {
	Results []*Dc `protobuf:"bytes,1,rep,name=results" json:"results,omitempty"`
}

func (m *DcResult) Reset()                    { *m = DcResult{} }
func (m *DcResult) String() string            { return proto.CompactTextString(m) }
func (*DcResult) ProtoMessage()               {}
func (*DcResult) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *DcResult) GetResults() []*Dc {
	if m != nil {
		return m.Results
	}
	return nil
}

type UpdateDcRequest struct {
	OldName string `protobuf:"bytes,1,opt,name=old_name,json=oldName" json:"old_name,omitempty"`
	NewName string `protobuf:"bytes,2,opt,name=new_name,json=newName" json:"new_name,omitempty"`
}

func (m *UpdateDcRequest) Reset()                    { *m = UpdateDcRequest{} }
func (m *UpdateDcRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateDcRequest) ProtoMessage()               {}
func (*UpdateDcRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *UpdateDcRequest) GetOldName() string {
	if m != nil {
		return m.OldName
	}
	return ""
}

func (m *UpdateDcRequest) GetNewName() string {
	if m != nil {
		return m.NewName
	}
	return ""
}

type WriteResponse struct {
	Id string `protobuf:"bytes,3,opt,name=id" json:"id,omitempty"`
}

func (m *WriteResponse) Reset()                    { *m = WriteResponse{} }
func (m *WriteResponse) String() string            { return proto.CompactTextString(m) }
func (*WriteResponse) ProtoMessage()               {}
func (*WriteResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *WriteResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type EmptyRequest struct {
}

func (m *EmptyRequest) Reset()                    { *m = EmptyRequest{} }
func (m *EmptyRequest) String() string            { return proto.CompactTextString(m) }
func (*EmptyRequest) ProtoMessage()               {}
func (*EmptyRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func init() {
	proto.RegisterType((*Event)(nil), "eventmaster.Event")
	proto.RegisterType((*Query)(nil), "eventmaster.Query")
	proto.RegisterType((*Topic)(nil), "eventmaster.Topic")
	proto.RegisterType((*TopicResult)(nil), "eventmaster.TopicResult")
	proto.RegisterType((*UpdateTopicRequest)(nil), "eventmaster.UpdateTopicRequest")
	proto.RegisterType((*DeleteTopicRequest)(nil), "eventmaster.DeleteTopicRequest")
	proto.RegisterType((*Dc)(nil), "eventmaster.Dc")
	proto.RegisterType((*DcResult)(nil), "eventmaster.DcResult")
	proto.RegisterType((*UpdateDcRequest)(nil), "eventmaster.UpdateDcRequest")
	proto.RegisterType((*WriteResponse)(nil), "eventmaster.WriteResponse")
	proto.RegisterType((*EmptyRequest)(nil), "eventmaster.EmptyRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for EventMaster service

type EventMasterClient interface {
	AddEvent(ctx context.Context, in *Event, opts ...grpc.CallOption) (*WriteResponse, error)
	GetEvents(ctx context.Context, in *Query, opts ...grpc.CallOption) (EventMaster_GetEventsClient, error)
	AddTopic(ctx context.Context, in *Topic, opts ...grpc.CallOption) (*WriteResponse, error)
	UpdateTopic(ctx context.Context, in *UpdateTopicRequest, opts ...grpc.CallOption) (*WriteResponse, error)
	DeleteTopic(ctx context.Context, in *DeleteTopicRequest, opts ...grpc.CallOption) (*WriteResponse, error)
	GetTopics(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*TopicResult, error)
	AddDc(ctx context.Context, in *Dc, opts ...grpc.CallOption) (*WriteResponse, error)
	UpdateDc(ctx context.Context, in *UpdateDcRequest, opts ...grpc.CallOption) (*WriteResponse, error)
	GetDcs(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*DcResult, error)
}

type eventMasterClient struct {
	cc *grpc.ClientConn
}

func NewEventMasterClient(cc *grpc.ClientConn) EventMasterClient {
	return &eventMasterClient{cc}
}

func (c *eventMasterClient) AddEvent(ctx context.Context, in *Event, opts ...grpc.CallOption) (*WriteResponse, error) {
	out := new(WriteResponse)
	err := grpc.Invoke(ctx, "/eventmaster.EventMaster/AddEvent", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventMasterClient) GetEvents(ctx context.Context, in *Query, opts ...grpc.CallOption) (EventMaster_GetEventsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_EventMaster_serviceDesc.Streams[0], c.cc, "/eventmaster.EventMaster/GetEvents", opts...)
	if err != nil {
		return nil, err
	}
	x := &eventMasterGetEventsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type EventMaster_GetEventsClient interface {
	Recv() (*Event, error)
	grpc.ClientStream
}

type eventMasterGetEventsClient struct {
	grpc.ClientStream
}

func (x *eventMasterGetEventsClient) Recv() (*Event, error) {
	m := new(Event)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *eventMasterClient) AddTopic(ctx context.Context, in *Topic, opts ...grpc.CallOption) (*WriteResponse, error) {
	out := new(WriteResponse)
	err := grpc.Invoke(ctx, "/eventmaster.EventMaster/AddTopic", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventMasterClient) UpdateTopic(ctx context.Context, in *UpdateTopicRequest, opts ...grpc.CallOption) (*WriteResponse, error) {
	out := new(WriteResponse)
	err := grpc.Invoke(ctx, "/eventmaster.EventMaster/UpdateTopic", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventMasterClient) DeleteTopic(ctx context.Context, in *DeleteTopicRequest, opts ...grpc.CallOption) (*WriteResponse, error) {
	out := new(WriteResponse)
	err := grpc.Invoke(ctx, "/eventmaster.EventMaster/DeleteTopic", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventMasterClient) GetTopics(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*TopicResult, error) {
	out := new(TopicResult)
	err := grpc.Invoke(ctx, "/eventmaster.EventMaster/GetTopics", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventMasterClient) AddDc(ctx context.Context, in *Dc, opts ...grpc.CallOption) (*WriteResponse, error) {
	out := new(WriteResponse)
	err := grpc.Invoke(ctx, "/eventmaster.EventMaster/AddDc", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventMasterClient) UpdateDc(ctx context.Context, in *UpdateDcRequest, opts ...grpc.CallOption) (*WriteResponse, error) {
	out := new(WriteResponse)
	err := grpc.Invoke(ctx, "/eventmaster.EventMaster/UpdateDc", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventMasterClient) GetDcs(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*DcResult, error) {
	out := new(DcResult)
	err := grpc.Invoke(ctx, "/eventmaster.EventMaster/GetDcs", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for EventMaster service

type EventMasterServer interface {
	AddEvent(context.Context, *Event) (*WriteResponse, error)
	GetEvents(*Query, EventMaster_GetEventsServer) error
	AddTopic(context.Context, *Topic) (*WriteResponse, error)
	UpdateTopic(context.Context, *UpdateTopicRequest) (*WriteResponse, error)
	DeleteTopic(context.Context, *DeleteTopicRequest) (*WriteResponse, error)
	GetTopics(context.Context, *EmptyRequest) (*TopicResult, error)
	AddDc(context.Context, *Dc) (*WriteResponse, error)
	UpdateDc(context.Context, *UpdateDcRequest) (*WriteResponse, error)
	GetDcs(context.Context, *EmptyRequest) (*DcResult, error)
}

func RegisterEventMasterServer(s *grpc.Server, srv EventMasterServer) {
	s.RegisterService(&_EventMaster_serviceDesc, srv)
}

func _EventMaster_AddEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Event)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventMasterServer).AddEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eventmaster.EventMaster/AddEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventMasterServer).AddEvent(ctx, req.(*Event))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventMaster_GetEvents_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Query)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EventMasterServer).GetEvents(m, &eventMasterGetEventsServer{stream})
}

type EventMaster_GetEventsServer interface {
	Send(*Event) error
	grpc.ServerStream
}

type eventMasterGetEventsServer struct {
	grpc.ServerStream
}

func (x *eventMasterGetEventsServer) Send(m *Event) error {
	return x.ServerStream.SendMsg(m)
}

func _EventMaster_AddTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Topic)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventMasterServer).AddTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eventmaster.EventMaster/AddTopic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventMasterServer).AddTopic(ctx, req.(*Topic))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventMaster_UpdateTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTopicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventMasterServer).UpdateTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eventmaster.EventMaster/UpdateTopic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventMasterServer).UpdateTopic(ctx, req.(*UpdateTopicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventMaster_DeleteTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTopicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventMasterServer).DeleteTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eventmaster.EventMaster/DeleteTopic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventMasterServer).DeleteTopic(ctx, req.(*DeleteTopicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventMaster_GetTopics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventMasterServer).GetTopics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eventmaster.EventMaster/GetTopics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventMasterServer).GetTopics(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventMaster_AddDc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Dc)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventMasterServer).AddDc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eventmaster.EventMaster/AddDc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventMasterServer).AddDc(ctx, req.(*Dc))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventMaster_UpdateDc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDcRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventMasterServer).UpdateDc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eventmaster.EventMaster/UpdateDc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventMasterServer).UpdateDc(ctx, req.(*UpdateDcRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventMaster_GetDcs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventMasterServer).GetDcs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eventmaster.EventMaster/GetDcs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventMasterServer).GetDcs(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _EventMaster_serviceDesc = grpc.ServiceDesc{
	ServiceName: "eventmaster.EventMaster",
	HandlerType: (*EventMasterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddEvent",
			Handler:    _EventMaster_AddEvent_Handler,
		},
		{
			MethodName: "AddTopic",
			Handler:    _EventMaster_AddTopic_Handler,
		},
		{
			MethodName: "UpdateTopic",
			Handler:    _EventMaster_UpdateTopic_Handler,
		},
		{
			MethodName: "DeleteTopic",
			Handler:    _EventMaster_DeleteTopic_Handler,
		},
		{
			MethodName: "GetTopics",
			Handler:    _EventMaster_GetTopics_Handler,
		},
		{
			MethodName: "AddDc",
			Handler:    _EventMaster_AddDc_Handler,
		},
		{
			MethodName: "UpdateDc",
			Handler:    _EventMaster_UpdateDc_Handler,
		},
		{
			MethodName: "GetDcs",
			Handler:    _EventMaster_GetDcs_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetEvents",
			Handler:       _EventMaster_GetEvents_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "eventmaster.proto",
}

func init() { proto.RegisterFile("eventmaster.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 721 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x55, 0xdd, 0x6e, 0xd3, 0x4c,
	0x10, 0x6d, 0xe2, 0x38, 0x76, 0xc6, 0xf9, 0x69, 0xf7, 0xeb, 0x27, 0xdc, 0x08, 0xd4, 0xca, 0x02,
	0x54, 0x10, 0xaa, 0x50, 0x2b, 0xe0, 0x82, 0xde, 0x14, 0x39, 0x2d, 0x48, 0x80, 0x84, 0x5b, 0xc4,
	0x65, 0x64, 0xe2, 0xa5, 0xb5, 0x48, 0xe2, 0xe0, 0xdd, 0x14, 0xf5, 0x41, 0x78, 0x3f, 0x5e, 0x83,
	0x3b, 0x66, 0x67, 0xed, 0x62, 0xbb, 0x21, 0xa9, 0xc4, 0xdd, 0xee, 0x99, 0x33, 0xbf, 0x7b, 0xc6,
	0x86, 0x0d, 0x7e, 0xc9, 0xa7, 0x72, 0x12, 0x0a, 0xc9, 0xd3, 0xbd, 0x59, 0x9a, 0xc8, 0x84, 0x39,
	0x05, 0xc8, 0xfb, 0x51, 0x07, 0x73, 0xa0, 0xee, 0x6c, 0x0b, 0x6c, 0x32, 0x0c, 0xe3, 0xc8, 0xad,
	0xed, 0xd4, 0x76, 0x5b, 0x81, 0x45, 0xf7, 0x37, 0x11, 0x7b, 0x08, 0xbd, 0x59, 0x98, 0x2a, 0xdb,
	0x35, 0xa3, 0x4e, 0x8c, 0x8e, 0x86, 0x07, 0x19, 0xef, 0x1e, 0x80, 0x26, 0xc8, 0x78, 0xc2, 0x5d,
	0x03, 0x29, 0x46, 0xd0, 0x22, 0xe4, 0x0c, 0x01, 0xd6, 0x85, 0x7a, 0x34, 0x72, 0x1b, 0xe4, 0x89,
	0x27, 0x45, 0x97, 0xc9, 0x2c, 0x1e, 0x0d, 0xa7, 0x21, 0xd2, 0x4d, 0xc2, 0x5b, 0x84, 0xbc, 0x47,
	0x80, 0xdd, 0x01, 0x4b, 0x86, 0xe7, 0x43, 0xc1, 0xa5, 0xdb, 0xdc, 0x31, 0xd0, 0xd6, 0xc4, 0xeb,
	0x29, 0x97, 0x8c, 0x41, 0xe3, 0x22, 0x11, 0xd2, 0xb5, 0xc8, 0x83, 0xce, 0xaa, 0x44, 0x19, 0xa6,
	0xe7, 0x5c, 0x0e, 0xd5, 0x95, 0x9c, 0x6c, 0x72, 0xea, 0x68, 0xf8, 0x35, 0xa2, 0x99, 0xef, 0x5c,
	0xf0, 0xd4, 0x6d, 0x69, 0x5f, 0x75, 0x56, 0x58, 0x14, 0xca, 0xd0, 0x05, 0xc4, 0xda, 0x01, 0x9d,
	0xbd, 0x5f, 0x06, 0x98, 0x1f, 0xe6, 0x3c, 0xbd, 0xca, 0xaa, 0xae, 0x51, 0x30, 0x55, 0x75, 0x9e,
	0xbd, 0x4e, 0xc8, 0x5f, 0xb3, 0x1b, 0xcb, 0xb2, 0x37, 0xb4, 0x2f, 0x65, 0xaf, 0x4e, 0xc1, 0xb8,
	0xe5, 0x14, 0x16, 0x3c, 0x8a, 0xa5, 0x73, 0x96, 0x1f, 0x25, 0xef, 0xce, 0xd6, 0x1d, 0xab, 0x33,
	0xdb, 0x85, 0x75, 0x81, 0x95, 0xe5, 0xae, 0xf4, 0x5c, 0x2d, 0x7a, 0xae, 0x2e, 0xe1, 0x83, 0xeb,
	0x37, 0xbb, 0x0f, 0x5d, 0x3e, 0x8d, 0x8a, 0x3c, 0x20, 0x5e, 0x1b, 0xd1, 0x3f, 0xac, 0x3d, 0xf8,
	0x4f, 0xc7, 0x4b, 0xf9, 0x88, 0xc7, 0x97, 0x3c, 0xd2, 0x54, 0x87, 0xa8, 0x1b, 0x64, 0x0a, 0x32,
	0x0b, 0xf1, 0x1f, 0xa3, 0x2e, 0x31, 0x6a, 0x99, 0xdd, 0x26, 0x76, 0x0f, 0x0d, 0x25, 0x2e, 0xce,
	0x47, 0x24, 0x18, 0xfa, 0x4b, 0xcc, 0xc7, 0x91, 0xdb, 0xd1, 0xf3, 0x51, 0xc8, 0xb1, 0x02, 0xd8,
	0x03, 0xe8, 0x92, 0x39, 0x14, 0x23, 0xf4, 0x8c, 0xa7, 0xe7, 0x6e, 0x17, 0x29, 0x76, 0xd0, 0x51,
	0xe8, 0x51, 0x0e, 0xb2, 0x4d, 0x30, 0xa9, 0x0c, 0xb7, 0x87, 0x59, 0xcc, 0x40, 0x5f, 0x14, 0x3a,
	0x8e, 0x27, 0xb1, 0x74, 0xd7, 0x35, 0x4a, 0x17, 0xef, 0x04, 0xcc, 0x33, 0x35, 0xff, 0xca, 0xd3,
	0xd4, 0xaa, 0x02, 0xdd, 0x06, 0x47, 0x4d, 0x73, 0x28, 0x46, 0x17, 0x7c, 0x12, 0xd2, 0x4a, 0xb4,
	0x03, 0x50, 0xd0, 0x29, 0x21, 0xde, 0x4b, 0x70, 0x28, 0x50, 0xc0, 0xc5, 0x7c, 0x2c, 0xd9, 0x13,
	0xb0, 0x52, 0x3a, 0x09, 0x92, 0x93, 0xb3, 0xcf, 0xf6, 0x8a, 0xdb, 0xa9, 0xa9, 0x39, 0xc5, 0xfb,
	0x0a, 0xec, 0xe3, 0x0c, 0x83, 0xf1, 0x2c, 0xc4, 0xb7, 0x39, 0x17, 0xb4, 0xa5, 0xc9, 0x38, 0x2a,
	0x16, 0x64, 0xe1, 0x9d, 0xca, 0x41, 0xd3, 0x94, 0x7f, 0xd7, 0x26, 0xbd, 0x9e, 0x16, 0xde, 0x17,
	0x55, 0x6a, 0xdc, 0xa8, 0xf4, 0x00, 0x98, 0xcf, 0xc7, 0xbc, 0x92, 0x6c, 0x79, 0xff, 0xde, 0x26,
	0xd4, 0xfd, 0xd1, 0xf5, 0x7e, 0x64, 0x5b, 0xed, 0x3d, 0x03, 0xdb, 0xcf, 0x3b, 0x7e, 0x54, 0xed,
	0xb8, 0x57, 0xea, 0xd8, 0x2f, 0xb4, 0x7b, 0x02, 0x3d, 0xdd, 0xae, 0xff, 0x6f, 0xbd, 0x7a, 0xdb,
	0xd0, 0xf9, 0x94, 0xc6, 0x92, 0x63, 0x09, 0xb3, 0x64, 0x2a, 0xe8, 0xb3, 0x83, 0xbb, 0x61, 0xe8,
	0x02, 0xe3, 0xc8, 0xeb, 0x42, 0x7b, 0x30, 0x99, 0xc9, 0xab, 0x2c, 0xcd, 0xfe, 0xcf, 0x06, 0x38,
	0x24, 0xe5, 0x77, 0x54, 0x15, 0x3b, 0x04, 0xfb, 0x28, 0xd2, 0xe2, 0x66, 0xe5, 0x17, 0x22, 0xac,
	0xdf, 0x2f, 0x61, 0xa5, 0x5c, 0xde, 0x1a, 0x7b, 0x01, 0xad, 0x13, 0xae, 0x17, 0x48, 0x54, 0xdc,
	0xe9, 0x7b, 0xd2, 0x5f, 0x10, 0xd2, 0x5b, 0x7b, 0x5a, 0xcb, 0xd2, 0x6a, 0xe1, 0x2d, 0x10, 0xc6,
	0x8a, 0xb4, 0x6f, 0xc1, 0x29, 0xa8, 0x85, 0x6d, 0x97, 0xc8, 0x37, 0x75, 0xb4, 0x3a, 0x5a, 0x41,
	0x0e, 0x95, 0x68, 0x37, 0x85, 0xb2, 0x22, 0xda, 0x2b, 0x1a, 0x09, 0x39, 0x08, 0xb6, 0x55, 0x6e,
	0xbf, 0xf0, 0x10, 0x7d, 0x77, 0xc1, 0x3a, 0x90, 0x3a, 0x30, 0xc6, 0x73, 0x30, 0x71, 0x3a, 0x28,
	0xb7, 0xaa, 0x82, 0x56, 0xe4, 0x3e, 0x06, 0x3b, 0x97, 0x15, 0xbb, 0xbb, 0x60, 0x28, 0xfe, 0x2d,
	0x7b, 0x38, 0x84, 0x26, 0xf6, 0xe0, 0x2f, 0x6f, 0xe0, 0xff, 0xaa, 0xba, 0xb3, 0xea, 0x3f, 0x37,
	0xe9, 0xcf, 0x7b, 0xf0, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x50, 0x24, 0x70, 0x48, 0x8e, 0x07, 0x00,
	0x00,
}
