import 'package:flutter/material.dart';

class NavBar extends StatefulWidget {
  const NavBar({
    Key? key,
  }) : super(key: key);

  @override
  _NavBarState createState() => _NavBarState();
}

class _NavBarState extends State<NavBar> {
  int _seletedItem = 0;

  @override
  Widget build(BuildContext context) {
    return BottomNavigationBar(
      items: [
        BottomNavigationBarItem(icon: Icon(Icons.map_outlined), label: 'Spoty'),
        BottomNavigationBarItem(icon: Icon(Icons.add_circle_outline), label: "Dodaj nowy"),
        BottomNavigationBarItem(icon: Icon(Icons.person_outline), label: "Profil")
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
