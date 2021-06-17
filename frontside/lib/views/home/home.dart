import 'package:flutter/material.dart';

class Home extends StatelessWidget {
  Home({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Container(
        color: Colors.black,
        child: Row(mainAxisAlignment: MainAxisAlignment.center, children: [
          Container(
              color: Theme.of(context).primaryColor, width: 50, height: 50),
          Container(width: 50, height: 50),
          Container(color: Colors.white, width: 50, height: 50),
        ]));
  }
}
