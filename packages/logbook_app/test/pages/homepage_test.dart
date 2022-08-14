import 'dart:io';

import 'package:flutter_test/flutter_test.dart';
import 'package:logbook/pages/homepage/index.dart';
import 'package:logbook_core/logbook_core.dart';
import 'package:get_it/get_it.dart';

import '../widget_test_app.dart';

void main() {
  late _DummySearchService searchService;

  setUp(() {
    GetIt.I.registerSingleton<SearchService>(_DummySearchService());
  });

  tearDown(() async {
    await GetIt.I.reset();
  });

  testWidgets('render homepage', (tester) async {
    await tester.pumpWidget(const WidgetTestApp(Homepage()));

    expect(find.text('Logbook'), findsOneWidget);
  });

  testWidgets('show all log entries after opening the app', (tester) async {
    searchService = GetIt.I.get<SearchService>() as _DummySearchService;
    searchService.searchResults.add(
      LogEntry(
        dateTime: DateTime.now(),
        title: 'Hello, World!',
        directory: '',
      ),
    );

    await tester.pumpWidget(const WidgetTestApp(Homepage()));

    await tester.pumpAndSettle();
    expect(find.text('Hello, World!'), findsOneWidget);
  });
}

class _DummySearchService implements SearchService {
  final List<LogEntry> searchResults = [];

  @override
  Future<List<LogEntry>> search(Directory _, String query) async {
    return searchResults;
  }
}
