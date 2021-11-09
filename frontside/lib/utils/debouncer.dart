import 'dart:async';

import 'package:flutter/material.dart';

class Debouncer {
  Duration delay;
  Timer? _timer;

  Debouncer({this.delay = const Duration(milliseconds: 500)});

  void debounce(VoidCallback callback) {
    if (_timer?.isActive ?? false) {
      _timer?.cancel();
    }
    _timer = Timer(delay, callback);
  }
}
