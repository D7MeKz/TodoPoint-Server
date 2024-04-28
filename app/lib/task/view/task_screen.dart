import 'package:Todopoint/calendar/component/main_calendar.dart';
import 'package:Todopoint/common/const/colors.dart';
import 'package:Todopoint/task/component/task_card.dart';
import 'package:flutter/material.dart';

class TaskScreen extends StatefulWidget {

  const TaskScreen({super.key});

  @override
  State<TaskScreen> createState() => _TaskScreenState();
}

class _TaskScreenState extends State<TaskScreen> {
  DateTime selectedDate = DateTime.utc(
      DateTime.now().year,
      DateTime.now().month,
      DateTime.now().day
  );

  @override
  Widget build(BuildContext context) {
    return Column(
      children: [
        // Calendar
        Padding(
          padding: const EdgeInsets.symmetric(horizontal: 16.0, vertical: 8.0),
          child: Container(
            child: MainCalendar(
              onDaySelected: onDaySelected,
              selectedDate: selectedDate,
            )
          ),
        ),
        // Card
        TaskCard(isChecked: false, content: "hello", date: selectedDate.toString()),
      ],
    );
  }

  void onDaySelected(DateTime selectedDate, DateTime foucusedDate){
    setState(() {
      this.selectedDate = selectedDate;
    });
  }
}


