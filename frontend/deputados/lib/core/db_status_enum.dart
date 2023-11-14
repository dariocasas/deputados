enum DbStatus {
  uninitializated,
  indexCreated,
  dbCreated;

  const DbStatus();

  @override
  String toString() {
    switch (index) {
      case 1:
        return 'Índice criado';
      case 2:
        return 'db criado';
      default:
        return 'Não inicializado';
    }
  }

  factory DbStatus.fromInt(int index) {
    switch (index) {
      case 1:
        return indexCreated;
      case 2:
        return dbCreated;
      default:
        return uninitializated;
    }
  }

  String route() {
    switch (index) {
      case 1:
        return '/idxok/';
      case 2:
        return '/dbok/';
      default:
        return '/nodb/';
    }
  }
}
