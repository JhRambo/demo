// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: msgpack.proto

/*
Package msgpack is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package msgpack

import (
	"context"
	"io"
	"net/http"

	"github.com/golang/protobuf/descriptor"
	"github.com/golang/protobuf/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// Suppress "imported and not used" errors
var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray
var _ = descriptor.ForMessage
var _ = metadata.Join

func request_MsgpackHttp_Binary_0(ctx context.Context, marshaler runtime.Marshaler, client MsgpackHttpClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq MsgpackRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.Binary(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_MsgpackHttp_Binary_0(ctx context.Context, marshaler runtime.Marshaler, server MsgpackHttpServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq MsgpackRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.Binary(ctx, &protoReq)
	return msg, metadata, err

}

// RegisterMsgpackHttpHandlerServer registers the http handlers for service MsgpackHttp to "mux".
// UnaryRPC     :call MsgpackHttpServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterMsgpackHttpHandlerFromEndpoint instead.
func RegisterMsgpackHttpHandlerServer(ctx context.Context, mux *runtime.ServeMux, server MsgpackHttpServer) error {

	mux.Handle("POST", pattern_MsgpackHttp_Binary_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_MsgpackHttp_Binary_0(rctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_MsgpackHttp_Binary_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterMsgpackHttpHandlerFromEndpoint is same as RegisterMsgpackHttpHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterMsgpackHttpHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterMsgpackHttpHandler(ctx, mux, conn)
}

// RegisterMsgpackHttpHandler registers the http handlers for service MsgpackHttp to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterMsgpackHttpHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterMsgpackHttpHandlerClient(ctx, mux, NewMsgpackHttpClient(conn))
}

// RegisterMsgpackHttpHandlerClient registers the http handlers for service MsgpackHttp
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "MsgpackHttpClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "MsgpackHttpClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "MsgpackHttpClient" to call the correct interceptors.
func RegisterMsgpackHttpHandlerClient(ctx context.Context, mux *runtime.ServeMux, client MsgpackHttpClient) error {

	mux.Handle("POST", pattern_MsgpackHttp_Binary_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_MsgpackHttp_Binary_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_MsgpackHttp_Binary_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_MsgpackHttp_Binary_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0}, []string{"binary"}, "", runtime.AssumeColonVerbOpt(true)))
)

var (
	forward_MsgpackHttp_Binary_0 = runtime.ForwardResponseMessage
)
