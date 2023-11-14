// ignore_for_file: avoid_init_to_null

import 'package:rx_notifier/rx_notifier.dart';
import '../db_status_enum.dart';

part 'db_status_store.g.dart';

@RxStore()
abstract class _DbStatusStore {
  @RxValue()
  DbStatus dbStatus = DbStatus.uninitializated;
  @RxValue()
  int indexCount = 0;
  @RxValue()
  bool loading = false;
  @RxValue()
  bool ready = false;
  @RxValue()
  bool askClose = false;
}
