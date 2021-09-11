import 'package:flutter/material.dart';
// import 'package:beamer/beamer.dart';

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
      bottomNavigationBar: CustomNavBar(),
    );
  }
}

class CustomNavBar extends StatefulWidget {
  const CustomNavBar({
    Key? key,
  }) : super(key: key);

  @override
  _CustomNavBarState createState() => _CustomNavBarState();
}

class _CustomNavBarState extends State<CustomNavBar> {
  int _seletedItem = 0;

  @override
  Widget build(BuildContext context) {
    return BottomNavigationBar(
      items: [
        BottomNavigationBarItem(icon: Icon(Icons.map_outlined), title: Text("Spoty")),
        BottomNavigationBarItem(icon: Icon(Icons.add_circle_outline), title: Text("Dodaj nowy")),
        BottomNavigationBarItem(icon: Icon(Icons.person_outline), title: Text("Profil"))
      ],
      currentIndex: _seletedItem,
      onTap: (index) {
        setState(() {
          _seletedItem = index;
        });
      },
    );
  }
}
