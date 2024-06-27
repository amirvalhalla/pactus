// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: network.proto

package pactus

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

// Request message for retrieving overall network information.
type GetNetworkInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Only returns the peers with connected status
	OnlyConnected bool `protobuf:"varint,1,opt,name=only_connected,json=onlyConnected,proto3" json:"only_connected,omitempty"`
}

func (x *GetNetworkInfoRequest) Reset() {
	*x = GetNetworkInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_network_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNetworkInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNetworkInfoRequest) ProtoMessage() {}

func (x *GetNetworkInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_network_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNetworkInfoRequest.ProtoReflect.Descriptor instead.
func (*GetNetworkInfoRequest) Descriptor() ([]byte, []int) {
	return file_network_proto_rawDescGZIP(), []int{0}
}

func (x *GetNetworkInfoRequest) GetOnlyConnected() bool {
	if x != nil {
		return x.OnlyConnected
	}
	return false
}

// Response message containing information about the overall network.
type GetNetworkInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the network.
	NetworkName string `protobuf:"bytes,1,opt,name=network_name,json=networkName,proto3" json:"network_name,omitempty"`
	// Total bytes sent across the network.
	TotalSentBytes int64 `protobuf:"varint,2,opt,name=total_sent_bytes,json=totalSentBytes,proto3" json:"total_sent_bytes,omitempty"`
	// Total bytes received across the network.
	TotalReceivedBytes int64 `protobuf:"varint,3,opt,name=total_received_bytes,json=totalReceivedBytes,proto3" json:"total_received_bytes,omitempty"`
	// Number of connected peers.
	ConnectedPeersCount uint32 `protobuf:"varint,4,opt,name=connected_peers_count,json=connectedPeersCount,proto3" json:"connected_peers_count,omitempty"`
	// List of connected peers.
	ConnectedPeers []*PeerInfo `protobuf:"bytes,5,rep,name=connected_peers,json=connectedPeers,proto3" json:"connected_peers,omitempty"`
	// Bytes sent per peer ID.
	SentBytes map[int32]int64 `protobuf:"bytes,6,rep,name=sent_bytes,json=sentBytes,proto3" json:"sent_bytes,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	// Bytes received per peer ID.
	ReceivedBytes map[int32]int64 `protobuf:"bytes,7,rep,name=received_bytes,json=receivedBytes,proto3" json:"received_bytes,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *GetNetworkInfoResponse) Reset() {
	*x = GetNetworkInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_network_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNetworkInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNetworkInfoResponse) ProtoMessage() {}

func (x *GetNetworkInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_network_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNetworkInfoResponse.ProtoReflect.Descriptor instead.
func (*GetNetworkInfoResponse) Descriptor() ([]byte, []int) {
	return file_network_proto_rawDescGZIP(), []int{1}
}

func (x *GetNetworkInfoResponse) GetNetworkName() string {
	if x != nil {
		return x.NetworkName
	}
	return ""
}

func (x *GetNetworkInfoResponse) GetTotalSentBytes() int64 {
	if x != nil {
		return x.TotalSentBytes
	}
	return 0
}

func (x *GetNetworkInfoResponse) GetTotalReceivedBytes() int64 {
	if x != nil {
		return x.TotalReceivedBytes
	}
	return 0
}

func (x *GetNetworkInfoResponse) GetConnectedPeersCount() uint32 {
	if x != nil {
		return x.ConnectedPeersCount
	}
	return 0
}

func (x *GetNetworkInfoResponse) GetConnectedPeers() []*PeerInfo {
	if x != nil {
		return x.ConnectedPeers
	}
	return nil
}

func (x *GetNetworkInfoResponse) GetSentBytes() map[int32]int64 {
	if x != nil {
		return x.SentBytes
	}
	return nil
}

func (x *GetNetworkInfoResponse) GetReceivedBytes() map[int32]int64 {
	if x != nil {
		return x.ReceivedBytes
	}
	return nil
}

// Request message for retrieving information about a specific node in the
// network.
type GetNodeInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetNodeInfoRequest) Reset() {
	*x = GetNodeInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_network_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNodeInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNodeInfoRequest) ProtoMessage() {}

func (x *GetNodeInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_network_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNodeInfoRequest.ProtoReflect.Descriptor instead.
func (*GetNodeInfoRequest) Descriptor() ([]byte, []int) {
	return file_network_proto_rawDescGZIP(), []int{2}
}

// Response message containing information about the overall network.
type ConnectionInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Total number of the connection.
	Connections uint64 `protobuf:"varint,1,opt,name=connections,proto3" json:"connections,omitempty"`
	// Number of inbound connections.
	InboundConnections uint64 `protobuf:"varint,2,opt,name=inbound_connections,json=inboundConnections,proto3" json:"inbound_connections,omitempty"`
	// Number of outbound connections.
	OutboundConnections uint64 `protobuf:"varint,3,opt,name=outbound_connections,json=outboundConnections,proto3" json:"outbound_connections,omitempty"`
}

func (x *ConnectionInfo) Reset() {
	*x = ConnectionInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_network_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectionInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectionInfo) ProtoMessage() {}

func (x *ConnectionInfo) ProtoReflect() protoreflect.Message {
	mi := &file_network_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectionInfo.ProtoReflect.Descriptor instead.
func (*ConnectionInfo) Descriptor() ([]byte, []int) {
	return file_network_proto_rawDescGZIP(), []int{3}
}

func (x *ConnectionInfo) GetConnections() uint64 {
	if x != nil {
		return x.Connections
	}
	return 0
}

func (x *ConnectionInfo) GetInboundConnections() uint64 {
	if x != nil {
		return x.InboundConnections
	}
	return 0
}

func (x *ConnectionInfo) GetOutboundConnections() uint64 {
	if x != nil {
		return x.OutboundConnections
	}
	return 0
}

// Response message containing information about a specific node in the network.
type GetNodeInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Moniker of the node.
	Moniker string `protobuf:"bytes,1,opt,name=moniker,proto3" json:"moniker,omitempty"`
	// Agent information of the node.
	Agent string `protobuf:"bytes,2,opt,name=agent,proto3" json:"agent,omitempty"`
	// Peer ID of the node.
	PeerId string `protobuf:"bytes,3,opt,name=peer_id,json=peerId,proto3" json:"peer_id,omitempty"`
	// Timestamp when the node started.
	StartedAt uint64 `protobuf:"varint,4,opt,name=started_at,json=startedAt,proto3" json:"started_at,omitempty"`
	// Reachability status of the node.
	Reachability string `protobuf:"bytes,5,opt,name=reachability,proto3" json:"reachability,omitempty"`
	// List of services provided by the node.
	Services []int32 `protobuf:"varint,6,rep,packed,name=services,proto3" json:"services,omitempty"`
	// Names of services provided by the node.
	ServicesNames []string `protobuf:"bytes,7,rep,name=services_names,json=servicesNames,proto3" json:"services_names,omitempty"`
	// List of addresses associated with the node.
	LocalAddrs []string `protobuf:"bytes,8,rep,name=local_addrs,json=localAddrs,proto3" json:"local_addrs,omitempty"`
	// List of protocols supported by the node.
	Protocols []string `protobuf:"bytes,9,rep,name=protocols,proto3" json:"protocols,omitempty"`
	// Clock offset
	ClockOffset float64 `protobuf:"fixed64,13,opt,name=clock_offset,json=clockOffset,proto3" json:"clock_offset,omitempty"`
	// Connection information
	ConnectionInfo *ConnectionInfo `protobuf:"bytes,14,opt,name=connection_info,json=connectionInfo,proto3" json:"connection_info,omitempty"`
}

func (x *GetNodeInfoResponse) Reset() {
	*x = GetNodeInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_network_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNodeInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNodeInfoResponse) ProtoMessage() {}

func (x *GetNodeInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_network_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNodeInfoResponse.ProtoReflect.Descriptor instead.
func (*GetNodeInfoResponse) Descriptor() ([]byte, []int) {
	return file_network_proto_rawDescGZIP(), []int{4}
}

func (x *GetNodeInfoResponse) GetMoniker() string {
	if x != nil {
		return x.Moniker
	}
	return ""
}

func (x *GetNodeInfoResponse) GetAgent() string {
	if x != nil {
		return x.Agent
	}
	return ""
}

func (x *GetNodeInfoResponse) GetPeerId() string {
	if x != nil {
		return x.PeerId
	}
	return ""
}

func (x *GetNodeInfoResponse) GetStartedAt() uint64 {
	if x != nil {
		return x.StartedAt
	}
	return 0
}

func (x *GetNodeInfoResponse) GetReachability() string {
	if x != nil {
		return x.Reachability
	}
	return ""
}

func (x *GetNodeInfoResponse) GetServices() []int32 {
	if x != nil {
		return x.Services
	}
	return nil
}

func (x *GetNodeInfoResponse) GetServicesNames() []string {
	if x != nil {
		return x.ServicesNames
	}
	return nil
}

func (x *GetNodeInfoResponse) GetLocalAddrs() []string {
	if x != nil {
		return x.LocalAddrs
	}
	return nil
}

func (x *GetNodeInfoResponse) GetProtocols() []string {
	if x != nil {
		return x.Protocols
	}
	return nil
}

func (x *GetNodeInfoResponse) GetClockOffset() float64 {
	if x != nil {
		return x.ClockOffset
	}
	return 0
}

func (x *GetNodeInfoResponse) GetConnectionInfo() *ConnectionInfo {
	if x != nil {
		return x.ConnectionInfo
	}
	return nil
}

// Information about a peer in the network.
type PeerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Status of the peer.
	Status int32 `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	// Moniker of the peer.
	Moniker string `protobuf:"bytes,2,opt,name=moniker,proto3" json:"moniker,omitempty"`
	// Agent information of the peer.
	Agent string `protobuf:"bytes,3,opt,name=agent,proto3" json:"agent,omitempty"`
	// Peer ID of the peer.
	PeerId string `protobuf:"bytes,4,opt,name=peer_id,json=peerId,proto3" json:"peer_id,omitempty"`
	// Consensus keys used by the peer.
	ConsensusKeys []string `protobuf:"bytes,5,rep,name=consensus_keys,json=consensusKeys,proto3" json:"consensus_keys,omitempty"`
	// Consensus address of the peer.
	ConsensusAddress []string `protobuf:"bytes,6,rep,name=consensus_address,json=consensusAddress,proto3" json:"consensus_address,omitempty"`
	// Services provided by the peer.
	Services uint32 `protobuf:"varint,7,opt,name=services,proto3" json:"services,omitempty"`
	// Hash of the last block the peer knows.
	LastBlockHash string `protobuf:"bytes,8,opt,name=last_block_hash,json=lastBlockHash,proto3" json:"last_block_hash,omitempty"`
	// Height of the peer in the blockchain.
	Height uint32 `protobuf:"varint,9,opt,name=height,proto3" json:"height,omitempty"`
	// Count of received bundles.
	ReceivedBundles int32 `protobuf:"varint,10,opt,name=received_bundles,json=receivedBundles,proto3" json:"received_bundles,omitempty"`
	// Count of invalid bundles received.
	InvalidBundles int32 `protobuf:"varint,11,opt,name=invalid_bundles,json=invalidBundles,proto3" json:"invalid_bundles,omitempty"`
	// Timestamp of the last sent bundle.
	LastSent int64 `protobuf:"varint,12,opt,name=last_sent,json=lastSent,proto3" json:"last_sent,omitempty"`
	// Timestamp of the last received bundle.
	LastReceived int64 `protobuf:"varint,13,opt,name=last_received,json=lastReceived,proto3" json:"last_received,omitempty"`
	// Bytes sent per message type.
	SentBytes map[int32]int64 `protobuf:"bytes,14,rep,name=sent_bytes,json=sentBytes,proto3" json:"sent_bytes,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	// Bytes received per message type.
	ReceivedBytes map[int32]int64 `protobuf:"bytes,15,rep,name=received_bytes,json=receivedBytes,proto3" json:"received_bytes,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	// Network address of the peer.
	Address string `protobuf:"bytes,16,opt,name=address,proto3" json:"address,omitempty"`
	// Direction of connection with the peer.
	Direction string `protobuf:"bytes,17,opt,name=direction,proto3" json:"direction,omitempty"`
	// List of protocols supported by the peer.
	Protocols []string `protobuf:"bytes,18,rep,name=protocols,proto3" json:"protocols,omitempty"`
	// Total sessions with the peer.
	TotalSessions int32 `protobuf:"varint,19,opt,name=total_sessions,json=totalSessions,proto3" json:"total_sessions,omitempty"`
	// Completed sessions with the peer.
	CompletedSessions int32 `protobuf:"varint,20,opt,name=completed_sessions,json=completedSessions,proto3" json:"completed_sessions,omitempty"`
}

func (x *PeerInfo) Reset() {
	*x = PeerInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_network_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeerInfo) ProtoMessage() {}

func (x *PeerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_network_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeerInfo.ProtoReflect.Descriptor instead.
func (*PeerInfo) Descriptor() ([]byte, []int) {
	return file_network_proto_rawDescGZIP(), []int{5}
}

func (x *PeerInfo) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *PeerInfo) GetMoniker() string {
	if x != nil {
		return x.Moniker
	}
	return ""
}

func (x *PeerInfo) GetAgent() string {
	if x != nil {
		return x.Agent
	}
	return ""
}

func (x *PeerInfo) GetPeerId() string {
	if x != nil {
		return x.PeerId
	}
	return ""
}

func (x *PeerInfo) GetConsensusKeys() []string {
	if x != nil {
		return x.ConsensusKeys
	}
	return nil
}

func (x *PeerInfo) GetConsensusAddress() []string {
	if x != nil {
		return x.ConsensusAddress
	}
	return nil
}

func (x *PeerInfo) GetServices() uint32 {
	if x != nil {
		return x.Services
	}
	return 0
}

func (x *PeerInfo) GetLastBlockHash() string {
	if x != nil {
		return x.LastBlockHash
	}
	return ""
}

func (x *PeerInfo) GetHeight() uint32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *PeerInfo) GetReceivedBundles() int32 {
	if x != nil {
		return x.ReceivedBundles
	}
	return 0
}

func (x *PeerInfo) GetInvalidBundles() int32 {
	if x != nil {
		return x.InvalidBundles
	}
	return 0
}

func (x *PeerInfo) GetLastSent() int64 {
	if x != nil {
		return x.LastSent
	}
	return 0
}

func (x *PeerInfo) GetLastReceived() int64 {
	if x != nil {
		return x.LastReceived
	}
	return 0
}

func (x *PeerInfo) GetSentBytes() map[int32]int64 {
	if x != nil {
		return x.SentBytes
	}
	return nil
}

func (x *PeerInfo) GetReceivedBytes() map[int32]int64 {
	if x != nil {
		return x.ReceivedBytes
	}
	return nil
}

func (x *PeerInfo) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *PeerInfo) GetDirection() string {
	if x != nil {
		return x.Direction
	}
	return ""
}

func (x *PeerInfo) GetProtocols() []string {
	if x != nil {
		return x.Protocols
	}
	return nil
}

func (x *PeerInfo) GetTotalSessions() int32 {
	if x != nil {
		return x.TotalSessions
	}
	return 0
}

func (x *PeerInfo) GetCompletedSessions() int32 {
	if x != nil {
		return x.CompletedSessions
	}
	return 0
}

var File_network_proto protoreflect.FileDescriptor

var file_network_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x06, 0x70, 0x61, 0x63, 0x74, 0x75, 0x73, 0x22, 0x3e, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x4e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x25, 0x0a, 0x0e, 0x6f, 0x6e, 0x6c, 0x79, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x6f, 0x6e, 0x6c, 0x79, 0x43, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x22, 0xae, 0x04, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x4e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x73,
	0x65, 0x6e, 0x74, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0e, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12,
	0x30, 0x0a, 0x14, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65,
	0x64, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x12, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x42, 0x79, 0x74, 0x65,
	0x73, 0x12, 0x32, 0x0a, 0x15, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x70,
	0x65, 0x65, 0x72, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x13, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x50, 0x65, 0x65, 0x72, 0x73,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x39, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x65, 0x64, 0x5f, 0x70, 0x65, 0x65, 0x72, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x70, 0x61, 0x63, 0x74, 0x75, 0x73, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x0e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x50, 0x65, 0x65, 0x72, 0x73,
	0x12, 0x4c, 0x0a, 0x0a, 0x73, 0x65, 0x6e, 0x74, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x06,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x70, 0x61, 0x63, 0x74, 0x75, 0x73, 0x2e, 0x47, 0x65,
	0x74, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x74, 0x65, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x09, 0x73, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x58,
	0x0a, 0x0e, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73,
	0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x70, 0x61, 0x63, 0x74, 0x75, 0x73, 0x2e,
	0x47, 0x65, 0x74, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x42,
	0x79, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0d, 0x72, 0x65, 0x63, 0x65, 0x69,
	0x76, 0x65, 0x64, 0x42, 0x79, 0x74, 0x65, 0x73, 0x1a, 0x3c, 0x0a, 0x0e, 0x53, 0x65, 0x6e, 0x74,
	0x42, 0x79, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x40, 0x0a, 0x12, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76,
	0x65, 0x64, 0x42, 0x79, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x14, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x4e,
	0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x96,
	0x01, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x2f, 0x0a, 0x13, 0x69, 0x6e, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x12, 0x69, 0x6e, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x31, 0x0a, 0x14, 0x6f, 0x75, 0x74, 0x62, 0x6f, 0x75, 0x6e, 0x64,
	0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x13, 0x6f, 0x75, 0x74, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x87, 0x03, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x4e,
	0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x6f, 0x6e, 0x69, 0x6b, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x6f, 0x6e, 0x69, 0x6b, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x67, 0x65,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x12,
	0x17, 0x0a, 0x07, 0x70, 0x65, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x70, 0x65, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65, 0x61, 0x63, 0x68,
	0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72,
	0x65, 0x61, 0x63, 0x68, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x05, 0x52, 0x08, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x0d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x1f,
	0x0a, 0x0b, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x73, 0x18, 0x08, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x0a, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x73, 0x12,
	0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x18, 0x09, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x12, 0x21, 0x0a,
	0x0c, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x0b, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x12, 0x3f, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69,
	0x6e, 0x66, 0x6f, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x61, 0x63, 0x74,
	0x75, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x0e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66,
	0x6f, 0x22, 0xe9, 0x06, 0x0a, 0x08, 0x50, 0x65, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x6f, 0x6e, 0x69, 0x6b, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x6f, 0x6e, 0x69, 0x6b, 0x65, 0x72,
	0x12, 0x14, 0x0a, 0x05, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x65, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x65, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x5f, 0x6b, 0x65, 0x79,
	0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73,
	0x75, 0x73, 0x4b, 0x65, 0x79, 0x73, 0x12, 0x2b, 0x0a, 0x11, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e,
	0x73, 0x75, 0x73, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x10, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12,
	0x26, 0x0a, 0x0f, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x68, 0x61,
	0x73, 0x68, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12,
	0x29, 0x0a, 0x10, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x5f, 0x62, 0x75, 0x6e, 0x64,
	0x6c, 0x65, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x72, 0x65, 0x63, 0x65, 0x69,
	0x76, 0x65, 0x64, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x69, 0x6e,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0e, 0x69, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x42, 0x75, 0x6e, 0x64,
	0x6c, 0x65, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x6e, 0x74,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x53, 0x65, 0x6e, 0x74,
	0x12, 0x23, 0x0a, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65,
	0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x52, 0x65, 0x63,
	0x65, 0x69, 0x76, 0x65, 0x64, 0x12, 0x3e, 0x0a, 0x0a, 0x73, 0x65, 0x6e, 0x74, 0x5f, 0x62, 0x79,
	0x74, 0x65, 0x73, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x70, 0x61, 0x63, 0x74,
	0x75, 0x73, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x53, 0x65, 0x6e, 0x74,
	0x42, 0x79, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x73, 0x65, 0x6e, 0x74,
	0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x4a, 0x0a, 0x0e, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65,
	0x64, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e,
	0x70, 0x61, 0x63, 0x74, 0x75, 0x73, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e,
	0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x42, 0x79, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x0d, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x42, 0x79, 0x74, 0x65,
	0x73, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x10, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x64,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x18, 0x12, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x13, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0d, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2d,
	0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x73, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x14, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x63, 0x6f, 0x6d, 0x70,
	0x6c, 0x65, 0x74, 0x65, 0x64, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x3c, 0x0a,
	0x0e, 0x53, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x40, 0x0a, 0x12, 0x52,
	0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x42, 0x79, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0xa2, 0x01,
	0x0a, 0x07, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x4f, 0x0a, 0x0e, 0x47, 0x65, 0x74,
	0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1d, 0x2e, 0x70, 0x61,
	0x63, 0x74, 0x75, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x61, 0x63,
	0x74, 0x75, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x0b, 0x47, 0x65,
	0x74, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1a, 0x2e, 0x70, 0x61, 0x63, 0x74,
	0x75, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x61, 0x63, 0x74, 0x75, 0x73, 0x2e, 0x47,
	0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x42, 0x0a, 0x0e, 0x70, 0x61, 0x63, 0x74, 0x75, 0x73, 0x2e, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x70, 0x61, 0x63, 0x74, 0x75, 0x73, 0x2d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f,
	0x70, 0x61, 0x63, 0x74, 0x75, 0x73, 0x2f, 0x77, 0x77, 0x77, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f,
	0x70, 0x61, 0x63, 0x74, 0x75, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_network_proto_rawDescOnce sync.Once
	file_network_proto_rawDescData = file_network_proto_rawDesc
)

func file_network_proto_rawDescGZIP() []byte {
	file_network_proto_rawDescOnce.Do(func() {
		file_network_proto_rawDescData = protoimpl.X.CompressGZIP(file_network_proto_rawDescData)
	})
	return file_network_proto_rawDescData
}

var file_network_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_network_proto_goTypes = []any{
	(*GetNetworkInfoRequest)(nil),  // 0: pactus.GetNetworkInfoRequest
	(*GetNetworkInfoResponse)(nil), // 1: pactus.GetNetworkInfoResponse
	(*GetNodeInfoRequest)(nil),     // 2: pactus.GetNodeInfoRequest
	(*ConnectionInfo)(nil),         // 3: pactus.ConnectionInfo
	(*GetNodeInfoResponse)(nil),    // 4: pactus.GetNodeInfoResponse
	(*PeerInfo)(nil),               // 5: pactus.PeerInfo
	nil,                            // 6: pactus.GetNetworkInfoResponse.SentBytesEntry
	nil,                            // 7: pactus.GetNetworkInfoResponse.ReceivedBytesEntry
	nil,                            // 8: pactus.PeerInfo.SentBytesEntry
	nil,                            // 9: pactus.PeerInfo.ReceivedBytesEntry
}
var file_network_proto_depIdxs = []int32{
	5, // 0: pactus.GetNetworkInfoResponse.connected_peers:type_name -> pactus.PeerInfo
	6, // 1: pactus.GetNetworkInfoResponse.sent_bytes:type_name -> pactus.GetNetworkInfoResponse.SentBytesEntry
	7, // 2: pactus.GetNetworkInfoResponse.received_bytes:type_name -> pactus.GetNetworkInfoResponse.ReceivedBytesEntry
	3, // 3: pactus.GetNodeInfoResponse.connection_info:type_name -> pactus.ConnectionInfo
	8, // 4: pactus.PeerInfo.sent_bytes:type_name -> pactus.PeerInfo.SentBytesEntry
	9, // 5: pactus.PeerInfo.received_bytes:type_name -> pactus.PeerInfo.ReceivedBytesEntry
	0, // 6: pactus.Network.GetNetworkInfo:input_type -> pactus.GetNetworkInfoRequest
	2, // 7: pactus.Network.GetNodeInfo:input_type -> pactus.GetNodeInfoRequest
	1, // 8: pactus.Network.GetNetworkInfo:output_type -> pactus.GetNetworkInfoResponse
	4, // 9: pactus.Network.GetNodeInfo:output_type -> pactus.GetNodeInfoResponse
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_network_proto_init() }
func file_network_proto_init() {
	if File_network_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_network_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GetNetworkInfoRequest); i {
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
		file_network_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GetNetworkInfoResponse); i {
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
		file_network_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GetNodeInfoRequest); i {
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
		file_network_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*ConnectionInfo); i {
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
		file_network_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GetNodeInfoResponse); i {
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
		file_network_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*PeerInfo); i {
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
			RawDescriptor: file_network_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_network_proto_goTypes,
		DependencyIndexes: file_network_proto_depIdxs,
		MessageInfos:      file_network_proto_msgTypes,
	}.Build()
	File_network_proto = out.File
	file_network_proto_rawDesc = nil
	file_network_proto_goTypes = nil
	file_network_proto_depIdxs = nil
}
