import 'package:flutter/material.dart';

import '../../common/const/color.dart';
import '../../common/layout/custom_text_from_field.dart';
import '../../common/layout/default_layout.dart';

class RegisterScreen extends StatefulWidget {
  const RegisterScreen({super.key});

  @override
  State<RegisterScreen> createState() => _RegisterScreenState();
}

class _RegisterScreenState extends State<RegisterScreen> {
  String username = '';
  String password = '';
  String email = '';

  @override
  Widget build(BuildContext context) {
    return DefaultLayout(
      title: '',
      backgroundColor: Colors.white,
      child: SafeArea(
        top: true,
        bottom: false,
        child: Padding(
          padding: const EdgeInsets.symmetric(horizontal: 24.0),
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.stretch,
            children: [
              _Title(),
              const SizedBox(
                height: 16.0,
              ),
              CustomTextFormField(
                hintText: '이메일을 입력해주세요.',
                onChanged: (String value) {
                  email = value;
                },
              ),
              const SizedBox(
                height: 12.0,
              ),
              CustomTextFormField(
                hintText: "비밀번호를 입력해주세요.",
                onChanged: (String value) {
                  password = value;
                  print("Password : " + password);
                },
                obscureText: true,
              ),
              const SizedBox(
                height: 12.0,
              ),
              CustomTextFormField(
                hintText: '닉네임을 입력해주세요.',
                onChanged: (String value) {
                  username = value;
                },
              ),
              const SizedBox(
                height: 16.0,
              ),
              ElevatedButton(
                onPressed: () {},
                style: ElevatedButton.styleFrom(
                  foregroundColor: Colors.white,
                  backgroundColor: PRIMARY_COLOR,
                  shape: RoundedRectangleBorder(
                    borderRadius: BorderRadius.circular(8.0),
                  ),
                ),
                child: Text("회원가입"),
              ),
            ],
          ),
        ),
      ),
    );
  }
}

class _Title extends StatelessWidget {
  const _Title({super.key});

  @override
  Widget build(BuildContext context) {
    return Text(
      "회원가입",
      style: TextStyle(
          fontSize: 34, fontWeight: FontWeight.w500, color: Colors.black),
    );
  }
}


