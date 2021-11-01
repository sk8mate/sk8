// GENERATED CODE - DO NOT MODIFY BY HAND
import 'package:flutter/material.dart';
import 'package:intl/intl.dart';
import 'intl/messages_all.dart';

// **************************************************************************
// Generator: Flutter Intl IDE plugin
// Made by Localizely
// **************************************************************************

// ignore_for_file: non_constant_identifier_names, lines_longer_than_80_chars
// ignore_for_file: join_return_with_assignment, prefer_final_in_for_each
// ignore_for_file: avoid_redundant_argument_values, avoid_escaping_inner_quotes

class AppI18n {
  AppI18n();

  static AppI18n? _current;

  static AppI18n get current {
    assert(_current != null,
        'No instance of AppI18n was loaded. Try to initialize the AppI18n delegate before accessing AppI18n.current.');
    return _current!;
  }

  static const AppLocalizationDelegate delegate = AppLocalizationDelegate();

  static Future<AppI18n> load(Locale locale) {
    final name = (locale.countryCode?.isEmpty ?? false)
        ? locale.languageCode
        : locale.toString();
    final localeName = Intl.canonicalizedLocale(name);
    return initializeMessages(localeName).then((_) {
      Intl.defaultLocale = localeName;
      final instance = AppI18n();
      AppI18n._current = instance;

      return instance;
    });
  }

  static AppI18n of(BuildContext context) {
    final instance = AppI18n.maybeOf(context);
    assert(instance != null,
        'No instance of AppI18n present in the widget tree. Did you add AppI18n.delegate in localizationsDelegates?');
    return instance!;
  }

  static AppI18n? maybeOf(BuildContext context) {
    return Localizations.of<AppI18n>(context, AppI18n);
  }

  /// `Add spot`
  String get addSpot {
    return Intl.message(
      'Add spot',
      name: 'addSpot',
      desc: '',
      args: [],
    );
  }

  /// `Home`
  String get homeTitle {
    return Intl.message(
      'Home',
      name: 'homeTitle',
      desc: '',
      args: [],
    );
  }

  /// `Profil`
  String get profilTitle {
    return Intl.message(
      'Profil',
      name: 'profilTitle',
      desc: '',
      args: [],
    );
  }

  /// `Spots`
  String get spotsTitle {
    return Intl.message(
      'Spots',
      name: 'spotsTitle',
      desc: '',
      args: [],
    );
  }
}

class AppLocalizationDelegate extends LocalizationsDelegate<AppI18n> {
  const AppLocalizationDelegate();

  List<Locale> get supportedLocales {
    return const <Locale>[
      Locale.fromSubtags(languageCode: 'en'),
      Locale.fromSubtags(languageCode: 'pl'),
    ];
  }

  @override
  bool isSupported(Locale locale) => _isSupported(locale);
  @override
  Future<AppI18n> load(Locale locale) => AppI18n.load(locale);
  @override
  bool shouldReload(AppLocalizationDelegate old) => false;

  bool _isSupported(Locale locale) {
    for (var supportedLocale in supportedLocales) {
      if (supportedLocale.languageCode == locale.languageCode) {
        return true;
      }
    }
    return false;
  }
}
