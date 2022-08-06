import 'package:flutter/material.dart';

import 'pages/homepage/index.dart';

void main() => runApp(const LogbookApp());

class NoAnimationPageTransitionsBuilder extends PageTransitionsBuilder {
  const NoAnimationPageTransitionsBuilder();

  @override
  Widget buildTransitions<T>(
    PageRoute<T> route,
    BuildContext context,
    Animation<double> animation,
    Animation<double> secondaryAnimation,
    Widget child,
  ) {
    return child;
  }
}

const pageTransitionsTheme = PageTransitionsTheme(
  builders: <TargetPlatform, PageTransitionsBuilder>{
    TargetPlatform.linux: NoAnimationPageTransitionsBuilder(),
    TargetPlatform.macOS: NoAnimationPageTransitionsBuilder(),
  },
);

class LogbookApp extends StatelessWidget {
  const LogbookApp({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      home: const Homepage(),
      debugShowCheckedModeBanner: false,
      theme: ThemeData.light()
          .copyWith(pageTransitionsTheme: pageTransitionsTheme),
    );
  }
}
