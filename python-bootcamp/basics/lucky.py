for x in range(1,22):
	if x == 12 or x == 21:
		status = 'lucky'
	elif x % 2 == 0:
		status = 'even'
	else:
		status = 'odd'
	print(f'{x} is {status}')