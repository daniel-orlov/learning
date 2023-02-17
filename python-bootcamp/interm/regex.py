import re


source = 'My credit card number is 4444 5458 5986 3526.'

pattern = re.compile(r'(\b\d{4}\b)')
res = pattern.findall(source)
# match = res.group()
# print(match)
print(res)

import re

#define parse_date below
def parse_date(input):
    pattern = re.compile('^(\d\d)[,/.](\d\d)[,/.](\d{4})$')
    # pattern = re.compile('^(\d{2})[.,/](\d{2})[.,/](\d{4})$')
    res = pattern.search(input)
    day = res.group(1)
    month = res.group(2)
    year = res.group(3)
    if res:
        return {'d':day, 'm':month, 'y':year}
    else:
        return None
