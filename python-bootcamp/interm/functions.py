#practicing functions

def nested_sum(a_list):
    total = 0
    for elem in a_list:
        total += sum(elem)
    return total

t = [[123,45,678], [76,56,123], [54,11,45]]
t1 = [1,2,3,4,5,6,9,7,8,9,8,4,2]
# print(nested_sum(t))

def middle(a_list):
    return a_list[1:-1]

# print(middle(t1))

def chop(a_list):
    a_list.pop()
    a_list.pop(0)

# print(chop(t1))
# print(t1)


