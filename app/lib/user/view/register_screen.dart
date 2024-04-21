import 'package:Todopoint/common/component/custom_text_form_field.dart';
import 'package:Todopoint/common/const/colors.dart';
import 'package:Todopoint/common/layout/default_layout.dart';
import 'package:Todopoint/user/view/login_screen.dart';
import 'package:dio/dio.dart';
import 'package:flutter/material.dart';

class RegisterScreen extends StatefulWidget {
  const RegisterScreen({super.key});

  @override
  State<RegisterScreen> createState() => _RegisterScreenState();
}

class _RegisterScreenState extends State<RegisterScreen> {
  String email = '';
  String username = '';
  String password = '';

  @override
  Widget build(BuildContext context) {
    final dio = Dio();
    return DefaultLayout(
      title: '',
      child: SingleChildScrollView(
        keyboardDismissBehavior: ScrollViewKeyboardDismissBehavior.onDrag,
        child: SafeArea(
          top: true,
          bottom: false,
          child: Padding(
            padding: EdgeInsets.symmetric(horizontal: 16.0),
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.stretch,
              children: [
                const _Title(),
                const SizedBox(
                  height: 16.0,
                ),
                CustomTextFormField(
                  hintText: "이메일을 입력하세요.",
                  onChanged: (String value) {
                    email = value;
                  },
                ),
                const SizedBox(
                  height: 12.0,
                ),
                CustomTextFormField(
                  hintText: "닉네임을 입력하세요.",
                  onChanged: (String value) {
                    username = value;
                  },
                ),
                const SizedBox(
                  height: 12.0,
                ),
                CustomTextFormField(
                  hintText: "비밀번호를을 입력하세요.",
                  onChanged: (String value) {
                    password = value;
                  },
                  obscureText: true,
                ),
                const SizedBox(
                  height: 12.0,
                ),
                ElevatedButton(
                    onPressed: () async {
                      final resp = await dio.post('http://localhost:3000/auth/register',
                          data: {'email': email, 'username':username, 'password': password}
                      );
                      // If success register, go to login page
                      if(resp.statusCode == 201){
                        Navigator.of(context).push(
                            MaterialPageRoute(
                              builder: (_) => const LoginScreen(),
                            )
                        );
                      }
                    },
                    style: ElevatedButton.styleFrom(
                      foregroundColor: Colors.white,
                      backgroundColor: PRIMARY_COLOR,
                    ),
                    child: Text('회원가입')
                )
              ],
            ),
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
    return const Column(
      mainAxisAlignment: MainAxisAlignment.center,
      children: [
        Text(
          "Register",
          style: TextStyle(
              fontSize: 34,
              fontWeight: FontWeight.w500,
              color: Colors.black
          ),
        ),
      ],
    );
  }
}

