from random import shuffle
# import pdb; pdb.set_trace()


class Card:

    def __init__(self, suit, value):
        self.suit = suit
        self.value = value

    def __repr__(self):
        return f'{self.value} of {self.suit}'


class Deck:

    def __init__(self):
        suits = ['Hearts', 'Diamonds', 'Clubs', 'Spades']
        values = ['A', '2', '3', '4', '5', '6', '7', '8', '9', '10', 'J', 'Q', 'K']
        self.cards = [Card(suit, value) for suit in suits for value in values]

    def __repr__(self):
        return f'Deck of {self.count()} cards'

    def __iter__(self):
        return iter(self.cards)

    def count(self):
        return len(self.cards)

    def _deal(self, num=1):
        count = self.count()
        turns = 0
        hand = []
        while turns < num:
            if num <= count:
                hand.append(self.cards.pop())
                turns += 1
            else:
                raise ValueError('All cards have been dealt')
        return(hand)

    def deal_card(self):
        return self._deal()[0]

    def deal_hand(self, hand=2):
        return self._deal(hand)

    def shuffle(self):
        if self.count() == 52:
            shuffle(self.cards)
            return self
        else:
            raise ValueError('Only full decks can be shuffled')

# d = Deck()
# print(d)
# d.shuffle()
# card = d.deal_card()
# print(card)
# hand = d.deal_hand(7)
# print(hand)
# print(d)

# for c in d:
#     print(c)
# # card1 = Card()
# # print(card1)

# # my_deck = Deck()
# # print(my_deck)
# # my_deck.shuffle()
# # card = my_deck.deal_card()
# # print(card)
# # hand = my_deck.deal_hand(5)
# # prit(hand)
# # print(len(my_deck))
