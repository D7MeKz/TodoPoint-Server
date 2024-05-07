import 'package:Todopoint/common/const/color.dart';
import 'package:flutter/material.dart';

class CustomTextFormField extends StatelessWidget {
  final String? hintText; // Get from params
  final String? errorText;
  final bool obscureText;
  final bool autoFocus;
  final ValueChanged<String>? onChanged;

  const CustomTextFormField({
    required this.onChanged,
    this.autoFocus = false,
    this.obscureText = false,
    this.hintText,
    this.errorText,
    Key? key,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    final baseBorder = OutlineInputBorder(
        borderSide: BorderSide(
            color: INPUT_BORDER_COLOR,
            width: 1.0)); // 테두리가 있는 입력하는 border,default는 underline

    return TextFormField(
      cursorColor: PRIMARY_COLOR,
      // 비밀번호 입력할때만
      obscureText: obscureText,
      autofocus: autoFocus, // 눌러야 focusing이 된다.
      onChanged: onChanged,
      decoration: InputDecoration(
          contentPadding: EdgeInsets.all(20),
          hintText: hintText,
          hintStyle: TextStyle(
            color: BODY_TEXT_COLOR,
            fontSize: 14.0,
          ),
          errorText: errorText,
          fillColor: INPUT_BG_COLOR,
          // true 배경색 존재
          filled: true,
          enabledBorder: baseBorder,
          focusedBorder: baseBorder.copyWith(
              borderSide:
              baseBorder.borderSide.copyWith(color: PRIMARY_COLOR))),
    );
  }
}
