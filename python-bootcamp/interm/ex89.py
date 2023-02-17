from functools import wraps


def show_args(func):
    @wraps(func)
    def wrapper(*args, **kwargs):
        # args = tuple(args)
        # kwargs = dict(kwargs)
        print('Here are the args: {}'.format(args))
        print('Here are the kwargs: {}'.format(kwargs))
        return func(*args, **kwargs)
    return wrapper


@show_args
def do_nothing(*args, **kwargs):
    pass


do_nothing(1, 2, 3, a="hi", b="bye")

# Should print (on two lines):
# Here are the args: (1, 2, 3)
# Here are the kwargs: {"a": "hi", "b": "bye"}
