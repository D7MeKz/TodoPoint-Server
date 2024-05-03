import 'package:flutter/material.dart';


class DefaultLayout2 extends StatelessWidget {
  const DefaultLayout({required this.body, this.actions, required this.title,Key? key}): super(key: key);
  final String title;
  final Widget body;
  final List<Widget>? actions;

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text(title),
        actions: actions,
      ),
      body: Padding(
        padding: EdgeInsets.symmetric(horizontal: 16.0),
        child: body,
      ),
    );
  }
}
