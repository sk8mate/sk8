import 'package:dio/dio.dart';

class ConfigService {
  var dio = Dio()
    ..interceptors.add(
      InterceptorsWrapper(onRequest:
          (RequestOptions options, RequestInterceptorHandler handler) async {
        options.headers.addAll({
          'content-type': 'application/json',
        });
        options.baseUrl = 'https://api.sk8.town';
        options.receiveTimeout = 15000;
        options.connectTimeout = 15000;
        options.followRedirects = false;
        options.validateStatus = (status) {
          return status! < 300;
        };
        return handler.next(options);
      }),
    );
}
