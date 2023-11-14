import 'package:flutter/material.dart';
import 'package:flutter_modular/flutter_modular.dart';

import '../../core/service/server_service.dart';
import '../../core/store/db_status_store.dart';
import 'confirm_dialog.dart';

void closeDialog(BuildContext context) {
  final serverService = Modular.get<ServerService>();
  final dbStatusStore = Modular.get<DbStatusStore>();
  dbStatusStore.askClose = false;

  showDialog(
    context: context,
    builder: (BuildContext context) => ConfirmDialog(
      onConfirmed: () {
        serverService.shutdown();
        return;
      },
      onDenied: () {
        Modular.to.pop();
        return;
      },
      title: 'Fechar Deputados?',
    ),
  );
}
