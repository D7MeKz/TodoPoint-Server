import 'package:Todopoint/common/const/colors.dart';
import 'package:flutter/material.dart';
import 'package:table_calendar/table_calendar.dart';


class MainCalendar extends StatelessWidget {
  final OnDaySelected onDaySelected; // 날짜 선택 시 실행할 함수
  final DateTime selectedDate; // 선택한 날짜.

  const MainCalendar({super.key, required this.onDaySelected, required this.selectedDate});

  @override
  Widget build(BuildContext context) {
    return TableCalendar(
      onDaySelected: onDaySelected,
      // 날짜 선택 시 실행할 함수
      selectedDayPredicate: (date) =>
        date.year == selectedDate.year && date.month == selectedDate.month && date.day == selectedDate.day,
      focusedDay: DateTime.now(),
      firstDay: DateTime(2000,1,1),
      lastDay: DateTime(3000,1,1),
      // Calendar Header
      headerStyle: const HeaderStyle(
        titleCentered: true, // 제목 중앙 위치
        titleTextStyle: TextStyle(
          fontWeight: FontWeight.w700,
          fontSize: 16.0
        ),
      ),
      calendarStyle: const CalendarStyle(
        isTodayHighlighted: false,
        selectedDecoration: BoxDecoration(
          shape: BoxShape.circle,
          color: SECONDARY_COLOR,
        ),

        // 캘린더 글꼴
        defaultTextStyle: TextStyle(
          fontWeight: FontWeight.w500,
          color: BODY_TEXT_COLOR,
        ),
        weekendTextStyle: TextStyle(
          fontWeight: FontWeight.w500,
          color: PRIMARY_COLOR,
        ),
        selectedTextStyle: TextStyle(
          fontWeight: FontWeight.w500
        )
      ),
    );
  }
}
