import 'package:beamer/beamer.dart';
import 'package:flutter/material.dart';
import 'package:frontside/beamer_router.dart';
import 'package:frontside/theme.dart';

void main() {
  Beamer.setPathUrlStrategy();
  runApp(Sk8());
}

class Sk8 extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp.router(
      title: 'sk8',
      theme: defaultTheme,
      routeInformationParser: BeamerParser(),
      routerDelegate: BeamerRouter.routerDelegate,
      backButtonDispatcher:
          BeamerBackButtonDispatcher(delegate: BeamerRouter.routerDelegate),
    );
  }
}
