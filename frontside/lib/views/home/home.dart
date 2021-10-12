import 'dart:async';
import 'dart:convert';

import 'package:frontside/views/home/autocomplete_request.pb.dart';
import 'package:http/http.dart' as http;
import 'package:flutter_dotenv/flutter_dotenv.dart';

import 'package:flutter/material.dart';
import 'package:frontside/parts/navbar.dart';
import 'package:frontside/views/home/autocomplete_response.pb.dart';
// import 'package:frontside/views/home/autocomplete_request.pb.dart';

// class Debouncer<T> {
//   Debouncer(this.duration, this.onValue);
//   final Duration duration;
//   void Function(T value) onValue;
//   late T _value;
//   late Timer? _timer;
//   T get value => _value;
//   set value(T val) {
//     _value = val;
//     _timer?.cancel();
//     _timer = Timer(duration, () => onValue(_value));
//   }
// }

// final _completer = Completer<AutocompleteResponse?>();

// var _debouncer = Debouncer<String>(Duration(milliseconds: 500), (value) {
//   _completer.complete(_getSuggestions(query));
// });

class PlacesSearchDelegate extends SearchDelegate<AutocompleteEntry?> {
  @override
  Widget buildLeading(BuildContext context) {
    return IconButton(
      icon: Icon(Icons.arrow_back),
      onPressed: () {
        close(context, null);
      },
    );
  }

  Future<AutocompleteResponse?> _getSuggestions(String value) async {
    if (value.isEmpty) {
      return null;
    }

    // TODO: Debounce a request
    // TODO: Parametrize a language
    // TODO: Error handling
    var request = AutocompleteRequest(language: "pl", search: value);
    final queryParameters = {
      "language": request.language,
      "search": request.search
    };
    final uri = Uri.parse(
        '${dotenv.env['SK8_BACKSIDE_URL']}/places/autocomplete?${Uri(queryParameters: queryParameters).query}');
    var response = await http.get(uri);

    return AutocompleteResponse.create()
      ..mergeFromProto3Json(jsonDecode(utf8.decode(response.bodyBytes)));
  }

  @override
  Widget buildResults(BuildContext context) {
    return FutureBuilder(
        builder: (BuildContext context,
            AsyncSnapshot<AutocompleteResponse?> response) {
          var suggestions = response.data?.data ?? [];

          return ListView.builder(
            itemBuilder: (context, index) {
              var suggestion = suggestions[index];

              return ListTile(
                title: Text(suggestion.name),
                subtitle: Text(suggestion.address),
                onTap: () {
                  close(context, suggestion);
                },
              );
            },
            itemCount: suggestions.length,
          );
        },
        future: _getSuggestions(query));
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
              print("result: ${result}");
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
