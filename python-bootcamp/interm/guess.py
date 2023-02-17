import random

random_num = random.randint(1,10)
guess = None

while guess != random_num:
	guess = input('Pick a number from 1 to 10: ')
	guess = int(guess)
	# if guess != '':
	if guess < random_num:
		print('TOO LOW!')
	elif guess > random_num:
		print('TOO HIGH!')
	else:
		print('YOU WIN!')
		replay = input('PLAY AGAIN? (y/n) ')
		if replay == 'y':
			random_num = random.randint(1,10)
			guess = None
		else:
			print('GG! BIE!')
			break 			
	# else:
	# 	print('You have not picked a number!')
# print(random_num)