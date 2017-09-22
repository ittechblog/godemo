// Code generated by protoc-gen-go.
// source: MultiRowMutation.proto
// DO NOT EDIT!

package pb

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type MultiRowMutationProcessorRequest struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *MultiRowMutationProcessorRequest) Reset()         { *m = MultiRowMutationProcessorRequest{} }
func (m *MultiRowMutationProcessorRequest) String() string { return proto.CompactTextString(m) }
func (*MultiRowMutationProcessorRequest) ProtoMessage()    {}

type MultiRowMutationProcessorResponse struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *MultiRowMutationProcessorResponse) Reset()         { *m = MultiRowMutationProcessorResponse{} }
func (m *MultiRowMutationProcessorResponse) String() string { return proto.CompactTextString(m) }
func (*MultiRowMutationProcessorResponse) ProtoMessage()    {}

type MutateRowsRequest struct {
	MutationRequest  []*MutationProto `protobuf:"bytes,1,rep,name=mutation_request" json:"mutation_request,omitempty"`
	NonceGroup       *uint64          `protobuf:"varint,2,opt,name=nonce_group" json:"nonce_group,omitempty"`
	Nonce            *uint64          `protobuf:"varint,3,opt,name=nonce" json:"nonce,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *MutateRowsRequest) Reset()         { *m = MutateRowsRequest{} }
func (m *MutateRowsRequest) String() string { return proto.CompactTextString(m) }
func (*MutateRowsRequest) ProtoMessage()    {}

func (m *MutateRowsRequest) GetMutationRequest() []*MutationProto {
	if m != nil {
		return m.MutationRequest
	}
	return nil
}

func (m *MutateRowsRequest) GetNonceGroup() uint64 {
	if m != nil && m.NonceGroup != nil {
		return *m.NonceGroup
	}
	return 0
}

func (m *MutateRowsRequest) GetNonce() uint64 {
	if m != nil && m.Nonce != nil {
		return *m.Nonce
	}
	return 0
}

type MutateRowsResponse struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *MutateRowsResponse) Reset()         { *m = MutateRowsResponse{} }
func (m *MutateRowsResponse) String() string { return proto.CompactTextString(m) }
func (*MutateRowsResponse) ProtoMessage()    {}

func init() {
}
