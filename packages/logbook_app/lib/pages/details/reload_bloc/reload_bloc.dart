import 'dart:io';

import 'package:equatable/equatable.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:logbook_core/logbook_core.dart';

// ignore: depend_on_referenced_packages
import 'package:uuid/uuid.dart';

part 'reload_event.dart';

part 'reload_state.dart';

class ReloadBloc extends Bloc<ReloadEvent, ReloadState> {
  ReloadBloc() : super(Loading()) {
    on<ReloadEvent>((event, emit) {
      emit(Loading());
    });

    on<NoteSelected>((event, emit) {
      emit(Loading(Directory(event.note.directory)));
    });

    on<LogEntryEdited>((event, emit) {
      emit(Reloading(event.logEntryPath));
    });
  }
}
