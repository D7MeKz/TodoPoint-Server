import 'package:app/common/model/cursor_pagination_model.dart';
import 'package:app/common/provider/pagination_provider.dart';
import 'package:app/restaurant/model/restaurant_model.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:collection/collection.dart';

import '../repository/restaurant_repository.dart';

// null contain -> ?
final restaurantDetailProvider =
    Provider.family<RestaurantModel?, String>((ref, id) {
  final state = ref.watch(restaurantProvider);

  // 사실상 데이터가 존재하지 않는다.

  if (state is! CursorPagination) {
    return null;
  }
  // If data did not exsit, they return error
  return state.data.firstWhereOrNull((element) => element.id == id);
});

final restaurantProvider =
    StateNotifierProvider<RestaurantStateNotifier, CursorPaginationBase>((ref) {
  final repository = ref.watch(restaurantRepositoryProvider);
  final notifier = RestaurantStateNotifier(repository: repository);

  return notifier;
});

class RestaurantStateNotifier
    extends PaginationProvider<RestaurantModel, RestaurantRepository> {
  RestaurantStateNotifier({
    required super.repository,
  });
  // 리스트의 값이 변경되면 새로 렌더링된다.

  void getDetail({
    required String id,
  }) async {
    //  만약 데이터가 없는 상태라면, 데이터를 가져오는 시도를한다.
    if (state is! CursorPagination) {
      await paginate();
    }

    if (state is! CursorPagination) {
      return;
    }

    final pState = state as CursorPagination;

    final resp = await repository.getRestaurantDetail(id: id);
    // 데이터가 없을 때는 그냥 캐시의 끝에 데이터를 추가한다.
    if(pState.data.where((element) => element.id == id).isEmpty){
      state = pState.copyWith(
        data: <RestaurantModel>[
          ... pState.data,
          resp,
        ]
      );
    }
    // Restaurant Detail 모델로 변경
    state = pState.copyWith(
      data: pState.data
          .map<RestaurantModel>((e) => e.id == id ? resp : e)
          .toList(),
    );
  }
}
