# def product(a, b):
#     """
#     >>> product(1, 2)
#     2
#     >>> product(1, 2.0)
#     2.0
#     """
#     return a * b

# def return_day(day):
#     """
#     >>> return_day(1)
#     'Sunday'
#     >>> return_day(0)

#     """
#     week = {
#         1: 'Sunday',
#         2: 'Monday',
#         3: 'Tuesday',
#         4: 'Wednesday',
#         5: 'Thursday',
#         6: 'Friday',
#         7: 'Saturday',
#     }
#     if day in range(1, 8):
#         return week.get(day)
#     return None


# def last_element(list_here):
#     """
#     >>> last_element([])

#     >>> last_element([1, 2, 3, 4])
#     4
#     >>> last_element(['a', 'b'])
#     'b'
#     >>> last_element([True])
#     True
#     >>> last_element('some string')
#     'g'
#     """
#     elements = list(list_here)
#     if len(elements) > 0:
#         return elements.pop()
#     return None

def frequency(the_list, val):
    """
    >>> frequency([], None)
    0
    >>> frequency([1, 1, 2, 3, 5], 1)
    2
    >>> frequency('string here', 0)
    Traceback (most recent call last):
    TypeError: must be str, not int
    """
    return the_list.count(val)
