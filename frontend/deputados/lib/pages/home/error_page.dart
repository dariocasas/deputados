import 'package:flutter/material.dart';

import '../widget/check_close.dart';
import '../widget/custom_app_bar.dart';
import 'main_drawer.dart';

class ErrorPage extends StatefulWidget {
  final String error;
  const ErrorPage({super.key, required this.error});

  @override
  State<ErrorPage> createState() => _ErrorPageState();
}

class _ErrorPageState extends State<ErrorPage> {
  @override
  Widget build(BuildContext context) {
    final theme = Theme.of(context);

    return CheckClose(
      child: Scaffold(
        drawer: const MainDrawer(),
        appBar: const CustomAppBar(title: 'Deputados. Error.'),
        body: Center(
          child: Column(
            mainAxisAlignment: MainAxisAlignment.center,
            children: <Widget>[
              Text(
                Uri.decodeComponent(widget.error),
                style: theme.textTheme.bodyMedium,
              ),
            ],
          ),
        ),
      ),
    );
  }
}
