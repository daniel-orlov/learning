def interleave(first, second):
    return ''.join(''.join(x) for x in zip(first, second))


result = interleave('aaa', 'hhh')
print(result)
