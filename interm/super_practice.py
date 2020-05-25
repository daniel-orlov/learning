class Animal:

    def __init__(self, name, species):
        self.name = name
        self.species = species

    def __repr__(self):
        return f'{self.name} is a {self.species}'

    def make_sound(self, sound):
        print(f'This animal says {sound.upper()}!')

class Cat(Animal):

    def __init__(self, name, species, breed, toy):
        # self.name = name
        # self.species = species
        super().__init__(name, species = 'Cat')
        self.breed = breed
        self.toy = toy

    def play(self):
        print (f'{self.name} plays with {self.toy}')

haski = Cat('Haski', 'Cat', 'Red Super-Intelegent', 'Dad\'s tummy')
print(haski)
haski.play()
