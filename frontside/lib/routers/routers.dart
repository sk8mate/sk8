import 'package:beamer/beamer.dart';
import 'package:frontside/views/about/about.dart';
import 'package:frontside/views/home/home.dart';



class RouteManager{
  static const String homeRoute = "/";
  static const String aboutRoute = "/about";

  static final routerDelegate = BeamerDelegate(
    locationBuilder: SimpleLocationBuilder(
        routes: {

          homeRoute: (context, state) => Home(),
          aboutRoute: (context, state) => About(),
        }

    ),
  );

}