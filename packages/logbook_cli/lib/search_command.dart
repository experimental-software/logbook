import 'package:args/command_runner.dart';
import 'package:get_it/get_it.dart';
import 'package:logbook_core/logbook_core.dart';

class SearchCommand extends Command {
  final SearchService searchService = GetIt.I.get();

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
    argParser.addFlag(
      'regular-expression',
      abbr: 'r',
      help: 'Search with a regular expression',
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
      logEntries = await searchService.search(
        System.archiveDir,
        searchTerm,
        isRegularExpression: results['regular-expression'],
      );
    } else {
      logEntries = await searchService.search(
        System.baseDir,
        searchTerm,
        isRegularExpression: results['regular-expression'],
      );
    }
    for (var logEntry in logEntries.reversed) {
      print(logEntry.directory);
    }
  }
}
