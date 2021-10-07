import 'package:beamer/beamer.dart';
import 'package:flutter/cupertino.dart';
import 'package:sk8/views/about/about.dart';
import 'package:sk8/views/home/home.dart';
import 'package:sk8/views/login/login.dart';

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
      Login.route: (context, state) => BeamPage(
            key: ValueKey('about'),
            title: 'About',
            child: Login(),
          ),
    }),
  );
}
