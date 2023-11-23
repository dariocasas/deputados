import 'package:flutter/material.dart';
import 'package:flutter_modular/flutter_modular.dart';
import 'package:rx_notifier/rx_notifier.dart';

import 'package:deputados/core/store/db_status_store.dart';

import '../../core/service/database_service.dart';
import '../widget/custom_app_bar.dart';
import '../widget/check_close.dart';
import 'main_drawer.dart';

class CreateIdxPage extends StatefulWidget {
  const CreateIdxPage({super.key});

  @override
  State<CreateIdxPage> createState() => _CreateIdxPageState();
}

class _CreateIdxPageState extends State<CreateIdxPage> {
  final grpcService = Modular.get<DatabaseService>();
  final dbStatusStore = Modular.get<DbStatusStore>();
  final TextEditingController _maxItmesController =
      TextEditingController(text: "0");

  var done = false;

  @override
  Widget build(BuildContext context) {
    final theme = Theme.of(context);
    final loading = context.select(() => dbStatusStore.loading);

    return Scaffold(
      drawer: const MainDrawer(),
      appBar: const CustomAppBar(title: 'Deputados. Criar índice.'),
      body: CheckClose(
        child: Center(
          child: SizedBox(
            width: 500,
            child: Column(
              mainAxisAlignment: MainAxisAlignment.start,
              crossAxisAlignment: CrossAxisAlignment.center,
              children: <Widget>[
                const SizedBox(
                  height: 200,
                ),
                Row(
                  mainAxisAlignment: MainAxisAlignment.start,
                  mainAxisSize: MainAxisSize.min,
                  children: [
                    const Padding(
                      padding: EdgeInsets.only(right: 16),
                      child: Text('Limitar qtde. deputados (0=ilimitado) :'),
                    ),
                    SizedBox(
                      width: 40,
                      child: TextField(
                        controller: _maxItmesController,
                      ),
                    )
                  ],
                ),
                Row(
                  mainAxisAlignment: MainAxisAlignment.center,
                  children: [
                    ElevatedButton(
                      onPressed: loading == true
                          ? null
                          : () => grpcService.createIdx(
                              maxItems:
                                  int.parse(_maxItmesController.value.text)),
                      child: Text(
                        'Criar Index',
                        style: theme.textTheme.bodyLarge,
                      ),
                    ),
                  ],
                ),
                const SizedBox(
                  height: 40,
                ),
                if (loading) const CircularProgressIndicator(),
              ],
            ),
          ),
        ),
      ),
    );
  }
}
