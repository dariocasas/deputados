//
//  Generated code. Do not modify.
//  source: proto/server_service.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:async' as $async;
import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'package:protobuf/protobuf.dart' as $pb;

import 'server_service.pb.dart' as $0;

export 'server_service.pb.dart';

@$pb.GrpcServiceName('pb.ServerService')
class ServerServiceClient extends $grpc.Client {
  static final _$shutdown = $grpc.ClientMethod<$0.ShutdownRequest, $0.ShutdownResponse>(
      '/pb.ServerService/Shutdown',
      ($0.ShutdownRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.ShutdownResponse.fromBuffer(value));

  ServerServiceClient($grpc.ClientChannel channel,
      {$grpc.CallOptions? options,
      $core.Iterable<$grpc.ClientInterceptor>? interceptors})
      : super(channel, options: options,
        interceptors: interceptors);

  $grpc.ResponseFuture<$0.ShutdownResponse> shutdown($0.ShutdownRequest request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$shutdown, request, options: options);
  }
}

@$pb.GrpcServiceName('pb.ServerService')
abstract class ServerServiceBase extends $grpc.Service {
  $core.String get $name => 'pb.ServerService';

  ServerServiceBase() {
    $addMethod($grpc.ServiceMethod<$0.ShutdownRequest, $0.ShutdownResponse>(
        'Shutdown',
        shutdown_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.ShutdownRequest.fromBuffer(value),
        ($0.ShutdownResponse value) => value.writeToBuffer()));
  }

  $async.Future<$0.ShutdownResponse> shutdown_Pre($grpc.ServiceCall call, $async.Future<$0.ShutdownRequest> request) async {
    return shutdown(call, await request);
  }

  $async.Future<$0.ShutdownResponse> shutdown($grpc.ServiceCall call, $0.ShutdownRequest request);
}
