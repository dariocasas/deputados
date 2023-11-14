//
//  Generated code. Do not modify.
//  source: proto/fotos_service.proto
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

import 'fotos_service.pb.dart' as $0;

export 'fotos_service.pb.dart';

@$pb.GrpcServiceName('pb.FotosService')
class FotosServiceClient extends $grpc.Client {
  static final _$getFotos = $grpc.ClientMethod<$0.GetFotosRequest, $0.FotoResponse>(
      '/pb.FotosService/GetFotos',
      ($0.GetFotosRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.FotoResponse.fromBuffer(value));

  FotosServiceClient($grpc.ClientChannel channel,
      {$grpc.CallOptions? options,
      $core.Iterable<$grpc.ClientInterceptor>? interceptors})
      : super(channel, options: options,
        interceptors: interceptors);

  $grpc.ResponseStream<$0.FotoResponse> getFotos($0.GetFotosRequest request, {$grpc.CallOptions? options}) {
    return $createStreamingCall(_$getFotos, $async.Stream.fromIterable([request]), options: options);
  }
}

@$pb.GrpcServiceName('pb.FotosService')
abstract class FotosServiceBase extends $grpc.Service {
  $core.String get $name => 'pb.FotosService';

  FotosServiceBase() {
    $addMethod($grpc.ServiceMethod<$0.GetFotosRequest, $0.FotoResponse>(
        'GetFotos',
        getFotos_Pre,
        false,
        true,
        ($core.List<$core.int> value) => $0.GetFotosRequest.fromBuffer(value),
        ($0.FotoResponse value) => value.writeToBuffer()));
  }

  $async.Stream<$0.FotoResponse> getFotos_Pre($grpc.ServiceCall call, $async.Future<$0.GetFotosRequest> request) async* {
    yield* getFotos(call, await request);
  }

  $async.Stream<$0.FotoResponse> getFotos($grpc.ServiceCall call, $0.GetFotosRequest request);
}
