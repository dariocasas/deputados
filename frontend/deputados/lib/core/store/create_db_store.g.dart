// GENERATED CODE - DO NOT MODIFY BY HAND

// ignore_for_file: library_private_types_in_public_api, no_leading_underscores_for_local_identifiers

part of 'create_db_store.dart';

// **************************************************************************
// RxGenerator
// **************************************************************************

class CreateDbStore = _CreateDbStore with _CreateDbStoreMixin;

mixin _CreateDbStoreMixin on _CreateDbStore {
  ///
  /// GENERATED deputados(RxList<CreateDbData>)
  ///

  late final _deputadosRx = RxNotifier<RxList<CreateDbData>>(super.deputados);
  RxValueListenable<RxList<CreateDbData>> get deputadosListenable =>
      _deputadosRx;

  @override
  set deputados(RxList<CreateDbData> _deputadosValue) =>
      _deputadosRx.value = _deputadosValue;
  @override
  RxList<CreateDbData> get deputados => _deputadosRx.value;

  ///
  /// GENERATED totalMilliseconds(int)
  ///

  late final _totalMillisecondsRx = RxNotifier<int>(super.totalMilliseconds);
  RxValueListenable<int> get totalMillisecondsListenable =>
      _totalMillisecondsRx;

  @override
  set totalMilliseconds(int _totalMillisecondsValue) =>
      _totalMillisecondsRx.value = _totalMillisecondsValue;
  @override
  int get totalMilliseconds => _totalMillisecondsRx.value;

  ///
  /// GENERATED concurrency(int)
  ///

  late final _concurrencyRx = RxNotifier<int>(super.concurrency);
  RxValueListenable<int> get concurrencyListenable => _concurrencyRx;

  @override
  set concurrency(int _concurrencyValue) =>
      _concurrencyRx.value = _concurrencyValue;
  @override
  int get concurrency => _concurrencyRx.value;

  ///
  /// GENERATED totalPartialTime(PartialTime?)
  ///

  late final _totalPartialTimeRx =
      RxNotifier<PartialTime?>(super.totalPartialTime);
  RxValueListenable<PartialTime?> get totalPartialTimeListenable =>
      _totalPartialTimeRx;

  @override
  set totalPartialTime(PartialTime? _totalPartialTimeValue) =>
      _totalPartialTimeRx.value = _totalPartialTimeValue;
  @override
  PartialTime? get totalPartialTime => _totalPartialTimeRx.value;
}
