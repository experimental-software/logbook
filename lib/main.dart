import 'package:flutter/material.dart';
import 'package:logbook/pages/details/index.dart';
import 'package:routemaster/routemaster.dart';

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

final routes = RouteMap(
  routes: {
    '/': (_) => const MaterialPage(child: Homepage()),
    '/log-entry/:dir': (info) {
      return MaterialPage(
        child: AsyncDetailsPage(encodedDir: info.pathParameters['dir']!),
      );
    },
  },
);

class LogbookApp extends StatelessWidget {
  const LogbookApp({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return MaterialApp.router(
      routerDelegate: RoutemasterDelegate(routesBuilder: (context) => routes),
      routeInformationParser: const RoutemasterParser(),
      debugShowCheckedModeBanner: false,
      theme: ThemeData.light().copyWith(
        pageTransitionsTheme: pageTransitionsTheme
      ),
    );
  }
}
