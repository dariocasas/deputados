import 'package:deputados/pages/widget/check_close.dart';
import 'package:flutter/material.dart';
import 'package:flutter_modular/flutter_modular.dart';

import '../core/service/database_service.dart';
import 'home/main_drawer.dart';
import 'widget/custom_app_bar.dart';

class HomePage extends StatefulWidget {
  const HomePage({super.key});

  @override
  State<HomePage> createState() => _HomePageState();
}

class _HomePageState extends State<HomePage> {
  final grpcService = Modular.get<DatabaseService>();

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
        appBar: CustomAppBar(
          title: "",
        ),
        body: Center(
          child: CircularProgressIndicator(),
        ),
      ),
    );
  }
}
