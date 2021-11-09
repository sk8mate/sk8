import 'package:flutter/material.dart';
import 'package:frontside/parts/navbar.dart';

import './places_search_delegate.dart';

class Home extends StatelessWidget {
  static const String route = "/";

  Home({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Column(children: [
        Padding(
          padding: EdgeInsets.all(16),
          child: TextField(
            readOnly: true,
            onTap: () async {
              var result = await showSearch(
                  context: context, delegate: PlacesSearchDelegate());
              print("result: $result");
            },
            decoration: InputDecoration(
                border: OutlineInputBorder(
                    borderSide: BorderSide(color: Colors.black),
                    borderRadius: BorderRadius.circular(16))),
          ),
        ),
      ]),
      bottomNavigationBar: NavBar(),
    );
  }
}
