import 'dart:io';

import 'package:deputados/proto/server_service.pbgrpc.dart';
import 'package:flutter/foundation.dart';
import 'package:grpc/grpc.dart';

import 'grpc_service.dart';

abstract class ServerService {
  void init();

  void shutdown();
}

class ServerServiceImpl implements ServerService {
  late final ClientChannel channel;
  final GrpcService grpcService;

  ServerServiceImpl(this.grpcService);

  @override
  void init() {
    debugPrint('ServerService.init()');
    channel = grpcService.getClientChannel();
  }

  @override
  Future<void> shutdown() async {
    debugPrint('ServerService.shutdown()');
    final serverServiceClient = ServerServiceClient(channel);

    try {
      await serverServiceClient.shutdown(ShutdownRequest());
      await grpcService.shutdown();
    } catch (e) {
      debugPrint(e.toString());
    }
    debugPrint('t+');

    exit(0);
  }
}
