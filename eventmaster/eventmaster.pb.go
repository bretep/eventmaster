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
	EventId           string   `protobuf:"bytes,17,opt,name=event_id,json=eventId" json:"event_id,omitempty"`
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

func (m *Query) GetEventId() string {
	if m != nil {
		return m.EventId
	}
	return ""
}

type Topic struct {
	Id         string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	TopicName  string `protobuf:"bytes,2,opt,name=topic_name,json=topicName" json:"topic_name,omitempty"`
	DataSchema []byte `protobuf:"bytes,3,opt,name=data_schema,json=dataSchema,proto3" json:"data_schema,omitempty"`
}

func (m *Topic) Reset()                    { *m = Topic{} }
func (m *Topic) String() string            { return proto.CompactTextString(m) }
func (*Topic) ProtoMessage()               {}
func (*Topic) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Topic) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

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
	Id     string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	DcName string `protobuf:"bytes,2,opt,name=dc_name,json=dcName" json:"dc_name,omitempty"`
}

func (m *Dc) Reset()                    { *m = Dc{} }
func (m *Dc) String() string            { return proto.CompactTextString(m) }
func (*Dc) ProtoMessage()               {}
func (*Dc) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Dc) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Dc) GetDcName() string {
	if m != nil {
		return m.DcName
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
	// 732 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x55, 0xdb, 0x6e, 0xd3, 0x40,
	0x10, 0x6d, 0x2e, 0x4e, 0x9c, 0x71, 0x2e, 0xed, 0x02, 0xc2, 0x8d, 0x40, 0xad, 0x2c, 0x40, 0x05,
	0x41, 0x85, 0x5a, 0x01, 0x0f, 0xf4, 0xa5, 0xc8, 0x69, 0x41, 0x02, 0x24, 0xdc, 0xa2, 0x3e, 0x46,
	0x26, 0x5e, 0x5a, 0x8b, 0x24, 0x0e, 0xde, 0x4d, 0x51, 0xff, 0x03, 0xfe, 0x8f, 0x4f, 0x61, 0x76,
	0xd6, 0x2e, 0x5e, 0x37, 0x34, 0x95, 0x78, 0xdb, 0x3d, 0x7b, 0xe6, 0x7e, 0xc6, 0x86, 0x35, 0x7e,
	0xce, 0xa7, 0x72, 0x12, 0x0a, 0xc9, 0xd3, 0xed, 0x59, 0x9a, 0xc8, 0x84, 0x39, 0x05, 0xc8, 0xfb,
	0x55, 0x05, 0x6b, 0xa0, 0xee, 0x6c, 0x1d, 0x6c, 0x7a, 0x18, 0xc6, 0x91, 0x5b, 0xd9, 0xac, 0x6c,
	0xb5, 0x82, 0x26, 0xdd, 0xdf, 0x45, 0xec, 0x11, 0xf4, 0x66, 0x61, 0xaa, 0xde, 0x2e, 0x19, 0x55,
	0x62, 0x74, 0x34, 0x3c, 0xc8, 0x78, 0xf7, 0x01, 0x34, 0x41, 0xc6, 0x13, 0xee, 0xd6, 0x90, 0x52,
	0x0b, 0x5a, 0x84, 0x1c, 0x23, 0xc0, 0xba, 0x50, 0x8d, 0x46, 0x6e, 0x9d, 0x2c, 0xf1, 0xa4, 0xe8,
	0x32, 0x99, 0xc5, 0xa3, 0xe1, 0x34, 0x44, 0xba, 0x45, 0x78, 0x8b, 0x90, 0x8f, 0x08, 0xb0, 0xbb,
	0xd0, 0x94, 0xe1, 0xe9, 0x50, 0x70, 0xe9, 0x36, 0x36, 0x6b, 0xf8, 0xd6, 0xc0, 0xeb, 0x11, 0x97,
	0x8c, 0x41, 0xfd, 0x2c, 0x11, 0xd2, 0x6d, 0x92, 0x05, 0x9d, 0x55, 0x8a, 0x32, 0x4c, 0x4f, 0xb9,
	0x1c, 0xaa, 0x2b, 0x19, 0xd9, 0x64, 0xd4, 0xd1, 0xf0, 0x5b, 0x44, 0x33, 0xdb, 0xb9, 0xe0, 0xa9,
	0xdb, 0xd2, 0xb6, 0xea, 0xac, 0xb0, 0x28, 0x94, 0xa1, 0x0b, 0x88, 0xb5, 0x03, 0x3a, 0x7b, 0x3f,
	0xeb, 0x60, 0x7d, 0x9a, 0xf3, 0xf4, 0x22, 0xcb, 0xba, 0x42, 0xce, 0x54, 0xd6, 0x79, 0xf4, 0x2a,
	0x21, 0xff, 0x8c, 0x5e, 0xbb, 0x2e, 0x7a, 0x5d, 0xdb, 0x52, 0xf4, 0x72, 0x17, 0x6a, 0x37, 0xec,
	0xc2, 0x82, 0xa1, 0x34, 0x75, 0x4c, 0x73, 0x28, 0x79, 0x75, 0xb6, 0xae, 0x58, 0x9d, 0xd9, 0x16,
	0xac, 0x0a, 0xcc, 0x2c, 0x37, 0xa5, 0x71, 0xb5, 0x68, 0x5c, 0x5d, 0xc2, 0x07, 0x97, 0x33, 0x7b,
	0x00, 0x5d, 0x3e, 0x8d, 0x8a, 0x3c, 0x20, 0x5e, 0x1b, 0xd1, 0xbf, 0xac, 0x6d, 0xb8, 0xa5, 0xfd,
	0xa5, 0x7c, 0xc4, 0xe3, 0x73, 0x1e, 0x69, 0xaa, 0x43, 0xd4, 0x35, 0x7a, 0x0a, 0xb2, 0x17, 0xe2,
	0x3f, 0x41, 0x5d, 0xa2, 0x57, 0x93, 0xdd, 0x26, 0x76, 0x0f, 0x1f, 0x0c, 0x2e, 0xf6, 0x47, 0x24,
	0xe8, 0xfa, 0x6b, 0xcc, 0xc7, 0x91, 0xdb, 0xd1, 0xfd, 0x51, 0xc8, 0x81, 0x02, 0xd8, 0x43, 0xe8,
	0xd2, 0x73, 0x28, 0x46, 0x68, 0x19, 0x4f, 0x4f, 0xdd, 0x2e, 0x52, 0xec, 0xa0, 0xa3, 0xd0, 0xfd,
	0x1c, 0x64, 0xb7, 0xc1, 0xa2, 0x34, 0xdc, 0x1e, 0x46, 0xb1, 0x02, 0x7d, 0x51, 0xe8, 0x38, 0x9e,
	0xc4, 0xd2, 0x5d, 0xd5, 0x28, 0x5d, 0x8c, 0x4d, 0x58, 0x33, 0x36, 0xc1, 0x3b, 0x01, 0xeb, 0x58,
	0x8d, 0x46, 0xa9, 0xe2, 0x72, 0x4f, 0xf0, 0x54, 0x9a, 0x62, 0xb5, 0xac, 0xe5, 0x0d, 0x70, 0x54,
	0xe3, 0x87, 0x62, 0x74, 0xc6, 0x27, 0x21, 0xad, 0x46, 0x3b, 0x00, 0x05, 0x1d, 0x11, 0xe2, 0xbd,
	0x06, 0x87, 0x1c, 0x07, 0x5c, 0xcc, 0xc7, 0x92, 0x3d, 0x85, 0x66, 0x4a, 0x27, 0x41, 0xca, 0x73,
	0x76, 0xd8, 0x76, 0x71, 0x91, 0x35, 0x35, 0xa7, 0x78, 0xdf, 0x80, 0x7d, 0x9e, 0xa1, 0x33, 0x9e,
	0xb9, 0xf8, 0x3e, 0xe7, 0x82, 0xca, 0x48, 0xc6, 0x91, 0x4e, 0x28, 0x5b, 0x68, 0xbc, 0x53, 0x3a,
	0xf8, 0x34, 0xe5, 0x3f, 0x8a, 0xb9, 0x36, 0xf1, 0x7e, 0xb3, 0x4c, 0x77, 0x81, 0xf9, 0x7c, 0xcc,
	0x4b, 0xc1, 0xcc, 0xfa, 0x2b, 0xa5, 0xfa, 0xbd, 0x67, 0x50, 0xf5, 0xaf, 0x36, 0x0d, 0xb5, 0x1d,
	0x19, 0x1d, 0x6b, 0x44, 0x9a, 0xfe, 0x02, 0x6c, 0x3f, 0x6f, 0xc5, 0xe3, 0x72, 0x2b, 0x7a, 0x46,
	0x2b, 0xfc, 0x42, 0x1f, 0x0e, 0xa1, 0xa7, 0xfb, 0xe0, 0xff, 0x5f, 0x13, 0xbc, 0x0d, 0xe8, 0x9c,
	0xa4, 0xb1, 0xe4, 0x98, 0xc2, 0x2c, 0x99, 0x0a, 0x9e, 0x65, 0x5e, 0xcb, 0x33, 0xf7, 0xba, 0xd0,
	0x1e, 0x4c, 0x66, 0xf2, 0x22, 0x0b, 0xb3, 0xf3, 0xbb, 0x0e, 0x0e, 0xad, 0xc3, 0x07, 0xca, 0x8a,
	0xed, 0x81, 0xbd, 0x1f, 0xe9, 0x05, 0x61, 0xe6, 0xe8, 0x08, 0xeb, 0xf7, 0x0d, 0xcc, 0x88, 0xe5,
	0xad, 0xb0, 0x57, 0xd0, 0x3a, 0xe4, 0x7a, 0x09, 0x45, 0xc9, 0x9c, 0xbe, 0x49, 0xfd, 0x05, 0x2e,
	0xbd, 0x95, 0xe7, 0x95, 0x2c, 0xac, 0x56, 0xe8, 0x02, 0xc5, 0x2c, 0x09, 0xfb, 0x1e, 0x9c, 0x82,
	0x8c, 0xd8, 0x86, 0x41, 0xbe, 0x2a, 0xb0, 0xe5, 0xde, 0x0a, 0x3a, 0x29, 0x79, 0xbb, 0xaa, 0xa0,
	0x25, 0xde, 0xde, 0x50, 0x4b, 0xc8, 0x40, 0xb0, 0x75, 0xb3, 0xfc, 0xc2, 0x20, 0xfa, 0xee, 0x82,
	0x3d, 0x21, 0x75, 0xa0, 0x8f, 0x97, 0x60, 0x61, 0x77, 0x50, 0x87, 0x65, 0x05, 0x2d, 0x89, 0x7d,
	0x00, 0x76, 0x2e, 0x2b, 0x76, 0x6f, 0x41, 0x53, 0xfc, 0x1b, 0xd6, 0xb0, 0x07, 0x0d, 0xac, 0xc1,
	0xbf, 0xbe, 0x80, 0x3b, 0x65, 0x75, 0x67, 0xd9, 0x7f, 0x69, 0xd0, 0xdf, 0x7b, 0xf7, 0x4f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x95, 0xe1, 0x94, 0x47, 0xd2, 0x07, 0x00, 0x00,
}
