import 'dart:io';

import 'package:path/path.dart' as path;
import 'package:system/system.dart' as sys;

class SystemService {
  Future<void> archive(String originalDirectoryPath) async {
    await System.archive(originalDirectoryPath);
  }

  void shutdownApp() {
    exit(0);
  }
}

class System {
  System._();

  static Directory? _baseDir;
  static Directory? _archiveDir;

  static String get macVsCodePath =>
      '/Applications/Visual\\ Studio\\ Code.app/Contents/MacOS/Electron';

  static Future<void> openDirectory(String directory) async {
    sys.System.invoke('open $directory');
  }

  static Future<void> openInEditor(String directory) async {
    if (Platform.isMacOS) {
      sys.System.invoke('$macVsCodePath $directory > /dev/null 2>&1 &');
    }
    if (Platform.isLinux) {
      sys.System.invoke('code $directory');
    }
  }

  static Future<void> openInApp([String? directory = '']) async {
    if (Platform.isMacOS) {
      sys.System.invoke(
        '/Applications/logbook.app/Contents/MacOS/logbook $directory > /dev/null 2>&1  &',
      );
      return;
    }
    if (Platform.isLinux) {
      var home = path.absolute(Platform.environment['HOME']!);
      sys.System.invoke(
        '$home/bin/logbook/logbook $directory > /dev/null 2>&1 &',
      );
      return;
    }
    throw 'Unsupported OS';
  }

  static Future<void> archive(String originalDirectoryPath) async {
    final Directory originalDirectory = Directory(originalDirectoryPath);
    if (!originalDirectory.existsSync()) {
      return;
    }
    var archivedDirectoryPath = originalDirectoryPath.replaceFirst(
      baseDir.path,
      archiveDir.path,
    );
    final Directory archivedDirectory = Directory(archivedDirectoryPath);
    await archivedDirectory.create(recursive: true);
    await originalDirectory.rename(archivedDirectoryPath);

    final parentDayDirectory = originalDirectory.parent;
    final parentMonthDirectory = parentDayDirectory.parent;
    final parentYearDirectory = parentMonthDirectory.parent;

    if (await parentDayDirectory.list().isEmpty) {
      await parentDayDirectory.delete();
    }
    if (await parentMonthDirectory.list().isEmpty) {
      await parentMonthDirectory.delete();
    }
    if (await parentYearDirectory.list().isEmpty) {
      await parentYearDirectory.delete();
    }
  }

  static set baseDir(Directory? baseDir) {
    _baseDir = baseDir;
  }

  static Directory get baseDir {
    if (_baseDir != null) {
      return _baseDir!;
    }

    if (Platform.isMacOS) {
      return Directory('/Users/jmewes/doc/Notizen');
    }
    if (Platform.isLinux) {
      return Directory('/home/janux/doc/Notizen');
    }
    throw 'Unsupported OS: ${Platform.operatingSystem}';
  }

  static Directory get archiveDir {
    if (_archiveDir != null) {
      return _archiveDir!;
    }

    if (Platform.isMacOS) {
      return Directory('/Users/jmewes/doc/Archiv');
    }
    if (Platform.isLinux) {
      return Directory('/home/janux/doc/Archiv');
    }
    throw 'Unsupported OS: ${Platform.operatingSystem}';
  }

  static set archiveDir(Directory? archiveDir) {
    _archiveDir = archiveDir;
  }
}
