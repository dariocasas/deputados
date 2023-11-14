import 'package:deputados/core/store/fotos_store.dart';
import 'package:flutter/material.dart';
import 'package:font_awesome_flutter/font_awesome_flutter.dart';
import 'package:url_launcher/url_launcher.dart';

class Foto extends StatefulWidget {
  final InfoFoto infoFoto;
  final Size size;
  const Foto({super.key, required this.infoFoto, required this.size});

  @override
  State<Foto> createState() => _FotoState();
}

class _FotoState extends State<Foto> {
  @override
  Widget build(BuildContext context) {
    // final theme = Theme.of(context);
    final f2 = Size(widget.size.width - 10, widget.size.height - 6);
    final f1 = Size(f2.width - 8, f2.height - 8);

    return Padding(
      padding: const EdgeInsets.only(top: 8, left: 4, right: 4),
      child: SizedBox(
        width: f1.width,
        height: f1.height,
        child: Container(
          color: Colors.blueGrey,
          child: Column(
            children: [
              SizedBox(
                height: f1.height - 52,
                child: Row(
                  children: [
                    Image.network(
                      widget.infoFoto.fotoUrl2,
                      errorBuilder: (context, error, stackTrace) {
                        return Image.network(
                          widget.infoFoto.fotoUrl,
                          errorBuilder: (context, error, stackTrace) {
                            return Image.asset(
                                "assets/images/composicao-congresso-600x400.webp");
                          },
                        );
                      },
                    ),
                    Expanded(
                      child: Container(
                        color: Colors.blueGrey.shade200,
                        child: Column(
                          crossAxisAlignment: CrossAxisAlignment.start,
                          children: [
                            Wrap(
                              children: socialNetworks(widget.infoFoto),
                            ),
                            Expanded(child: Container()),
                          ],
                        ),
                      ),
                    )
                  ],
                ),
              ),
              const SizedBox(height: 6),
              Row(
                mainAxisAlignment: MainAxisAlignment.center,
                children: [
                  Text(widget.infoFoto.nome,
                      overflow: TextOverflow.clip,
                      style: const TextStyle(
                          color: Colors.white,
                          fontWeight: FontWeight.w500,
                          fontSize: 16)),
                ],
              ),
              Row(
                mainAxisAlignment: MainAxisAlignment.center,
                children: [
                  Text("${widget.infoFoto.partido} - ${widget.infoFoto.uf}",
                      style: const TextStyle(
                          color: Colors.yellow,
                          fontWeight: FontWeight.w600,
                          fontSize: 14)),
                ],
              )
            ],
          ),
        ),
      ),
    );
  }
}

Future<void> _launchUrl(url) async {
  final Uri url0 = Uri.parse(url);
  if (!await launchUrl(url0)) {
    throw Exception('Could not launch $url0');
  }
}

List<Widget> socialNetworks(InfoFoto foto) {
  final List<Widget> res = [];

  res.add(IconButton(
    tooltip: 'Página web',
    onPressed: () {
      _launchUrl('https://www.camara.leg.br/deputados/${foto.id}');
    },
    icon: const FaIcon(FontAwesomeIcons.globe),
  ));

  for (var red in foto.redes) {
    res.add(IconButton(
      tooltip: red.network.toString(),
      onPressed: () {
        _launchUrl(red.url);
      },
      icon: red.network.icon(),
    ));
  }

  return res;
}
