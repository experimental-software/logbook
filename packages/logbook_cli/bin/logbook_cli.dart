#!/usr/bin/env dart

import 'package:args/command_runner.dart';
import 'package:logbook_cli/add_command.dart';
import 'package:logbook_cli/search_command.dart';
import 'package:get_it/get_it.dart';
import 'package:logbook_core/logbook_core.dart';

void main(List<String> args) {
  GetIt.I.registerSingleton(SearchService());
  GetIt.I.registerSingleton(WriteService());

  CommandRunner('logbook', 'A markdown-based engineering logbook.')
    ..addCommand(SearchCommand())
    ..addCommand(AddCommand())
    ..run(args);
}
