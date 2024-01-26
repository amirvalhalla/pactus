# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: network.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\rnetwork.proto\x12\x06pactus\"\x17\n\x15GetNetworkInfoRequest\"\xae\x04\n\x16GetNetworkInfoResponse\x12!\n\x0cnetwork_name\x18\x01 \x01(\tR\x0bnetworkName\x12(\n\x10total_sent_bytes\x18\x02 \x01(\rR\x0etotalSentBytes\x12\x30\n\x14total_received_bytes\x18\x03 \x01(\rR\x12totalReceivedBytes\x12\x32\n\x15\x63onnected_peers_count\x18\x04 \x01(\rR\x13\x63onnectedPeersCount\x12\x39\n\x0f\x63onnected_peers\x18\x05 \x03(\x0b\x32\x10.pactus.PeerInfoR\x0e\x63onnectedPeers\x12L\n\nsent_bytes\x18\x06 \x03(\x0b\x32-.pactus.GetNetworkInfoResponse.SentBytesEntryR\tsentBytes\x12X\n\x0ereceived_bytes\x18\x07 \x03(\x0b\x32\x31.pactus.GetNetworkInfoResponse.ReceivedBytesEntryR\rreceivedBytes\x1a<\n\x0eSentBytesEntry\x12\x10\n\x03key\x18\x01 \x01(\rR\x03key\x12\x14\n\x05value\x18\x02 \x01(\x04R\x05value:\x02\x38\x01\x1a@\n\x12ReceivedBytesEntry\x12\x10\n\x03key\x18\x01 \x01(\rR\x03key\x12\x14\n\x05value\x18\x02 \x01(\x04R\x05value:\x02\x38\x01\"\x14\n\x12GetNodeInfoRequest\"\x98\x02\n\x13GetNodeInfoResponse\x12\x18\n\x07moniker\x18\x01 \x01(\tR\x07moniker\x12\x14\n\x05\x61gent\x18\x02 \x01(\tR\x05\x61gent\x12\x17\n\x07peer_id\x18\x03 \x01(\x0cR\x06peerId\x12\x1d\n\nstarted_at\x18\x04 \x01(\x04R\tstartedAt\x12\"\n\x0creachability\x18\x05 \x01(\tR\x0creachability\x12\x1a\n\x08services\x18\x06 \x03(\x05R\x08services\x12%\n\x0eservices_names\x18\x07 \x03(\tR\rservicesNames\x12\x14\n\x05\x61\x64\x64rs\x18\x08 \x03(\tR\x05\x61\x64\x64rs\x12\x1c\n\tprotocols\x18\t \x03(\tR\tprotocols\"\xed\x06\n\x08PeerInfo\x12\x16\n\x06status\x18\x01 \x01(\x05R\x06status\x12\x18\n\x07moniker\x18\x02 \x01(\tR\x07moniker\x12\x14\n\x05\x61gent\x18\x03 \x01(\tR\x05\x61gent\x12\x17\n\x07peer_id\x18\x04 \x01(\x0cR\x06peerId\x12%\n\x0e\x63onsensus_keys\x18\x05 \x03(\tR\rconsensusKeys\x12+\n\x11\x63onsensus_address\x18\x06 \x03(\tR\x10\x63onsensusAddress\x12\x1a\n\x08services\x18\x07 \x01(\rR\x08services\x12&\n\x0flast_block_hash\x18\x08 \x01(\x0cR\rlastBlockHash\x12\x16\n\x06height\x18\t \x01(\rR\x06height\x12+\n\x11received_messages\x18\n \x01(\x05R\x10receivedMessages\x12)\n\x10invalid_messages\x18\x0b \x01(\x05R\x0finvalidMessages\x12\x1b\n\tlast_sent\x18\x0c \x01(\x03R\x08lastSent\x12#\n\rlast_received\x18\r \x01(\x03R\x0clastReceived\x12>\n\nsent_bytes\x18\x0e \x03(\x0b\x32\x1f.pactus.PeerInfo.SentBytesEntryR\tsentBytes\x12J\n\x0ereceived_bytes\x18\x0f \x03(\x0b\x32#.pactus.PeerInfo.ReceivedBytesEntryR\rreceivedBytes\x12\x18\n\x07\x61\x64\x64ress\x18\x10 \x01(\tR\x07\x61\x64\x64ress\x12\x1c\n\tdirection\x18\x11 \x01(\tR\tdirection\x12\x1c\n\tprotocols\x18\x12 \x03(\tR\tprotocols\x12%\n\x0etotal_sessions\x18\x13 \x01(\x05R\rtotalSessions\x12-\n\x12\x63ompleted_sessions\x18\x14 \x01(\x05R\x11\x63ompletedSessions\x1a<\n\x0eSentBytesEntry\x12\x10\n\x03key\x18\x01 \x01(\x05R\x03key\x12\x14\n\x05value\x18\x02 \x01(\x03R\x05value:\x02\x38\x01\x1a@\n\x12ReceivedBytesEntry\x12\x10\n\x03key\x18\x01 \x01(\x05R\x03key\x12\x14\n\x05value\x18\x02 \x01(\x03R\x05value:\x02\x38\x01\x32\xa2\x01\n\x07Network\x12O\n\x0eGetNetworkInfo\x12\x1d.pactus.GetNetworkInfoRequest\x1a\x1e.pactus.GetNetworkInfoResponse\x12\x46\n\x0bGetNodeInfo\x12\x1a.pactus.GetNodeInfoRequest\x1a\x1b.pactus.GetNodeInfoResponseBB\n\x0epactus.networkZ0github.com/pactus-project/pactus/www/grpc/pactusb\x06proto3')

_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, globals())
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'network_pb2', globals())
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'\n\016pactus.networkZ0github.com/pactus-project/pactus/www/grpc/pactus'
  _GETNETWORKINFORESPONSE_SENTBYTESENTRY._options = None
  _GETNETWORKINFORESPONSE_SENTBYTESENTRY._serialized_options = b'8\001'
  _GETNETWORKINFORESPONSE_RECEIVEDBYTESENTRY._options = None
  _GETNETWORKINFORESPONSE_RECEIVEDBYTESENTRY._serialized_options = b'8\001'
  _PEERINFO_SENTBYTESENTRY._options = None
  _PEERINFO_SENTBYTESENTRY._serialized_options = b'8\001'
  _PEERINFO_RECEIVEDBYTESENTRY._options = None
  _PEERINFO_RECEIVEDBYTESENTRY._serialized_options = b'8\001'
  _GETNETWORKINFOREQUEST._serialized_start=25
  _GETNETWORKINFOREQUEST._serialized_end=48
  _GETNETWORKINFORESPONSE._serialized_start=51
  _GETNETWORKINFORESPONSE._serialized_end=609
  _GETNETWORKINFORESPONSE_SENTBYTESENTRY._serialized_start=483
  _GETNETWORKINFORESPONSE_SENTBYTESENTRY._serialized_end=543
  _GETNETWORKINFORESPONSE_RECEIVEDBYTESENTRY._serialized_start=545
  _GETNETWORKINFORESPONSE_RECEIVEDBYTESENTRY._serialized_end=609
  _GETNODEINFOREQUEST._serialized_start=611
  _GETNODEINFOREQUEST._serialized_end=631
  _GETNODEINFORESPONSE._serialized_start=634
  _GETNODEINFORESPONSE._serialized_end=914
  _PEERINFO._serialized_start=917
  _PEERINFO._serialized_end=1794
  _PEERINFO_SENTBYTESENTRY._serialized_start=1668
  _PEERINFO_SENTBYTESENTRY._serialized_end=1728
  _PEERINFO_RECEIVEDBYTESENTRY._serialized_start=1730
  _PEERINFO_RECEIVEDBYTESENTRY._serialized_end=1794
  _NETWORK._serialized_start=1797
  _NETWORK._serialized_end=1959
# @@protoc_insertion_point(module_scope)
