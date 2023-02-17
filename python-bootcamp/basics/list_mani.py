# def list_manipulation(a_list, command, loc, val=None):
#     actual_list = list(actual_list)
#     index = 0
#     loc = str(loc)
#     if loc == "end":
#         index = -1
#     else:
#         index = 0
#     if command == "add":
#         return actual_list.insert(index, val)
#     else:
#         popped = actual_list.pop(index)
#         return popped

#     if c == 'remove':
#         return l.pop() if loc == 'end' else l.pop(0)

#     return [val] + l if loc == 'beginning' else l + [val]


print(list_manipulation([1,2,3], "add", "beginning", 20))

