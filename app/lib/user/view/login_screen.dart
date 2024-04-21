import 'package:Todopoint/user/view/register_screen.dart';
import 'package:dio/dio.dart';
import 'package:flutter/material.dart';

import '../../common/component/custom_text_form_field.dart';
import '../../common/const/colors.dart';
import '../../common/const/data.dart';
import '../../common/layout/default_layout.dart';
import '../../common/view/root_tab.dart';

class LoginScreen extends StatefulWidget {

  const LoginScreen({super.key});

  @override
  State<LoginScreen> createState() => _LoginScreenState();
}

class _LoginScreenState extends State<LoginScreen> {
  String email = '';
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
            padding:  EdgeInsets.symmetric(horizontal: 16.0),
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.stretch,
              children: [
                _Title(),
                const SizedBox(
                  height: 16.0,
                ),
                CustomTextFormField(
                  hintText: "이메일을 입력해주세요.",
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
                    },
                    obscureText: true
                ),
                const SizedBox(
                  height: 12.0,
                ),
                ElevatedButton(
                    onPressed: () async {
                      // Codec<String, String> stringToBase64 = utf8.fuse(base64);
                      // String token = stringToBase64.encode(rawString);
                      final resp = await dio.post('http://localhost:3000/auth/login',
                          data: {'email': email, 'password': password}
                      );
                      final accessToken = resp.data['access_token'];
                      final refreshToken = resp.data['refresh_token'];

                      await storage.write(key: REFRESH_TOKEN_KEY, value: refreshToken);
                      await storage.write(key: ACCESS_TOKNE_KEY, value: accessToken);

                      if (resp.statusCode == 200){
                        Navigator.of(context).push(
                            MaterialPageRoute(
                              builder: (_) =>RootTab(),
                            )
                        );
                      }
                    },
                    style: ElevatedButton.styleFrom(
                      foregroundColor: Colors.white,
                      backgroundColor: PRIMARY_COLOR,
                    ),
                    child: Text('로그인')
                ),
                TextButton(
                    onPressed: () async {

                      Navigator.of(context).push(
                        MaterialPageRoute(
                            builder: (_) => RegisterScreen(),
                        )
                      );

                    },
                    style: TextButton.styleFrom(
                      foregroundColor: BODY_TEXT_COLOR,
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
          "Login",
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

