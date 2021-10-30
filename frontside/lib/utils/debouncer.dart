import 'dart:async';

import 'package:flutter/material.dart';

class Debouncer {
  Duration delay;
  Timer? _timer;

  Debouncer({this.delay = const Duration(milliseconds: 500)});

  void debounce(VoidCallback callback) {
    if (_timer != null && _timer!.isActive) {
      _timer!.cancel();
    }
    _timer = new Timer(delay, callback);
  }
}
