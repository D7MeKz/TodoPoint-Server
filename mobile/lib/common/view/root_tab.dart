import 'package:app/common/const/colors.dart';
import 'package:app/common/layout/default_layout.dart';
import 'package:app/restaurant/view/restaurant_screen.dart';
import 'package:flutter/material.dart';


class RootTab extends StatefulWidget {
  const RootTab({super.key});

  @override
  State<RootTab> createState() => _RootTabState();
}

class _RootTabState extends State<RootTab> with SingleTickerProviderStateMixin {
  int index = 0;
  late TabController controller; // 나중에 이 값이 입력될거야.. 하지만 무조건 값을 넣을거야.!

  @override
  void initState() {
    super.initState();
    // vsync ->state
    // this는 특정 기능을 가지고 있어야 한다.
    controller = TabController(length: 4, vsync: this);
    controller.addListener(tabListener);
  }

  @override
  void dispose() {
    controller.removeListener(tabListener);
    super.dispose();
  }

  void tabListener(){
    setState(() {
      index = controller.index;
    });
  }

  @override
  Widget build(BuildContext context) {
    return DefaultLayout(
      title: 'Todo',
      bottomNavigationBar: BottomNavigationBar(
        selectedItemColor: PRIMARY_COLOR,
        unselectedItemColor: BODY_TEXT_COLOR,
        selectedFontSize: 10,
        unselectedFontSize: 10,
        type: BottomNavigationBarType.fixed,
        onTap: (int index){
          controller.animateTo(index);
        },
        currentIndex: index,
        items: const [
          BottomNavigationBarItem(
              icon: Icon(Icons.home_outlined),
              label:'Home'
          ),
          BottomNavigationBarItem(
              icon: Icon(Icons.fastfood_outlined),
              label:'Food'
          ),
          BottomNavigationBarItem(
              icon: Icon(Icons.receipt_long_outlined),
              label:'Receipt'
          ),
          BottomNavigationBarItem(
              icon: Icon(Icons.person_outlined),
              label:'Profile',
          ),
        ],
      ),
      child: TabBarView(
        physics: const NeverScrollableScrollPhysics(),
        controller: controller,
        children: [
          RestaurantScreen(),
          Container(child: Text('Food'),),
          Container(child: Text('Order'),),
          Container(child: Text('Profile'),),
        ],
      ),
    );
  }
}

