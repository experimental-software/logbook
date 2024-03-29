import 'dart:io';

import 'package:logbook_core/src/system_service.dart';
import 'package:uuid/uuid.dart';

import 'log_entry.dart';

class WriteService {
  Future<LogEntry> createLogEntry({
    required String title,
    String description = '',
  }) {
    return _createLogEntry(title: title, description: description);
  }

  Future<LogEntry> updateLogEntry({
    required LogEntry logEntry,
    required String title,
    required String description,
  }) async {
    var logEntryFile = await File(logEntry.path).create();
    var sink = logEntryFile.openWrite();
    sink.write('# $title\n\n');
    sink.write(description);
    sink.write('\n');
    await sink.flush();
    await sink.close();

    return LogEntry(
      dateTime: logEntry.dateTime,
      title: title,
      directory: logEntry.directory,
    );
  }

  Future<Directory> createNoteEntry({
    required String title,
    required String description,
    required Directory baseDir,
  }) {
    return _createNoteEntry(
      title: title,
      description: description,
      baseDir: baseDir,
    );
  }
}

Future<LogEntry> createLogEntry({
  required String title,
  required String description,
}) {
  return _createLogEntry(title: title, description: description);
}

Future<LogEntry> _createLogEntry({
  required String title,
  required String description,
}) async {
  var now = DateTime.now();
  var baseDir = System.baseDir;
  var slug = slugify(title);

  var logEntryDirectory =
      await WriterUtils.createLogEntryDirectory(baseDir, now, slug);
  var logEntryPath = '${logEntryDirectory.path}/$slug.md';

  var logEntryFile = await File(logEntryPath).create();
  var sink = logEntryFile.openWrite();
  sink.write('# $title\n\n');
  sink.write(description);
  sink.write('\n');
  await sink.flush();
  await sink.close();

  return LogEntry(
    dateTime: now,
    title: title,
    directory: logEntryDirectory.path,
  );
}

Future<Directory> createNoteEntry({
  required String title,
  required String description,
  required Directory baseDir,
}) {
  return _createNoteEntry(
    title: title,
    description: description,
    baseDir: baseDir,
  );
}

Future<Directory> _createNoteEntry({
  required String title,
  required String description,
  required Directory baseDir,
}) async {
  var slug = slugify(title);

  int biggestExistingIndex = 0;
  final Directory parentDirectory = Directory(baseDir.path);
  for (var f in parentDirectory.listSync()) {
    var name = f.path.split('/').last;
    var index = WriterUtils.parseIndexFromDirectory(name);
    if (index != null && index > biggestExistingIndex) {
      biggestExistingIndex = index;
    }
  }

  late String noteEntryDirectory;

  var id = (biggestExistingIndex + 1).toString().padLeft(2, '0');
  noteEntryDirectory = '${baseDir.path}/${id}_$slug';

  var logEntryPath = '$noteEntryDirectory/index.md';

  var logEntryFile = await File(logEntryPath).create(recursive: true);
  var sink = logEntryFile.openWrite();
  sink.write('# $title\n\n');
  sink.write(description);
  sink.write('\n');
  await sink.flush();
  await sink.close();

  return Directory(noteEntryDirectory);
}

String slugify(String s) {
  var result = s.toLowerCase();
  final maxSlugLength = 35;
  if (result.length > maxSlugLength) {
    result = result.substring(0, maxSlugLength);
  }
  result = result.replaceAll(RegExp(r'[^\w\s-]'), '');
  result = result.replaceAll(RegExp(r'[-\s]+'), '-');
  result = result.replaceAll(RegExp(r'-$'), '');
  return result;
}

class WriterUtils {
  static int? parseIndexFromDirectory(String name) {
    final regex = RegExp(r'^(\d+)_.*');
    final match = regex.firstMatch(name);
    if (match == null) {
      return null;
    } else {
      return int.parse(match.group(1)!);
    }
  }

  static Future<Directory> createLogEntryDirectory(
    Directory baseDir,
    DateTime time,
    String slug,
  ) async {
    var year = time.year.toString();
    var month = time.month.toString().padLeft(2, '0');
    var day = time.day.toString().padLeft(2, '0');
    var hour = time.hour.toString().padLeft(2, '0');
    var minute = time.minute.toString().padLeft(2, '0');

    var logEntryDirectory =
        '${baseDir.path}/$year/$month/$day/$hour.${minute}_$slug';

    var result = Directory(logEntryDirectory);
    if (result.existsSync()) {
      logEntryDirectory += '_${const Uuid().v4()}';
      result = Directory(logEntryDirectory);
    }

    return result.create(recursive: true);
  }
}
