// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'db_status_store.dart';

// **************************************************************************
// RxGenerator
// **************************************************************************

class DbStatusStore = _DbStatusStore with _DbStatusStoreMixin;

mixin _DbStatusStoreMixin on _DbStatusStore {
  ///
  /// GENERATED dbStatus(DbStatus)
  ///

  late final _dbStatusRx = RxNotifier<DbStatus>(super.dbStatus);
  RxValueListenable<DbStatus> get dbStatusListenable => _dbStatusRx;

  @override
  set dbStatus(DbStatus _dbStatusValue) => _dbStatusRx.value = _dbStatusValue;
  @override
  DbStatus get dbStatus => _dbStatusRx.value;

  ///
  /// GENERATED serverStatus(ServerStatus)
  ///

  late final _serverStatusRx = RxNotifier<ServerStatus>(super.serverStatus);
  RxValueListenable<ServerStatus> get serverStatusListenable => _serverStatusRx;

  @override
  set serverStatus(ServerStatus _serverStatusValue) =>
      _serverStatusRx.value = _serverStatusValue;
  @override
  ServerStatus get serverStatus => _serverStatusRx.value;

  ///
  /// GENERATED indexCount(int)
  ///

  late final _indexCountRx = RxNotifier<int>(super.indexCount);
  RxValueListenable<int> get indexCountListenable => _indexCountRx;

  @override
  set indexCount(int _indexCountValue) =>
      _indexCountRx.value = _indexCountValue;
  @override
  int get indexCount => _indexCountRx.value;

  ///
  /// GENERATED loading(bool)
  ///

  late final _loadingRx = RxNotifier<bool>(super.loading);
  RxValueListenable<bool> get loadingListenable => _loadingRx;

  @override
  set loading(bool _loadingValue) => _loadingRx.value = _loadingValue;
  @override
  bool get loading => _loadingRx.value;

  ///
  /// GENERATED ready(bool)
  ///

  late final _readyRx = RxNotifier<bool>(super.ready);
  RxValueListenable<bool> get readyListenable => _readyRx;

  @override
  set ready(bool _readyValue) => _readyRx.value = _readyValue;
  @override
  bool get ready => _readyRx.value;

  ///
  /// GENERATED askClose(bool)
  ///

  late final _askCloseRx = RxNotifier<bool>(super.askClose);
  RxValueListenable<bool> get askCloseListenable => _askCloseRx;

  @override
  set askClose(bool _askCloseValue) => _askCloseRx.value = _askCloseValue;
  @override
  bool get askClose => _askCloseRx.value;
}
