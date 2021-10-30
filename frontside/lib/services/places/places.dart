import 'dart:convert';

import 'package:http/http.dart' as http;
import 'package:flutter_dotenv/flutter_dotenv.dart';

import 'package:frontside/services/places/autocomplete_request.pb.dart';
import 'package:frontside/services/places/autocomplete_response.pb.dart';

Future<AutocompleteResponse> getPlaces(String value) async {
  if (value.isEmpty) {
    return AutocompleteResponse(status: "success", data: []);
  }

  // TODO: Parametrize a language
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
