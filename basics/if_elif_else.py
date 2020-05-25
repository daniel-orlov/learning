#Practicing nested conditionals

age = input('How old are you?')

if age:
	age = int(age)
	if age >= 18 and age < 21:
		print('You can enter, but need to wear a wristband')
	elif age > 21:
		print('You are good to enter and drink, if you want')
	else:
		print('No enter, sorry. Grow up first!')
else:
	print('Please enter your age.')