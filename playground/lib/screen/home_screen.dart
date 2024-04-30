import 'package:flutter/material.dart';
import 'package:playground/layout/default_layout.dart';

class HomeScreen extends StatelessWidget {
  const HomeScreen({super.key});

  @override
  Widget build(BuildContext context) {
      return DefaultLayout(
          body: ListView(
            children: [
              ElevatedButton(onPressed: () {}, child: Text('SateProviderScreen'))
            ],
          ),
          title: 'HomeScreen',
      );
  }
}
