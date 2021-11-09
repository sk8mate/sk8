import 'dart:async';
import 'package:flutter/material.dart';

import '../../services/places/autocomplete_response.pb.dart';
import '../../services/places/places.dart';
import '../../utils/debouncer.dart';

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
