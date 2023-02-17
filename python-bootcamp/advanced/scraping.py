import requests
from bs4 import BeautifulSoup
from csv import writer

# GRAB TITLE, LINK AND DATE
# STORE IN CSV FILE

url = 'https://www.rithmschool.com/blog'
response = requests.get(url)
soup = BeautifulSoup(response.text, 'html.parser')
articles = soup.find_all('article')


    # print(x.find('a').get_text())
    # print(url[:-5:] + (x.find('a')['href']))
    # print(x.find('time')['datetime'])
    # print('\n')


# def tit_lin_da(x):
#     global url
#     a_tag = x.find('a')
#     title = a_tag.get_text()
#     link = url[:-5:] + (a_tag['href'])
#     date = x.find('time')['datetime']
#     l = [title, link, date]
#     return l

with open('scraping_results.csv', 'w') as file:
    csv_writer = writer(file)
    csv_writer.writerow(['title', 'link', 'date & time'])
    for item in articles:
        a_tag = item.find('a')
        title = a_tag.get_text()
        link = url[:-5:] + (a_tag['href'])
        date = item.find('time')['datetime']
        l = [title, link, date]
        csv_writer.writerow(l)


# titles = soup.article.find_all('a')
# # dates = articles.find('[datetime]')

# .get_text()

# print(articles)
