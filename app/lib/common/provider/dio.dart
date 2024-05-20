
import 'package:Todopoint/common/provider/secure_storage.dart';
import 'package:dio/dio.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:flutter_secure_storage/flutter_secure_storage.dart';

import '../const/data.dart';

final dioProvider = Provider<Dio>((ref) {
  final dio = Dio();

  final storage =  ref.watch(secureStorageProvider);

  dio.interceptors.add(CustomInterceptor(storage, ref));

  return dio;
});

class CustomInterceptor extends Interceptor{
  final FlutterSecureStorage storage;
  final Ref ref;

  CustomInterceptor(this.storage, this.ref);

  // 요청 보내기 전에
  @override
  Future<void> onRequest(RequestOptions options, RequestInterceptorHandler handler) async {
    if(options.headers['accessToken'] == 'true'){
      options.headers.remove('accessToken');

      final token = await storage.read(key: ACCESS_TOKEN_KEY);

      options.headers.addAll({
        'authorization' :'Bearer $token'
      });
    }

    // Refresh Token
    if(options.headers['refreshToken'] == 'true'){
      options.headers.remove('refreshToken');

      final refreshToken = await storage.read(key: REFRESH_TOKEN_KEY);

      options.headers.addAll({
        'authorization' :'Bearer $refreshToken'
      });
    }
    super.onRequest(options, handler);
  }

  @override
  void onError(DioException err, ErrorInterceptorHandler handler) async {
    // 401 Error
    // 토큰이 재발급
    // 새 토큰 발급 다시 요청

    final refreshToken = await storage.read(key: REFRESH_TOKEN_KEY);

    if(refreshToken == null){
      return handler.reject(err); // 에러 반환
      // handler.resolve(response) // 에러가 나지 않음.
    }

    final isStatus401 = err.response?.statusCode == 401;
    final isPathRefresh = err.requestOptions.path == '/auth/token'; // 토큰을 refresh하려다가 오류가 발생한건지?
    if(isStatus401 && isPathRefresh){
      final dio = Dio();

      try{
        final resp = await dio.post(
          'http://$ip/auth/token',
          options: Options(
              headers: {
                'authorization' : 'Bearer $refreshToken',
              }
          ),
        );

        final accessToken = resp.data['accessToken'];

        final options = err.requestOptions;
        options.headers.addAll({
          'authorization': 'Bearer $accessToken'
        });
        await storage.write(key: ACCESS_TOKEN_KEY, value: accessToken);

        // 요청 재전송
        final response = await dio.fetch(options);
        // 성공적인 응답에 대한 반응을 response
        return handler.resolve(response);
      }catch(e){
        // Logout
        // ref.read(authProvider.notifier).logout();
        return handler.reject(err); // 에러 반환
      }
    }
  }
  @override
  void onResponse(Response response, ResponseInterceptorHandler handler) {
    return super.onResponse(response, handler);
  }
}
