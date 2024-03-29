import 'package:beamer/beamer.dart';
import 'package:flutter/cupertino.dart';
import 'package:frontside/views/about/about.dart';
import 'package:frontside/views/home/home.dart';

class BeamerRouter {
  static final routerDelegate = BeamerDelegate(
    locationBuilder: SimpleLocationBuilder(routes: {
      Home.route: (context, state) => BeamPage(
            key: ValueKey('home'),
            title: 'Home',
            child: Home(),
          ),
      About.route: (context, state) => BeamPage(
            key: ValueKey('about'),
            title: 'About',
            child: About(),
          ),
    }),
  );
}
