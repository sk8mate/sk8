import 'package:beamer/beamer.dart';
import 'package:flutter_dotenv/flutter_dotenv.dart';
import 'package:flutter/material.dart';

import './beamer_router.dart';
import './theme.dart';

Future main() async {
  await dotenv.load(fileName: ".env");
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
