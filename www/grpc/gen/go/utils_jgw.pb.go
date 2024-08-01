// Code generated by protoc-gen-jrpc-gateway. DO NOT EDIT.
// source: utils.proto

/*
Package pactus is a reverse proxy.

It translates gRPC into JSON-RPC 2.0
*/
package pactus

import (
	"context"
	"encoding/json"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/encoding/protojson"
)

type UtilsJsonRPC struct {
	client UtilsClient
}

type paramsAndHeadersUtils struct {
	Headers metadata.MD     `json:"headers,omitempty"`
	Params  json.RawMessage `json:"params"`
}

// RegisterUtilsJsonRPC register the grpc client Utils for json-rpc.
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterUtilsJsonRPC(conn *grpc.ClientConn) *UtilsJsonRPC {
	return &UtilsJsonRPC{
		client: NewUtilsClient(conn),
	}
}

func (s *UtilsJsonRPC) Methods() map[string]func(ctx context.Context, message json.RawMessage) (any, error) {
	return map[string]func(ctx context.Context, params json.RawMessage) (any, error){

		"pactus.utils.sign_message_with_private_key": func(ctx context.Context, data json.RawMessage) (any, error) {
			req := new(SignMessageWithPrivateKeyRequest)

			var jrpcData paramsAndHeadersUtils

			if err := json.Unmarshal(data, &jrpcData); err != nil {
				return nil, err
			}

			err := protojson.Unmarshal(jrpcData.Params, req)
			if err != nil {
				return nil, err
			}

			return s.client.SignMessageWithPrivateKey(metadata.NewOutgoingContext(ctx, jrpcData.Headers), req)
		},

		"pactus.utils.verify_message": func(ctx context.Context, data json.RawMessage) (any, error) {
			req := new(VerifyMessageRequest)

			var jrpcData paramsAndHeadersUtils

			if err := json.Unmarshal(data, &jrpcData); err != nil {
				return nil, err
			}

			err := protojson.Unmarshal(jrpcData.Params, req)
			if err != nil {
				return nil, err
			}

			return s.client.VerifyMessage(metadata.NewOutgoingContext(ctx, jrpcData.Headers), req)
		},
	}
}