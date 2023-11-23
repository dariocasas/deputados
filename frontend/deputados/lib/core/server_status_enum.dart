import 'package:flutter/material.dart';

enum ServerStatus {
  unknown,
  serving,
  notServing,
  serviceUnknown;

  const ServerStatus();

  @override
  String toString() {
    switch (index) {
      case 1:
        return 'Serving';
      case 2:
        return 'Not Serving';
      case 3:
        return 'Service Unknown';
      default:
        return 'Unknown';
    }
  }

  factory ServerStatus.fromInt(int index) {
    switch (index) {
      case 1:
        return serving;
      case 2:
        return notServing;
      case 3:
        return serviceUnknown;
      default:
        return unknown;
    }
  }

  Widget icon() {
    return Container(
        // color: Colors.transparent,
        width: 28.0,
        height: 28.0,
        decoration: BoxDecoration(
          border: Border.all(
            color: Colors.white54,
            width: 2,
          ),
          color: getColor(),
          shape: BoxShape.circle,
        ));
  }

  Color getColor() {
    switch (index) {
      case 1:
        return Colors.green;
      case 2:
        return Colors.yellow;
      case 3:
        return Colors.red;
      default:
        return Colors.red;
    }
  }
}
