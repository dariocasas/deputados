import 'package:deputados/core/store/db_status_store.dart';
import 'package:flutter/material.dart';
import 'package:flutter_modular/flutter_modular.dart';
import 'package:rx_notifier/rx_notifier.dart';

import '../widget/check_close.dart';
import '../widget/custom_app_bar.dart';
import 'create_db_board.dart';
import '../../core/service/database_service.dart';
import 'main_drawer.dart';

class CreateDbPage extends StatefulWidget {
  const CreateDbPage({super.key});

  @override
  State<CreateDbPage> createState() => _CreateDbPageState();
}

class _CreateDbPageState extends State<CreateDbPage> {
  final dbStatusStore = Modular.get<DbStatusStore>();
  final grpcService = Modular.get<DatabaseService>();
  double _currentSliderValue = 5;
  final TextEditingController _timeOutController =
      TextEditingController(text: "5000");

  @override
  Widget build(BuildContext context) {
    final theme = Theme.of(context);
    final indexCount = context.select(() => dbStatusStore.indexCount);

    return CheckClose(
      child: Scaffold(
        drawer: const MainDrawer(),
        appBar: const CustomAppBar(title: 'Deputados. Criar banco de dados'),
        body: Center(
          child: SizedBox(
            width: 320,
            child: Column(
              mainAxisAlignment: MainAxisAlignment.start,
              crossAxisAlignment: CrossAxisAlignment.start,
              children: <Widget>[
                const SizedBox(
                  height: 100,
                ),
                Row(
                  mainAxisAlignment: MainAxisAlignment.center,
                  children: [
                    Text(
                      'Deputados $indexCount',
                      style: theme.textTheme.bodyLarge,
                    ),
                  ],
                ),
                const SizedBox(
                  height: 50,
                ),
                Row(
                  mainAxisAlignment: MainAxisAlignment.start,
                  mainAxisSize: MainAxisSize.min,
                  children: [
                    const Padding(
                      padding: EdgeInsets.only(right: 16),
                      child: Text('concorrência:'),
                    ),
                    Text(_currentSliderValue.round().toString()),
                    Slider(
                      value: _currentSliderValue,
                      max: 64,
                      min: 1,
                      divisions: 64,
                      label: _currentSliderValue.round().toString(),
                      onChanged: (double value) {
                        setState(() {
                          _currentSliderValue = value;
                        });
                      },
                    ),
                  ],
                ),
                Row(
                  mainAxisAlignment: MainAxisAlignment.start,
                  mainAxisSize: MainAxisSize.min,
                  children: [
                    const Padding(
                      padding: EdgeInsets.only(right: 16),
                      child: Text('timeout (milisegundos):'),
                    ),
                    SizedBox(
                      width: 100,
                      child: TextField(
                        controller: _timeOutController,
                      ),
                    )
                  ],
                ),
                const SizedBox(
                  height: 20,
                ),
                const Divider(),
                const SizedBox(
                  height: 20,
                ),
                Row(
                  mainAxisAlignment: MainAxisAlignment.center,
                  children: [
                    ElevatedButton(
                      onPressed: () {
                        var value = int.tryParse(_timeOutController.text);
                        if (value != null) {
                          grpcService.createDB(
                            concurrency: _currentSliderValue.round(),
                            timeout: value,
                          );
                          showDialog(
                            context: context,
                            builder: (BuildContext context) =>
                                const CreateDBBoard(),
                          );
                        }
                      },
                      child: Text(
                        'Criar e povoar DB',
                        style: theme.textTheme.bodyLarge,
                      ),
                    ),
                  ],
                ),
              ],
            ),
          ),
        ),
      ),
    );
  }
}
