import 'package:beamer/beamer.dart';
import 'package:flutter/material.dart';

import '../../generated/l10n.dart';
import '../../screens/index.dart';

class BottomNavigation extends StatefulWidget {
  const BottomNavigation({
    Key? key,
  }) : super(key: key);

  @override
  _BottomNavigationState createState() => _BottomNavigationState();
}

class _BottomNavigationState extends State<BottomNavigation> {
  int _seletedItem = 0;

  @override
  Widget build(BuildContext context) {
    return BottomNavigationBar(
      items: [
        BottomNavigationBarItem(
          icon: const Icon(Icons.map_outlined),
          label: AppI18n.of(context).spotsTitle,
        ),
        BottomNavigationBarItem(
          icon: const Icon(Icons.add_circle_outline),
          label: AppI18n.of(context).addSpot,
        ),
        BottomNavigationBarItem(
          icon: const Icon(Icons.person_outline),
          label: AppI18n.of(context).profilTitle,
        )
      ],
      currentIndex: _seletedItem,
      onTap: (index) {
        setState(() {
          _seletedItem = index;
        });
        String path = HomeScreen.route;
        switch (index) {
          case 0:
            path = HomeScreen.route;
            break;
          case 1:
            path = AddSpotScreen.route;
            break;
          case 2:
            path = ProfileScreen.route;
            break;
        }
        Beamer.of(context).update(
          state: BeamState.fromUriString(path),
        );
      },
    );
  }
}
