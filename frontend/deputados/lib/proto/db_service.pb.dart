//
//  Generated code. Do not modify.
//  source: proto/db_service.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

import 'db_service.pbenum.dart';

export 'db_service.pbenum.dart';

class DropDbRequest extends $pb.GeneratedMessage {
  factory DropDbRequest() => create();
  DropDbRequest._() : super();
  factory DropDbRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory DropDbRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'DropDbRequest', package: const $pb.PackageName(_omitMessageNames ? '' : 'pb'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  DropDbRequest clone() => DropDbRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  DropDbRequest copyWith(void Function(DropDbRequest) updates) => super.copyWith((message) => updates(message as DropDbRequest)) as DropDbRequest;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static DropDbRequest create() => DropDbRequest._();
  DropDbRequest createEmptyInstance() => create();
  static $pb.PbList<DropDbRequest> createRepeated() => $pb.PbList<DropDbRequest>();
  @$core.pragma('dart2js:noInline')
  static DropDbRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<DropDbRequest>(create);
  static DropDbRequest? _defaultInstance;
}

class DropDbResponse extends $pb.GeneratedMessage {
  factory DropDbResponse() => create();
  DropDbResponse._() : super();
  factory DropDbResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory DropDbResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'DropDbResponse', package: const $pb.PackageName(_omitMessageNames ? '' : 'pb'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  DropDbResponse clone() => DropDbResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  DropDbResponse copyWith(void Function(DropDbResponse) updates) => super.copyWith((message) => updates(message as DropDbResponse)) as DropDbResponse;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static DropDbResponse create() => DropDbResponse._();
  DropDbResponse createEmptyInstance() => create();
  static $pb.PbList<DropDbResponse> createRepeated() => $pb.PbList<DropDbResponse>();
  @$core.pragma('dart2js:noInline')
  static DropDbResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<DropDbResponse>(create);
  static DropDbResponse? _defaultInstance;
}

class DbStatusRequest extends $pb.GeneratedMessage {
  factory DbStatusRequest() => create();
  DbStatusRequest._() : super();
  factory DbStatusRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory DbStatusRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'DbStatusRequest', package: const $pb.PackageName(_omitMessageNames ? '' : 'pb'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  DbStatusRequest clone() => DbStatusRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  DbStatusRequest copyWith(void Function(DbStatusRequest) updates) => super.copyWith((message) => updates(message as DbStatusRequest)) as DbStatusRequest;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static DbStatusRequest create() => DbStatusRequest._();
  DbStatusRequest createEmptyInstance() => create();
  static $pb.PbList<DbStatusRequest> createRepeated() => $pb.PbList<DbStatusRequest>();
  @$core.pragma('dart2js:noInline')
  static DbStatusRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<DbStatusRequest>(create);
  static DbStatusRequest? _defaultInstance;
}

class DbStatusResponse extends $pb.GeneratedMessage {
  factory DbStatusResponse({
    DbStatusEnum? status,
    $core.int? indexCount,
  }) {
    final $result = create();
    if (status != null) {
      $result.status = status;
    }
    if (indexCount != null) {
      $result.indexCount = indexCount;
    }
    return $result;
  }
  DbStatusResponse._() : super();
  factory DbStatusResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory DbStatusResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'DbStatusResponse', package: const $pb.PackageName(_omitMessageNames ? '' : 'pb'), createEmptyInstance: create)
    ..e<DbStatusEnum>(1, _omitFieldNames ? '' : 'status', $pb.PbFieldType.OE, defaultOrMaker: DbStatusEnum.UNINITIALIZATED, valueOf: DbStatusEnum.valueOf, enumValues: DbStatusEnum.values)
    ..a<$core.int>(2, _omitFieldNames ? '' : 'indexCount', $pb.PbFieldType.O3, protoName: 'indexCount')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  DbStatusResponse clone() => DbStatusResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  DbStatusResponse copyWith(void Function(DbStatusResponse) updates) => super.copyWith((message) => updates(message as DbStatusResponse)) as DbStatusResponse;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static DbStatusResponse create() => DbStatusResponse._();
  DbStatusResponse createEmptyInstance() => create();
  static $pb.PbList<DbStatusResponse> createRepeated() => $pb.PbList<DbStatusResponse>();
  @$core.pragma('dart2js:noInline')
  static DbStatusResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<DbStatusResponse>(create);
  static DbStatusResponse? _defaultInstance;

  @$pb.TagNumber(1)
  DbStatusEnum get status => $_getN(0);
  @$pb.TagNumber(1)
  set status(DbStatusEnum v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasStatus() => $_has(0);
  @$pb.TagNumber(1)
  void clearStatus() => clearField(1);

  @$pb.TagNumber(2)
  $core.int get indexCount => $_getIZ(1);
  @$pb.TagNumber(2)
  set indexCount($core.int v) { $_setSignedInt32(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasIndexCount() => $_has(1);
  @$pb.TagNumber(2)
  void clearIndexCount() => clearField(2);
}

class PopulateIndexRequest extends $pb.GeneratedMessage {
  factory PopulateIndexRequest({
    $core.int? maxItems,
  }) {
    final $result = create();
    if (maxItems != null) {
      $result.maxItems = maxItems;
    }
    return $result;
  }
  PopulateIndexRequest._() : super();
  factory PopulateIndexRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory PopulateIndexRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'PopulateIndexRequest', package: const $pb.PackageName(_omitMessageNames ? '' : 'pb'), createEmptyInstance: create)
    ..a<$core.int>(1, _omitFieldNames ? '' : 'maxItems', $pb.PbFieldType.O3, protoName: 'maxItems')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  PopulateIndexRequest clone() => PopulateIndexRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  PopulateIndexRequest copyWith(void Function(PopulateIndexRequest) updates) => super.copyWith((message) => updates(message as PopulateIndexRequest)) as PopulateIndexRequest;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static PopulateIndexRequest create() => PopulateIndexRequest._();
  PopulateIndexRequest createEmptyInstance() => create();
  static $pb.PbList<PopulateIndexRequest> createRepeated() => $pb.PbList<PopulateIndexRequest>();
  @$core.pragma('dart2js:noInline')
  static PopulateIndexRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<PopulateIndexRequest>(create);
  static PopulateIndexRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $core.int get maxItems => $_getIZ(0);
  @$pb.TagNumber(1)
  set maxItems($core.int v) { $_setSignedInt32(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasMaxItems() => $_has(0);
  @$pb.TagNumber(1)
  void clearMaxItems() => clearField(1);
}

class PopulateIndexResponse extends $pb.GeneratedMessage {
  factory PopulateIndexResponse({
    $core.int? indexCount,
  }) {
    final $result = create();
    if (indexCount != null) {
      $result.indexCount = indexCount;
    }
    return $result;
  }
  PopulateIndexResponse._() : super();
  factory PopulateIndexResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory PopulateIndexResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'PopulateIndexResponse', package: const $pb.PackageName(_omitMessageNames ? '' : 'pb'), createEmptyInstance: create)
    ..a<$core.int>(1, _omitFieldNames ? '' : 'indexCount', $pb.PbFieldType.O3, protoName: 'indexCount')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  PopulateIndexResponse clone() => PopulateIndexResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  PopulateIndexResponse copyWith(void Function(PopulateIndexResponse) updates) => super.copyWith((message) => updates(message as PopulateIndexResponse)) as PopulateIndexResponse;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static PopulateIndexResponse create() => PopulateIndexResponse._();
  PopulateIndexResponse createEmptyInstance() => create();
  static $pb.PbList<PopulateIndexResponse> createRepeated() => $pb.PbList<PopulateIndexResponse>();
  @$core.pragma('dart2js:noInline')
  static PopulateIndexResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<PopulateIndexResponse>(create);
  static PopulateIndexResponse? _defaultInstance;

  @$pb.TagNumber(1)
  $core.int get indexCount => $_getIZ(0);
  @$pb.TagNumber(1)
  set indexCount($core.int v) { $_setSignedInt32(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasIndexCount() => $_has(0);
  @$pb.TagNumber(1)
  void clearIndexCount() => clearField(1);
}

class PopulateDbRequest extends $pb.GeneratedMessage {
  factory PopulateDbRequest({
    $core.int? concurrency,
    $core.int? timeout,
  }) {
    final $result = create();
    if (concurrency != null) {
      $result.concurrency = concurrency;
    }
    if (timeout != null) {
      $result.timeout = timeout;
    }
    return $result;
  }
  PopulateDbRequest._() : super();
  factory PopulateDbRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory PopulateDbRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'PopulateDbRequest', package: const $pb.PackageName(_omitMessageNames ? '' : 'pb'), createEmptyInstance: create)
    ..a<$core.int>(1, _omitFieldNames ? '' : 'concurrency', $pb.PbFieldType.O3)
    ..a<$core.int>(2, _omitFieldNames ? '' : 'timeout', $pb.PbFieldType.O3)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  PopulateDbRequest clone() => PopulateDbRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  PopulateDbRequest copyWith(void Function(PopulateDbRequest) updates) => super.copyWith((message) => updates(message as PopulateDbRequest)) as PopulateDbRequest;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static PopulateDbRequest create() => PopulateDbRequest._();
  PopulateDbRequest createEmptyInstance() => create();
  static $pb.PbList<PopulateDbRequest> createRepeated() => $pb.PbList<PopulateDbRequest>();
  @$core.pragma('dart2js:noInline')
  static PopulateDbRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<PopulateDbRequest>(create);
  static PopulateDbRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $core.int get concurrency => $_getIZ(0);
  @$pb.TagNumber(1)
  set concurrency($core.int v) { $_setSignedInt32(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasConcurrency() => $_has(0);
  @$pb.TagNumber(1)
  void clearConcurrency() => clearField(1);

  @$pb.TagNumber(2)
  $core.int get timeout => $_getIZ(1);
  @$pb.TagNumber(2)
  set timeout($core.int v) { $_setSignedInt32(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasTimeout() => $_has(1);
  @$pb.TagNumber(2)
  void clearTimeout() => clearField(2);
}

class PartialTime extends $pb.GeneratedMessage {
  factory PartialTime({
    $core.int? dbread,
    $core.int? getapi,
    $core.int? getpage,
    $core.int? scanpage,
    $core.int? dbwrite,
  }) {
    final $result = create();
    if (dbread != null) {
      $result.dbread = dbread;
    }
    if (getapi != null) {
      $result.getapi = getapi;
    }
    if (getpage != null) {
      $result.getpage = getpage;
    }
    if (scanpage != null) {
      $result.scanpage = scanpage;
    }
    if (dbwrite != null) {
      $result.dbwrite = dbwrite;
    }
    return $result;
  }
  PartialTime._() : super();
  factory PartialTime.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory PartialTime.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'PartialTime', package: const $pb.PackageName(_omitMessageNames ? '' : 'pb'), createEmptyInstance: create)
    ..a<$core.int>(1, _omitFieldNames ? '' : 'dbread', $pb.PbFieldType.O3)
    ..a<$core.int>(2, _omitFieldNames ? '' : 'getapi', $pb.PbFieldType.O3)
    ..a<$core.int>(3, _omitFieldNames ? '' : 'getpage', $pb.PbFieldType.O3)
    ..a<$core.int>(4, _omitFieldNames ? '' : 'scanpage', $pb.PbFieldType.O3)
    ..a<$core.int>(5, _omitFieldNames ? '' : 'dbwrite', $pb.PbFieldType.O3)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  PartialTime clone() => PartialTime()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  PartialTime copyWith(void Function(PartialTime) updates) => super.copyWith((message) => updates(message as PartialTime)) as PartialTime;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static PartialTime create() => PartialTime._();
  PartialTime createEmptyInstance() => create();
  static $pb.PbList<PartialTime> createRepeated() => $pb.PbList<PartialTime>();
  @$core.pragma('dart2js:noInline')
  static PartialTime getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<PartialTime>(create);
  static PartialTime? _defaultInstance;

  @$pb.TagNumber(1)
  $core.int get dbread => $_getIZ(0);
  @$pb.TagNumber(1)
  set dbread($core.int v) { $_setSignedInt32(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasDbread() => $_has(0);
  @$pb.TagNumber(1)
  void clearDbread() => clearField(1);

  @$pb.TagNumber(2)
  $core.int get getapi => $_getIZ(1);
  @$pb.TagNumber(2)
  set getapi($core.int v) { $_setSignedInt32(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasGetapi() => $_has(1);
  @$pb.TagNumber(2)
  void clearGetapi() => clearField(2);

  @$pb.TagNumber(3)
  $core.int get getpage => $_getIZ(2);
  @$pb.TagNumber(3)
  set getpage($core.int v) { $_setSignedInt32(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasGetpage() => $_has(2);
  @$pb.TagNumber(3)
  void clearGetpage() => clearField(3);

  @$pb.TagNumber(4)
  $core.int get scanpage => $_getIZ(3);
  @$pb.TagNumber(4)
  set scanpage($core.int v) { $_setSignedInt32(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasScanpage() => $_has(3);
  @$pb.TagNumber(4)
  void clearScanpage() => clearField(4);

  @$pb.TagNumber(5)
  $core.int get dbwrite => $_getIZ(4);
  @$pb.TagNumber(5)
  set dbwrite($core.int v) { $_setSignedInt32(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasDbwrite() => $_has(4);
  @$pb.TagNumber(5)
  void clearDbwrite() => clearField(5);
}

class DeputadoResponse extends $pb.GeneratedMessage {
  factory DeputadoResponse({
    $core.bool? error,
    $core.String? errorString,
    $core.int? id,
    $core.String? name,
    $core.int? milliseconds,
    PartialTime? partialTime,
  }) {
    final $result = create();
    if (error != null) {
      $result.error = error;
    }
    if (errorString != null) {
      $result.errorString = errorString;
    }
    if (id != null) {
      $result.id = id;
    }
    if (name != null) {
      $result.name = name;
    }
    if (milliseconds != null) {
      $result.milliseconds = milliseconds;
    }
    if (partialTime != null) {
      $result.partialTime = partialTime;
    }
    return $result;
  }
  DeputadoResponse._() : super();
  factory DeputadoResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory DeputadoResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'DeputadoResponse', package: const $pb.PackageName(_omitMessageNames ? '' : 'pb'), createEmptyInstance: create)
    ..aOB(1, _omitFieldNames ? '' : 'error')
    ..aOS(2, _omitFieldNames ? '' : 'errorString', protoName: 'errorString')
    ..a<$core.int>(3, _omitFieldNames ? '' : 'id', $pb.PbFieldType.O3)
    ..aOS(4, _omitFieldNames ? '' : 'name')
    ..a<$core.int>(5, _omitFieldNames ? '' : 'milliseconds', $pb.PbFieldType.O3)
    ..aOM<PartialTime>(6, _omitFieldNames ? '' : 'partialTime', protoName: 'partialTime', subBuilder: PartialTime.create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  DeputadoResponse clone() => DeputadoResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  DeputadoResponse copyWith(void Function(DeputadoResponse) updates) => super.copyWith((message) => updates(message as DeputadoResponse)) as DeputadoResponse;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static DeputadoResponse create() => DeputadoResponse._();
  DeputadoResponse createEmptyInstance() => create();
  static $pb.PbList<DeputadoResponse> createRepeated() => $pb.PbList<DeputadoResponse>();
  @$core.pragma('dart2js:noInline')
  static DeputadoResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<DeputadoResponse>(create);
  static DeputadoResponse? _defaultInstance;

  @$pb.TagNumber(1)
  $core.bool get error => $_getBF(0);
  @$pb.TagNumber(1)
  set error($core.bool v) { $_setBool(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasError() => $_has(0);
  @$pb.TagNumber(1)
  void clearError() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get errorString => $_getSZ(1);
  @$pb.TagNumber(2)
  set errorString($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasErrorString() => $_has(1);
  @$pb.TagNumber(2)
  void clearErrorString() => clearField(2);

  @$pb.TagNumber(3)
  $core.int get id => $_getIZ(2);
  @$pb.TagNumber(3)
  set id($core.int v) { $_setSignedInt32(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasId() => $_has(2);
  @$pb.TagNumber(3)
  void clearId() => clearField(3);

  @$pb.TagNumber(4)
  $core.String get name => $_getSZ(3);
  @$pb.TagNumber(4)
  set name($core.String v) { $_setString(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasName() => $_has(3);
  @$pb.TagNumber(4)
  void clearName() => clearField(4);

  @$pb.TagNumber(5)
  $core.int get milliseconds => $_getIZ(4);
  @$pb.TagNumber(5)
  set milliseconds($core.int v) { $_setSignedInt32(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasMilliseconds() => $_has(4);
  @$pb.TagNumber(5)
  void clearMilliseconds() => clearField(5);

  @$pb.TagNumber(6)
  PartialTime get partialTime => $_getN(5);
  @$pb.TagNumber(6)
  set partialTime(PartialTime v) { setField(6, v); }
  @$pb.TagNumber(6)
  $core.bool hasPartialTime() => $_has(5);
  @$pb.TagNumber(6)
  void clearPartialTime() => clearField(6);
  @$pb.TagNumber(6)
  PartialTime ensurePartialTime() => $_ensure(5);
}


const _omitFieldNames = $core.bool.fromEnvironment('protobuf.omit_field_names');
const _omitMessageNames = $core.bool.fromEnvironment('protobuf.omit_message_names');
