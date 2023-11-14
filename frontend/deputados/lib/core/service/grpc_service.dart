import 'package:flutter/foundation.dart';
import 'package:grpc/grpc.dart';
import 'package:flutter_dotenv/flutter_dotenv.dart';

abstract class GrpcService {
  void init();
  ClientChannel getClientChannel();
  Future<void> shutdown();
}

class GrpcServiceImpl implements GrpcService {
  late final ClientChannel channel;

  GrpcServiceImpl();

  @override
  void init() {
    debugPrint('GrpcService.init()');

    final sGrpcServerPort = dotenv.env["DEPS_GRPC_PORT"];
    int? grpcServerPort = int.tryParse(sGrpcServerPort ?? "");
    grpcServerPort ??= 50051;
    final grpcServerAddr = dotenv.env["DEPS_GRPC_ADDR"] ?? "localhost";

    channel = ClientChannel(
      grpcServerAddr,
      port: grpcServerPort,
      options: const ChannelOptions(
        credentials: ChannelCredentials.insecure(),
        // CodecRegistry(codecs: const [GzipCodec(), IdentityCodec()]),
      ),
    );

    channel.onConnectionStateChanged.listen((state) {
      debugPrint("grpc:onConnectionStateChanged=$state");
    }, onError: (err) {
      debugPrint("grpc:onError occur=$err");
    }, onDone: () {
      debugPrint("grpc:ondone called");
    }, cancelOnError: true);
  }

  @override
  ClientChannel getClientChannel() {
    return channel;
  }

  @override
  Future<void> shutdown() async {
    debugPrint('GrpcService.shutdown()');
    return await channel.shutdown();
  }
}
