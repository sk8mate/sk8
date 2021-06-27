import 'package:flutter/material.dart';
import 'package:frontside/views/about/about.dart';
import 'package:frontside/views/home/home.dart';



class RouteManager{
  static const String home = "/";
  static const String about = "/about";

  static Route<dynamic> generateRoute(RouteSettings routeSettings){

    switch(routeSettings.name){
      case home:
        return MaterialPageRoute(
          builder: (context) => Home(),

        );
      case about:
        return MaterialPageRoute(
            builder: (context) => About()

        );

      default:
        throw FormatException("ups");

    }

  }

}