import 'dart:io';

import 'package:logbook_core/logbook_core.dart';

class FakeSystemService implements SystemService {
  final List<String> archivedDirectories = [];

  @override
  Future<void> archive(String originalDirectoryPath) async {
    archivedDirectories.add(originalDirectoryPath);
  }

  @override
  void shutdownApp() {
    // Don't do this in tests
  }
}

class FakeReadService implements ReadService {
  @override
  Future<String> readDescriptionLogOrNoteDescriptionFile(Directory dir) async {
    return dir.path;
  }
}
