// ignore_for_file: avoid_init_to_null

import 'package:deputados/core/social_networks_enum.dart';
import 'package:rx_notifier/rx_notifier.dart';

part 'fotos_store.g.dart';

class Red {
  String url = '';
  SocialNetwork network = SocialNetwork.unknown;

  @override
  String toString() {
    return '(${network.name}) $url';
  }
}

class InfoFoto {
  int id = 0;
  String fotoUrl = '';
  String fotoUrl2 = '';
  String nome = '';
  String partido = '';
  String uf = '';
  String url = '';
  List<Red> redes = [];
}

@RxStore()
abstract class _FotosStore {
  @RxValue()
  RxList<InfoFoto> fotos = RxList<InfoFoto>([]);
}
