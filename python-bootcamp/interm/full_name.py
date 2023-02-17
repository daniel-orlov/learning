def extract_full_name(lst):
    return list(map(lambda val: f"{val['first']} {val['last']}", lst))
    # lst1 = [x for x in lst]
    # return {:v.upper() for k,v in lst1}
    # (k:v for k,v in x)
    # list(zip(*lst)


names = [{'first': 'Elie', 'last': 'Schoppik'},
         {'first': 'Colt', 'last': 'Steele'}]
print(extract_full_name(names))  # ['Elie Schoppik', 'Colt Steele']
