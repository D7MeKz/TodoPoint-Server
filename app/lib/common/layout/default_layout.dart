import 'package:flutter/material.dart';

import '../const/color.dart';

class DefaultLayout extends StatelessWidget {
  final Widget child;
  final String? title;
  final Widget? bottomNavigationBar;
  final Color? backgroundColor;

  const DefaultLayout({required this.child,super.key, required this.title, this.bottomNavigationBar, this.backgroundColor});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: backgroundColor ?? BACKGROUND_COLOR,
      appBar: renderAppBar(),
      body: child,
      bottomNavigationBar: bottomNavigationBar,
    );
  }

  AppBar? renderAppBar(){
    if(title  == null){
      return null;
    }else{
      return AppBar(
        backgroundColor: backgroundColor ?? BACKGROUND_COLOR,
        elevation: 0, // 앞으로 튀어나오는 효과
        title: Text(
          title!, // Not null
          style: const TextStyle(
              fontSize: 16.0,
              fontWeight: FontWeight.w500
          ),
        ),
        foregroundColor: PRIMARY_TEXT_COLOR,
      );
    }
  }
}
