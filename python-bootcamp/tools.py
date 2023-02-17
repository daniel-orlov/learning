#get rid of all the punctuation in string with REGEX
import re, string
regex = re.compile('[%s]' % re.escape(string.punctuation))
regex.sub('', s)
#get rid of all the punctuation in string with set
exclude = set(string.punctuation)
''.join(ch for ch in s if ch not in exclude)


def prepare_text(file_name):
    '''Takes file as input,
    gets rid of newlines and punctuation, converts into lower case and
    returns a list.
    '''
    import re, string
    regex = re.compile('[%s]' % re.escape(string.punctuation))
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
