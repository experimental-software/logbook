import 'package:engineering_logbook/core/log_entry.dart';
import 'package:engineering_logbook/core/search.dart';
import 'package:flutter/material.dart';

import 'util/system.dart';

void main() => runApp(const EngineeringLogbookApp());

class EngineeringLogbookApp extends StatelessWidget {
  const EngineeringLogbookApp({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return const MaterialApp(
      home: Homepage(),
    );
  }
}

class Homepage extends StatefulWidget {
  const Homepage({Key? key}) : super(key: key);

  @override
  State<Homepage> createState() => _HomepageState();
}

class _HomepageState extends State<Homepage> {
  final TextEditingController _searchTermController = TextEditingController();
  final ScrollController _scrollController = ScrollController();

  late Future<List<LogEntry>> _logEntries;

  @override
  void initState() {
    _updateLogEntryList();
    super.initState();
  }

  void _updateLogEntryList() {
    _logEntries = search(System.baseDir, _searchTermController.text);
    setState(() {});
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('Engineering Logbook')),
      floatingActionButton: FloatingActionButton(
        onPressed: () {},
        tooltip: 'Add log entry',
        child: const Icon(Icons.add),
      ),
      body: Padding(
        padding: const EdgeInsets.all(25),
        child: Column(
          children: [
            _buildSearchRow(),
            const SizedBox(height: 20),
            FutureBuilder<List<LogEntry>>(
              future: _logEntries,
              builder: (context, snapshot) {
                if (snapshot.connectionState != ConnectionState.done) {
                  return const CircularProgressIndicator();
                }
                if (snapshot.hasData) {
                  return _buildLogEntryTable(snapshot.data!);
                } else {
                  return _buildLogEntryTable([]);
                }
              }
            ),
          ],
        ),
      ),
    );
  }

  Widget _buildSearchRow() {
    return Row(
      mainAxisAlignment: MainAxisAlignment.spaceBetween,
      children: [
        Expanded(
          child: SizedBox(
            width: 200,
            child: TextField(
              controller: _searchTermController,
              decoration: const InputDecoration(
                border: OutlineInputBorder(),
                hintText: 'Search log entries...',
                suffixIcon: Icon(Icons.search),
              ),
              onSubmitted: (query) {
                _updateLogEntryList();
                setState(() {});
              },
              autofocus: true,
            ),
          ),
        )
      ],
    );
  }

  Widget _buildLogEntryTable(List<LogEntry> logEntries) {
    return Flexible(
      child: Row(
        children: [
          Expanded(
            child: Scrollbar(
              thumbVisibility: true,
              thickness: 8,
              controller: _scrollController,
              child: SingleChildScrollView(
                scrollDirection: Axis.vertical,
                controller: _scrollController,
                child: DataTable(
                  showCheckboxColumn: false,
                  columns: const <DataColumn>[
                    DataColumn(
                      label: Text('Date/Time'),
                    ),
                    DataColumn(
                      label: Text('Titel'),
                    ),
                    DataColumn(
                      label: Text('Actions'),
                    ),
                  ],
                  rows: _buildTableRows(context, logEntries),
                ),
              ),
            ),
          ),
        ],
      ),
    );
  }

  List<DataRow> _buildTableRows(
    BuildContext context,
    List<LogEntry> logEntries,
  ) {
    var deviceInfo = MediaQuery.of(context);
    const widthDateTimeColumn = 220.0;
    const widthActionsColumn = 400.0;
    final widthTitleColumn =
        deviceInfo.size.width - widthDateTimeColumn - widthActionsColumn;

    return logEntries
        .map((logEntry) => DataRow(
              onSelectChanged: (_) {},
              cells: [
                DataCell(SizedBox(
                  width: widthDateTimeColumn,
                  child: Text(logEntry.dateTime.toIso8601String()),
                )),
                DataCell(SizedBox(
                  width: widthTitleColumn,
                  child: Text(logEntry.title),
                )),
                DataCell(
                  SizedBox(
                      width: widthActionsColumn,
                      child: Row(
                        children: [
                          IconButton(
                            icon: const Icon(Icons.edit),
                            onPressed: () {
                              System.openInEditor("/tmp");
                            },
                          ),
                          IconButton(
                            icon: const Icon(Icons.folder),
                            onPressed: () {
                              System.openInExplorer("/tmp");
                            },
                          ),
                          IconButton(
                            icon: const Icon(Icons.copy),
                            onPressed: () {
                              System.copyToClipboard("/tmp");
                            },
                          ),
                          const SizedBox(width: 30),
                          IconButton(
                            icon: const Icon(Icons.delete),
                            onPressed: () {},
                          ),
                        ],
                      )),
                ),
              ],
            ))
        .toList();
  }

  @override
  void dispose() {
    _searchTermController.dispose();
    _scrollController.dispose();
    super.dispose();
  }
}
