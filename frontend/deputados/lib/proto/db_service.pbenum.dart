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

class DbStatusEnum extends $pb.ProtobufEnum {
  static const DbStatusEnum UNINITIALIZATED = DbStatusEnum._(0, _omitEnumNames ? '' : 'UNINITIALIZATED');
  static const DbStatusEnum INDEX_CREATED = DbStatusEnum._(1, _omitEnumNames ? '' : 'INDEX_CREATED');
  static const DbStatusEnum DB_CREATED = DbStatusEnum._(2, _omitEnumNames ? '' : 'DB_CREATED');

  static const $core.List<DbStatusEnum> values = <DbStatusEnum> [
    UNINITIALIZATED,
    INDEX_CREATED,
    DB_CREATED,
  ];

  static final $core.Map<$core.int, DbStatusEnum> _byValue = $pb.ProtobufEnum.initByValue(values);
  static DbStatusEnum? valueOf($core.int value) => _byValue[value];

  const DbStatusEnum._($core.int v, $core.String n) : super(v, n);
}


const _omitEnumNames = $core.bool.fromEnvironment('protobuf.omit_enum_names');
