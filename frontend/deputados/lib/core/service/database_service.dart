import 'dart:async';

import 'package:flutter/foundation.dart';
import 'package:flutter_dotenv/flutter_dotenv.dart';
import 'package:flutter_modular/flutter_modular.dart';
import 'package:grpc/grpc.dart';

import 'package:deputados/core/server_status_enum.dart';
import 'package:deputados/core/db_status_enum.dart';
import 'package:deputados/core/store/create_db_store.dart' as createdbstore;
import 'package:deputados/core/store/db_status_store.dart' as dbstatusstore;
import 'package:deputados/proto/db_service.pbgrpc.dart';
import 'package:deputados/proto/health_service.pbgrpc.dart';

import 'grpc_service.dart';

abstract class DatabaseService {
  void init();

  Future<DbStatus> dbStatus();
  Future<bool> createIdx({int maxItems = 0});
  Future<bool> createDB({int concurrency = 5, int timeout = 15000});
  Future<bool> cancelCreateDB();
  Future<bool> dropDB();
}

class DatabaseServiceImpl implements DatabaseService {
  late final ClientChannel channel;
  late final HealthServiceClient healthClient;
  final GrpcService grpcService;

  final dbstatusstore.DbStatusStore dbStatusStore;
  final createdbstore.CreateDbStore createDbStore;

  DatabaseServiceImpl(this.dbStatusStore, this.createDbStore, this.grpcService);

  @override
  void init() {
    debugPrint('DatabaseService.init()');
    channel = grpcService.getClientChannel();

    healthClient = HealthServiceClient(channel);

    final sGrpcHealthCheckSecs = dotenv.env["DEPS_GRPC_HEALTH_CHECK_SECS"];
    int? grpcHealthCheckSecs = int.tryParse(sGrpcHealthCheckSecs ?? "3");
    if (grpcHealthCheckSecs != null) {
      Timer.periodic(Duration(seconds: grpcHealthCheckSecs), (timer) async {
        final request = HealthCheckRequest();
        try {
          final response = await healthClient
              .check(request)
              .timeout(const Duration(seconds: 1));

          dbStatusStore.serverStatus =
              ServerStatus.fromInt(response.status.value);

          debugPrint('grpc server ${response.status.name}');
        } on TimeoutException catch (e) {
          debugPrint('grpc server TimeoutException $e');
          dbStatusStore.serverStatus = ServerStatus.unknown;
        }
      });
    }
  }

  @override
  Future<DbStatus> dbStatus() async {
    final dbServiceClient = DbServiceClient(channel);
    try {
      final response = await dbServiceClient.dbStatus(
        DbStatusRequest(),
        // options: CallOptions(compression: const GzipCodec()),
      );

      dbStatusStore.indexCount = response.indexCount;
      dbStatusStore.dbStatus = DbStatus.fromInt(response.status.value);
      debugPrint('DatabaseService.dbStatus(): ${response.status.toString()}');

      debugPrint(dbStatusStore.dbStatus.route());
      Modular.to.navigate(dbStatusStore.dbStatus.route());

      return DbStatus.fromInt(response.status.value);
    } on GrpcError catch (e) {
      Modular.to.navigate('/error/${e.message}');
      debugPrint('error: ${e.message}');

      Future.delayed(const Duration(seconds: 3), () => dbStatus());

      return Future.value(DbStatus.uninitializated);
    }
  }

  @override
  Future<bool> cancelCreateDB() async {
    final dbServiceClient = DbServiceClient(channel);
    await dbServiceClient.cancelPopulateDb(CancelPopulateDbRequest());
    dbStatusStore.dbStatus = DbStatus.indexCreated;
    dbStatusStore.loading = false;
    debugPrint(dbStatusStore.dbStatus.route());
    Modular.to.navigate(dbStatusStore.dbStatus.route());
    return Future.value(true);
  }

  @override
  Future<bool> createDB({int concurrency = 5, int timeout = 15000}) {
    dbStatusStore.indexCount = 0;
    dbStatusStore.loading = false;
    createDbStore.totalMilliseconds = 0;
    createDbStore.concurrency = concurrency;
    createDbStore.deputados.clear();
    createDbStore.totalPartialTime = createdbstore.PartialTime(
      dbread: 0,
      getapi: 0,
      getpage: 0,
      scanpage: 0,
      dbwrite: 0,
    );

    dbStatusStore.loading = true;

    final dbServiceClient = DbServiceClient(channel);
    final responsStream = dbServiceClient.populateDb(PopulateDbRequest(
      concurrency: concurrency,
      timeout: timeout,
    ));

    responsStream.listen(
      (value) {
        dbStatusStore.indexCount = dbStatusStore.indexCount - 1;

        createDbStore.totalMilliseconds =
            createDbStore.totalMilliseconds + value.milliseconds;

        createDbStore.totalPartialTime = createdbstore.PartialTime(
          dbread:
              createDbStore.totalPartialTime!.dbread + value.partialTime.dbread,
          getapi:
              createDbStore.totalPartialTime!.getapi + value.partialTime.getapi,
          getpage: createDbStore.totalPartialTime!.getpage +
              value.partialTime.getpage,
          scanpage: createDbStore.totalPartialTime!.scanpage +
              value.partialTime.scanpage,
          dbwrite: createDbStore.totalPartialTime!.dbwrite +
              value.partialTime.dbwrite,
        );

        var data = createdbstore.CreateDbData();
        data.id = value.id;
        data.name = value.name;
        data.error = value.error;
        data.errorString = value.errorString;
        data.milliseconds = value.milliseconds;
        data.partialTime = createdbstore.PartialTime(
          dbread: value.partialTime.dbread,
          getapi: value.partialTime.getapi,
          getpage: value.partialTime.getpage,
          scanpage: value.partialTime.scanpage,
          dbwrite: value.partialTime.dbwrite,
        );

        createDbStore.deputados.add(data);

        debugPrint(
            "id: ${value.id} name: ${value.name} milliseconds: ${value.milliseconds} error: ${value.error.toString()} ${value.errorString}");
        debugPrint(
            "dbread: ${value.partialTime.dbread} getapi: ${value.partialTime.getapi} getpage: ${value.partialTime.getpage} scanpage: ${value.partialTime.scanpage} dbwrite ${value.partialTime.dbwrite}");
      },
      onDone: () {
        dbStatusStore.loading = false;
        dbStatusStore.ready = true;
      },
    );
    return Future.value(true);
  }

  @override
  Future<bool> createIdx({int maxItems = 0}) async {
    final dbServiceClient = DbServiceClient(channel);
    dbStatusStore.indexCount = 0;

    dbStatusStore.loading = true;
    final response = await dbServiceClient
        .populateIndex(PopulateIndexRequest(maxItems: maxItems));
    dbStatusStore.indexCount = response.indexCount;
    dbStatusStore.dbStatus = DbStatus.indexCreated;

    dbStatusStore.loading = false;

    debugPrint(dbStatusStore.dbStatus.route());
    Modular.to.navigate(dbStatusStore.dbStatus.route());

    return Future.value(true);
  }

  @override
  Future<bool> dropDB() async {
    final dbServiceClient = DbServiceClient(channel);

    dbStatusStore.loading = true;
    await dbServiceClient.dropDb(DropDbRequest());

    dbStatusStore.dbStatus = DbStatus.uninitializated;
    dbStatusStore.indexCount = 0;
    dbStatusStore.loading = false;

    debugPrint(dbStatusStore.dbStatus.route());
    Modular.to.navigate(dbStatusStore.dbStatus.route());

    return Future.value(true);
  }
}
