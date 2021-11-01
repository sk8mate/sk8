import 'package:beamer/beamer.dart';
import 'package:flutter/material.dart';
import 'package:flutter_localizations/flutter_localizations.dart';

import 'configuration/beamer_router.dart';
import 'configuration/theme.dart';
import 'generated/l10n.dart';

void main() async {
  ErrorWidget.builder = (FlutterErrorDetails details) {
    bool inDebug = false;
    assert(() {
      inDebug = true;
      return true;
    }());
    if (inDebug) {
      return ErrorWidget(details.exception);
    }
    return Container(
      alignment: Alignment.center,
      child: const Text(
        'Something went wrong :(',
        style: TextStyle(color: Colors.red),
        textDirection: TextDirection.ltr,
      ),
    );
  };
  Beamer.setPathUrlStrategy();
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return MaterialApp.router(
      title: 'sk8',
      localizationsDelegates: const [
        AppI18n.delegate,
        GlobalMaterialLocalizations.delegate,
        GlobalWidgetsLocalizations.delegate,
        GlobalCupertinoLocalizations.delegate,
      ],
      supportedLocales: AppI18n.delegate.supportedLocales,
      theme: defaultTheme,
      routeInformationParser: BeamerParser(),
      routerDelegate: BeamerRouter.routerDelegate,
      backButtonDispatcher: BeamerBackButtonDispatcher(
        delegate: BeamerRouter.routerDelegate,
      ),
    );
  }
}
