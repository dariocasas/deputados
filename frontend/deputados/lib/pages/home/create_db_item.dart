import 'package:flutter/material.dart';

import 'package:deputados/core/store/create_db_store.dart';

class CreateDbItem extends StatelessWidget {
  final CreateDbData deputado;
  const CreateDbItem({
    super.key,
    required this.deputado,
  });

  @override
  Widget build(BuildContext context) {
    final theme = Theme.of(context);
    final totalms = deputado.partialTime!.dbread +
        deputado.partialTime!.getapi +
        deputado.partialTime!.dbread +
        deputado.partialTime!.getpage +
        deputado.partialTime!.scanpage +
        deputado.partialTime!.dbwrite;

    return SizedBox(
      height: 20,
      child: Row(
        children: [
          SizedBox(
            width: 70,
            child: Text(deputado.id.toString()),
          ),
          SizedBox(
            width: 80,
            child: Text('${deputado.milliseconds.toString()} ms'),
          ),
          SizedBox(
            width: 95,
            child: Text('total = ${totalms.toString()} ms',
                style: theme.textTheme.bodySmall),
          ),
          SizedBox(
            width: 85,
            child: Text('dbread ${deputado.partialTime?.dbread.toString()} ms',
                style: theme.textTheme.bodySmall),
          ),
          SizedBox(
            width: 85,
            child: Text('getapi ${deputado.partialTime?.getapi.toString()} ms',
                style: theme.textTheme.bodySmall),
          ),
          SizedBox(
            width: 85,
            child: Text(
                'getpage ${deputado.partialTime?.getpage.toString()} ms',
                style: theme.textTheme.bodySmall),
          ),
          SizedBox(
            width: 90,
            child: Text(
                'scanpage ${deputado.partialTime?.scanpage.toString()} ms',
                style: theme.textTheme.bodySmall),
          ),
          SizedBox(
            width: 85,
            child: Text(
                'dbwrite ${deputado.partialTime?.dbwrite.toString()} ms',
                style: theme.textTheme.bodySmall),
          ),
          const SizedBox(width: 10),
          Text(deputado.name),
          const SizedBox(width: 10),
          Expanded(
            child: Text(
              deputado.errorString,
              style: theme.textTheme.bodySmall,
              overflow: TextOverflow.ellipsis,
            ),
          ),
        ],
      ),
    );
  }
}
