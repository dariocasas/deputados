import 'package:deputados/core/service/fotos_service.dart';
import 'package:deputados/pages/fotos/foto.dart';
import 'package:flutter/material.dart';
import 'package:flutter_modular/flutter_modular.dart';
import 'package:rx_notifier/rx_notifier.dart';

import '../../core/store/fotos_store.dart';
import '../home/main_drawer.dart';
import '../widget/custom_app_bar.dart';
import '../widget/check_close.dart';

const kX = 320.0;
const kY = 200.0;

class FotosPage extends StatefulWidget {
  const FotosPage({super.key});

  @override
  State<FotosPage> createState() => _FotosPageState();
}

class _FotosPageState extends State<FotosPage> {
  final fotosStore = Modular.get<FotosStore>();
  final fotosService = Modular.get<FotosService>();

  @override
  void initState() {
    fotosService.getFotos();

    super.initState();
  }

  @override
  Widget build(BuildContext context) {
    final fotos = context.select(() => fotosStore.fotos);

    return Scaffold(
      drawer: const MainDrawer(),
      appBar: const CustomAppBar(title: "Deputados. Fotos"),
      body: CheckClose(
        child: Column(
          children: [
            Expanded(
                child: GridView.builder(
              itemCount: fotos.length,
              gridDelegate: const SliverGridDelegateWithMaxCrossAxisExtent(
                maxCrossAxisExtent: kX,
                mainAxisExtent: kY,
                childAspectRatio: 1,
              ),
              itemBuilder: (BuildContext context, int index) {
                return Center(
                  child: Foto(
                    infoFoto: fotos[index],
                    size: const Size(kX, kY),
                  ),
                );
              },
            )),
          ],
        ),
      ),
    );
  }
}
