import 'dart:convert';

import 'package:app/common/const/data.dart';
import 'package:app/common/view/root_tab.dart';
import 'package:flutter/material.dart';
import 'package:app/common/layout/default_layout.dart';
import 'package:app/common/const/colors.dart';
import 'package:app/common/component/custom_text_form_field.dart';
import 'package:dio/dio.dart';

class LoginScreen extends StatefulWidget {
  const LoginScreen({Key? key}) : super(key: key);

  @override
  State<LoginScreen> createState() => _LoginScreenState();
}

class _LoginScreenState extends State<LoginScreen> {
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
                  padding: const EdgeInsets.symmetric(horizontal: 16.0),
                  child: Column(
                    crossAxisAlignment: CrossAxisAlignment.stretch,
                    children: [
                      _Title(),
                      const SizedBox(
                        height: 10.0,
                      ),
                      _SubTitle(),
                      Image.asset(
                        'asset/img/misc/logo.png',
                        width: MediaQuery.of(context).size.width / 3 * 2,
                      ),
                      CustomTextFormField(
                        hintText: '이메일을 입력해주세요.',
                        onChanged: (String value) {
                          username = value;
                          print("Username : " + username);
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
                          obscureText: true),

                      const SizedBox(
                        height: 12.0,
                      ),
                      ElevatedButton(
                          onPressed: () async {
                            final rawString = '$username:$password';
                            Codec<String, String> stringToBase64 = utf8.fuse(base64);

                            String token = stringToBase64.encode(rawString);
                            final resp = await dio.post(
                              'http://$ip/auth/login',
                              options: Options(
                                headers: {
                                  'authorization': 'Basic $token',
                                },
                              ),
                            );
                            final refreshToken = resp.data['refreshToken'];
                            final accessToken = resp.data['accessToken'];

                            await storage.write(key: REFRESH_TOKEN_KEY, value: refreshToken);
                            await storage.write(key: ACCESS_TOKEN_KEY, value: accessToken);

                            if (resp.statusCode != 404){
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
                          child: Text('로그인')),
                      TextButton(
                          onPressed: () async {

                            final resp = await dio.post('http://localhost:3000/auth/token',
                                data: {'email': username, 'password': password}
                            );

                          },
                          style: TextButton.styleFrom(
                            foregroundColor: BODY_TEXT_COLOR,
                          ),
                          child: Text('회원가입'))
                    ],
                  ),
                )
            )
        )
    );
  }
}

class _Title extends StatelessWidget {
  const _Title({super.key});

  @override
  Widget build(BuildContext context) {
    return Text(
      "환영합니다.",
      style: TextStyle(
          fontSize: 34, fontWeight: FontWeight.w500, color: Colors.black),
    );
  }
}

class _SubTitle extends StatelessWidget {
  const _SubTitle({super.key});

  @override
  Widget build(BuildContext context) {
    return Text(
      "로그인해주세요!\n오늘도 즐거운 하루를 보내기를!",
      style: TextStyle(
        fontSize: 16,
        color: BODY_TEXT_COLOR,
      ),
    );
  }
}
