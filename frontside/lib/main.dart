import 'package:flutter/material.dart';
import 'package:frontside/theme.dart';
import 'package:frontside/views/home/home.dart';

void main() {
  runApp(Sk8());
}

class Sk8 extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'sk8',
      theme: defaultTheme,
      home: Home(),
    );
  }
}
