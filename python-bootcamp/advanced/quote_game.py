from bs4 import BeautifulSoup
import requests
from random import choice
from csv import DictReader

def read_quotes(filename):
    with open(filename, encoding='utf-8') as file:
        csv_reader = DictReader(file)
        return list(csv_reader)

quotes = read_quotes('quotes_db.csv')

def play(quotes):
    chosen_quote = choice(quotes)
    chosen_quote_text = chosen_quote['text']
    author = chosen_quote['author']
    def get_hints(author_bio_link):
        url = 'http://quotes.toscrape.com'
        page = url + author_bio_link
        resp = requests.get(page)
        soup = BeautifulSoup(resp.text, 'html.parser')
        date = soup.find(class_='author-born-date').get_text()
        place = soup.find(class_='author-born-location').get_text()
        name = soup.find(class_='author-title').get_text()
        initials = ''
        for c in name:
            if c in 'ABCDEFGHIJKLMNOPQRSTUVWXYZ':
                initials += c
        return [date, place, initials]
    hints = get_hints(chosen_quote['bio'])
    print(f'\n\nLet\'s play a game - Guess the Author of a Quote! \n\nHere is one:\n\n{chosen_quote_text}')
    num_attempts = 4
    num_hints = 0
    guess = input(f'\nWho do you think said that? You have {num_attempts} attempts left to guess: ')
    while num_attempts > 0:
        if guess.lower() == author.lower():
            print('\nGood guess. Congrats! You won!')
            break
        else:
            num_attempts -= 1
            if num_attempts == 3:
                hint = f"The author of this quote was born on {hints[0]}."
                guess = input(f'\n\nNot exactly. Here is a hint:\n\n{hint}\n\nTry again. You have {num_attempts} atempts left ')
            elif num_attempts == 2:
                hint = f"The author of this quote was born {hints[1]}."
                guess = input(f'\n\nNot quite. Here is a hint:\n\n{hint}\n\nTry again: ')
            elif num_attempts == 1:
                hint = f"The author's initials are {hints[2]}."
                guess = input(f'\n\nNo, sorry. Here is a hint:\n\n{hint}\n\nTry again: ')
            else:
                print(f'It was {author}. Better luck next time!')
                break

play(quotes)
while True:
    again = ''
    while again.lower() not in ('y', 'n', 'yes', 'no'):
        again = input('\n\nPlay again? (y/n)')
    if again.lower() in ('y', 'yes'):
        play(quotes)
    else:
        print('\n\nGood game. Bye!')
        break
