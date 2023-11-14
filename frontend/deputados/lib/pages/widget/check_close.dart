import 'package:flutter/material.dart';
import 'package:flutter_modular/flutter_modular.dart';
import 'package:rx_notifier/rx_notifier.dart';

import '../../core/store/db_status_store.dart';
import 'close_dialog.dart';

class CheckClose extends StatefulWidget {
  final Widget child;
  const CheckClose({super.key, required this.child});

  @override
  State<CheckClose> createState() => _CheckCloseState();
}

class _CheckCloseState extends State<CheckClose> {
  final dbStatusStore = Modular.get<DbStatusStore>();

  @override
  void initState() {
    Future.delayed(
        const Duration(milliseconds: 10), () => dbStatusStore.askClose = false);
    super.initState();
  }

  @override
  Widget build(BuildContext context) {
    final askClose = context.select(() => dbStatusStore.askClose);

    if (askClose == true) {
      Future.delayed(
          const Duration(milliseconds: 10), () => closeDialog(context));
    }

    return widget.child;
  }
}
