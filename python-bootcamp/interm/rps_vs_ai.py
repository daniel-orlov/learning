from random import randint
pl_wins = 0
pc_wins = 0
w_score = input('\n\tWe play till: \n\n\t(enter number of wins) ')
w_score = int(w_score)
while w_score <= 0:
    w_score = input(
        '\n\tWe play till: \n\n\t(enter number of wins greater than zero) ')
    w_score = int(w_score)
while pl_wins < w_score and pc_wins < w_score:
    print(f'\n\t\tSCORE \n\n\t    YOU {pl_wins} : {pc_wins} PC')
    # input('Welcome to Rock, Paper, Scissors! Press Enter to start')
    print('\nChoose your weapon: rock, paper or scissors (r/p/s) \n\nTo exit type \'e\'\n')
    pl1 = input().lower()
    if pl1 == 'e':
        break
    elif pl1:
        input('\n\tPress Enter, to make computer choose')
        pc_move = randint(0, 2)
        if pc_move == 0:
            pl2 = 'r'
            print('\n\tComputer plays ROCK')
        elif pc_move == 1:
            pl2 = 'p'
            print('\n\tComputer plays PAPER')
        else:
            pl2 = 's'
            print('\n\tComputer plays SCISSORS')
    else:
        print('\n\tChoose your weapon, Player')
    if pl1 == pl2:
        print('\n\tDraw! Do it again!')
    elif pl1 == 's' and pl2 == 'p':
        print('\n\tYou win!')
        pl_wins += 1
    elif pl1 == 'p' and pl2 == 'r':
        print('\n\tYou win!')
        pl_wins += 1
    elif pl1 == 'r' and pl2 == 's':
        print('\n\tYou win!')
        pl_wins += 1
    else:
        print('\n\tComputer wins =)')
        pc_wins += 1
print(f'\n\n\t    FINAL  SCORE \n\n\t    YOU {pl_wins} : {pc_wins} PC\n')
