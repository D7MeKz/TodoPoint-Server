import 'package:flutter/cupertino.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:flutter_secure_storage/flutter_secure_storage.dart';

import '../../common/const/data.dart';
import '../../common/provider/secure_storage.dart';
import '../repository/auth_repository.dart';

final authProvider = ChangeNotifierProvider<AuthProvider>((ref) {
  final authRepository = ref.watch(authRepositoryProvider);
  final storage = ref.watch(secureStorageProvider);

  return AuthProvider(authRepository: authRepository, storage: storage);
});

class AuthProvider extends ChangeNotifier{
  final AuthRepository authRepository;
  final FlutterSecureStorage storage;

  AuthProvider( {required this.authRepository, required this.storage});

  Future<void> login({
    required String username,
    required String password,
  }) async {
    try {
      final resp = await authRepository.login(
        username: username,
        password: password,
      );

      await storage.write(key: ACCESS_TOKEN_KEY, value: resp.accessToken);
      await storage.write(key: REFRESH_TOKEN_KEY, value: resp.refreshToken);
      return;
      // final userResp = await repository.getMe();

      // state = userResp;
      //
      // return userResp;
    } catch (e) {
      // state = UserModelError(e.toString());
      // return Future.value(state);
      print(e.toString());
      return;
    }
  }



}