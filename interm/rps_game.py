input('Welcome to Rock, Paper, Scissors! Press Enter to start')
print('Player one, choose your weapon: rock, paper or scissors')
pl1 = input()
print('*** NO CHEATING! *** \n\n\n' * 20)
print('Player two, choose your weapon: rock, paper or scissors')
pl2 = input()
if pl1 and pl2:
    if pl1 == pl2:
        print('Draw! Do it again!')
    elif pl1 == 'scissors' and pl2 == 'paper':
        print('Player one wins')
    elif pl1 == 'paper' and pl2 == 'rock':
        print('Player one wins')
    elif pl1 == 'rock' and pl2 == 'scissors':
        print('Player one wins')
    else:
        print('Player two wins')
else:
    print('Choose your weapons')
