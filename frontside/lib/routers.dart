import 'package:beamer/beamer.dart';
import 'package:flutter/cupertino.dart';
import 'package:frontside/views/about/about.dart';
import 'package:frontside/views/home/home.dart';



class RouteManager{
  static const String home = "/";
  static const String about = "/about";

  static final routerDelegate = BeamerDelegate(
    locationBuilder: SimpleLocationBuilder(
        routes: {

          home: (context, state) =>  BeamPage(
            key: ValueKey('home'),
            title: 'Home',
            child: Home(),
          ),
          about: (context, state) => BeamPage(
              key: ValueKey('about'),
              title: 'About',
              child: About(),
          ),

    }

    ),
  );

}