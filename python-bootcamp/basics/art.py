# Creating cute art with a combination of while and for loop

# 1-9
# u"\U0001F60D"

# for times in range(10):
# 	print('<3' * times)

times = 1
while times < 11:
    for reps in range(8, 0, -1):
        print('__' * reps + '<3' * times + '__' * reps)
        times += 2
