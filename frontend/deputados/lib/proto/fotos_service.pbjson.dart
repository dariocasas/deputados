//
//  Generated code. Do not modify.
//  source: proto/fotos_service.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:convert' as $convert;
import 'dart:core' as $core;
import 'dart:typed_data' as $typed_data;

@$core.Deprecated('Use getFotosRequestDescriptor instead')
const GetFotosRequest$json = {
  '1': 'GetFotosRequest',
};

/// Descriptor for `GetFotosRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List getFotosRequestDescriptor = $convert.base64Decode(
    'Cg9HZXRGb3Rvc1JlcXVlc3Q=');

@$core.Deprecated('Use redesDescriptor instead')
const Redes$json = {
  '1': 'Redes',
  '2': [
    {'1': 'red', '3': 1, '4': 3, '5': 9, '10': 'red'},
  ],
};

/// Descriptor for `Redes`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List redesDescriptor = $convert.base64Decode(
    'CgVSZWRlcxIQCgNyZWQYASADKAlSA3JlZA==');

@$core.Deprecated('Use fotoResponseDescriptor instead')
const FotoResponse$json = {
  '1': 'FotoResponse',
  '2': [
    {'1': 'id', '3': 1, '4': 1, '5': 5, '10': 'id'},
    {'1': 'fotoUrl', '3': 2, '4': 1, '5': 9, '10': 'fotoUrl'},
    {'1': 'nome', '3': 3, '4': 1, '5': 9, '10': 'nome'},
    {'1': 'partido', '3': 4, '4': 1, '5': 9, '10': 'partido'},
    {'1': 'uf', '3': 5, '4': 1, '5': 9, '10': 'uf'},
    {'1': 'url', '3': 6, '4': 1, '5': 9, '10': 'url'},
    {'1': 'redes', '3': 7, '4': 1, '5': 11, '6': '.pb.Redes', '10': 'redes'},
    {'1': 'fotoUrl2', '3': 8, '4': 1, '5': 9, '10': 'fotoUrl2'},
  ],
};

/// Descriptor for `FotoResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List fotoResponseDescriptor = $convert.base64Decode(
    'CgxGb3RvUmVzcG9uc2USDgoCaWQYASABKAVSAmlkEhgKB2ZvdG9VcmwYAiABKAlSB2ZvdG9Vcm'
    'wSEgoEbm9tZRgDIAEoCVIEbm9tZRIYCgdwYXJ0aWRvGAQgASgJUgdwYXJ0aWRvEg4KAnVmGAUg'
    'ASgJUgJ1ZhIQCgN1cmwYBiABKAlSA3VybBIfCgVyZWRlcxgHIAEoCzIJLnBiLlJlZGVzUgVyZW'
    'RlcxIaCghmb3RvVXJsMhgIIAEoCVIIZm90b1VybDI=');

