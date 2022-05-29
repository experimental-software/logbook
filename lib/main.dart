import 'package:engineering_logbook/core/log_entry.dart';
import 'package:engineering_logbook/core/search.dart';
import 'package:flutter/material.dart';

import 'util/system.dart';
import 'widgets/create_log_dialog.dart';

void main() => runApp(const EngineeringLogbookApp());

class EngineeringLogbookApp extends StatelessWidget {
  const EngineeringLogbookApp({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return const MaterialApp(
      home: Homepage(),
      debugShowCheckedModeBanner: false,
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

  final Set<LogEntry> _markedForDeletion = {};

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
      // floatingActionButtonLocation: FloatingActionButtonLocation.startDocked,
      floatingActionButton: FloatingActionButton(
        onPressed: () async {
          await _showCreateLogDialog(context);
          _updateLogEntryList();
        },
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
              },
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
                  columns: <DataColumn>[
                    DataColumn(
                      label: _markedForDeletion.isNotEmpty ? IconButton(
                        icon: const Icon(Icons.delete),
                        onPressed: () {
                          for (var logEntry in _markedForDeletion) {
                            logEntries.remove(logEntry);
                            System.delete(logEntry.directory);
                          }
                          _markedForDeletion.clear();
                          _updateLogEntryList();
                          setState(() {});
                        },
                      ) : const Text(''),
                    ),
                    const DataColumn(
                      label: Text('Date/Time'),
                    ),
                    const DataColumn(
                      label: Text('Titel'),
                    ),
                    const DataColumn(
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
                DataCell(
                  MarkDeletedCheckbox(
                    key: ValueKey(logEntry.dateTime),
                    markedForDeletion: _markedForDeletion,
                    logEntry: logEntry,
                    notifyParent: () {
                      setState(() {});
                    },
                  ),
                ),
                DataCell(SizedBox(
                  width: widthDateTimeColumn,
                  child: Text(_formatTime(logEntry.dateTime)),
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
                              System.openInEditor(logEntry.directory);
                            },
                          ),
                          IconButton(
                            icon: const Icon(Icons.folder),
                            onPressed: () {
                              System.openInExplorer(logEntry.directory);
                            },
                          ),
                          IconButton(
                            icon: const Icon(Icons.copy),
                            onPressed: () {
                              System.copyToClipboard(logEntry.directory);
                            },
                          ),
                        ],
                      )),
                ),
              ],
            ))
        .toList();
  }

  // TODO Does this work without the "context" argument?
  Future<void> _showCreateLogDialog(BuildContext context) async {
    await showDialog(
        context: context,
        builder: (context) {
          return const CreateLogDialog();
        });
  }

  @override
  void dispose() {
    _searchTermController.dispose();
    _scrollController.dispose();
    super.dispose();
  }
}

String _formatTime(DateTime t) {
  var year = t.year.toString();
  var month = t.month.toString().padLeft(2, '0');
  var day = t.day.toString().padLeft(2, '0');
  var hour = t.hour.toString().padLeft(2, '0');
  var minute = t.minute.toString().padLeft(2, '0');

  return '$year-$month-$day $hour:$minute';
}

class MarkDeletedCheckbox extends StatefulWidget {
  final LogEntry logEntry;
  final Set<LogEntry> markedForDeletion;
  final Function notifyParent;

  const MarkDeletedCheckbox({
    Key? key,
    required this.markedForDeletion,
    required this.logEntry,
    required this.notifyParent,
  }) : super(
          key: key,
        );

  @override
  State<MarkDeletedCheckbox> createState() => _MarkDeletedCheckboxState();
}

class _MarkDeletedCheckboxState extends State<MarkDeletedCheckbox> {
  bool isChecked = false;

  @override
  Widget build(BuildContext context) {
    return Checkbox(
      key: ValueKey(widget.logEntry.dateTime.toIso8601String()),
      value: isChecked,
      onChanged: (bool? value) {
        if (value == null) {
          return;
        }

        if (value) {
          widget.markedForDeletion.add(widget.logEntry);
        } else {
          widget.markedForDeletion.remove(widget.logEntry);
        }

        setState(() {
          isChecked = value;
          widget.notifyParent();
        });
      },
    );
  }
}
