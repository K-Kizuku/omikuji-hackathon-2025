// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        (unknown)
// source: event.proto

package protobuf

import (
	account "github.com/K-Kizuku/pymon-command-api/pkg/protobuf/account"
	pymon "github.com/K-Kizuku/pymon-command-api/pkg/protobuf/pymon"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type EventMetadata struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	SeqNr         uint64                 `protobuf:"varint,2,opt,name=seqNr,proto3" json:"seqNr,omitempty"`
	AggregateId   string                 `protobuf:"bytes,3,opt,name=aggregate_id,json=aggregateId,proto3" json:"aggregate_id,omitempty"`
	ExecutorId    string                 `protobuf:"bytes,4,opt,name=executor_id,json=executorId,proto3" json:"executor_id,omitempty"`
	OccurredAt    uint64                 `protobuf:"varint,5,opt,name=occurred_at,json=occurredAt,proto3" json:"occurred_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EventMetadata) Reset() {
	*x = EventMetadata{}
	mi := &file_event_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EventMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventMetadata) ProtoMessage() {}

func (x *EventMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_event_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventMetadata.ProtoReflect.Descriptor instead.
func (*EventMetadata) Descriptor() ([]byte, []int) {
	return file_event_proto_rawDescGZIP(), []int{0}
}

func (x *EventMetadata) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *EventMetadata) GetSeqNr() uint64 {
	if x != nil {
		return x.SeqNr
	}
	return 0
}

func (x *EventMetadata) GetAggregateId() string {
	if x != nil {
		return x.AggregateId
	}
	return ""
}

func (x *EventMetadata) GetExecutorId() string {
	if x != nil {
		return x.ExecutorId
	}
	return ""
}

func (x *EventMetadata) GetOccurredAt() uint64 {
	if x != nil {
		return x.OccurredAt
	}
	return 0
}

type EventEnvelope struct {
	state    protoimpl.MessageState `protogen:"open.v1"`
	Metadata *EventMetadata         `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// Types that are valid to be assigned to Event:
	//
	//	*EventEnvelope_UserCreated
	//	*EventEnvelope_UserNameChanged
	//	*EventEnvelope_UserDeleted
	//	*EventEnvelope_UserLoggedIn
	//	*EventEnvelope_ExpGottenByLoginBonus
	//	*EventEnvelope_ExpGottenByTraining
	//	*EventEnvelope_PymonCreated
	//	*EventEnvelope_PymonNameChanged
	//	*EventEnvelope_PymonExpChanged
	//	*EventEnvelope_PymonStatusChanged
	//	*EventEnvelope_PymonSkillChanged
	//	*EventEnvelope_PymonDeleted
	//	*EventEnvelope_SkillLearned
	//	*EventEnvelope_SkillForgotten
	//	*EventEnvelope_SpAssigned
	//	*EventEnvelope_TrainingDid
	Event         isEventEnvelope_Event `protobuf_oneof:"event"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EventEnvelope) Reset() {
	*x = EventEnvelope{}
	mi := &file_event_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EventEnvelope) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventEnvelope) ProtoMessage() {}

func (x *EventEnvelope) ProtoReflect() protoreflect.Message {
	mi := &file_event_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventEnvelope.ProtoReflect.Descriptor instead.
func (*EventEnvelope) Descriptor() ([]byte, []int) {
	return file_event_proto_rawDescGZIP(), []int{1}
}

func (x *EventEnvelope) GetMetadata() *EventMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *EventEnvelope) GetEvent() isEventEnvelope_Event {
	if x != nil {
		return x.Event
	}
	return nil
}

func (x *EventEnvelope) GetUserCreated() *account.UserCreated {
	if x != nil {
		if x, ok := x.Event.(*EventEnvelope_UserCreated); ok {
			return x.UserCreated
		}
	}
	return nil
}

func (x *EventEnvelope) GetUserNameChanged() *account.UserNameChanged {
	if x != nil {
		if x, ok := x.Event.(*EventEnvelope_UserNameChanged); ok {
			return x.UserNameChanged
		}
	}
	return nil
}

func (x *EventEnvelope) GetUserDeleted() *account.UserDeleted {
	if x != nil {
		if x, ok := x.Event.(*EventEnvelope_UserDeleted); ok {
			return x.UserDeleted
		}
	}
	return nil
}

func (x *EventEnvelope) GetUserLoggedIn() *account.LoggedIn {
	if x != nil {
		if x, ok := x.Event.(*EventEnvelope_UserLoggedIn); ok {
			return x.UserLoggedIn
		}
	}
	return nil
}

func (x *EventEnvelope) GetExpGottenByLoginBonus() *pymon.ExpGottenByLoginBonus {
	if x != nil {
		if x, ok := x.Event.(*EventEnvelope_ExpGottenByLoginBonus); ok {
			return x.ExpGottenByLoginBonus
		}
	}
	return nil
}

func (x *EventEnvelope) GetExpGottenByTraining() *pymon.ExpGottenByTraining {
	if x != nil {
		if x, ok := x.Event.(*EventEnvelope_ExpGottenByTraining); ok {
			return x.ExpGottenByTraining
		}
	}
	return nil
}

func (x *EventEnvelope) GetPymonCreated() *pymon.PymonCreated {
	if x != nil {
		if x, ok := x.Event.(*EventEnvelope_PymonCreated); ok {
			return x.PymonCreated
		}
	}
	return nil
}

func (x *EventEnvelope) GetPymonNameChanged() *pymon.PymonNameChanged {
	if x != nil {
		if x, ok := x.Event.(*EventEnvelope_PymonNameChanged); ok {
			return x.PymonNameChanged
		}
	}
	return nil
}

func (x *EventEnvelope) GetPymonExpChanged() *pymon.PymonExpChanged {
	if x != nil {
		if x, ok := x.Event.(*EventEnvelope_PymonExpChanged); ok {
			return x.PymonExpChanged
		}
	}
	return nil
}

func (x *EventEnvelope) GetPymonStatusChanged() *pymon.PymonStatusChanged {
	if x != nil {
		if x, ok := x.Event.(*EventEnvelope_PymonStatusChanged); ok {
			return x.PymonStatusChanged
		}
	}
	return nil
}

func (x *EventEnvelope) GetPymonSkillChanged() *pymon.PymonSkillChanged {
	if x != nil {
		if x, ok := x.Event.(*EventEnvelope_PymonSkillChanged); ok {
			return x.PymonSkillChanged
		}
	}
	return nil
}

func (x *EventEnvelope) GetPymonDeleted() *pymon.PymonDeleted {
	if x != nil {
		if x, ok := x.Event.(*EventEnvelope_PymonDeleted); ok {
			return x.PymonDeleted
		}
	}
	return nil
}

func (x *EventEnvelope) GetSkillLearned() *pymon.SkillLearned {
	if x != nil {
		if x, ok := x.Event.(*EventEnvelope_SkillLearned); ok {
			return x.SkillLearned
		}
	}
	return nil
}

func (x *EventEnvelope) GetSkillForgotten() *pymon.SkillForgotten {
	if x != nil {
		if x, ok := x.Event.(*EventEnvelope_SkillForgotten); ok {
			return x.SkillForgotten
		}
	}
	return nil
}

func (x *EventEnvelope) GetSpAssigned() *pymon.SPAssigned {
	if x != nil {
		if x, ok := x.Event.(*EventEnvelope_SpAssigned); ok {
			return x.SpAssigned
		}
	}
	return nil
}

func (x *EventEnvelope) GetTrainingDid() *pymon.TrainingDid {
	if x != nil {
		if x, ok := x.Event.(*EventEnvelope_TrainingDid); ok {
			return x.TrainingDid
		}
	}
	return nil
}

type isEventEnvelope_Event interface {
	isEventEnvelope_Event()
}

type EventEnvelope_UserCreated struct {
	UserCreated *account.UserCreated `protobuf:"bytes,2,opt,name=user_created,json=userCreated,proto3,oneof"`
}

type EventEnvelope_UserNameChanged struct {
	UserNameChanged *account.UserNameChanged `protobuf:"bytes,3,opt,name=user_name_changed,json=userNameChanged,proto3,oneof"`
}

type EventEnvelope_UserDeleted struct {
	UserDeleted *account.UserDeleted `protobuf:"bytes,4,opt,name=user_deleted,json=userDeleted,proto3,oneof"`
}

type EventEnvelope_UserLoggedIn struct {
	UserLoggedIn *account.LoggedIn `protobuf:"bytes,5,opt,name=user_logged_in,json=userLoggedIn,proto3,oneof"`
}

type EventEnvelope_ExpGottenByLoginBonus struct {
	ExpGottenByLoginBonus *pymon.ExpGottenByLoginBonus `protobuf:"bytes,6,opt,name=exp_gotten_by_login_bonus,json=expGottenByLoginBonus,proto3,oneof"`
}

type EventEnvelope_ExpGottenByTraining struct {
	ExpGottenByTraining *pymon.ExpGottenByTraining `protobuf:"bytes,7,opt,name=exp_gotten_by_training,json=expGottenByTraining,proto3,oneof"`
}

type EventEnvelope_PymonCreated struct {
	PymonCreated *pymon.PymonCreated `protobuf:"bytes,8,opt,name=pymon_created,json=pymonCreated,proto3,oneof"`
}

type EventEnvelope_PymonNameChanged struct {
	PymonNameChanged *pymon.PymonNameChanged `protobuf:"bytes,9,opt,name=pymon_name_changed,json=pymonNameChanged,proto3,oneof"`
}

type EventEnvelope_PymonExpChanged struct {
	PymonExpChanged *pymon.PymonExpChanged `protobuf:"bytes,10,opt,name=pymon_exp_changed,json=pymonExpChanged,proto3,oneof"`
}

type EventEnvelope_PymonStatusChanged struct {
	PymonStatusChanged *pymon.PymonStatusChanged `protobuf:"bytes,11,opt,name=pymon_status_changed,json=pymonStatusChanged,proto3,oneof"`
}

type EventEnvelope_PymonSkillChanged struct {
	PymonSkillChanged *pymon.PymonSkillChanged `protobuf:"bytes,12,opt,name=pymon_skill_changed,json=pymonSkillChanged,proto3,oneof"`
}

type EventEnvelope_PymonDeleted struct {
	PymonDeleted *pymon.PymonDeleted `protobuf:"bytes,13,opt,name=pymon_deleted,json=pymonDeleted,proto3,oneof"`
}

type EventEnvelope_SkillLearned struct {
	SkillLearned *pymon.SkillLearned `protobuf:"bytes,14,opt,name=skill_learned,json=skillLearned,proto3,oneof"`
}

type EventEnvelope_SkillForgotten struct {
	SkillForgotten *pymon.SkillForgotten `protobuf:"bytes,15,opt,name=skill_forgotten,json=skillForgotten,proto3,oneof"`
}

type EventEnvelope_SpAssigned struct {
	SpAssigned *pymon.SPAssigned `protobuf:"bytes,16,opt,name=sp_assigned,json=spAssigned,proto3,oneof"`
}

type EventEnvelope_TrainingDid struct {
	TrainingDid *pymon.TrainingDid `protobuf:"bytes,17,opt,name=training_did,json=trainingDid,proto3,oneof"`
}

func (*EventEnvelope_UserCreated) isEventEnvelope_Event() {}

func (*EventEnvelope_UserNameChanged) isEventEnvelope_Event() {}

func (*EventEnvelope_UserDeleted) isEventEnvelope_Event() {}

func (*EventEnvelope_UserLoggedIn) isEventEnvelope_Event() {}

func (*EventEnvelope_ExpGottenByLoginBonus) isEventEnvelope_Event() {}

func (*EventEnvelope_ExpGottenByTraining) isEventEnvelope_Event() {}

func (*EventEnvelope_PymonCreated) isEventEnvelope_Event() {}

func (*EventEnvelope_PymonNameChanged) isEventEnvelope_Event() {}

func (*EventEnvelope_PymonExpChanged) isEventEnvelope_Event() {}

func (*EventEnvelope_PymonStatusChanged) isEventEnvelope_Event() {}

func (*EventEnvelope_PymonSkillChanged) isEventEnvelope_Event() {}

func (*EventEnvelope_PymonDeleted) isEventEnvelope_Event() {}

func (*EventEnvelope_SkillLearned) isEventEnvelope_Event() {}

func (*EventEnvelope_SkillForgotten) isEventEnvelope_Event() {}

func (*EventEnvelope_SpAssigned) isEventEnvelope_Event() {}

func (*EventEnvelope_TrainingDid) isEventEnvelope_Event() {}

var File_event_proto protoreflect.FileDescriptor

var file_event_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x6c, 0x6f,
	0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x70, 0x79, 0x6d, 0x6f, 0x6e,
	0x2f, 0x65, 0x78, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x70, 0x79, 0x6d, 0x6f,
	0x6e, 0x2f, 0x70, 0x79, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x70,
	0x79, 0x6d, 0x6f, 0x6e, 0x2f, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x12, 0x70, 0x79, 0x6d, 0x6f, 0x6e, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x70, 0x79, 0x6d, 0x6f, 0x6e, 0x2f, 0x74, 0x72, 0x61, 0x69,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9a, 0x01, 0x0a, 0x0d, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x65, 0x71,
	0x4e, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x73, 0x65, 0x71, 0x4e, 0x72, 0x12,
	0x21, 0x0a, 0x0c, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65,
	0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f,
	0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x63, 0x63, 0x75, 0x72, 0x72, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x6f, 0x63, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x64, 0x41, 0x74, 0x22, 0xe2, 0x09, 0x0a, 0x0d, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x45, 0x6e,
	0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x12, 0x30, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x41, 0x0a, 0x0c, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x48, 0x00, 0x52, 0x0b,
	0x75, 0x73, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x4e, 0x0a, 0x11, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d,
	0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x48, 0x00, 0x52, 0x0f, 0x75, 0x73, 0x65, 0x72,
	0x4e, 0x61, 0x6d, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x12, 0x41, 0x0a, 0x0c, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x48,
	0x00, 0x52, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x12, 0x3f,
	0x0a, 0x0e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x64, 0x5f, 0x69, 0x6e,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x64, 0x49, 0x6e, 0x48,
	0x00, 0x52, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x64, 0x49, 0x6e, 0x12,
	0x5c, 0x0a, 0x19, 0x65, 0x78, 0x70, 0x5f, 0x67, 0x6f, 0x74, 0x74, 0x65, 0x6e, 0x5f, 0x62, 0x79,
	0x5f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x62, 0x6f, 0x6e, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x20, 0x2e, 0x70, 0x79, 0x6d, 0x6f, 0x6e, 0x2e, 0x65, 0x78, 0x70, 0x2e, 0x45,
	0x78, 0x70, 0x47, 0x6f, 0x74, 0x74, 0x65, 0x6e, 0x42, 0x79, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x42,
	0x6f, 0x6e, 0x75, 0x73, 0x48, 0x00, 0x52, 0x15, 0x65, 0x78, 0x70, 0x47, 0x6f, 0x74, 0x74, 0x65,
	0x6e, 0x42, 0x79, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x42, 0x6f, 0x6e, 0x75, 0x73, 0x12, 0x55, 0x0a,
	0x16, 0x65, 0x78, 0x70, 0x5f, 0x67, 0x6f, 0x74, 0x74, 0x65, 0x6e, 0x5f, 0x62, 0x79, 0x5f, 0x74,
	0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e,
	0x70, 0x79, 0x6d, 0x6f, 0x6e, 0x2e, 0x65, 0x78, 0x70, 0x2e, 0x45, 0x78, 0x70, 0x47, 0x6f, 0x74,
	0x74, 0x65, 0x6e, 0x42, 0x79, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x48, 0x00, 0x52,
	0x13, 0x65, 0x78, 0x70, 0x47, 0x6f, 0x74, 0x74, 0x65, 0x6e, 0x42, 0x79, 0x54, 0x72, 0x61, 0x69,
	0x6e, 0x69, 0x6e, 0x67, 0x12, 0x40, 0x0a, 0x0d, 0x70, 0x79, 0x6d, 0x6f, 0x6e, 0x5f, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x79,
	0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x79, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x79, 0x6d, 0x6f, 0x6e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x48, 0x00, 0x52, 0x0c, 0x70, 0x79, 0x6d, 0x6f, 0x6e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x4d, 0x0a, 0x12, 0x70, 0x79, 0x6d, 0x6f, 0x6e, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70, 0x79, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x79, 0x6d, 0x6f, 0x6e,
	0x2e, 0x50, 0x79, 0x6d, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x64, 0x48, 0x00, 0x52, 0x10, 0x70, 0x79, 0x6d, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x64, 0x12, 0x4a, 0x0a, 0x11, 0x70, 0x79, 0x6d, 0x6f, 0x6e, 0x5f, 0x65,
	0x78, 0x70, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x70, 0x79, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x79, 0x6d, 0x6f, 0x6e, 0x2e, 0x50,
	0x79, 0x6d, 0x6f, 0x6e, 0x45, 0x78, 0x70, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x48, 0x00,
	0x52, 0x0f, 0x70, 0x79, 0x6d, 0x6f, 0x6e, 0x45, 0x78, 0x70, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x64, 0x12, 0x53, 0x0a, 0x14, 0x70, 0x79, 0x6d, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1f, 0x2e, 0x70, 0x79, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x79, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x79,
	0x6d, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64,
	0x48, 0x00, 0x52, 0x12, 0x70, 0x79, 0x6d, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x12, 0x50, 0x0a, 0x13, 0x70, 0x79, 0x6d, 0x6f, 0x6e, 0x5f,
	0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x70, 0x79, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x79, 0x6d, 0x6f,
	0x6e, 0x2e, 0x50, 0x79, 0x6d, 0x6f, 0x6e, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x64, 0x48, 0x00, 0x52, 0x11, 0x70, 0x79, 0x6d, 0x6f, 0x6e, 0x53, 0x6b, 0x69, 0x6c,
	0x6c, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x12, 0x40, 0x0a, 0x0d, 0x70, 0x79, 0x6d, 0x6f,
	0x6e, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x70, 0x79, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x79, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x79,
	0x6d, 0x6f, 0x6e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x48, 0x00, 0x52, 0x0c, 0x70, 0x79,
	0x6d, 0x6f, 0x6e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x12, 0x40, 0x0a, 0x0d, 0x73, 0x6b,
	0x69, 0x6c, 0x6c, 0x5f, 0x6c, 0x65, 0x61, 0x72, 0x6e, 0x65, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x70, 0x79, 0x6d, 0x6f, 0x6e, 0x2e, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x2e,
	0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x4c, 0x65, 0x61, 0x72, 0x6e, 0x65, 0x64, 0x48, 0x00, 0x52, 0x0c,
	0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x4c, 0x65, 0x61, 0x72, 0x6e, 0x65, 0x64, 0x12, 0x46, 0x0a, 0x0f,
	0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x66, 0x6f, 0x72, 0x67, 0x6f, 0x74, 0x74, 0x65, 0x6e, 0x18,
	0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x79, 0x6d, 0x6f, 0x6e, 0x2e, 0x73, 0x6b,
	0x69, 0x6c, 0x6c, 0x2e, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x46, 0x6f, 0x72, 0x67, 0x6f, 0x74, 0x74,
	0x65, 0x6e, 0x48, 0x00, 0x52, 0x0e, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x46, 0x6f, 0x72, 0x67, 0x6f,
	0x74, 0x74, 0x65, 0x6e, 0x12, 0x3b, 0x0a, 0x0b, 0x73, 0x70, 0x5f, 0x61, 0x73, 0x73, 0x69, 0x67,
	0x6e, 0x65, 0x64, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x79, 0x6d, 0x6f,
	0x6e, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x53, 0x50, 0x41, 0x73, 0x73, 0x69, 0x67,
	0x6e, 0x65, 0x64, 0x48, 0x00, 0x52, 0x0a, 0x73, 0x70, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65,
	0x64, 0x12, 0x3d, 0x0a, 0x0c, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x64, 0x69,
	0x64, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x79, 0x6d, 0x6f, 0x6e, 0x2e,
	0x74, 0x72, 0x61, 0x69, 0x6e, 0x2e, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x44, 0x69,
	0x64, 0x48, 0x00, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x44, 0x69, 0x64,
	0x42, 0x07, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x42, 0x7f, 0x0a, 0x09, 0x63, 0x6f, 0x6d,
	0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x42, 0x0a, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x4b, 0x2d, 0x4b, 0x69, 0x7a, 0x75, 0x6b, 0x75, 0x2f, 0x70, 0x79, 0x6d, 0x6f, 0x6e, 0x2d,
	0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0xa2, 0x02, 0x03, 0x45, 0x58, 0x58, 0xaa, 0x02,
	0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0xca, 0x02, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0xe2, 0x02,
	0x11, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_event_proto_rawDescOnce sync.Once
	file_event_proto_rawDescData = file_event_proto_rawDesc
)

func file_event_proto_rawDescGZIP() []byte {
	file_event_proto_rawDescOnce.Do(func() {
		file_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_event_proto_rawDescData)
	})
	return file_event_proto_rawDescData
}

var file_event_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_event_proto_goTypes = []any{
	(*EventMetadata)(nil),               // 0: event.EventMetadata
	(*EventEnvelope)(nil),               // 1: event.EventEnvelope
	(*account.UserCreated)(nil),         // 2: account.account.UserCreated
	(*account.UserNameChanged)(nil),     // 3: account.account.UserNameChanged
	(*account.UserDeleted)(nil),         // 4: account.account.UserDeleted
	(*account.LoggedIn)(nil),            // 5: account.login.LoggedIn
	(*pymon.ExpGottenByLoginBonus)(nil), // 6: pymon.exp.ExpGottenByLoginBonus
	(*pymon.ExpGottenByTraining)(nil),   // 7: pymon.exp.ExpGottenByTraining
	(*pymon.PymonCreated)(nil),          // 8: pymon.pymon.PymonCreated
	(*pymon.PymonNameChanged)(nil),      // 9: pymon.pymon.PymonNameChanged
	(*pymon.PymonExpChanged)(nil),       // 10: pymon.pymon.PymonExpChanged
	(*pymon.PymonStatusChanged)(nil),    // 11: pymon.pymon.PymonStatusChanged
	(*pymon.PymonSkillChanged)(nil),     // 12: pymon.pymon.PymonSkillChanged
	(*pymon.PymonDeleted)(nil),          // 13: pymon.pymon.PymonDeleted
	(*pymon.SkillLearned)(nil),          // 14: pymon.skill.SkillLearned
	(*pymon.SkillForgotten)(nil),        // 15: pymon.skill.SkillForgotten
	(*pymon.SPAssigned)(nil),            // 16: pymon.status.SPAssigned
	(*pymon.TrainingDid)(nil),           // 17: pymon.train.TrainingDid
}
var file_event_proto_depIdxs = []int32{
	0,  // 0: event.EventEnvelope.metadata:type_name -> event.EventMetadata
	2,  // 1: event.EventEnvelope.user_created:type_name -> account.account.UserCreated
	3,  // 2: event.EventEnvelope.user_name_changed:type_name -> account.account.UserNameChanged
	4,  // 3: event.EventEnvelope.user_deleted:type_name -> account.account.UserDeleted
	5,  // 4: event.EventEnvelope.user_logged_in:type_name -> account.login.LoggedIn
	6,  // 5: event.EventEnvelope.exp_gotten_by_login_bonus:type_name -> pymon.exp.ExpGottenByLoginBonus
	7,  // 6: event.EventEnvelope.exp_gotten_by_training:type_name -> pymon.exp.ExpGottenByTraining
	8,  // 7: event.EventEnvelope.pymon_created:type_name -> pymon.pymon.PymonCreated
	9,  // 8: event.EventEnvelope.pymon_name_changed:type_name -> pymon.pymon.PymonNameChanged
	10, // 9: event.EventEnvelope.pymon_exp_changed:type_name -> pymon.pymon.PymonExpChanged
	11, // 10: event.EventEnvelope.pymon_status_changed:type_name -> pymon.pymon.PymonStatusChanged
	12, // 11: event.EventEnvelope.pymon_skill_changed:type_name -> pymon.pymon.PymonSkillChanged
	13, // 12: event.EventEnvelope.pymon_deleted:type_name -> pymon.pymon.PymonDeleted
	14, // 13: event.EventEnvelope.skill_learned:type_name -> pymon.skill.SkillLearned
	15, // 14: event.EventEnvelope.skill_forgotten:type_name -> pymon.skill.SkillForgotten
	16, // 15: event.EventEnvelope.sp_assigned:type_name -> pymon.status.SPAssigned
	17, // 16: event.EventEnvelope.training_did:type_name -> pymon.train.TrainingDid
	17, // [17:17] is the sub-list for method output_type
	17, // [17:17] is the sub-list for method input_type
	17, // [17:17] is the sub-list for extension type_name
	17, // [17:17] is the sub-list for extension extendee
	0,  // [0:17] is the sub-list for field type_name
}

func init() { file_event_proto_init() }
func file_event_proto_init() {
	if File_event_proto != nil {
		return
	}
	file_event_proto_msgTypes[1].OneofWrappers = []any{
		(*EventEnvelope_UserCreated)(nil),
		(*EventEnvelope_UserNameChanged)(nil),
		(*EventEnvelope_UserDeleted)(nil),
		(*EventEnvelope_UserLoggedIn)(nil),
		(*EventEnvelope_ExpGottenByLoginBonus)(nil),
		(*EventEnvelope_ExpGottenByTraining)(nil),
		(*EventEnvelope_PymonCreated)(nil),
		(*EventEnvelope_PymonNameChanged)(nil),
		(*EventEnvelope_PymonExpChanged)(nil),
		(*EventEnvelope_PymonStatusChanged)(nil),
		(*EventEnvelope_PymonSkillChanged)(nil),
		(*EventEnvelope_PymonDeleted)(nil),
		(*EventEnvelope_SkillLearned)(nil),
		(*EventEnvelope_SkillForgotten)(nil),
		(*EventEnvelope_SpAssigned)(nil),
		(*EventEnvelope_TrainingDid)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_event_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_event_proto_goTypes,
		DependencyIndexes: file_event_proto_depIdxs,
		MessageInfos:      file_event_proto_msgTypes,
	}.Build()
	File_event_proto = out.File
	file_event_proto_rawDesc = nil
	file_event_proto_goTypes = nil
	file_event_proto_depIdxs = nil
}
