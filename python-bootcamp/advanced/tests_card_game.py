import unittest
from python_works.card_game import Card, Deck


class CardTest(unittest.TestCase):
    """this contains tests for Card class"""

    def setUp(self):
        self.card = Card('Diamonds', 'King')

    def test_init(self):
        """each card should have a suit and a value"""
        self.assertEqual(self.card.suit, 'Diamonds')
        self.assertEqual(self.card.value, 'King')

    def test_repr(self):
        """when printed card should indicate its value then suit"""
        self.assertEqual(repr(self.card), 'King of Diamonds')


class DeckTest(unittest.TestCase):
    """this contains tests for Deck class"""
    def setUp(self):
        self.deck = Deck()

    def test_init(self):
        """each deck should be a list of 52 items (cards)"""
        self.assertEqual(len(self.deck.cards), 52)
        self.assertTrue(isinstance(self.deck.cards, list))

    def test_repr(self):
        """deck should be repr by the number of cards it holds"""
        self.assertEqual(repr(self.deck), 'Deck of 52 cards')

    def test_count(self):
        """counts the number of cards in the deck"""
        self.assertEqual(self.deck.count(), 52)
        self.deck.cards.pop()
        self.assertTrue(self.deck.count(), 51)

    def test__deal_for_accuracy(self):
        """_deal should deal cards into hands unless deck is empty"""
        cards = self.deck._deal(7)
        self.assertEqual(len(cards), 7)
        self.assertEqual(self.deck.count(), 45)

    def test__deal_no_cards(self):
        """should raise ValueError when trying to deal from empty deck"""
        self.deck._deal(self.deck.count())
        with self.assertRaises(ValueError):
            self.deck._deal(1)

    def test_deal_card(self):
        """should deal 1 card"""
        self.assertNotIn(self.deck.deal_card(), self.deck.cards)
        self.assertEqual(len(self.deck.cards), 51)

    def test_deal_hand(self):
        """should deal hand of n>2 cards"""
        self.assertNotIn(self.deck.deal_hand()[0], self.deck.cards)
        self.assertEqual(len(self.deck.cards), 52 - len(self.deck.deal_hand()))

    def test_shuffle(self):
        """shuffle should only work on full deck and randomize it"""
        self.deck.shuffle()
        self.assertEqual(len(self.deck.cards), 52)
        with self.assertRaises(ValueError):
            self.deck.deal_hand(3)
            self.deck.shuffle()

if __name__ == '__main__':
    unittest.main()
