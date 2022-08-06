import 'package:args/command_runner.dart';
import 'package:logbook_core/logbook_core.dart';

class SearchCommand extends Command {
  @override
  final name = 'search';

  @override
  final description = 'Searches for log entries.';

  SearchCommand() {
    argParser.addFlag(
      'archive',
      abbr: 'a',
      help: 'Search in archived log entries',
    );
  }

  @override
  void run() async {
    var results = argResults;
    if (results == null) {
      return;
    }
    var searchTerm = '';
    if (results.rest.isNotEmpty) {
      searchTerm = results.rest.first;
    }
    List logEntries;
    if (results['archive']) {
      logEntries = await search(System.archiveDir, searchTerm);
    } else {
      logEntries = await search(System.baseDir, searchTerm);
    }
    for (var logEntry in logEntries) {
      print(logEntry.directory);
    }
  }
}
