import 'package:Todopoint/common/const/color.dart';
import 'package:Todopoint/common/layout/default_layout.dart';
import 'package:Todopoint/user/provider/auth_provider.dart';
import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';

import '../../common/layout/custom_text_from_field.dart';

class LoginScreen extends ConsumerStatefulWidget {
  static const String routeName = 'login';

  const LoginScreen({super.key});

  @override
  ConsumerState<LoginScreen> createState() => _LoginScreenState();
}

class _LoginScreenState extends ConsumerState<LoginScreen> {
  String email = '';
  String password = '';

  @override
  Widget build(BuildContext context) {
    final state = ref.watch(authProvider);
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
                  height: 10.0,
                ),
                _SubTitle(),
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
                  height: 16.0,
                ),
                ElevatedButton(
                  onPressed: () {
                    state.login(username: email, password: password);

                  },
                  style: ElevatedButton.styleFrom(
                    foregroundColor: Colors.white,
                    backgroundColor: PRIMARY_COLOR,
                    shape: RoundedRectangleBorder(
                      borderRadius: BorderRadius.circular(8.0),
                    ),
                  ),
                  child: Text("로그인"),
                ),
                const SizedBox(
                  height: 8.0,
                ),
                _RegisterTextButton(),
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
      "투두뽀인트",
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
      "환영합니다.\n수많은 버그를 기대해주세요.",
      style: TextStyle(
        fontSize: 16,
        color: PRIMARY_TEXT_COLOR,
      ),
    );
  }
}

class _RegisterTextButton extends StatelessWidget {
  const _RegisterTextButton({super.key});

  @override
  Widget build(BuildContext context) {
    return TextButton(
      onPressed: () {

      },
      child: Text('회원가입'),
      style: TextButton.styleFrom(
        textStyle: TextStyle(
          fontSize: 14.0,
          fontWeight: FontWeight.w500,
        ),
        foregroundColor: GREY_TWO,
      ),
    );
  }
}
