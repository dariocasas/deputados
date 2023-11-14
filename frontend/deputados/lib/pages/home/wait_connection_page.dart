// ignore_for_file: use_build_context_synchronously

import 'package:flutter/material.dart';
import 'package:flutter_modular/flutter_modular.dart';

import 'package:deputados/core/store/db_status_store.dart';

import '../../core/service/database_service.dart';
import '../widget/check_close.dart';
import '../widget/custom_app_bar.dart';
import 'main_drawer.dart';

class WaitConnectionPage extends StatefulWidget {
  const WaitConnectionPage({super.key});

  @override
  State<WaitConnectionPage> createState() => _WaitConnectionPageState();
}

class _WaitConnectionPageState extends State<WaitConnectionPage> {
  final grpcService = Modular.get<DatabaseService>();
  final dbStatusStore = Modular.get<DbStatusStore>();

  @override
  void initState() {
    grpcService.dbStatus();
    super.initState();
  }

  @override
  Widget build(BuildContext context) {
    return const CheckClose(
      child: Scaffold(
        drawer: MainDrawer(),
        appBar: CustomAppBar(title: 'Conectando...'),
        body: Center(
          child: CircularProgressIndicator(),
        ),
      ),
    );
  }
}
