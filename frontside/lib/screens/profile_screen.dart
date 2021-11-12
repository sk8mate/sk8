import 'package:flutter/material.dart';
import 'package:sk8/generated/l10n.dart';

import '/widgets/index.dart';

class ProfileScreen extends StatelessWidget {
  static String route = '/profile';
  const ProfileScreen({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return AppScaffold(
      title: AppI18n.of(context).profilTitle,
      child: Center(
        child: Text(
          AppI18n.of(context).profilTitle,
        ),
      ),
    );
  }
}
