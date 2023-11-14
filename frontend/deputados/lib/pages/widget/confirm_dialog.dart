import 'package:flutter/material.dart';

class ConfirmDialog extends StatelessWidget {
  final VoidCallback? onConfirmed;
  final VoidCallback? onDenied;
  final String title;
  const ConfirmDialog({
    super.key,
    required this.onConfirmed,
    this.onDenied,
    required this.title,
  });

  @override
  Widget build(BuildContext context) {
    return AlertDialog(
      title: Text(title),
      actions: <Widget>[
        TextButton(
          onPressed: () => onConfirmed?.call(),
          child: const Text('Sim'),
        ),
        TextButton(
          onPressed: () => onDenied?.call(),
          child: const Text('Não'),
        ),
      ],
    );
  }
}
