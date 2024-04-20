import 'package:Todopoint/common/const/colors.dart';
import 'package:Todopoint/common/layout/default_layout.dart';
import 'package:Todopoint/task/view/task_screen.dart';
import 'package:flutter/material.dart';

class RootTab extends StatefulWidget {
  const RootTab({super.key});

  @override
  State<RootTab> createState() => _RootTabState();
}

class _RootTabState extends State<RootTab> with SingleTickerProviderStateMixin{
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
      title: 'todo',
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
            icon: Icon(Icons.check),
            label: 'Todo'
          ),
          BottomNavigationBarItem(
              icon: Icon(Icons.calendar_today_rounded),
              label: 'Calendar'
          ),
          BottomNavigationBarItem(
              icon: Icon(Icons.person_outline),
              label: 'Community'
          ),
          BottomNavigationBarItem(
              icon: Icon(Icons.settings),
              label: 'Settings'
          ),
        ],
      ),
      child: TabBarView(
        physics: const NeverScrollableScrollPhysics(),
        controller: controller,
        children: [
          TaskScreen(),
          Container(child:Text('Helloworld')),
          Container(child:Text('Helloworld')),
         Container(child:Text('Helloworld')),
        ],
      ),
    );
  }
}
