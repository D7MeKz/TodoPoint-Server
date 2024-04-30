import 'package:app/common/model/cursor_pagination_model.dart';
import 'package:app/common/model/pagination_params.dart';
import 'package:app/restaurant/model/restaurant_model.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';

import '../repository/restaurant_repository.dart';

final restaurantProvider =
    StateNotifierProvider<RestaurantStateNotifier, CursorPaginationBase>((ref) {
  final repository = ref.watch(restaurantRepositoryProvider);
  final notifier = RestaurantStateNotifier(repository: repository);

  return notifier;
});

class RestaurantStateNotifier extends StateNotifier<CursorPaginationBase> {
  final RestaurantRepository repository;

  RestaurantStateNotifier({
    required this.repository,
  }) : super(CursorPaginationLoading()) {
    paginate(); // 인스턴스를 만들자마자 pagination 실행해주세요. 추가적인 검증이 필요 없다.
  }
  // 리스트의 값이 변경되면 새로 렌더링된다.

  void paginate({
    int fetchCount = 20,
    // true 추가로 데이터를 더 가져옴
    // false - 새로고침(현재 상태를 덮어씌움)
    bool fetchMore = false, // 추가로 데이터 가져오기

    // 강제로 다시 로딩하기
    // true - Loading
    bool forceRefetch = false,
  }) async {
    try{
      // 5개의 상태 ( model에서 클래스로 정의한다.)

      // 1.바로 반환하는 상황
      // hasMore = false(기존 상태에서 다음 데이터가 없다면)
      if (state is CursorPagination && !forceRefetch) {
        final pState = state as CursorPagination;

        if (!pState.meta.hasMore == false) {
          return;
        }
      }

      // 로딩 중 1- fetchMore = true
      //  fetchMore = false -> 함수 그대로 실행 (중간에 데이터가 추가될 수 있음)
      final isLoading = state is CursorPaginationLoading;
      final isRefetching = state is CursorPaginationRefetching;
      final isFetchingMore = state is CursorPaginationFetchingMore;

      if (fetchMore && (isLoading || isRefetching || isFetchingMore)) {
        return;
      }

      // Pagination Params
      PaginationParams paginationParams = PaginationParams(
        count: fetchCount,
      );

      // 데이터를 추가로 더 가져오는 상황
      if (fetchMore) {
        // CursorPagination이라는 것은 무조건적으로 데이터를 가지고 있음을 의미한다.
        final pState = state as CursorPagination;
        state =
            CursorPaginationFetchingMore(meta: pState.meta, data: pState.data);

        paginationParams = paginationParams.copyWith(
          after: pState.data.last.id,
        );
      }else{
        // 데이터를 처음부터 가져오는 상황
        // 만약에 데이터가 있는 상황이라면 기본 데이터로ㅓ 보존한 채로 fetch
        if(state is CursorPagination && !forceRefetch){
          final pState = state as CursorPagination;

          state = CursorPagination(meta: pState.meta, data: pState.data);
        }else{
          // 나머지 상황
          state = CursorPaginationLoading();
        }
      }
      final resp = await repository.paginate(
        paginationParams: paginationParams,
      );

      if(state is CursorPaginationFetchingMore){
        final pState = state as CursorPaginationFetchingMore;

        // 기존 20개 + 최근 20개
        state = resp.copyWith(
            data: [
              ...pState.data,
              ...resp.data
            ]
        );
      }else{
        // 맨 처음에 대한 페이지
        state = resp;
      }
    }catch(e){
      state = CursorPaginationError(message: "데이터를 가져오지 못했습니다.");
    }
  }
}
