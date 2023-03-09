class PluginRepository {
  Future<String> heavyCall() async {
    return Future<String>.delayed(
      const Duration(seconds: 2),
      () => 'Data Loaded',
    );
  }
}
