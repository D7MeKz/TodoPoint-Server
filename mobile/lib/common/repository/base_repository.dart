
import 'package:app/common/model/model_with_id.dart';
import 'package:retrofit/http.dart';

import '../model/cursor_pagination_model.dart';
import '../model/pagination_params.dart';

abstract class IBasePaginationRepository<T extends IModelWithId>{
  Future<CursorPagination<T>> paginate({
    @Queries() PaginationParams? paginationParams = const PaginationParams(), // 쿼리로 값 전송
  });
}