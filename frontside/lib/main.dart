import 'package:beamer/beamer.dart';
import 'package:flutter/material.dart';
import 'package:frontside/routers/routers.dart';
void main() {
  Beamer.setPathUrlStrategy();
  runApp(Sk8());
}

class Sk8 extends StatelessWidget {


  @override
  Widget build(BuildContext context) {
    return MaterialApp.router(
      routeInformationParser: BeamerParser(),
      routerDelegate: RouteManager.routerDelegate,
      backButtonDispatcher:
        BeamerBackButtonDispatcher(delegate:  RouteManager.routerDelegate),

    );

  }
}
