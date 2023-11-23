// ignore_for_file: lines_longer_than_80_chars
import 'package:flutter/material.dart';
import 'package:flutter_modular/flutter_modular.dart';
import 'package:rx_notifier/rx_notifier.dart';

import 'package:deputados/core/store/create_db_store.dart';
import 'package:deputados/core/store/db_status_store.dart';

import '../widget/custom_app_bar.dart';
import '../widget/check_close.dart';
import 'create_db_finished.dart';
import 'create_db_item.dart';
import 'create_db_working.dart';

class CreateDBBoard extends StatefulWidget {
  const CreateDBBoard({super.key});

  @override
  State<CreateDBBoard> createState() => _CreateDBBoardState();
}

class _CreateDBBoardState extends State<CreateDBBoard> {
  final createDbStore = Modular.get<CreateDbStore>();
  final dbStatusStore = Modular.get<DbStatusStore>();
  final startTime = DateTime.now();

  @override
  void initState() {
    dbStatusStore.ready = false;
    super.initState();
  }

  @override
  Widget build(BuildContext context) {
    final concurrency = context.select(() => createDbStore.concurrency);
    final deputados =
        context.select(() => createDbStore.deputados.reversed.toList());
    final dbStatusItems = -context.select(() => dbStatusStore.indexCount);
    final ready = context.select(() => dbStatusStore.ready);

    final deltaTime = DateTime.now().difference(startTime);
    final seconds = deltaTime.inSeconds.toString();
    final avg = dbStatusItems / deltaTime.inSeconds;
    final sAvg = avg.toStringAsFixed(avg.truncateToDouble() == avg ? 0 : 3);

    if (deputados.isEmpty) {
      return Container();
    } else {
      return Scaffold(
        appBar: CustomAppBar(
            title:
                "Criando db: $dbStatusItems deputados  | tempo: $seconds seg  | média: $sAvg deputados/segundo | concorrência: $concurrency"),
        body: CheckClose(
          child: Column(
            children: [
              Expanded(
                child: ListView.builder(
                  itemExtent: 30,
                  itemCount: deputados.length,
                  itemBuilder: (context, index) {
                    return ListTile(
                      title: CreateDbItem(deputado: deputados[index]),
                    );
                  },
                ),
              ),
              if (ready == true)
                const CreateDbFinished()
              else
                const CreateDbWorking()
            ],
          ),
        ),
      );
    }
  }
}
