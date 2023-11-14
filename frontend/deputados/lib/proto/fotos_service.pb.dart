//
//  Generated code. Do not modify.
//  source: proto/fotos_service.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

class GetFotosRequest extends $pb.GeneratedMessage {
  factory GetFotosRequest() => create();
  GetFotosRequest._() : super();
  factory GetFotosRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory GetFotosRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'GetFotosRequest', package: const $pb.PackageName(_omitMessageNames ? '' : 'pb'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  GetFotosRequest clone() => GetFotosRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  GetFotosRequest copyWith(void Function(GetFotosRequest) updates) => super.copyWith((message) => updates(message as GetFotosRequest)) as GetFotosRequest;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static GetFotosRequest create() => GetFotosRequest._();
  GetFotosRequest createEmptyInstance() => create();
  static $pb.PbList<GetFotosRequest> createRepeated() => $pb.PbList<GetFotosRequest>();
  @$core.pragma('dart2js:noInline')
  static GetFotosRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<GetFotosRequest>(create);
  static GetFotosRequest? _defaultInstance;
}

class Redes extends $pb.GeneratedMessage {
  factory Redes({
    $core.Iterable<$core.String>? red,
  }) {
    final $result = create();
    if (red != null) {
      $result.red.addAll(red);
    }
    return $result;
  }
  Redes._() : super();
  factory Redes.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Redes.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'Redes', package: const $pb.PackageName(_omitMessageNames ? '' : 'pb'), createEmptyInstance: create)
    ..pPS(1, _omitFieldNames ? '' : 'red')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Redes clone() => Redes()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Redes copyWith(void Function(Redes) updates) => super.copyWith((message) => updates(message as Redes)) as Redes;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static Redes create() => Redes._();
  Redes createEmptyInstance() => create();
  static $pb.PbList<Redes> createRepeated() => $pb.PbList<Redes>();
  @$core.pragma('dart2js:noInline')
  static Redes getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Redes>(create);
  static Redes? _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<$core.String> get red => $_getList(0);
}

class FotoResponse extends $pb.GeneratedMessage {
  factory FotoResponse({
    $core.int? id,
    $core.String? fotoUrl,
    $core.String? nome,
    $core.String? partido,
    $core.String? uf,
    $core.String? url,
    Redes? redes,
    $core.String? fotoUrl2,
  }) {
    final $result = create();
    if (id != null) {
      $result.id = id;
    }
    if (fotoUrl != null) {
      $result.fotoUrl = fotoUrl;
    }
    if (nome != null) {
      $result.nome = nome;
    }
    if (partido != null) {
      $result.partido = partido;
    }
    if (uf != null) {
      $result.uf = uf;
    }
    if (url != null) {
      $result.url = url;
    }
    if (redes != null) {
      $result.redes = redes;
    }
    if (fotoUrl2 != null) {
      $result.fotoUrl2 = fotoUrl2;
    }
    return $result;
  }
  FotoResponse._() : super();
  factory FotoResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory FotoResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'FotoResponse', package: const $pb.PackageName(_omitMessageNames ? '' : 'pb'), createEmptyInstance: create)
    ..a<$core.int>(1, _omitFieldNames ? '' : 'id', $pb.PbFieldType.O3)
    ..aOS(2, _omitFieldNames ? '' : 'fotoUrl', protoName: 'fotoUrl')
    ..aOS(3, _omitFieldNames ? '' : 'nome')
    ..aOS(4, _omitFieldNames ? '' : 'partido')
    ..aOS(5, _omitFieldNames ? '' : 'uf')
    ..aOS(6, _omitFieldNames ? '' : 'url')
    ..aOM<Redes>(7, _omitFieldNames ? '' : 'redes', subBuilder: Redes.create)
    ..aOS(8, _omitFieldNames ? '' : 'fotoUrl2', protoName: 'fotoUrl2')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  FotoResponse clone() => FotoResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  FotoResponse copyWith(void Function(FotoResponse) updates) => super.copyWith((message) => updates(message as FotoResponse)) as FotoResponse;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static FotoResponse create() => FotoResponse._();
  FotoResponse createEmptyInstance() => create();
  static $pb.PbList<FotoResponse> createRepeated() => $pb.PbList<FotoResponse>();
  @$core.pragma('dart2js:noInline')
  static FotoResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<FotoResponse>(create);
  static FotoResponse? _defaultInstance;

  @$pb.TagNumber(1)
  $core.int get id => $_getIZ(0);
  @$pb.TagNumber(1)
  set id($core.int v) { $_setSignedInt32(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get fotoUrl => $_getSZ(1);
  @$pb.TagNumber(2)
  set fotoUrl($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasFotoUrl() => $_has(1);
  @$pb.TagNumber(2)
  void clearFotoUrl() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get nome => $_getSZ(2);
  @$pb.TagNumber(3)
  set nome($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasNome() => $_has(2);
  @$pb.TagNumber(3)
  void clearNome() => clearField(3);

  @$pb.TagNumber(4)
  $core.String get partido => $_getSZ(3);
  @$pb.TagNumber(4)
  set partido($core.String v) { $_setString(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasPartido() => $_has(3);
  @$pb.TagNumber(4)
  void clearPartido() => clearField(4);

  @$pb.TagNumber(5)
  $core.String get uf => $_getSZ(4);
  @$pb.TagNumber(5)
  set uf($core.String v) { $_setString(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasUf() => $_has(4);
  @$pb.TagNumber(5)
  void clearUf() => clearField(5);

  @$pb.TagNumber(6)
  $core.String get url => $_getSZ(5);
  @$pb.TagNumber(6)
  set url($core.String v) { $_setString(5, v); }
  @$pb.TagNumber(6)
  $core.bool hasUrl() => $_has(5);
  @$pb.TagNumber(6)
  void clearUrl() => clearField(6);

  @$pb.TagNumber(7)
  Redes get redes => $_getN(6);
  @$pb.TagNumber(7)
  set redes(Redes v) { setField(7, v); }
  @$pb.TagNumber(7)
  $core.bool hasRedes() => $_has(6);
  @$pb.TagNumber(7)
  void clearRedes() => clearField(7);
  @$pb.TagNumber(7)
  Redes ensureRedes() => $_ensure(6);

  @$pb.TagNumber(8)
  $core.String get fotoUrl2 => $_getSZ(7);
  @$pb.TagNumber(8)
  set fotoUrl2($core.String v) { $_setString(7, v); }
  @$pb.TagNumber(8)
  $core.bool hasFotoUrl2() => $_has(7);
  @$pb.TagNumber(8)
  void clearFotoUrl2() => clearField(8);
}


const _omitFieldNames = $core.bool.fromEnvironment('protobuf.omit_field_names');
const _omitMessageNames = $core.bool.fromEnvironment('protobuf.omit_message_names');
