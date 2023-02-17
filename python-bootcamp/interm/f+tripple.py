def triple_and_filter(lst):
    return list(map(lambda y: y*3, filter(lambda x: x % 4 == 0, lst)))
