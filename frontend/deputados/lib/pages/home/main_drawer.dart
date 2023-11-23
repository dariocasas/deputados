import 'package:flutter/material.dart';
import 'package:flutter_modular/flutter_modular.dart';
import '../../core/service/database_service.dart';
import '../../core/service/server_service.dart';
import '../../core/store/db_status_store.dart';
import 'main_drawer_header.dart';
import '../widget/confirm_dialog.dart';

class MainDrawer extends StatefulWidget {
  const MainDrawer({super.key});

  @override
  State<MainDrawer> createState() => _MainDrawerState();
}

class _MainDrawerState extends State<MainDrawer> {
  final grpcService = Modular.get<DatabaseService>();
  final serverService = Modular.get<ServerService>();
  final dbStatusStore = Modular.get<DbStatusStore>();

  @override
  Widget build(BuildContext context) {
    return NavigationDrawer(
      selectedIndex: -1,
      onDestinationSelected: (index) {
        if (index == 0) {
          showDialog(
            context: context,
            builder: (BuildContext context) => ConfirmDialog(
              onConfirmed: () async {
                Modular.to.pop();
                await grpcService.dropDB();
                Modular.to.navigate('/');
                Modular.to.pop();
                return;
              },
              onDenied: () {
                Modular.to.pop();
                return;
              },
              title: 'Zerar banco de dados?',
            ),
          );
        }
        if (index == 1) {
          Modular.to.navigate('/fotos/');
        }
        if (index == 2) {
          dbStatusStore.askClose = true;
          Modular.to.pop();
        }
      },
      children: const [
        MainDrawerHeader(),
        NavigationDrawerDestination(
          icon: Icon(Icons.delete),
          label: Text('Zerar DB'),
        ),
        NavigationDrawerDestination(
          icon: Icon(Icons.photo),
          label: Text('Fotos'),
        ),
        Divider(),
        NavigationDrawerDestination(
          icon: Icon(Icons.close),
          label: Text('Fechar'),
        ),
      ],
    );
  }
}
