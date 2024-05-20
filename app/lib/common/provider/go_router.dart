
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:go_router/go_router.dart';

import '../../user/view/login_screen.dart';

final routerProvider = Provider<GoRouter>((ref){
  return GoRouter(
    routes: GoRouterNotifier().routes,
    initialLocation: '/login',
  );
});


class GoRouterNotifier{
  List<GoRoute> get routes => [
    GoRoute(
      path: '/login',
      name: LoginScreen.routeName,
      builder: (context, state) => LoginScreen(),
    ),
  ];
}