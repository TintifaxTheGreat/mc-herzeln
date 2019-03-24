package main

/*

import numpy as np
from typing import List

CARDSTRINGS = [
['HA', 'HK', 'HO', 'HU', 'HX', 'H9', 'H8', 'H7'],
['SA', 'SK', 'SO', 'SU', 'SX', 'S9', 'S8', 'S7'],
['PA', 'PK', 'PO', 'PU', 'PX', 'P9', 'P8', 'P7'],
['EA', 'EK', 'EO', 'EU', 'EX', 'E9', 'E8', 'E7'],
]


def card_to_string(card: List[int]) -> str:
return CARDSTRINGS[card[0]][card[1]]


def string_to_card(s: str) -> List[int]:
for i, suit in enumerate(CARDSTRINGS):
try:
index = CARDSTRINGS[i].index(s)
except ValueError:
continue
return [i, index]
raise ValueError


def show_cards(cards: np.ndarray) -> str:
s = ''
for (line, column), value in np.ndenumerate(cards):
if value:
s += CARDSTRINGS[line][column] + ' '
return s

*/
