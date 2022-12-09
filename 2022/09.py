import common

PUZZLE = common.string(9)

TEST_PUZZLE = """R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2
"""

TEST_PUZZLE_2 = """R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20
"""

delta = {
    "U": (0, 1),
    "D": (0, -1),
    "L": (-1, 0),
    "R": (1, 0),
}


def first(puzzle):
    return unique_tails_in_rope(puzzle, 2)


def second(puzzle):
    return unique_tails_in_rope(puzzle, 10)


def unique_tails_in_rope(puzzle, rope_length):
    rope = [(0, 0) for _ in range(rope_length)]
    tails = set()
    for l in puzzle.split("\n")[:-1]:
        direction, count = l.split(" ")
        for _ in range(int(count)):
            rope = move_rope(direction, rope)
            tails.add(rope[-1])
    return len(tails)


def move_rope(direction, rope):
    # move head
    d, h = delta[direction], rope[0]
    rope[0] = (h[0] + d[0], h[1] + d[1])
    # adjust tail
    for i in range(0, len(rope) - 1):
        rope[i + 1] = follow(rope[i], rope[i + 1])
    return rope


def follow(h, t):
    d = (h[0] - t[0], h[1] - t[1])
    if abs(d[0]) < 2 and abs(d[1]) < 2:
        return t
    return t[0] + sign(d[0]), t[1] + sign(d[1])


def sign(x):
    if x == 0:
        return 0
    return 1 if x > 0 else -1


# `pytest *`
def test():
    assert first(TEST_PUZZLE) == 13
    assert first(PUZZLE) == 6522

    assert second(TEST_PUZZLE) == 1
    assert second(TEST_PUZZLE_2) == 36
    assert second(PUZZLE) == 2717


if __name__ == "__main__":
    print(first(PUZZLE))
    print(second(PUZZLE))
