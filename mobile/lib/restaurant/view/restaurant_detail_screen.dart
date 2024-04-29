
import 'package:app/common/layout/default_layout.dart';
import 'package:app/product/component/product_card.dart';
import 'package:app/restaurant/component/restaurant_card.dart';
import 'package:app/restaurant/model/restaurant_detail_model.dart';
import 'package:dio/dio.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';

import '../../common/const/data.dart';

class RestaurantDetailScreen extends StatelessWidget {
  final String id;

  const RestaurantDetailScreen({super.key, required this.id});

  Future<Map<String, dynamic>> getRestaurantDetail() async {
    final dio = Dio();

    // Get Access token key
    final accessToken = await storage.read(key: ACCESS_TOKEN_KEY);
    // get data
    final resp = await dio.get('http://$ip/restaurant/$id',
        options: Options(
            headers:{
              'authorization' : 'Bearer $accessToken',
            }
        )
    );

    // Return data
    return resp.data;
  }

  @override
  Widget build(BuildContext context) {
    return DefaultLayout(
        title: 'Fire',
        child: FutureBuilder<Map<String, dynamic>>(
          builder: (_, AsyncSnapshot<Map<String, dynamic>> snapshot){
            if(!snapshot.hasData){
              return const Center(
                child: CircularProgressIndicator(),
              );
            }

            final item = RestaurantDetailModel.fromJson(json: snapshot.data!);
            return CustomScrollView(
              slivers: [
                renderTop(
                  model: item,
                ),
                renderLabel(),
                renderProducts(
                  products : item.products,
                ),
              ],
            );
          },
          future: getRestaurantDetail(),
        )
    );
  }

  SliverPadding renderLabel(){
    return SliverPadding(
      padding: EdgeInsets.symmetric(horizontal: 16.0),
      sliver: SliverToBoxAdapter(
        child: Text(
          '메뉴',
          style: TextStyle(
            fontSize: 18.0,
            fontWeight: FontWeight.w500,
          ),
        ),
      ),
    );
  }
  SliverPadding renderProducts({
    required List<RestaurantProductModel> products,
  }) {
    return SliverPadding(
      padding: EdgeInsets.symmetric(horizontal: 16.0),
      sliver: SliverList(
        delegate: SliverChildBuilderDelegate(
              (context, index) {
            final model = products[index];

            return Padding(
              padding: const EdgeInsets.only(top: 16.0),
              child: ProductCard.fromModel(model: model),
            );
          },
          childCount: products.length,
        ),
      ),
    );
  }
  SliverToBoxAdapter renderTop({
    required RestaurantDetailModel model,
}){
    return SliverToBoxAdapter(
      child: RestaurantCard.fromModel(model: model, isDetail:true,)
    );
  }
}


