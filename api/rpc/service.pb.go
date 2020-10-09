// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: api/rpc/service.proto

package rpc

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

type Play struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GameId   string `protobuf:"bytes,1,opt,name=gameId,proto3" json:"gameId,omitempty"`
	PitIndex int64  `protobuf:"varint,2,opt,name=pitIndex,proto3" json:"pitIndex,omitempty"`
}

func (x *Play) Reset() {
	*x = Play{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_rpc_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Play) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Play) ProtoMessage() {}

func (x *Play) ProtoReflect() protoreflect.Message {
	mi := &file_api_rpc_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Play.ProtoReflect.Descriptor instead.
func (*Play) Descriptor() ([]byte, []int) {
	return file_api_rpc_service_proto_rawDescGZIP(), []int{0}
}

func (x *Play) GetGameId() string {
	if x != nil {
		return x.GameId
	}
	return ""
}

func (x *Play) GetPitIndex() int64 {
	if x != nil {
		return x.PitIndex
	}
	return 0
}

type GameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GameRequest) Reset() {
	*x = GameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_rpc_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameRequest) ProtoMessage() {}

func (x *GameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_rpc_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameRequest.ProtoReflect.Descriptor instead.
func (*GameRequest) Descriptor() ([]byte, []int) {
	return file_api_rpc_service_proto_rawDescGZIP(), []int{1}
}

func (x *GameRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type NewGame struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Player1 string `protobuf:"bytes,1,opt,name=player1,proto3" json:"player1,omitempty"`
	Player2 string `protobuf:"bytes,2,opt,name=player2,proto3" json:"player2,omitempty"`
}

func (x *NewGame) Reset() {
	*x = NewGame{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_rpc_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewGame) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewGame) ProtoMessage() {}

func (x *NewGame) ProtoReflect() protoreflect.Message {
	mi := &file_api_rpc_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewGame.ProtoReflect.Descriptor instead.
func (*NewGame) Descriptor() ([]byte, []int) {
	return file_api_rpc_service_proto_rawDescGZIP(), []int{2}
}

func (x *NewGame) GetPlayer1() string {
	if x != nil {
		return x.Player1
	}
	return ""
}

func (x *NewGame) GetPlayer2() string {
	if x != nil {
		return x.Player2
	}
	return ""
}

type OngoingGame struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Result int64       `protobuf:"varint,2,opt,name=result,proto3" json:"result,omitempty"`
	Turn   int64       `protobuf:"varint,3,opt,name=turn,proto3" json:"turn,omitempty"`
	Board1 *BoardState `protobuf:"bytes,4,opt,name=board1,proto3" json:"board1,omitempty"`
	Board2 *BoardState `protobuf:"bytes,5,opt,name=board2,proto3" json:"board2,omitempty"`
}

func (x *OngoingGame) Reset() {
	*x = OngoingGame{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_rpc_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OngoingGame) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OngoingGame) ProtoMessage() {}

func (x *OngoingGame) ProtoReflect() protoreflect.Message {
	mi := &file_api_rpc_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OngoingGame.ProtoReflect.Descriptor instead.
func (*OngoingGame) Descriptor() ([]byte, []int) {
	return file_api_rpc_service_proto_rawDescGZIP(), []int{3}
}

func (x *OngoingGame) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *OngoingGame) GetResult() int64 {
	if x != nil {
		return x.Result
	}
	return 0
}

func (x *OngoingGame) GetTurn() int64 {
	if x != nil {
		return x.Turn
	}
	return 0
}

func (x *OngoingGame) GetBoard1() *BoardState {
	if x != nil {
		return x.Board1
	}
	return nil
}

func (x *OngoingGame) GetBoard2() *BoardState {
	if x != nil {
		return x.Board2
	}
	return nil
}

type BoardState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pits   []int64 `protobuf:"varint,1,rep,packed,name=pits,proto3" json:"pits,omitempty"`
	Store  int64   `protobuf:"varint,2,opt,name=store,proto3" json:"store,omitempty"`
	Player *Player `protobuf:"bytes,3,opt,name=player,proto3" json:"player,omitempty"`
}

func (x *BoardState) Reset() {
	*x = BoardState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_rpc_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BoardState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BoardState) ProtoMessage() {}

func (x *BoardState) ProtoReflect() protoreflect.Message {
	mi := &file_api_rpc_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BoardState.ProtoReflect.Descriptor instead.
func (*BoardState) Descriptor() ([]byte, []int) {
	return file_api_rpc_service_proto_rawDescGZIP(), []int{4}
}

func (x *BoardState) GetPits() []int64 {
	if x != nil {
		return x.Pits
	}
	return nil
}

func (x *BoardState) GetStore() int64 {
	if x != nil {
		return x.Store
	}
	return 0
}

func (x *BoardState) GetPlayer() *Player {
	if x != nil {
		return x.Player
	}
	return nil
}

type Player struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Score int64  `protobuf:"varint,2,opt,name=score,proto3" json:"score,omitempty"`
}

func (x *Player) Reset() {
	*x = Player{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_rpc_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Player) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Player) ProtoMessage() {}

func (x *Player) ProtoReflect() protoreflect.Message {
	mi := &file_api_rpc_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Player.ProtoReflect.Descriptor instead.
func (*Player) Descriptor() ([]byte, []int) {
	return file_api_rpc_service_proto_rawDescGZIP(), []int{5}
}

func (x *Player) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Player) GetScore() int64 {
	if x != nil {
		return x.Score
	}
	return 0
}

var File_api_rpc_service_proto protoreflect.FileDescriptor

var file_api_rpc_service_proto_rawDesc = []byte{
	0x0a, 0x15, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x72, 0x70, 0x63, 0x22, 0x3a, 0x0a, 0x04,
	0x50, 0x6c, 0x61, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x61, 0x6d, 0x65, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x61, 0x6d, 0x65, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x69, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x70, 0x69, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x22, 0x1d, 0x0a, 0x0b, 0x47, 0x61, 0x6d, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3d, 0x0a, 0x07, 0x4e, 0x65, 0x77, 0x47, 0x61,
	0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x31, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x31, 0x12, 0x18, 0x0a, 0x07,
	0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x32, 0x22, 0x9b, 0x01, 0x0a, 0x0b, 0x4f, 0x6e, 0x67, 0x6f, 0x69,
	0x6e, 0x67, 0x47, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x75, 0x72, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x74, 0x75,
	0x72, 0x6e, 0x12, 0x27, 0x0a, 0x06, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x31, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x52, 0x06, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x31, 0x12, 0x27, 0x0a, 0x06, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x32, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x72, 0x70,
	0x63, 0x2e, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x06, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x32, 0x22, 0x5b, 0x0a, 0x0a, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x69, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03,
	0x52, 0x04, 0x70, 0x69, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x23, 0x0a, 0x06,
	0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x72,
	0x70, 0x63, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x06, 0x70, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x22, 0x32, 0x0a, 0x06, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x73, 0x63, 0x6f, 0x72, 0x65, 0x32, 0x92, 0x01, 0x0a, 0x07, 0x4d, 0x61, 0x6e, 0x63, 0x61, 0x6c,
	0x61, 0x12, 0x2c, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x61, 0x6d, 0x65, 0x12,
	0x0c, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x4e, 0x65, 0x77, 0x47, 0x61, 0x6d, 0x65, 0x1a, 0x10, 0x2e,
	0x72, 0x70, 0x63, 0x2e, 0x4f, 0x6e, 0x67, 0x6f, 0x69, 0x6e, 0x67, 0x47, 0x61, 0x6d, 0x65, 0x12,
	0x2a, 0x0a, 0x0b, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x12, 0x09,
	0x2e, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x1a, 0x10, 0x2e, 0x72, 0x70, 0x63, 0x2e,
	0x4f, 0x6e, 0x67, 0x6f, 0x69, 0x6e, 0x67, 0x47, 0x61, 0x6d, 0x65, 0x12, 0x2d, 0x0a, 0x07, 0x47,
	0x65, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x61, 0x6d,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x4f,
	0x6e, 0x67, 0x6f, 0x69, 0x6e, 0x67, 0x47, 0x61, 0x6d, 0x65, 0x42, 0x09, 0x5a, 0x07, 0x61, 0x70,
	0x69, 0x2f, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_rpc_service_proto_rawDescOnce sync.Once
	file_api_rpc_service_proto_rawDescData = file_api_rpc_service_proto_rawDesc
)

func file_api_rpc_service_proto_rawDescGZIP() []byte {
	file_api_rpc_service_proto_rawDescOnce.Do(func() {
		file_api_rpc_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_rpc_service_proto_rawDescData)
	})
	return file_api_rpc_service_proto_rawDescData
}

var file_api_rpc_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_api_rpc_service_proto_goTypes = []interface{}{
	(*Play)(nil),        // 0: rpc.Play
	(*GameRequest)(nil), // 1: rpc.GameRequest
	(*NewGame)(nil),     // 2: rpc.NewGame
	(*OngoingGame)(nil), // 3: rpc.OngoingGame
	(*BoardState)(nil),  // 4: rpc.BoardState
	(*Player)(nil),      // 5: rpc.Player
}
var file_api_rpc_service_proto_depIdxs = []int32{
	4, // 0: rpc.OngoingGame.board1:type_name -> rpc.BoardState
	4, // 1: rpc.OngoingGame.board2:type_name -> rpc.BoardState
	5, // 2: rpc.BoardState.player:type_name -> rpc.Player
	2, // 3: rpc.Mancala.CreateGame:input_type -> rpc.NewGame
	0, // 4: rpc.Mancala.ExecutePlay:input_type -> rpc.Play
	1, // 5: rpc.Mancala.GetGame:input_type -> rpc.GameRequest
	3, // 6: rpc.Mancala.CreateGame:output_type -> rpc.OngoingGame
	3, // 7: rpc.Mancala.ExecutePlay:output_type -> rpc.OngoingGame
	3, // 8: rpc.Mancala.GetGame:output_type -> rpc.OngoingGame
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_api_rpc_service_proto_init() }
func file_api_rpc_service_proto_init() {
	if File_api_rpc_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_rpc_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Play); i {
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
		file_api_rpc_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GameRequest); i {
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
		file_api_rpc_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewGame); i {
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
		file_api_rpc_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OngoingGame); i {
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
		file_api_rpc_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BoardState); i {
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
		file_api_rpc_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Player); i {
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
			RawDescriptor: file_api_rpc_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_rpc_service_proto_goTypes,
		DependencyIndexes: file_api_rpc_service_proto_depIdxs,
		MessageInfos:      file_api_rpc_service_proto_msgTypes,
	}.Build()
	File_api_rpc_service_proto = out.File
	file_api_rpc_service_proto_rawDesc = nil
	file_api_rpc_service_proto_goTypes = nil
	file_api_rpc_service_proto_depIdxs = nil
}
