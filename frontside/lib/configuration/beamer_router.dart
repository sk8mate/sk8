import 'package:beamer/beamer.dart';
import 'package:flutter/material.dart';
import 'package:sk8/generated/l10n.dart';

import '../screens/index.dart';

class BeamerRouter {
  static final routerDelegate = BeamerDelegate(
    initialPath: HomeScreen.route,
    notFoundRedirectNamed: HomeScreen.route,
    locationBuilder: SimpleLocationBuilder(routes: {
      HomeScreen.route: (context, state) => BeamPage(
            key: const ValueKey('home'),
            title: AppI18n.of(context).homeTitle,
            child: const HomeScreen(),
          ),
      ProfileScreen.route: (context, state) => BeamPage(
            key: const ValueKey('about'),
            title: AppI18n.of(context).profilTitle,
            child: const ProfileScreen(),
          ),
      AddSpotScreen.route: (context, state) => BeamPage(
            key: const ValueKey('add-spot'),
            title: AppI18n.of(context).addSpot,
            child: const AddSpotScreen(),
          ),
    }),
  );
}
