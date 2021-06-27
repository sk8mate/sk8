import 'package:flutter/material.dart';


class About extends StatefulWidget {

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
        )
    );
  }
}