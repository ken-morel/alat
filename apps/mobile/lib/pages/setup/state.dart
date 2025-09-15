import 'package:alat/state.dart';
import 'package:flutter/widgets.dart';

class SetupState {
  final PageController pageController;
  final AppState appState;
  const SetupState(this.pageController, this.appState);
  void next() {
    pageController.nextPage(
      duration: Duration(milliseconds: 200),
      curve: Curves.easeOutExpo,
    );
  }

  void prev() {
    pageController.previousPage(
      duration: Duration(milliseconds: 200),
      curve: Curves.easeOutExpo,
    );
  }
}
