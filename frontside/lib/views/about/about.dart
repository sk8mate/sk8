import 'package:flutter/material.dart';

class About extends StatefulWidget {
  static const String route = "/about";

  @override
  _About createState() => _About();
}

class _About extends State<About> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
        appBar: AppBar(
      title: Text("About"),
      centerTitle: true,
    ));
  }
}
