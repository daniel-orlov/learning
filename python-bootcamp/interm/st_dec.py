from time import time
from functools import wraps


def speed_test(func):
    @wraps(func)
    def wrapper(*args, **kwargs):
        start = time()
        result = func(*args, **kwargs)
        end = time()
        print(f'Time Elapsed: {end - start}')
        return result
    return wrapper


@speed_test
def sum_num():
    return (sum(x for x in range(10000)))

print(sum_num())
