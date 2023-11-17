// ignore_for_file: public_member_api_docs, sort_constructors_first
// ignore_for_file: avoid_init_to_null

import 'package:rx_notifier/rx_notifier.dart';

part 'create_db_store.g.dart';

class PartialTime {
  int getapi = 0;
  int dbread = 0;
  int getpage = 0;
  int scanpage = 0;
  int dbwrite = 0;
  PartialTime({
    required this.getapi,
    required this.dbread,
    required this.getpage,
    required this.scanpage,
    required this.dbwrite,
  });
}

class CreateDbData {
  int id = 0;
  String name = '';
  bool error = false;
  String errorString = '';
  int milliseconds = 0;
  PartialTime? partialTime = null;
}

@RxStore()
abstract class _CreateDbStore {
  @RxValue()
  RxList<CreateDbData> deputados = RxList<CreateDbData>([]);
  @RxValue()
  int totalMilliseconds = 0;
  @RxValue()
  PartialTime? totalPartialTime = null;
}
