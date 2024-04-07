import 'package:flutter/material.dart';
import 'package:app/common/component/custom_text_form_field.dart';
import 'package:app/user/view/login_screen.dart';


void main() {
  runApp(
      _App(),
  );
}

// Private Widget
class _App extends StatelessWidget {
  const _App({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      theme: ThemeData(
        fontFamily: 'NotoSans'
      ),
      debugShowCheckedModeBanner: false,
      home: Scaffold(
        backgroundColor: Colors.white,
        body: LoginScreen()
      ),
    );
  }
}
