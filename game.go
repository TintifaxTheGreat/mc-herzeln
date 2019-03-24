package main

type Agent uint8

const (
	AGENT_HUMAN Agent = iota + 1
	AGENT_RANDOM
	AGENT_MCTS
)

const (
	NOT_DROPPED = 0
	ON_TABLE    = 1
	DROPPED     = 2
	IN_HAND     = 0
	IN_TRICKS   = 1
)

type Game struct {
	players     int
	colors      int
	figures     int
	cardsInHand int
	agent       [4]Agent
	//cardpool	*tensor.Dense
	cardpool [4][8][3]uint8
}

func NewGame() *Game {
	p := new(Game)
	p.players = 4
	p.colors = 4
	p.figures = 8
	p.cardsInHand = 8
	p.agent = [4]Agent{AGENT_HUMAN, AGENT_RANDOM, AGENT_RANDOM, AGENT_RANDOM}

	// create the cardpool
	p.cardpool = [4][8][3]uint8{}
	//p.cardpool = tensor.New(tensor.WithShape(p.colors, p.figures, 3), tensor.Of(tensor.Uint8))

	/*

		self.cardpool = np.zeros((self.COLORS, self.FIGURES, 3), dtype=bool)
		self.cardpool[:, :, NOT_DROPPED] = 1

	*/

	return p
}

/*

import numpy as np
from .helpers import string_to_card, card_to_string, show_cards
from random import randint, seed
from typing import Tuple, List




class Game:
    """
    a gard game
    """

    def __init__(self):

        seed()

        # create the cardpool
        self.cardpool = np.zeros((self.COLORS, self.FIGURES, 3), dtype=bool)
        self.cardpool[:, :, NOT_DROPPED] = 1

        # create the player's cards
        self.cards = np.zeros((self.COLORS, self.FIGURES, self.PLAYERS, 2), dtype=bool)

        # create helper arrays for all colors
        self.ALL_COLORS = np.zeros((self.COLORS, self.FIGURES, self.COLORS), dtype=bool)
        for color in range(self.COLORS):
            self.ALL_COLORS[color, :, color] = 1

        # create helper arrays for all figures
        self.ALL_FIGURS = np.zeros((self.COLORS, self.FIGURES, self.FIGURES), dtype=bool)
        for figur in range(self.FIGURES):
            self.ALL_FIGURS[:, figur, figur] = 1

        # negative points are given to all hearts in tricks
        self.POINTS = self.ALL_COLORS[:, :, 0]  # TODO change this later

        self.myAgent = []
        for agent in self.AGENT:
            print('agent ', agent)
            if agent == AGENT_RANDOM:
                self.myAgent.append(self.random_agent)
            elif agent == AGENT_HUMAN:
                self.myAgent.append(self.human_agent)
            else:
                raise NotImplementedError

    def start(self):
        """
        the actual game
        """
        self.deal()
        self.play()
        points = self.evaluate()
        for i, p in enumerate(points):
            print('Spieler ', i + 1, ': ', p, ' Punkte')

    def random_agent(self, player_id: int, is_lead: bool, lead_card: List[int]) -> Tuple[List[int], bool]:

        if not is_lead:
            print('Ausspiel: ', card_to_string(lead_card))

        if is_lead:
            return self.random(self.cards[:, :, player_id, IN_HAND]), False

        legal_cards, followed_suit = self.legal_cards(player_id, lead_card)
        print('followed suit: ', followed_suit)
        return self.random(legal_cards), followed_suit

    def human_agent(self, player_id: int, is_lead: bool, lead_card: List[int]) -> Tuple[List[int], bool]:

        if not is_lead:
            print('Ausspiel: ', card_to_string(lead_card))

        card_string = show_cards(self.cards[:, :, player_id, IN_HAND])
        print(card_string)

        followed_suit = 0

        if is_lead:
            legal_card_string = card_string

        else:
            legal_cards, followed_suit = self.legal_cards(player_id, lead_card)
            legal_card_string = show_cards(legal_cards)

        response_string = '$'
        while not (response_string in legal_card_string):
            response_string = input('Please choose:  ')

        return string_to_card(response_string), followed_suit


    def play(self):
        """
        play until no cards remain
        """
        lead_player = 0

        for trick in range(self.CARDS_IN_HAND):
            print('###############')
            print('Stich', trick + 1)

            for cur_player in range(self.PLAYERS):
                print(show_cards(self.cards[:, :, cur_player, IN_HAND]), ' [',
                      show_cards(self.cards[:, :, cur_player, IN_TRICKS]), ']')
            print('\nAusspiel Spieler ', lead_player + 1)

            cur_player = lead_player
            highest = None

            for i in range(self.PLAYERS):

                if not i:
                    play, _ = self.myAgent[cur_player](cur_player, True, None)
                    lead_play = play
                    highest = cur_player, play[1]

                else:
                    play, followed_suit = self.myAgent[cur_player](cur_player, False, lead_play)

                    # the highest card is the one with the smallest(!) index
                    if followed_suit and play[1] < highest[1]:
                        highest = cur_player, play[1]

                # move card from hand of player to the table
                self.move(
                    play,
                    self.cards[:, :, cur_player, IN_HAND],
                    self.cardpool[:, :, ON_TABLE]
                )
                print('Table', show_cards(self.cardpool[:, :, ON_TABLE]))
                cur_player += 1
                if cur_player == self.PLAYERS:
                    cur_player = 0

            # add cards from table into winner's tricks and clear table
            self.cards[:, :, highest[0], IN_TRICKS] |= self.cardpool[:, :, ON_TABLE]
            self.cardpool[:, :, ON_TABLE] = 0
            lead_player = highest[0]
            print('Trick won by player ', lead_player + 1)
        print('###############')
        print('The End')
        for cur_player in range(self.PLAYERS):
            print(show_cards(self.cards[:, :, cur_player, IN_HAND]), ' [',
                  show_cards(self.cards[:, :, cur_player, IN_TRICKS]), ']')

    def legal_cards(self, player: int, outplay: List[int]) -> Tuple[np.ndarray, bool]:
        """
        calculate legal cards to drop
        :param player: number of current player
        :param outplay: lead card
        :return: (legal cards, b_followed_suit)
        """
        res_cards = np.array(self.cards[:, :, player, IN_HAND], copy=True) & self.ALL_COLORS[:, :, outplay[0]]
        if np.count_nonzero(res_cards):
            return res_cards, True
        return self.cards[:, :, player, IN_HAND], False

    @staticmethod
    def random(cards_from: np.ndarray) -> List[int]:
        """
        draw a card by random from a hand of cards
        :param cards_from: card to be drawn from
        :return: card in index form
        """
        i = randint(0, np.count_nonzero(cards_from) - 1)
        res = np.argwhere(cards_from)
        return res[i]

    @staticmethod
    def move(card: List[int], cards_from: np.ndarray, cards_to: np.ndarray) -> None:
        """
        move a card from a set of cards to another set of cards
        :param card: card in index form
        :param cards_from: card to be moved from
        :param cards_to: cards to be given to
        """
        cards_from[tuple(card)] = 0
        cards_to[tuple(card)] = 1

    def deal(self):
        """
        deal cards to each player
        """
        count_cards = self.CARDS_IN_HAND
        for player in range(self.PLAYERS):
            for i in range(self.CARDS_IN_HAND):
                card = self.random(self.cardpool[:, :, NOT_DROPPED])
                self.move(
                    card,
                    self.cardpool[:, :, NOT_DROPPED],
                    self.cards[:, :, player, IN_HAND]
                )
                count_cards -= 1

    def evaluate(self):
        """
        evalate game by assigning negative points according to the tricks of each player
        :return:
        """
        points = []
        for player in range(self.PLAYERS):
            points.append(np.count_nonzero(
                self.cards[:, :, player, IN_TRICKS] & self.POINTS
            ))
        return points

*/
