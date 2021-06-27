import 'dart:html';

import 'package:flutter/material.dart';
import 'package:frontside/routers/routers.dart';

class Home extends StatelessWidget {
  Home({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text("Home"),
        centerTitle: true,
      ),
      body: Column(
        children: <Widget>[
          Container(
            color: Colors.lightGreenAccent,

            width: double.infinity,
            padding: EdgeInsets.all(10),
            child: Column(

              children: <Widget>[
                SizedBox(height: 10),
                ElevatedButton.icon(
                  icon: Icon(Icons.home, size:32),
                  onPressed: (){Navigator.of(context).pushNamed(RouteManager.about);},
                  style: ElevatedButton.styleFrom(
                    primary: Colors.black,
                    onPrimary: Colors.lightGreenAccent,

                  ),
                  label: Text("About"),
               ),
                SizedBox(height: 10),
              ],
            )
          )

        ],
      ),

    );
  }
}
