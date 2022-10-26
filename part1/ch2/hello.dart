void main() {
  List<String> namesList = ['Seth', 'Kathy', 'Lars'];
  print(namesList);

  for (String name in namesList) {
    hello(name);
  }
}

void hello(String name) {
  print("Hello, $name!");
}