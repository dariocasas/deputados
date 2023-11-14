// ignore_for_file: lines_longer_than_80_chars

import 'package:flutter/material.dart';
import 'package:flutter_modular/flutter_modular.dart';
import 'package:rx_notifier/rx_notifier.dart';

import 'package:deputados/core/store/create_db_store.dart';
import 'package:deputados/core/store/db_status_store.dart';

class CreateDbFinished extends StatefulWidget {
  const CreateDbFinished({super.key});

  @override
  State<CreateDbFinished> createState() => _CreateDbFinishedState();
}

class _CreateDbFinishedState extends State<CreateDbFinished> {
  final createDbStore = Modular.get<CreateDbStore>();
  final dbStatusStore = Modular.get<DbStatusStore>();

  @override
  Widget build(BuildContext context) {
    final theme = Theme.of(context);

    final totalPartialTime = context.select(
      () => createDbStore.totalPartialTime,
    );

    final totalms = totalPartialTime!.dbread +
        totalPartialTime.getapi +
        totalPartialTime.dbread +
        totalPartialTime.getpage +
        totalPartialTime.scanpage +
        totalPartialTime.dbwrite;

    final milliseconds = context.select(
      () => createDbStore.totalMilliseconds,
    );
    final qDeputados = -context.select(
      () => dbStatusStore.indexCount,
    );

    return SizedBox(
      height: 130,
      child: ConstrainedBox(
        constraints:
            BoxConstraints(maxWidth: MediaQuery.of(context).size.width),
        child: Column(
          children: [
            const Divider(color: Colors.black),
            Padding(
              padding: const EdgeInsets.only(left: 16, top: 8),
              child: Column(
                children: [
                  Row(
                    //spacing: 10,
                    children: [
                      SizedBox(
                        width: 160,
                        child: Text('Total deputados: $qDeputados'),
                      ),
                      SizedBox(
                        width: 160,
                        child: Text('Total tempo: ${formatMs(milliseconds)}'),
                      ),
                      SizedBox(
                        width: 130,
                        child: Text('total = ${formatMs(totalms)}',
                            style: theme.textTheme.bodyMedium),
                      ),
                      SizedBox(
                        width: 130,
                        child: Text(
                            'dbread ${formatMs(totalPartialTime.dbread)}',
                            style: theme.textTheme.bodyMedium),
                      ),
                      SizedBox(
                        width: 130,
                        child: Text(
                            'getapi ${formatMs(totalPartialTime.getapi)}',
                            style: theme.textTheme.bodyMedium),
                      ),
                      SizedBox(
                        width: 130,
                        child: Text(
                            'getpage ${formatMs(totalPartialTime.getpage)}',
                            style: theme.textTheme.bodyMedium),
                      ),
                      SizedBox(
                        width: 130,
                        child: Text(
                            'scanpage ${formatMs(totalPartialTime.scanpage)}',
                            style: theme.textTheme.bodyMedium),
                      ),
                      SizedBox(
                        width: 130,
                        child: Text(
                            'dbwrite ${formatMs(totalPartialTime.dbwrite)}',
                            style: theme.textTheme.bodyMedium),
                      ),
                    ],
                  ),
                  const SizedBox(height: 10),
                  Row(
                    //spacing: 10,
                    children: [
                      const SizedBox(
                        width: 160,
                        child: Text('Média'),
                      ),
                      SizedBox(
                        width: 160,
                        child: Text(
                            'tempo: ${formatMs(milliseconds / qDeputados)}'),
                      ),
                      SizedBox(
                        width: 130,
                        child: Text('total = ${formatMs(totalms / qDeputados)}',
                            style: theme.textTheme.bodyMedium),
                      ),
                      SizedBox(
                        width: 130,
                        child: Text(
                            'dbread ${formatMs(totalPartialTime.dbread / qDeputados)}',
                            style: theme.textTheme.bodyMedium),
                      ),
                      SizedBox(
                        width: 130,
                        child: Text(
                            'getapi ${formatMs(totalPartialTime.getapi / qDeputados)}',
                            style: theme.textTheme.bodyMedium),
                      ),
                      SizedBox(
                        width: 130,
                        child: Text(
                            'getpage ${formatMs(totalPartialTime.getpage / qDeputados)}',
                            style: theme.textTheme.bodyMedium),
                      ),
                      SizedBox(
                        width: 130,
                        child: Text(
                            'scanpage ${formatMs(totalPartialTime.scanpage / qDeputados)}',
                            style: theme.textTheme.bodyMedium),
                      ),
                      SizedBox(
                        width: 130,
                        child: Text(
                            'dbwrite ${formatMs(totalPartialTime.dbwrite / qDeputados)}',
                            style: theme.textTheme.bodyMedium),
                      ),
                    ],
                  ),
                  Row(
                    mainAxisAlignment: MainAxisAlignment.end,
                    children: [
                      Padding(
                        padding: const EdgeInsets.only(top: 16, right: 24),
                        child: ElevatedButton(
                          onPressed: () => Modular.to.navigate('/'),
                          child: Text(
                            "fechar",
                            style: theme.textTheme.bodyLarge,
                          ),
                        ),
                      )
                    ],
                  )
                ],
              ),
            ),
          ],
        ),
      ),
    );
  }
}

String formatMs(num num) {
  if (num > 60000) {
    var d = Duration(milliseconds: num.toInt());
    int hour = d.inHours;
    int remainderMinutes = d.inMinutes.remainder(60);
    int remainderSeconds = d.inSeconds.remainder(60);

    return "${addZero(hour)}:${addZero(remainderMinutes)}:${addZero(remainderSeconds)}";
  } else if (num > 10000) {
    return "${(num / 1000).toStringAsFixed(0)} seg";
  } else {
    return "${num.toStringAsFixed(0)} ms";
  }
}

String addZero(int value) {
  return value < 10 ? "0$value" : "$value";
}
