// Code generated by protoc-gen-go.
// source: ZooKeeper.proto
// DO NOT EDIT!

package pb

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type SplitLogTask_State int32

const (
	SplitLogTask_UNASSIGNED SplitLogTask_State = 0
	SplitLogTask_OWNED      SplitLogTask_State = 1
	SplitLogTask_RESIGNED   SplitLogTask_State = 2
	SplitLogTask_DONE       SplitLogTask_State = 3
	SplitLogTask_ERR        SplitLogTask_State = 4
)

var SplitLogTask_State_name = map[int32]string{
	0: "UNASSIGNED",
	1: "OWNED",
	2: "RESIGNED",
	3: "DONE",
	4: "ERR",
}
var SplitLogTask_State_value = map[string]int32{
	"UNASSIGNED": 0,
	"OWNED":      1,
	"RESIGNED":   2,
	"DONE":       3,
	"ERR":        4,
}

func (x SplitLogTask_State) Enum() *SplitLogTask_State {
	p := new(SplitLogTask_State)
	*p = x
	return p
}
func (x SplitLogTask_State) String() string {
	return proto.EnumName(SplitLogTask_State_name, int32(x))
}
func (x *SplitLogTask_State) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(SplitLogTask_State_value, data, "SplitLogTask_State")
	if err != nil {
		return err
	}
	*x = SplitLogTask_State(value)
	return nil
}

type SplitLogTask_RecoveryMode int32

const (
	SplitLogTask_UNKNOWN       SplitLogTask_RecoveryMode = 0
	SplitLogTask_LOG_SPLITTING SplitLogTask_RecoveryMode = 1
	SplitLogTask_LOG_REPLAY    SplitLogTask_RecoveryMode = 2
)

var SplitLogTask_RecoveryMode_name = map[int32]string{
	0: "UNKNOWN",
	1: "LOG_SPLITTING",
	2: "LOG_REPLAY",
}
var SplitLogTask_RecoveryMode_value = map[string]int32{
	"UNKNOWN":       0,
	"LOG_SPLITTING": 1,
	"LOG_REPLAY":    2,
}

func (x SplitLogTask_RecoveryMode) Enum() *SplitLogTask_RecoveryMode {
	p := new(SplitLogTask_RecoveryMode)
	*p = x
	return p
}
func (x SplitLogTask_RecoveryMode) String() string {
	return proto.EnumName(SplitLogTask_RecoveryMode_name, int32(x))
}
func (x *SplitLogTask_RecoveryMode) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(SplitLogTask_RecoveryMode_value, data, "SplitLogTask_RecoveryMode")
	if err != nil {
		return err
	}
	*x = SplitLogTask_RecoveryMode(value)
	return nil
}

// Table's current state
type Table_State int32

const (
	Table_ENABLED   Table_State = 0
	Table_DISABLED  Table_State = 1
	Table_DISABLING Table_State = 2
	Table_ENABLING  Table_State = 3
)

var Table_State_name = map[int32]string{
	0: "ENABLED",
	1: "DISABLED",
	2: "DISABLING",
	3: "ENABLING",
}
var Table_State_value = map[string]int32{
	"ENABLED":   0,
	"DISABLED":  1,
	"DISABLING": 2,
	"ENABLING":  3,
}

func (x Table_State) Enum() *Table_State {
	p := new(Table_State)
	*p = x
	return p
}
func (x Table_State) String() string {
	return proto.EnumName(Table_State_name, int32(x))
}
func (x *Table_State) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Table_State_value, data, "Table_State")
	if err != nil {
		return err
	}
	*x = Table_State(value)
	return nil
}

type ReplicationState_State int32

const (
	ReplicationState_ENABLED  ReplicationState_State = 0
	ReplicationState_DISABLED ReplicationState_State = 1
)

var ReplicationState_State_name = map[int32]string{
	0: "ENABLED",
	1: "DISABLED",
}
var ReplicationState_State_value = map[string]int32{
	"ENABLED":  0,
	"DISABLED": 1,
}

func (x ReplicationState_State) Enum() *ReplicationState_State {
	p := new(ReplicationState_State)
	*p = x
	return p
}
func (x ReplicationState_State) String() string {
	return proto.EnumName(ReplicationState_State_name, int32(x))
}
func (x *ReplicationState_State) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ReplicationState_State_value, data, "ReplicationState_State")
	if err != nil {
		return err
	}
	*x = ReplicationState_State(value)
	return nil
}

// *
// Content of the meta-region-server znode.
type MetaRegionServer struct {
	// The ServerName hosting the meta region currently, or destination server,
	// if meta region is in transition.
	Server *ServerName `protobuf:"bytes,1,req,name=server" json:"server,omitempty"`
	// The major version of the rpc the server speaks.  This is used so that
	// clients connecting to the cluster can have prior knowledge of what version
	// to send to a RegionServer.  AsyncHBase will use this to detect versions.
	RpcVersion *uint32 `protobuf:"varint,2,opt,name=rpc_version" json:"rpc_version,omitempty"`
	// State of the region transition. OPEN means fully operational 'hbase:meta'
	State            *RegionState_State `protobuf:"varint,3,opt,name=state,enum=pb.RegionState_State" json:"state,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *MetaRegionServer) Reset()         { *m = MetaRegionServer{} }
func (m *MetaRegionServer) String() string { return proto.CompactTextString(m) }
func (*MetaRegionServer) ProtoMessage()    {}

func (m *MetaRegionServer) GetServer() *ServerName {
	if m != nil {
		return m.Server
	}
	return nil
}

func (m *MetaRegionServer) GetRpcVersion() uint32 {
	if m != nil && m.RpcVersion != nil {
		return *m.RpcVersion
	}
	return 0
}

func (m *MetaRegionServer) GetState() RegionState_State {
	if m != nil && m.State != nil {
		return *m.State
	}
	return RegionState_OFFLINE
}

// *
// Content of the master znode.
type Master struct {
	// The ServerName of the current Master
	Master *ServerName `protobuf:"bytes,1,req,name=master" json:"master,omitempty"`
	// Major RPC version so that clients can know what version the master can accept.
	RpcVersion       *uint32 `protobuf:"varint,2,opt,name=rpc_version" json:"rpc_version,omitempty"`
	InfoPort         *uint32 `protobuf:"varint,3,opt,name=info_port" json:"info_port,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Master) Reset()         { *m = Master{} }
func (m *Master) String() string { return proto.CompactTextString(m) }
func (*Master) ProtoMessage()    {}

func (m *Master) GetMaster() *ServerName {
	if m != nil {
		return m.Master
	}
	return nil
}

func (m *Master) GetRpcVersion() uint32 {
	if m != nil && m.RpcVersion != nil {
		return *m.RpcVersion
	}
	return 0
}

func (m *Master) GetInfoPort() uint32 {
	if m != nil && m.InfoPort != nil {
		return *m.InfoPort
	}
	return 0
}

// *
// Content of the '/hbase/running', cluster state, znode.
type ClusterUp struct {
	// If this znode is present, cluster is up.  Currently
	// the data is cluster start_date.
	StartDate        *string `protobuf:"bytes,1,req,name=start_date" json:"start_date,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ClusterUp) Reset()         { *m = ClusterUp{} }
func (m *ClusterUp) String() string { return proto.CompactTextString(m) }
func (*ClusterUp) ProtoMessage()    {}

func (m *ClusterUp) GetStartDate() string {
	if m != nil && m.StartDate != nil {
		return *m.StartDate
	}
	return ""
}

// *
// What we write under unassigned up in zookeeper as a region moves through
// open/close, etc., regions.  Details a region in transition.
type RegionTransition struct {
	// Code for EventType gotten by doing o.a.h.h.EventHandler.EventType.getCode()
	EventTypeCode *uint32 `protobuf:"varint,1,req,name=event_type_code" json:"event_type_code,omitempty"`
	// Full regionname in bytes
	RegionName []byte  `protobuf:"bytes,2,req,name=region_name" json:"region_name,omitempty"`
	CreateTime *uint64 `protobuf:"varint,3,req,name=create_time" json:"create_time,omitempty"`
	// The region server where the transition will happen or is happening
	ServerName       *ServerName `protobuf:"bytes,4,req,name=server_name" json:"server_name,omitempty"`
	Payload          []byte      `protobuf:"bytes,5,opt,name=payload" json:"payload,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *RegionTransition) Reset()         { *m = RegionTransition{} }
func (m *RegionTransition) String() string { return proto.CompactTextString(m) }
func (*RegionTransition) ProtoMessage()    {}

func (m *RegionTransition) GetEventTypeCode() uint32 {
	if m != nil && m.EventTypeCode != nil {
		return *m.EventTypeCode
	}
	return 0
}

func (m *RegionTransition) GetRegionName() []byte {
	if m != nil {
		return m.RegionName
	}
	return nil
}

func (m *RegionTransition) GetCreateTime() uint64 {
	if m != nil && m.CreateTime != nil {
		return *m.CreateTime
	}
	return 0
}

func (m *RegionTransition) GetServerName() *ServerName {
	if m != nil {
		return m.ServerName
	}
	return nil
}

func (m *RegionTransition) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

// *
// WAL SplitLog directory znodes have this for content.  Used doing distributed
// WAL splitting.  Holds current state and name of server that originated split.
type SplitLogTask struct {
	State            *SplitLogTask_State        `protobuf:"varint,1,req,name=state,enum=pb.SplitLogTask_State" json:"state,omitempty"`
	ServerName       *ServerName                `protobuf:"bytes,2,req,name=server_name" json:"server_name,omitempty"`
	Mode             *SplitLogTask_RecoveryMode `protobuf:"varint,3,opt,name=mode,enum=pb.SplitLogTask_RecoveryMode,def=0" json:"mode,omitempty"`
	XXX_unrecognized []byte                     `json:"-"`
}

func (m *SplitLogTask) Reset()         { *m = SplitLogTask{} }
func (m *SplitLogTask) String() string { return proto.CompactTextString(m) }
func (*SplitLogTask) ProtoMessage()    {}

const Default_SplitLogTask_Mode SplitLogTask_RecoveryMode = SplitLogTask_UNKNOWN

func (m *SplitLogTask) GetState() SplitLogTask_State {
	if m != nil && m.State != nil {
		return *m.State
	}
	return SplitLogTask_UNASSIGNED
}

func (m *SplitLogTask) GetServerName() *ServerName {
	if m != nil {
		return m.ServerName
	}
	return nil
}

func (m *SplitLogTask) GetMode() SplitLogTask_RecoveryMode {
	if m != nil && m.Mode != nil {
		return *m.Mode
	}
	return Default_SplitLogTask_Mode
}

// *
// The znode that holds state of table.
type Table struct {
	// This is the table's state.  If no znode for a table,
	// its state is presumed enabled.  See o.a.h.h.zookeeper.ZKTable class
	// for more.
	State            *Table_State `protobuf:"varint,1,req,name=state,enum=pb.Table_State,def=0" json:"state,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *Table) Reset()         { *m = Table{} }
func (m *Table) String() string { return proto.CompactTextString(m) }
func (*Table) ProtoMessage()    {}

const Default_Table_State Table_State = Table_ENABLED

func (m *Table) GetState() Table_State {
	if m != nil && m.State != nil {
		return *m.State
	}
	return Default_Table_State
}

// *
// Used by replication. Holds a replication peer key.
type ReplicationPeer struct {
	// clusterkey is the concatenation of the slave cluster's
	// hbase.zookeeper.quorum:hbase.zookeeper.property.clientPort:zookeeper.znode.parent
	Clusterkey              *string           `protobuf:"bytes,1,req,name=clusterkey" json:"clusterkey,omitempty"`
	ReplicationEndpointImpl *string           `protobuf:"bytes,2,opt,name=replicationEndpointImpl" json:"replicationEndpointImpl,omitempty"`
	Data                    []*BytesBytesPair `protobuf:"bytes,3,rep,name=data" json:"data,omitempty"`
	Configuration           []*NameStringPair `protobuf:"bytes,4,rep,name=configuration" json:"configuration,omitempty"`
	XXX_unrecognized        []byte            `json:"-"`
}

func (m *ReplicationPeer) Reset()         { *m = ReplicationPeer{} }
func (m *ReplicationPeer) String() string { return proto.CompactTextString(m) }
func (*ReplicationPeer) ProtoMessage()    {}

func (m *ReplicationPeer) GetClusterkey() string {
	if m != nil && m.Clusterkey != nil {
		return *m.Clusterkey
	}
	return ""
}

func (m *ReplicationPeer) GetReplicationEndpointImpl() string {
	if m != nil && m.ReplicationEndpointImpl != nil {
		return *m.ReplicationEndpointImpl
	}
	return ""
}

func (m *ReplicationPeer) GetData() []*BytesBytesPair {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *ReplicationPeer) GetConfiguration() []*NameStringPair {
	if m != nil {
		return m.Configuration
	}
	return nil
}

// *
// Used by replication. Holds whether enabled or disabled
type ReplicationState struct {
	State            *ReplicationState_State `protobuf:"varint,1,req,name=state,enum=pb.ReplicationState_State" json:"state,omitempty"`
	XXX_unrecognized []byte                  `json:"-"`
}

func (m *ReplicationState) Reset()         { *m = ReplicationState{} }
func (m *ReplicationState) String() string { return proto.CompactTextString(m) }
func (*ReplicationState) ProtoMessage()    {}

func (m *ReplicationState) GetState() ReplicationState_State {
	if m != nil && m.State != nil {
		return *m.State
	}
	return ReplicationState_ENABLED
}

// *
// Used by replication. Holds the current position in an WAL file.
type ReplicationHLogPosition struct {
	Position         *int64 `protobuf:"varint,1,req,name=position" json:"position,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ReplicationHLogPosition) Reset()         { *m = ReplicationHLogPosition{} }
func (m *ReplicationHLogPosition) String() string { return proto.CompactTextString(m) }
func (*ReplicationHLogPosition) ProtoMessage()    {}

func (m *ReplicationHLogPosition) GetPosition() int64 {
	if m != nil && m.Position != nil {
		return *m.Position
	}
	return 0
}

// *
// Used by replication. Used to lock a region server during failover.
type ReplicationLock struct {
	LockOwner        *string `protobuf:"bytes,1,req,name=lock_owner" json:"lock_owner,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ReplicationLock) Reset()         { *m = ReplicationLock{} }
func (m *ReplicationLock) String() string { return proto.CompactTextString(m) }
func (*ReplicationLock) ProtoMessage()    {}

func (m *ReplicationLock) GetLockOwner() string {
	if m != nil && m.LockOwner != nil {
		return *m.LockOwner
	}
	return ""
}

// *
// Metadata associated with a table lock in zookeeper
type TableLock struct {
	TableName        *TableName  `protobuf:"bytes,1,opt,name=table_name" json:"table_name,omitempty"`
	LockOwner        *ServerName `protobuf:"bytes,2,opt,name=lock_owner" json:"lock_owner,omitempty"`
	ThreadId         *int64      `protobuf:"varint,3,opt,name=thread_id" json:"thread_id,omitempty"`
	IsShared         *bool       `protobuf:"varint,4,opt,name=is_shared" json:"is_shared,omitempty"`
	Purpose          *string     `protobuf:"bytes,5,opt,name=purpose" json:"purpose,omitempty"`
	CreateTime       *int64      `protobuf:"varint,6,opt,name=create_time" json:"create_time,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *TableLock) Reset()         { *m = TableLock{} }
func (m *TableLock) String() string { return proto.CompactTextString(m) }
func (*TableLock) ProtoMessage()    {}

func (m *TableLock) GetTableName() *TableName {
	if m != nil {
		return m.TableName
	}
	return nil
}

func (m *TableLock) GetLockOwner() *ServerName {
	if m != nil {
		return m.LockOwner
	}
	return nil
}

func (m *TableLock) GetThreadId() int64 {
	if m != nil && m.ThreadId != nil {
		return *m.ThreadId
	}
	return 0
}

func (m *TableLock) GetIsShared() bool {
	if m != nil && m.IsShared != nil {
		return *m.IsShared
	}
	return false
}

func (m *TableLock) GetPurpose() string {
	if m != nil && m.Purpose != nil {
		return *m.Purpose
	}
	return ""
}

func (m *TableLock) GetCreateTime() int64 {
	if m != nil && m.CreateTime != nil {
		return *m.CreateTime
	}
	return 0
}

func init() {
	proto.RegisterEnum("pb.SplitLogTask_State", SplitLogTask_State_name, SplitLogTask_State_value)
	proto.RegisterEnum("pb.SplitLogTask_RecoveryMode", SplitLogTask_RecoveryMode_name, SplitLogTask_RecoveryMode_value)
	proto.RegisterEnum("pb.Table_State", Table_State_name, Table_State_value)
	proto.RegisterEnum("pb.ReplicationState_State", ReplicationState_State_name, ReplicationState_State_value)
}
