with open('words.txt') as file:
    fin = file.read()
    fin = fin.split('\n')
'''
def main(obj, func, letter=None):
    result = []
    for word in obj:
        if letter:
            for letter in letters:
                if func(word, letter) and not word in result:
                    result.append(word)
        else:
            if func(word) and not word in result:
                result.append(word)
    return result

def twentyplus(word):
    if len(word) > 19:
        return True
    return False

def has_no_letter(word, letter):
    if not letter in word:
        return True
    return False

def percentage(lst):
    global fin
    total = len(lst)
    per100 = (tot/len (fin))*100
    return f'Total words:\n{total} out of {len (fin)}, or {per100}%\n\n'

def has_no_letter2(word, letter):
    global fin
    no_letter_words = []
    abc = 'abcdefghijklmnopqrstuvwxyz'
    result = dict()
    for letter in abc:
        for word in fin:
            if not letter in word:
                no_letter_words.append(word)
        result.update({letter: len(no_letter_words)})
        no_letter_words.clear()
    return result

def avoids(word, letter):
    if letter in word:
        return False
    return True

def uses(word, letters):
    if letter in word:
         return True
    return False

def uses_all(file_name):
    used = input('Choose the obligatory letters: ')
    global fin
    result = fin.copy()
    for letter in used:
        for word in result:
            try:
                if not letter in word:
                    result.remove(word)
            except ValueError:
                continue
    return result

# print(uses_all('words.txt'))
'''


def in_bisect(fin, word):
    found = False
    half = fin[:(len(fin)//2)]
    scnd_half = fin[(len(fin)//2):]
    while not found:
        if word in half:
            return fin.index(word)
        else:
            print('here we go again')
            half = scnd_half[:(len(scnd_half)//2)]
            scnd_half = scnd_half[(len(scnd_half)//2):]

# print(in_bisect(fin, 'zymurgy'))


def reversi(fin):
    pairs = []
    for item in fin:
        meti = item[::-1]
        if meti in fin:  # 886 pairs w/o 2nd condition
            pairs.append((item, meti))
    return f'{len(pairs)} pairs found. Here they are:\n\n{pairs}'

# print(reversi(fin))


def fibonacci(n):
    known = {0: 0, 1: 1}
    if n in known:
        return known[n]
    res = fibonacci(n-1)+fibonacci(n-2)
    known[n] = res
    return res

# print(fibonacci(47))


def in_dict(fin, word):
    d = {}
    fin_d = d.fromkeys(fin, 0)
    if word in fin_d:
        return True
    else:
        return False

# print(in_dict(fin, 'volvo'))


def homophones(fin):
    res = []
    d = {}
    fin_d = d.fromkeys(fin, 0)
    for word in fin:
        if len(word) == 5 and word in fin_d and \
                word[1:] in fin_d and (word[:1]+word[2:]) in fin_d:
            res.append((word, word[1:], (word[:1]+word[2:])))
    return res

# res = homophones(fin)

# with open('result.txt', 'a') as f:
#     for item in res:
#         f.write(str(item) + '\n')


def most_frequent(fin):
    lettercount = []
    item = None
    text = str(fin)
    for char in text:
        item = text.count(char), char
        if item not in lettercount:
            lettercount.append(item)
    sortedlettercount = sorted(lettercount, reverse=True)
    # return dict(sortedlettercount)
    return sortedlettercount

# print(most_frequent(fin))


def signature(s):
    '''Returns the signature of this string.

    Signature is a string that contains all of the letters in order.

    s: string
    '''
    return ''.join(sorted(list(s)))


def signature2(s):
    return ''.join(sorted(list(s)))


text = 'arette'

# print(signature(text))
# print(signature2(text))


from string import punctuation as punct
import re
# def prepare_text(file_name):
#     with open(file_name) as f:
#         fin = f.read()
#         fin = fin.split(' ')
#         for word in fin:
#             for char in punct:
#                 if word.endswith(char):
#                     word = word[:-1]


def prepare_text(file_name):
    '''Takes file as input,
    gets rid of newlines and punctuation, converts into lower case and
    returns a list.
    '''
    regex = re.compile('[%s]' % re.escape(punct))
    with open(file_name) as f:
        fin = f.read()
        #Getting rid of the newline characters
        fin = ' '.join(fin.split('\n'))
        #Getting rid of the punctuation
        fin = regex.sub('', fin)
        #Converting into list with lower case
        fin = (fin.lower()).split(' ')
        fin.pop()

        return fin

def total_words(collection):
    '''Takes collection,
    counts the total number of words,
    returns an int.
    '''
    return len(collection)

def unique_words(prepared_book):
    '''Takes book processed with prepare_text,
    returns a set of unique words.
    '''
    return set(prepared_book)

def word_frequency(prepared_book, n=20, rarest=False):
    '''Takes a prepared_book (list) processed with prepare_text,
    maps unique words to their frequency,
    returns top-n used words as a list or bottom-n, if flag rarest.
    '''
    #Prepare a tuple of keys for dict, sorted in abc order
    keys = tuple(sorted(set(prepared_book)))
    #Create a dictionary
    hist = {}
    for word in keys:
        hist[word] = hist.get(word, 0) + 1

    return hist


the_prince = prepare_text('il-principe.txt')
print(words(the_prince))
print(unique_words(the_prince))
# print(unique_words_usage(the_prince, 25))
