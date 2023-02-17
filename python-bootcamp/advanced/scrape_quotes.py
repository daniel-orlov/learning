from bs4 import BeautifulSoup
import requests
from csv import DictWriter

url = 'http://quotes.toscrape.com'

def scrape_quotes(url):
    all_quotes = []
    response = requests.get(url)
    soup = BeautifulSoup(response.text, 'html.parser')
    quotes = soup.find_all(class_='quote')
    page_num = 1
    while soup.find(class_='next'):
        next_page_link = '/page/' + str(page_num) + '/'
        new_page = url + next_page_link
        new_resp = requests.get(new_page)
        soup = BeautifulSoup(new_resp.text, 'html.parser')
        new_q = soup.find_all(class_='quote')
        for q in new_q:
            all_quotes.append({
                'text': q.find(class_='text').get_text(),
                'author': q.find(class_='author').get_text(),
                'bio': q.find('a')['href']
            })
        page_num += 1
    return all_quotes

def write_quotes(quotes):
    with open('quotes_db.csv', 'w', encoding='utf-8') as file:
        headers = ['text', 'author', 'bio']
        csv_writer = DictWriter(file, fieldnames=headers)
        csv_writer.writeheader()
        for quote in quotes:
            csv_writer.writerow(quote)

quotes = scrape_quotes(url)
write_quotes(quotes)
print('Scraping DONE!')
