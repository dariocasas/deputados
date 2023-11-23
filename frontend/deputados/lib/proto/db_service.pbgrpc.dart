//
//  Generated code. Do not modify.
//  source: proto/db_service.proto
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

import 'db_service.pb.dart' as $0;

export 'db_service.pb.dart';

@$pb.GrpcServiceName('pb.DbService')
class DbServiceClient extends $grpc.Client {
  static final _$dropDb = $grpc.ClientMethod<$0.DropDbRequest, $0.DropDbResponse>(
      '/pb.DbService/DropDb',
      ($0.DropDbRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.DropDbResponse.fromBuffer(value));
  static final _$dbStatus = $grpc.ClientMethod<$0.DbStatusRequest, $0.DbStatusResponse>(
      '/pb.DbService/DbStatus',
      ($0.DbStatusRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.DbStatusResponse.fromBuffer(value));
  static final _$populateIndex = $grpc.ClientMethod<$0.PopulateIndexRequest, $0.PopulateIndexResponse>(
      '/pb.DbService/PopulateIndex',
      ($0.PopulateIndexRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.PopulateIndexResponse.fromBuffer(value));
  static final _$populateDb = $grpc.ClientMethod<$0.PopulateDbRequest, $0.DeputadoResponse>(
      '/pb.DbService/PopulateDb',
      ($0.PopulateDbRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.DeputadoResponse.fromBuffer(value));
  static final _$cancelPopulateDb = $grpc.ClientMethod<$0.CancelPopulateDbRequest, $0.CancelPopulateDbResponse>(
      '/pb.DbService/CancelPopulateDb',
      ($0.CancelPopulateDbRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.CancelPopulateDbResponse.fromBuffer(value));

  DbServiceClient($grpc.ClientChannel channel,
      {$grpc.CallOptions? options,
      $core.Iterable<$grpc.ClientInterceptor>? interceptors})
      : super(channel, options: options,
        interceptors: interceptors);

  $grpc.ResponseFuture<$0.DropDbResponse> dropDb($0.DropDbRequest request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$dropDb, request, options: options);
  }

  $grpc.ResponseFuture<$0.DbStatusResponse> dbStatus($0.DbStatusRequest request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$dbStatus, request, options: options);
  }

  $grpc.ResponseFuture<$0.PopulateIndexResponse> populateIndex($0.PopulateIndexRequest request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$populateIndex, request, options: options);
  }

  $grpc.ResponseStream<$0.DeputadoResponse> populateDb($0.PopulateDbRequest request, {$grpc.CallOptions? options}) {
    return $createStreamingCall(_$populateDb, $async.Stream.fromIterable([request]), options: options);
  }

  $grpc.ResponseFuture<$0.CancelPopulateDbResponse> cancelPopulateDb($0.CancelPopulateDbRequest request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$cancelPopulateDb, request, options: options);
  }
}

@$pb.GrpcServiceName('pb.DbService')
abstract class DbServiceBase extends $grpc.Service {
  $core.String get $name => 'pb.DbService';

  DbServiceBase() {
    $addMethod($grpc.ServiceMethod<$0.DropDbRequest, $0.DropDbResponse>(
        'DropDb',
        dropDb_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.DropDbRequest.fromBuffer(value),
        ($0.DropDbResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.DbStatusRequest, $0.DbStatusResponse>(
        'DbStatus',
        dbStatus_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.DbStatusRequest.fromBuffer(value),
        ($0.DbStatusResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.PopulateIndexRequest, $0.PopulateIndexResponse>(
        'PopulateIndex',
        populateIndex_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.PopulateIndexRequest.fromBuffer(value),
        ($0.PopulateIndexResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.PopulateDbRequest, $0.DeputadoResponse>(
        'PopulateDb',
        populateDb_Pre,
        false,
        true,
        ($core.List<$core.int> value) => $0.PopulateDbRequest.fromBuffer(value),
        ($0.DeputadoResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.CancelPopulateDbRequest, $0.CancelPopulateDbResponse>(
        'CancelPopulateDb',
        cancelPopulateDb_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.CancelPopulateDbRequest.fromBuffer(value),
        ($0.CancelPopulateDbResponse value) => value.writeToBuffer()));
  }

  $async.Future<$0.DropDbResponse> dropDb_Pre($grpc.ServiceCall call, $async.Future<$0.DropDbRequest> request) async {
    return dropDb(call, await request);
  }

  $async.Future<$0.DbStatusResponse> dbStatus_Pre($grpc.ServiceCall call, $async.Future<$0.DbStatusRequest> request) async {
    return dbStatus(call, await request);
  }

  $async.Future<$0.PopulateIndexResponse> populateIndex_Pre($grpc.ServiceCall call, $async.Future<$0.PopulateIndexRequest> request) async {
    return populateIndex(call, await request);
  }

  $async.Stream<$0.DeputadoResponse> populateDb_Pre($grpc.ServiceCall call, $async.Future<$0.PopulateDbRequest> request) async* {
    yield* populateDb(call, await request);
  }

  $async.Future<$0.CancelPopulateDbResponse> cancelPopulateDb_Pre($grpc.ServiceCall call, $async.Future<$0.CancelPopulateDbRequest> request) async {
    return cancelPopulateDb(call, await request);
  }

  $async.Future<$0.DropDbResponse> dropDb($grpc.ServiceCall call, $0.DropDbRequest request);
  $async.Future<$0.DbStatusResponse> dbStatus($grpc.ServiceCall call, $0.DbStatusRequest request);
  $async.Future<$0.PopulateIndexResponse> populateIndex($grpc.ServiceCall call, $0.PopulateIndexRequest request);
  $async.Stream<$0.DeputadoResponse> populateDb($grpc.ServiceCall call, $0.PopulateDbRequest request);
  $async.Future<$0.CancelPopulateDbResponse> cancelPopulateDb($grpc.ServiceCall call, $0.CancelPopulateDbRequest request);
}
