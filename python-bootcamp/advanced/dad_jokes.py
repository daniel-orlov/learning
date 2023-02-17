from requests import get
from pyfiglet import figlet_format as fig
from termcolor import colored as color

header = color(fig("DAD JOKES 2020"), 'yellow')
print(header)

url = 'https://icanhazdadjoke.com/search'

term = input('Let me tell you a joke! Give me topic, please: ')

res = get(url,
          headers={'Accept': 'application/json'},
          params={"term": term}
          ).json()

res_num = res['total_jokes']
results = res['results']
desired_num_jokes = 0
displayed_jokes = 0
result_num = 0
if res_num == 0:
    print(f'Sorry, I have no jokes about {term}. Please, try a different topic.')
elif res_num == 1:
    print(f'I have one joke about {term}. Here it is:')
    print(results[0]['joke'])
else:
    desired_num_jokes = input(f'I have {res_num} jokes about {term}.\nHow many do you want? MAX=20 ')
    if int(res_num) < int(desired_num_jokes):
        print("Sorry, I don't have that many jokes")
    else:
        while displayed_jokes < int(desired_num_jokes):
            print(results[result_num]['joke'] + '\n')
            displayed_jokes += 1
            result_num += 1
