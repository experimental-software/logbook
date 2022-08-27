import 'dart:io';

import 'package:args/command_runner.dart';
import 'package:get_it/get_it.dart';
import 'package:logbook_cli/archive_command.dart';
import 'package:logbook_core/logbook_core.dart';
import 'package:test/test.dart';

void main() {
  final WriteService writeService = WriteService();

  setUp(() async {
    GetIt.I.registerSingleton(SystemService());
    GetIt.I.registerSingleton(WriteService());

    System.baseDir = await Directory.systemTemp.createTemp('logs_');
    print('Log entry dir: ${System.baseDir}');
    System.archiveDir = await Directory.systemTemp.createTemp('archive_');
    print('Archive dir: ${System.archiveDir}');
  });

  test('archive with log entry dir', () async {
    var logEntry = await writeService.createLogEntry(
      title: 'Example log entry',
    );
    expect(System.archiveDir.listSync(), isEmpty);
    expect(System.baseDir.listSync(), isNotEmpty);

    List<String> args = ['archive', logEntry.directory];
    var runner = CommandRunner('logbook-test', 'Just a test runner');
    runner.addCommand(ArchiveCommand());
    await runner.runCommand(runner.parse(args));

    expect(System.archiveDir.listSync(), isNotEmpty);
    expect(System.baseDir.listSync(), isEmpty);
  });
}
