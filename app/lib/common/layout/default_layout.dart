import 'package:flutter/material.dart';

class DefaultLayout extends StatelessWidget {
  final Color? backgroundColor;
  final Widget child;
  final String? title;
  final Widget? bottomNavigationBar;

  const DefaultLayout({super.key, this.backgroundColor, required this.child, this.title, this.bottomNavigationBar});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: backgroundColor ?? Colors.white,
      appBar: renderAppBar(),
      body: child,
      bottomNavigationBar: bottomNavigationBar,
    );
  }

  AppBar? renderAppBar() {
    if (title == null) {
      return null;
    } else {
      return AppBar(
        backgroundColor: Colors.white,
        automaticallyImplyLeading: false, // 뒤로 가기 버튼 삭제
        elevation: 0, // 앞으로 튀어나오는 효과
        title: Text(
          title!, // Not null
          style: const TextStyle(
              fontSize: 16.0,
              fontWeight: FontWeight.w500
          ),
        ),
        foregroundColor: Colors.black,
      );
    }
  }
}
