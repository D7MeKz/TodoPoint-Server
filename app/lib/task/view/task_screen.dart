import 'package:Todopoint/common/const/colors.dart';
import 'package:Todopoint/task/component/task_card.dart';
import 'package:flutter/material.dart';

class TaskScreen extends StatelessWidget {
  const TaskScreen({super.key});

  @override
  Widget build(BuildContext context) {
    return Container(
      child: const Center(
        child: Padding(
          padding: EdgeInsets.symmetric(horizontal: 16.0),
          child: TaskCard(
            content: "Hello",
            isChecked: false,
            date: "2000",
          ),
        ),
      ),
    );
  }
}
