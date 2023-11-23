// ignore_for_file: lines_longer_than_80_chars

import 'package:flutter/material.dart';
import 'package:flutter_modular/flutter_modular.dart';

import '../../core/service/database_service.dart';

class CreateDbWorking extends StatefulWidget {
  const CreateDbWorking({super.key});

  @override
  State<CreateDbWorking> createState() => _CreateDbWorkingState();
}

class _CreateDbWorkingState extends State<CreateDbWorking> {
  final grpcService = Modular.get<DatabaseService>();

  @override
  Widget build(BuildContext context) {
    return SizedBox(
      height: 70,
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
                  ElevatedButton(
                      onPressed: () => grpcService.cancelCreateDB(),
                      child: const Text('Cancelar'))
                ],
              ),
            ),
          ],
        ),
      ),
    );
  }
}
