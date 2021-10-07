import 'package:flutter/material.dart';
import 'package:sk8/parts/navbar.dart';

class Home extends StatelessWidget {
  static const String route = "/home"; //default "/"

  Home({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text("Home"),
        centerTitle: true,
      ),
      body: SafeArea(
        child: Container(
          width: MediaQuery.of(context).size.width,
          child: Image(
            image: AssetImage('assets/example_screens/mapsicle-map.png'),
            fit: BoxFit.none,
          ),
        ),
      ),
      bottomNavigationBar: NavBar(),
    );
  }
}
