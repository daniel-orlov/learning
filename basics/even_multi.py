
def multiply_even_numbers(the_list):
    result = 1
    for x in the_list:
        if x % 2 == 0:
            result *= x
    return result

print(multiply_even_numbers([1, 2, 3, 4]))
