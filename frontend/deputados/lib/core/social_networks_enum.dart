import 'package:flutter/material.dart';
import 'package:font_awesome_flutter/font_awesome_flutter.dart';

enum SocialNetwork {
  unknown,
  twitter,
  facebook,
  instagram,
  youtube,
  linkedin;

  const SocialNetwork();

  @override
  String toString() {
    switch (index) {
      case 1:
        return 'Twitter';
      case 2:
        return 'Facebook';
      case 3:
        return 'Instagram';
      case 4:
        return 'YouTube';
      case 5:
        return 'LinkedIn';
      default:
        return 'unknown';
    }
  }

  factory SocialNetwork.fromInt(int index) {
    switch (index) {
      case 1:
        return twitter;
      case 2:
        return facebook;
      case 3:
        return instagram;
      case 4:
        return youtube;
      case 5:
        return linkedin;
      default:
        return unknown;
    }
  }

  factory SocialNetwork.fromUrl(String url) {
    for (var value in SocialNetwork.values) {
      if (url.contains('${value.name}.') == true) {
        return value;
      }
    }
    return unknown;
  }

  Widget icon() {
    switch (index) {
      case 1:
        return Image.asset("assets/images/x-black-48.png");
      case 2:
        return const Icon(FontAwesomeIcons.facebookF);
      case 3:
        return const Icon(FontAwesomeIcons.instagram);
      case 4:
        return const Icon(FontAwesomeIcons.youtube);
      case 5:
        return const Icon(FontAwesomeIcons.linkedin);
      default:
        return const Icon(FontAwesomeIcons.globe);
    }
  }
}
