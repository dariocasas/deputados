import 'package:flutter/foundation.dart';
import 'package:grpc/grpc.dart';

import 'package:deputados/proto/fotos_service.pbgrpc.dart';

import '../social_networks_enum.dart';
import '../store/fotos_store.dart';
import 'grpc_service.dart';

abstract class FotosService {
  void init();

  void getFotos();
}

class FotosServiceImpl implements FotosService {
  late final ClientChannel channel;
  final GrpcService grpcService;

  final FotosStore fotosStore;

  FotosServiceImpl(this.fotosStore, this.grpcService);

  @override
  void init() {
    debugPrint('FotosService.init()');
    channel = grpcService.getClientChannel();
  }

  @override
  void getFotos() {
    final fotosServiceClient = FotosServiceClient(channel);
    final responsStream = fotosServiceClient.getFotos(GetFotosRequest());
    responsStream.listen(
      (value) {
        var foto = InfoFoto();
        foto.id = value.id;
        foto.fotoUrl = value.fotoUrl;
        foto.fotoUrl2 = value.fotoUrl2;
        foto.nome = value.nome;
        foto.partido = value.partido;
        foto.uf = value.uf;
        foto.url = value.url;

        for (var element in value.redes.red) {
          final r = Red();
          r.url = element;
          r.network = SocialNetwork.fromUrl(element);
          foto.redes.add(r);
        }

        fotosStore.fotos.add(foto);
      },
      onDone: () {},
    );
  }
}
