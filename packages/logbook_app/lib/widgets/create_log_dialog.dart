import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:logbook_app/pages/details/index.dart';
import 'package:logbook_app/state.dart';
import 'package:logbook_core/logbook_core.dart';

class CreateLogDialog extends StatefulWidget {
  const CreateLogDialog({
    Key? key,
  }) : super(key: key);

  @override
  State<CreateLogDialog> createState() => _CreateLogDialogState();
}

class _CreateLogDialogState extends State<CreateLogDialog> {
  final TextEditingController _titleController = TextEditingController();
  final TextEditingController _descriptionController = TextEditingController();

  bool shouldOpenDetails = true;
  bool shouldCopyToClipboard = false;
  bool shouldOpenEditor = false;

  @override
  Widget build(BuildContext context) {
    return SimpleDialog(
      title: const Text('Add log entry'),
      children: [
        SizedBox(
          width: 900,
          height: 500,
          child: Padding(
            padding: const EdgeInsets.all(20),
            child: Column(
              children: [
                TextField(
                  controller: _titleController,
                  decoration: const InputDecoration(
                    border: OutlineInputBorder(),
                  ),
                ),
                const SizedBox(height: 20),
                Expanded(
                  child: TextField(
                    controller: _descriptionController,
                    keyboardType: TextInputType.multiline,
                    maxLines: null,
                    minLines: 40,
                    decoration: const InputDecoration(
                      border: OutlineInputBorder(),
                    ),
                  ),
                ),
                const SizedBox(height: 20),
                _buildButtons(),
              ],
            ),
          ),
        ),
      ],
    );
  }

  Widget _buildButtons() {
    var textEditor = LogbookConfig().textEditor;

    return Row(
      mainAxisAlignment: MainAxisAlignment.spaceBetween,
      children: [
        Row(
          children: [
            Row(
              children: [
                Checkbox(
                  value: shouldOpenDetails,
                  onChanged: (bool? value) {
                    setState(() {
                      shouldOpenDetails = value!;
                    });
                  },
                ),
                const Text('Open details'),
              ],
            ),
            Row(
              children: [
                Checkbox(
                  value: shouldCopyToClipboard,
                  onChanged: (bool? value) {
                    setState(() {
                      shouldCopyToClipboard = value!;
                    });
                  },
                ),
                const Text('Copy to clipboard'),
              ],
            ),
            if (textEditor != null)
              Row(
                children: [
                  Checkbox(
                    value: shouldOpenEditor,
                    onChanged: (bool? value) {
                      setState(() {
                        shouldOpenEditor = value!;
                      });
                    },
                  ),
                  const Text('Open in editor'),
                ],
              ),
          ],
        ),
        Row(
          mainAxisAlignment: MainAxisAlignment.end,
          children: [
            ElevatedButton(
              style: ElevatedButton.styleFrom(
                backgroundColor: Colors.teal,
              ),
              onPressed: () {
                Navigator.pop(context);
              },
              child: const Text('Cancel'),
            ),
            const SizedBox(width: 15),
            ElevatedButton(
              onPressed: () {
                var title = _titleController.text;
                var description = _descriptionController.text;

                createLogEntry(
                  title: title,
                  description: description,
                ).then((logEntry) {
                  final logbookBloc = context.read<LogbookBloc>();
                  logbookBloc.add(LogAdded());

                  if (shouldCopyToClipboard) {
                    Clipboard.setData(ClipboardData(text: logEntry.directory));
                  }

                  Navigator.pop(context);

                  if (shouldOpenDetails) {
                    Navigator.push(
                      context,
                      MaterialPageRoute(
                        builder: (context) => DetailsPage(
                          originalLogEntry: logEntry,
                        ),
                      ),
                    );
                  }

                  if (shouldOpenEditor) {
                    System.openInTextEditor(
                      textEditor!,
                      logEntry.directory,
                    );
                  }
                });
              },
              child: const Text('Save'),
            )
          ],
        ),
      ],
    );
  }

  @override
  void dispose() {
    _titleController.dispose();
    _descriptionController.dispose();
    super.dispose();
  }
}
