import 'package:flutter/material.dart';

import '../generated/l10n.dart';
import '../widgets/index.dart';

class AddSpotScreen extends StatelessWidget {
  static String route = '/add-spot';
  const AddSpotScreen({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return AppScaffold(
      title: AppI18n.of(context).addSpot,
      child: Center(
        child: Text(
          AppI18n.of(context).addSpot,
        ),
      ),
    );
  }
}
