import 'package:flutter/material.dart';

import '../constants/app_images.dart';
import '../generated/l10n.dart';
import '../widgets/index.dart';

class HomeScreen extends StatelessWidget {
  static String route = '/';
  const HomeScreen({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return AppScaffold(
      title: AppI18n.of(context).homeTitle,
      child: Image(
        image: AssetImage(AppImages.map),
        fit: BoxFit.none,
      ),
    );
  }
}
