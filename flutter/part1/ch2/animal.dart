class Animal {
  String name;
  String sound;

  Animal(this.name, this.sound);

  void makeSound() {
    print('The $name says $sound');
  }
}

class Dog extends Animal {
  Dog(String name) : super(name, 'Woof');
}