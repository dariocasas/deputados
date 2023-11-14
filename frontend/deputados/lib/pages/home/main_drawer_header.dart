import 'package:flutter/material.dart';

class MainDrawerHeader extends StatelessWidget {
  const MainDrawerHeader({super.key});

  @override
  Widget build(BuildContext context) {
    return const DrawerHeader(
      decoration: BoxDecoration(
        image: DecorationImage(
          fit: BoxFit.fill,
          image: AssetImage('assets/images/composicao-congresso-600x400.webp'),
        ),
      ),
      child: Text(''),
    );
  }
}
