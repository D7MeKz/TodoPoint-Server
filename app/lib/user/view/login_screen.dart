import 'package:flutter/material.dart';
import 'package:app/common/layout/default_layout.dart';
import 'package:app/common/const/colors.dart';
import 'package:app/common/component/custom_text_form_field.dart';

class LoginScreen extends StatelessWidget {
  const LoginScreen({super.key});

  @override
  Widget build(BuildContext context) {
    return DefaultLayout(
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
                      const SizedBox(height: 10.0,),
                      _SubTitle(),
                      Image.asset(
                        'asset/img/misc/logo.png',
                        width: MediaQuery.of(context).size.width/3 * 2,
                      ),
                      CustomTextFormField(
                        hintText:'이메일을 입력해주세요.',
                        onChanged: (String value){},
                      ),
                      const SizedBox(height: 12.0,),
                      CustomTextFormField(
                          hintText: "비밀번호를 입력해주세요.",
                          onChanged: (String value){},
                          obscureText: true
                      ),
                      const SizedBox(height: 12.0,),
                      ElevatedButton(
                          onPressed: (){},
                          style: ElevatedButton.styleFrom(
                            foregroundColor: Colors.white,
                            backgroundColor: PRIMARY_COLOR,
                          ),
                          child: Text(
                              '로그인'
                          )
                      ),
                      TextButton(
                          onPressed: (){},
                          style: TextButton.styleFrom(
                            foregroundColor: BODY_TEXT_COLOR,
                          ),
                          child: Text(
                              '회원가입'
                          )
                      )
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
        fontSize: 34,
        fontWeight: FontWeight.w500,
        color: Colors.black
      ),
    );
  }
}

class _SubTitle extends StatelessWidget {
  const _SubTitle({super.key});

  @override
  Widget build(BuildContext context) {
    return Text(
      "로그인해주세요!\n오늘도 즐거운 하루를 보내기를!" ,
      style: TextStyle(
        fontSize: 16,
        color: BODY_TEXT_COLOR,
      ),
    );
  }
}


