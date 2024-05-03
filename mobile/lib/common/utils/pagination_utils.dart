import 'package:app/common/provider/pagination_provider.dart';
import 'package:flutter/cupertino.dart';

abstract class PaginationUtils {
  static void pagination({
    required ScrollController controller,
    required PaginationProvider provider
}){
    if (controller.offset > controller.position.maxScrollExtent - 300) {
      provider.paginate(fetchMore: true);
    }
  }
}