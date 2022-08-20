part of 'reload_bloc.dart';

abstract class ReloadEvent extends Equatable {
  const ReloadEvent();
}

class UpdatedEvent extends ReloadEvent {
  final String id = const Uuid().v4();

  @override
  List<Object?> get props => [id];
}

class NoteSelected extends ReloadEvent {
  final String id = const Uuid().v4();

  @override
  List<Object?> get props => [id];
}
