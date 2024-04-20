import 'package:Todopoint/common/const/colors.dart';
import 'package:flutter/material.dart';

class TaskCard extends StatelessWidget {
  const TaskCard({super.key});

  @override
  Widget build(BuildContext context) {
    return Container(
      decoration: BoxDecoration(
        borderRadius: BorderRadius.circular(8),
        border: Border.all(
          color: BORDER_COLOR,
          width: 1.0,
        )
      ),
      width: double.infinity,
      child: const SizedBox(
        height: 10.0,
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.stretch,
          children: [
            Row(
              crossAxisAlignment: CrossAxisAlignment.stretch,
              children: [
                Icon(Icons.check_box_outline_blank),
                SizedBox(width: 8.0),
                Text('쿠버네티스 1장')
              ],
            ),
            Row(
              crossAxisAlignment: CrossAxisAlignment.stretch,
              children: [
                Text('Date'),
              ],
            )
          ],
        ),
      ),
    );
  }
}
