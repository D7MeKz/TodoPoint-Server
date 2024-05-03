import 'package:app/common/model/cursor_pagination_model.dart';
import 'package:app/common/provider/pagination_provider.dart';
import 'package:app/rating/model/rating_model.dart';
import 'package:app/restaurant/repository/restaurant_rating_repository.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';

final restaurantRatingProvider = StateNotifierProvider.family<RestaurantRatingNotifier,
    CursorPaginationBase, String>((ref, id) {
  final repo = ref.watch(restaurantRatingRepositoryProvider(id));
  return RestaurantRatingNotifier(repository: repo);
});

class RestaurantRatingNotifier
    extends PaginationProvider<RatingModel, RestaurantRatingRepository> {
  RestaurantRatingNotifier({required super.repository});
}
