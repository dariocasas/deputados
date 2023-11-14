// GENERATED CODE - DO NOT MODIFY BY HAND

// ignore_for_file: library_private_types_in_public_api, no_leading_underscores_for_local_identifiers

part of 'fotos_store.dart';

// **************************************************************************
// RxGenerator
// **************************************************************************

class FotosStore = _FotosStore with _FotosStoreMixin;

mixin _FotosStoreMixin on _FotosStore {
  ///
  /// GENERATED fotos(RxList<InfoFoto>)
  ///

  late final _fotosRx = RxNotifier<RxList<InfoFoto>>(super.fotos);
  RxValueListenable<RxList<InfoFoto>> get fotosListenable => _fotosRx;

  @override
  set fotos(RxList<InfoFoto> _fotosValue) => _fotosRx.value = _fotosValue;
  @override
  RxList<InfoFoto> get fotos => _fotosRx.value;
}
