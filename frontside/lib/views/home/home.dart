import 'dart:async';

import 'package:flutter/material.dart';
import 'package:frontside/parts/navbar.dart';

import 'package:frontside/services/places/autocomplete_response.pb.dart';
import 'package:frontside/services/places/places.dart';
import 'package:frontside/utils/debouncer.dart';

class PlacesSearchDelegate extends SearchDelegate<AutocompleteEntry?> {
  final _debouncer = Debouncer();

  @override
  Widget buildLeading(BuildContext context) {
    return IconButton(
      icon: Icon(Icons.arrow_back),
      onPressed: () {
        close(context, null);
      },
    );
  }

  Future<AutocompleteResponse> _getResults() async {
    final Completer<AutocompleteResponse> _completer = new Completer();

    _debouncer.debounce(() async {
      _completer.complete(await getPlaces(query));
    });

    return _completer.future;
  }

  @override
  Widget buildResults(BuildContext context) {
    return FutureBuilder(
        future: _getResults(),
        builder: (BuildContext context,
            AsyncSnapshot<AutocompleteResponse> response) {
          var places = response.data?.data ?? [];

          return ListView.builder(
            itemCount: places.length,
            itemBuilder: (context, index) {
              var place = places[index];

              return ListTile(
                title: Text(place.name),
                subtitle: Text(place.address),
                onTap: () {
                  close(context, place);
                },
              );
            },
          );
        });
  }

  @override
  Widget buildSuggestions(BuildContext context) {
    return buildResults(context);
  }

  @override
  List<Widget> buildActions(BuildContext context) => <Widget>[];
}

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
