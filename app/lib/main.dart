import 'package:Todopoint/user/view/login_screen.dart';
import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';

void main() {
  runApp(
    ProviderScope(
      child: _App(),
    ),
  );
}

// Private Widget
class _App extends ConsumerWidget {
  const _App({super.key});

  @override
  Widget build(BuildContext context, WidgetRef ref  ) {
    return MaterialApp(
      theme: ThemeData(fontFamily: 'NotoSans'),
      debugShowCheckedModeBanner: false,
      home: LoginScreen(),
    );
    // final router = ref.watch(routerProvider);
    //
    // return MaterialApp.router(
    //   theme: ThemeData(fontFamily: 'NotoSans'),
    //   debugShowCheckedModeBanner: false,
    //   routerDelegate: router.routerDelegate,
    //   routeInformationParser: router.routeInformationParser,
    //   routeInformationProvider: router.routeInformationProvider,
    // );
  }
}

