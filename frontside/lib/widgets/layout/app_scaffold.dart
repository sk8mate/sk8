import 'package:flutter/foundation.dart';
import 'package:flutter/material.dart';

import 'botom_navigation.dart';

class AppScaffold extends StatelessWidget {
  final String title;
  final bool showBottomNavigation;
  final Widget child;
  const AppScaffold({
    Key? key,
    required this.title,
    required this.child,
    this.showBottomNavigation = true,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        automaticallyImplyLeading: !kIsWeb,
        title: Text(title),
        centerTitle: true,
      ),
      body: SafeArea(
        child: SizedBox(
          width: MediaQuery.of(context).size.width,
          child: child,
        ),
      ),
      bottomNavigationBar:
          showBottomNavigation ? const BottomNavigation() : null,
    );
  }
}
