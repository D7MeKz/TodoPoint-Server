import 'package:app/common/const/colors.dart';
import 'package:app/common/layout/default_layout.dart';
import 'package:app/common/view/root_tab.dart';
import 'package:app/user/view/login_screen.dart';
import 'package:dio/dio.dart';
import 'package:flutter/material.dart';

import '../const/data.dart';


class SplashScreen extends StatefulWidget {
  const SplashScreen({super.key});

  @override
  State<SplashScreen> createState() => _SplashScreenState();
}

class _SplashScreenState extends State<SplashScreen> {
  @override
  void initState() {
    super.initState();
    //deleteToken();
    checkToken();
  }

  void deleteToken() async{
    await storage.deleteAll();
  }

  void checkToken() async{
    final refreshToken = await storage.read(key:REFRESH_TOKEN_KEY);
    final accessToken = await storage.read(key:ACCESS_TOKEN_KEY);
    final dio = Dio();

    // TODO Validate token
    try{
      final resp = await dio.post('http://localhost:3000/auth/token',
        options: Options(
          headers: {
            'authorization' : 'Bearer $refreshToken',
          }
        )
      );

      await storage.write(key: ACCESS_TOKEN_KEY, value: resp.data['accessToken']);

      // Go to Root Tab
      Navigator.of(context).pushAndRemoveUntil(
          MaterialPageRoute(
            builder: (_) => const RootTab(),
          ), (route) => false);
    }catch(e){
      // Go to login screen
      Navigator.of(context).pushAndRemoveUntil(
          MaterialPageRoute(
            builder: (_) => const LoginScreen(),
          ), (route) => false);
    }

  }
  @override
  Widget build(BuildContext context) {
    return DefaultLayout(
        backgroundColor: PRIMARY_COLOR,
        title: '',
        child: SizedBox(
          width: MediaQuery.of(context).size.width,
          child: Column(
            mainAxisAlignment: MainAxisAlignment.center,
            children: [
              Image.asset(
                'asset/img/logo/logo.png',
                width: MediaQuery.of(context).size.width/2,
              ),
              const SizedBox(height: 16.0,),
              const CircularProgressIndicator(
                color: Colors.white,
              )
            ],
          ),
        )
    );
  }
}
