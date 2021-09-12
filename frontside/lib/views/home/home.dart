import 'package:flutter/material.dart';
// import 'package:beamer/beamer.dart';
import 'package:frontside/parts/navbar.dart';

class Home extends StatelessWidget {
  static const String route = "/";

  Home({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text("Home"),
        centerTitle: true,
      ),
      body: Container(
        width: MediaQuery.of(context).size.width,
        child: Image(
          image: AssetImage('assets/example_screens/mapsicle-map.png'),
          fit: BoxFit.none,
        ),
      ),
      bottomNavigationBar: NavBar(),
    );
  }
}
