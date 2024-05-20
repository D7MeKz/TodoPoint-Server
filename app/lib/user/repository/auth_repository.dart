import 'package:Todopoint/user/model/login_response.dart';
import 'package:dio/dio.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';

import '../../common/provider/dio.dart';

final authRepositoryProvider = Provider<AuthRepository>(
    (ref) {
      final dio = ref.watch(dioProvider);
      return AuthRepository(dio: dio, baseUrl: 'http://localhost:3001/auth');
    }
);

class AuthRepository{
  final String baseUrl;
  final Dio dio;

  AuthRepository({required this.baseUrl, required this.dio});

  Future<LoginResponse> login({
    required String username,
    required String password,
  }) async {
    final serialized = '$username:$password';
    final resp = await dio.post(
      '$baseUrl/login',
      options: Options(headers: {
        'Authorization': 'Basic $serialized',
      }),
    );
    return LoginResponse.fromJson(resp.data);
  }
}