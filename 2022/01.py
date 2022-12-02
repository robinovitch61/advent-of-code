import common

PUZZLE = common.string(1)


def group(puzzle):
    elves = puzzle.split("\n\n")
    return tuple(sum(int(n) for n in e.split()) for e in elves)


def first(puzzle):
    return max(group(puzzle))


def second(puzzle):
    return sum(sorted(group(puzzle))[-3:])


TEST_PUZZLE = """1000
2000
3000

4000

5000
6000

7000
8000
9000

10000
"""


# `pytest *`
def test():
    assert first(TEST_PUZZLE) == 24000
    assert second(TEST_PUZZLE) == 45000


if __name__ == "__main__":
    print(first(PUZZLE))
    print(second(PUZZLE))
