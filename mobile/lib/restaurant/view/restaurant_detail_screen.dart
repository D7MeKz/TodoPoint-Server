import 'package:app/common/layout/default_layout.dart';
import 'package:app/product/component/product_card.dart';
import 'package:app/restaurant/component/restaurant_card.dart';
import 'package:flutter/material.dart';

class RestaurantDetailScreen extends StatelessWidget {
  const RestaurantDetailScreen({super.key});

  @override
  Widget build(BuildContext context) {
    return DefaultLayout(
        title: 'Fire',
        child: Column(
          children: [
            RestaurantCard(
              name: "우라나라에서 가장 맛있는 짜장면집",
              image: Image.asset("asset/img/food/ddeok_bok_gi.jpg",),
              tags: ["신규", "세일중"],
              ratings: 4.89,
              ratingsCount: 200,
              deliveryTime: 20,
              deliveryFee: 3000,
              isDetail: true,
              detail: "Hello world",
            ),
            Padding(
              padding: const EdgeInsets.symmetric(horizontal: 16.0),
              child: ProductCard(),
            )
          ],
        )
    );
  }
}
