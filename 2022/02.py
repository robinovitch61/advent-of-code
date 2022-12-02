import common

PUZZLE = common.string(2)

score_for_result = {
    "win": 6,
    "draw": 3,
    "lose": 0,
}

score_for_playing = {
    "rock": 1,
    "paper": 2,
    "scissors": 3,
}

to_play = {
    "A": "rock",
    "X": "rock",
    "B": "paper",
    "Y": "paper",
    "C": "scissors",
    "Z": "scissors",
}

i_won = {
    ("paper", "scissors"),
    ("rock", "paper"),
    ("scissors", "rock"),
}


def calc_score(rounds):
    score = 0
    for round in rounds:
        if round in i_won:
            score += score_for_result["win"]
        elif round[0] == round[1]:
            score += score_for_result["draw"]
        else:
            score += score_for_result["lose"]
        score += score_for_playing[round[1]]
    return score


def first(puzzle):
    splits = (line.split() for line in puzzle.splitlines())
    rounds = ((to_play[s[0]], to_play[s[1]]) for s in splits)
    return calc_score(rounds)


to_win = {
    "rock": "paper",
    "paper": "scissors",
    "scissors": "rock",
}

req_res = {
    "X": {v: k for k, v in to_win.items()},
    "Y": {v: v for k, v in to_win.items()},
    "Z": to_win,
}


def second(puzzle):
    splits = (line.split() for line in puzzle.splitlines())
    rounds = ()
    for s in splits:
        their_play = to_play[s[0]]
        rounds += ((their_play, req_res[s[1]][their_play]),)
    return calc_score(rounds)


TEST_PUZZLE = """A Y
B X
C Z
"""


# `pytest *`
def test():
    assert first(TEST_PUZZLE) == 15
    assert second(TEST_PUZZLE) == 12


if __name__ == "__main__":
    print(first(PUZZLE))
    print(second(PUZZLE))
