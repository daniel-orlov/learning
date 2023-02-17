def return_evens():
    return [x for x in range(1,50) if x % 2 == 0]
print(return_evens())

def return_odds():
    x = 0
    odds = []
    for x in range(1,50):
        if x % 2 != 0:
            odds.append(x)
    return odds

print(return_odds())
