import 'package:http/http.dart' as http;

import 'package:flutter/material.dart';
import 'package:frontside/parts/navbar.dart';
import 'package:frontside/views/home/autocomplete_response.pb.dart';
import 'package:frontside/views/home/autocomplete_request.pb.dart';

class PlacesSearchDelegate extends SearchDelegate<String> {
  @override
  Widget buildLeading(BuildContext context) {
    return IconButton(
      icon: Icon(Icons.arrow_back),
      onPressed: () {
        close(context, '');
      },
    );
  }

  Future<List<AutocompleteEntryResponse>> _getSuggestions() {}

  @override
  Widget buildSuggestions(BuildContext context) {
    return FutureBuilder(
      builder: (BuildContext context,
          AsyncSnapshot<List<AutocompleteEntryResponse>> response) {
        if (!response.hasData) {
          return Text('No results');
        }
        return ListView.builder(
          itemBuilder: (context, index) => ListTile(
            title: Text(response.data![index]),
            onTap: () {
              close(context, response.data![index]);
            },
          ),
          itemCount: suggestions.data!.length,
        );
      },
      future: _getSuggestions(),
    );
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
            onTap: () {
              showSearch(context: context, delegate: PlacesSearchDelegate());
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
