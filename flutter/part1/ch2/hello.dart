import 'dart:io';

void main() {
  stdout.writeln('Who would you like to say hello to?');
  String? name = stdin.readLineSync();
  return hello(name!);
}

void hello(String name) {
  print("Hello, $name!");
}

/// Adding a one-liner version of the function 'hello'
helloOneLiner(String name) => print('Hello, $name!');