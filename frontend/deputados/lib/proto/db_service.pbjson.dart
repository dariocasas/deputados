//
//  Generated code. Do not modify.
//  source: proto/db_service.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:convert' as $convert;
import 'dart:core' as $core;
import 'dart:typed_data' as $typed_data;

@$core.Deprecated('Use dbStatusEnumDescriptor instead')
const DbStatusEnum$json = {
  '1': 'DbStatusEnum',
  '2': [
    {'1': 'UNINITIALIZATED', '2': 0},
    {'1': 'INDEX_CREATED', '2': 1},
    {'1': 'DB_CREATED', '2': 2},
  ],
};

/// Descriptor for `DbStatusEnum`. Decode as a `google.protobuf.EnumDescriptorProto`.
final $typed_data.Uint8List dbStatusEnumDescriptor = $convert.base64Decode(
    'CgxEYlN0YXR1c0VudW0SEwoPVU5JTklUSUFMSVpBVEVEEAASEQoNSU5ERVhfQ1JFQVRFRBABEg'
    '4KCkRCX0NSRUFURUQQAg==');

@$core.Deprecated('Use dropDbRequestDescriptor instead')
const DropDbRequest$json = {
  '1': 'DropDbRequest',
};

/// Descriptor for `DropDbRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List dropDbRequestDescriptor = $convert.base64Decode(
    'Cg1Ecm9wRGJSZXF1ZXN0');

@$core.Deprecated('Use dropDbResponseDescriptor instead')
const DropDbResponse$json = {
  '1': 'DropDbResponse',
};

/// Descriptor for `DropDbResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List dropDbResponseDescriptor = $convert.base64Decode(
    'Cg5Ecm9wRGJSZXNwb25zZQ==');

@$core.Deprecated('Use dbStatusRequestDescriptor instead')
const DbStatusRequest$json = {
  '1': 'DbStatusRequest',
};

/// Descriptor for `DbStatusRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List dbStatusRequestDescriptor = $convert.base64Decode(
    'Cg9EYlN0YXR1c1JlcXVlc3Q=');

@$core.Deprecated('Use dbStatusResponseDescriptor instead')
const DbStatusResponse$json = {
  '1': 'DbStatusResponse',
  '2': [
    {'1': 'status', '3': 1, '4': 1, '5': 14, '6': '.pb.DbStatusEnum', '10': 'status'},
    {'1': 'indexCount', '3': 2, '4': 1, '5': 5, '10': 'indexCount'},
  ],
};

/// Descriptor for `DbStatusResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List dbStatusResponseDescriptor = $convert.base64Decode(
    'ChBEYlN0YXR1c1Jlc3BvbnNlEigKBnN0YXR1cxgBIAEoDjIQLnBiLkRiU3RhdHVzRW51bVIGc3'
    'RhdHVzEh4KCmluZGV4Q291bnQYAiABKAVSCmluZGV4Q291bnQ=');

@$core.Deprecated('Use populateIndexRequestDescriptor instead')
const PopulateIndexRequest$json = {
  '1': 'PopulateIndexRequest',
  '2': [
    {'1': 'maxItems', '3': 1, '4': 1, '5': 5, '10': 'maxItems'},
  ],
};

/// Descriptor for `PopulateIndexRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List populateIndexRequestDescriptor = $convert.base64Decode(
    'ChRQb3B1bGF0ZUluZGV4UmVxdWVzdBIaCghtYXhJdGVtcxgBIAEoBVIIbWF4SXRlbXM=');

@$core.Deprecated('Use populateIndexResponseDescriptor instead')
const PopulateIndexResponse$json = {
  '1': 'PopulateIndexResponse',
  '2': [
    {'1': 'indexCount', '3': 1, '4': 1, '5': 5, '10': 'indexCount'},
  ],
};

/// Descriptor for `PopulateIndexResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List populateIndexResponseDescriptor = $convert.base64Decode(
    'ChVQb3B1bGF0ZUluZGV4UmVzcG9uc2USHgoKaW5kZXhDb3VudBgBIAEoBVIKaW5kZXhDb3VudA'
    '==');

@$core.Deprecated('Use populateDbRequestDescriptor instead')
const PopulateDbRequest$json = {
  '1': 'PopulateDbRequest',
  '2': [
    {'1': 'concurrency', '3': 1, '4': 1, '5': 5, '10': 'concurrency'},
    {'1': 'timeout', '3': 2, '4': 1, '5': 5, '10': 'timeout'},
  ],
};

/// Descriptor for `PopulateDbRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List populateDbRequestDescriptor = $convert.base64Decode(
    'ChFQb3B1bGF0ZURiUmVxdWVzdBIgCgtjb25jdXJyZW5jeRgBIAEoBVILY29uY3VycmVuY3kSGA'
    'oHdGltZW91dBgCIAEoBVIHdGltZW91dA==');

@$core.Deprecated('Use partialTimeDescriptor instead')
const PartialTime$json = {
  '1': 'PartialTime',
  '2': [
    {'1': 'dbread', '3': 1, '4': 1, '5': 5, '10': 'dbread'},
    {'1': 'getapi', '3': 2, '4': 1, '5': 5, '10': 'getapi'},
    {'1': 'getpage', '3': 3, '4': 1, '5': 5, '10': 'getpage'},
    {'1': 'scanpage', '3': 4, '4': 1, '5': 5, '10': 'scanpage'},
    {'1': 'dbwrite', '3': 5, '4': 1, '5': 5, '10': 'dbwrite'},
  ],
};

/// Descriptor for `PartialTime`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List partialTimeDescriptor = $convert.base64Decode(
    'CgtQYXJ0aWFsVGltZRIWCgZkYnJlYWQYASABKAVSBmRicmVhZBIWCgZnZXRhcGkYAiABKAVSBm'
    'dldGFwaRIYCgdnZXRwYWdlGAMgASgFUgdnZXRwYWdlEhoKCHNjYW5wYWdlGAQgASgFUghzY2Fu'
    'cGFnZRIYCgdkYndyaXRlGAUgASgFUgdkYndyaXRl');

@$core.Deprecated('Use deputadoResponseDescriptor instead')
const DeputadoResponse$json = {
  '1': 'DeputadoResponse',
  '2': [
    {'1': 'error', '3': 1, '4': 1, '5': 8, '10': 'error'},
    {'1': 'errorString', '3': 2, '4': 1, '5': 9, '10': 'errorString'},
    {'1': 'id', '3': 3, '4': 1, '5': 5, '10': 'id'},
    {'1': 'name', '3': 4, '4': 1, '5': 9, '10': 'name'},
    {'1': 'milliseconds', '3': 5, '4': 1, '5': 5, '10': 'milliseconds'},
    {'1': 'partialTime', '3': 6, '4': 1, '5': 11, '6': '.pb.PartialTime', '10': 'partialTime'},
  ],
};

/// Descriptor for `DeputadoResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List deputadoResponseDescriptor = $convert.base64Decode(
    'ChBEZXB1dGFkb1Jlc3BvbnNlEhQKBWVycm9yGAEgASgIUgVlcnJvchIgCgtlcnJvclN0cmluZx'
    'gCIAEoCVILZXJyb3JTdHJpbmcSDgoCaWQYAyABKAVSAmlkEhIKBG5hbWUYBCABKAlSBG5hbWUS'
    'IgoMbWlsbGlzZWNvbmRzGAUgASgFUgxtaWxsaXNlY29uZHMSMQoLcGFydGlhbFRpbWUYBiABKA'
    'syDy5wYi5QYXJ0aWFsVGltZVILcGFydGlhbFRpbWU=');

