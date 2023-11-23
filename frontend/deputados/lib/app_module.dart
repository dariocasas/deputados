import 'package:deputados/core/service/fotos_service.dart';
import 'package:deputados/pages/home/create_idx_page.dart';
import 'package:flutter_modular/flutter_modular.dart';
import 'core/service/grpc_service.dart';
import 'core/service/server_service.dart';
import 'core/store/fotos_store.dart';
import 'pages/fotos/fotos.dart';
import 'pages/home/create_db_page.dart';
import 'pages/home/error_page.dart';
import 'core/service/database_service.dart';
import 'core/store/create_db_store.dart';
import 'core/store/db_status_store.dart';
import 'pages/home/wait_connection_page.dart';

class AppModule extends Module {
  @override
  void binds(Injector i) {
    i.addSingleton<CreateDbStore>(CreateDbStore.new);
    i.addSingleton<DbStatusStore>(DbStatusStore.new);
    i.addSingleton<GrpcService>(GrpcServiceImpl.new);
    i.addSingleton<DatabaseService>(DatabaseServiceImpl.new);
    i.addSingleton<FotosStore>(FotosStore.new);
    i.addSingleton<FotosService>(FotosServiceImpl.new);
    i.addSingleton<ServerService>(ServerServiceImpl.new);
  }

  @override
  void routes(RouteManager r) {
    const defTransition = TransitionType.noTransition;
    r.child(
      '/',
      child: (context) => const WaitConnectionPage(),
      transition: defTransition,
    );
    r.child(
      '/waitconn/',
      child: (context) => const WaitConnectionPage(),
      transition: defTransition,
    );
    r.child(
      '/nodb/',
      child: (context) => const CreateIdxPage(),
      transition: defTransition,
    );
    r.child(
      '/idxok/',
      child: (context) => const CreateDbPage(),
      transition: defTransition,
    );
    r.child(
      '/dbok/',
      child: (context) => const FotosPage(),
      transition: defTransition,
    );
    r.child(
      '/error/:error',
      child: (context) => ErrorPage(error: r.args.params['error']),
      transition: defTransition,
    );
    r.child(
      '/fotos/',
      child: (context) => const FotosPage(),
      transition: defTransition,
    );
  }
}
