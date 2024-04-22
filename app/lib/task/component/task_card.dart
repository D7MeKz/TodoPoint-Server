import 'dart:ui';

import 'package:Todopoint/common/const/colors.dart';
import 'package:flutter/material.dart';

class TaskCard extends StatelessWidget {
  final bool isChecked;
  final String content;
  final String date;

  const TaskCard({super.key, required this.isChecked, required this.content, required this.date});

  @override
  Widget build(BuildContext context) {
    return AnimatedContainer(
      margin: const EdgeInsets.symmetric(
        horizontal: 16,
        vertical: 8,
      ),
      decoration: BoxDecoration(
        border: Border.all(
          width: 1.0,
          color: PRIMARY_COLOR,
        ),
        color: Colors.white,
        borderRadius: BorderRadius.circular(8.0),
        boxShadow: [
          BoxShadow(
            color: Colors.black.withOpacity(.1),
            offset: const Offset(0, 4),
            blurRadius: 10
          )
        ]
      ),
      
      duration: const Duration(milliseconds: 600),
      child: Padding(
        padding: const EdgeInsets.all(8.0),
        child: ListTile(
          // Check Icon
          leading: GestureDetector(
            onTap: (){
              // TODO check or uncheck the tesxt
            },
            child: Container(
              margin: const EdgeInsets.all(4.0),
              height: 20,
              width: 20,
              // duration: const Duration(milliseconds: 600),
              decoration: BoxDecoration(
                color: PRIMARY_COLOR,
                shape: BoxShape.rectangle,
                border: Border.all(color: PRIMARY_COLOR, width: 2.0),
                borderRadius: BorderRadius.circular(4.0),
              ),
              child: const Icon(
                Icons.check_rounded,
                color: Colors.white,
                size: 15,
              ),
            ),
          ),

          // Content
          title: Padding(
            padding: const EdgeInsets.symmetric(vertical: 3),
            child: Row(
              mainAxisAlignment: MainAxisAlignment.spaceBetween,
              children: [
                Text(
                  content,
                  style: const TextStyle(
                    color: PRIMARY_COLOR,
                    fontWeight: FontWeight.w500,
                    decoration: TextDecoration.lineThrough,
                    decorationColor: PRIMARY_COLOR
                  ),
                ),
                Icon(Icons.document_scanner_rounded),
              ],
            ),
          ),

          // Tags and time
          subtitle: Padding(
            padding: const EdgeInsets.only(top: 5, bottom: 3),
            child: Row(
              mainAxisAlignment: MainAxisAlignment.spaceBetween,
              children: [
                Text("Tags"),
                Text(
                  date,
                  style: const TextStyle(
                    color: BODY_TEXT_COLOR,
                  ),
                )
              ],
            ),
          ),

          // Task Tag
        ),
      ),
    );
  }
}

class _Content extends StatelessWidget {
  final String content;

  const _Content({super.key, required this.content});

  @override
  Widget build(BuildContext context) {
    return Expanded(
        child: Text(
          content,
          style: const TextStyle(
            color: BODY_TEXT_COLOR,
          ),
        )
    );
  }
}



