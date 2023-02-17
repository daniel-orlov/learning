from requests import get
url = 'https://icanhazdadjoke.com/'
jokes = []
rep_count = 0

while len(jokes) < 602:
    res = get(url, headers={'Accept': 'text/plain'})
    if res.text not in jokes:
        jokes.append(res.text)

# print(len(jokes)) #568 when 1000, real total = 602
saving = input('save jokes? y/n')
if saving == 'y':
    with open('jokes.txt', 'a') as f:
        f.write(str(jokes))
