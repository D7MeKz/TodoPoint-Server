import 'package:app/user/model/login_response.dart';
import 'package:app/user/model/token_response.dart';
import 'package:dio/dio.dart';

import '../../common/utils/data_utils.dart';

class AuthRepository {
  final String baseUrl;
  final Dio dio;

  AuthRepository(this.dio, {required this.baseUrl});

  Future<LoginResponse> login({
    required String username,
    required String password,
  }) async {
    final serialized = DataUtils.plainToBase64('$username:$password');
    final resp = await dio.post(
      '$baseUrl/login',
      options: Options(headers: {
        'Authorization': 'Basic $serialized',
      }),
    );
  }

  Future<TokenResponse> token() async{
    final resp = await dio.post(
      '$baseUrl/token',
      options: Options(headers: {
        'refreshToken': true,
      }),
    );
    
    return TokenResponse.fromJson(resp.data);
  }
}
