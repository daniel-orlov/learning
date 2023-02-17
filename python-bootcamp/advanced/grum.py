class SociopathDict(dict):

    def __repr__(self):
        print('None of your business, leave me alone!')
        return super().__repr__()

    def __missing__(self, key):
        print(f'You want {key}? Well, it aint here! Now piss off!')

    def __setitem__(self, key, value):
        print('Why do you have to change things all the time?')
        print('Oh, fine...')
        super().__setitem__(key, value)

new_dict = SociopathDict({'first': 'Sherlock', 'last': 'Holmes'})
print(new_dict)
print(new_dict['city'])
new_dict['city'] = 'London'
print(new_dict)
