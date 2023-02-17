from copy import copy


class Human:
    def __init__(self, first, last, age):
        self.first = first
        self.last = last
        self.age = age

    def __repr__(self):
        retutf'Human named {self.first} {self.last} aged {self.age}'

    def __len__(self):
        return self.age

    def __add__(self, other):
        if isinstance(other, Human):
            return Human(first='Newborn', last=self.last, age=0)
        return "You can't add that!"

    def __mul__(self, other):
        if isinstance(other, int):
            return [copy(self) for x in range(other)]
        return "Can't multiply that!"


j = Human('Jenny', 'Larsson', 47)
k = Human('Karl', 'Mikaelson', 49)

print(j)
print(k)
print(j + k)
triplets = k * 3
triplets[2].first = 'Jessica'
print(triplets)
