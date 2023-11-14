import 'package:deputados/core/service/database_service.dart';
import 'package:flutter/material.dart';
import 'package:flutter_modular/flutter_modular.dart';
import 'package:window_manager/window_manager.dart';

import 'core/service/fotos_service.dart';
import 'core/service/grpc_service.dart';
import 'core/service/server_service.dart';
import 'core/store/db_status_store.dart';

class AppWidget extends StatefulWidget {
  const AppWidget({super.key});

  @override
  State<AppWidget> createState() => _AppWidgetState();
}

class _AppWidgetState extends State<AppWidget> with WindowListener {
  final grpcService = Modular.get<GrpcService>();
  final databaseService = Modular.get<DatabaseService>();
  final fotosService = Modular.get<FotosService>();
  final serverService = Modular.get<ServerService>();
  final dbStatusStore = Modular.get<DbStatusStore>();

  @override
  void initState() {
    super.initState();

    grpcService.init();
    databaseService.init();
    fotosService.init();
    serverService.init();

    windowManager.addListener(this);
    _init();
  }

  @override
  void dispose() {
    windowManager.removeListener(this);
    super.dispose();
  }

  void _init() async {
    await windowManager.setPreventClose(true);
    setState(() {});
  }

  @override
  void onWindowClose() async {
    bool isPreventClose = await windowManager.isPreventClose();
    if (isPreventClose) {
      dbStatusStore.askClose = true;
    }
  }

  @override
  Widget build(BuildContext context) {
    // Modular.setInitialRoute('/waitconn/');

    return MaterialApp.router(
      title: 'Deputados',
      theme: ThemeData(primarySwatch: Colors.blue),
      routerConfig: Modular.routerConfig,
    ); //added by extension
  }
}
