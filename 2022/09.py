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
    "-": (0, 0),
}


def move(direction, h, t):
    h = tuple(sum(t) for t in zip(h, delta[direction]))
    # tail up and left diagonally
    if t[1] - h[1] == 2 and h[0] - t[0] == 2:
        t = (h[0] - 1, h[1] + 1)
    # tail up and right diagonally
    if t[1] - h[1] == 2 and t[0] - h[0] == 2:
        t = (h[0] + 1, h[1] + 1)
    # tail down and left diagonally
    if h[1] - t[1] == 2 and h[0] - t[0] == 2:
        t = (h[0] - 1, h[1] - 1)
    # tail down and right diagonally
    if h[1] - t[1] == 2 and t[0] - h[0] == 2:
        t = (h[0] + 1, h[1] - 1)
    # tail two above
    if t[1] - h[1] == 2:
        t = (h[0], h[1] + 1)
    # tail two below
    if h[1] - t[1] == 2:
        t = (h[0], h[1] - 1)
    # tail two left
    if h[0] - t[0] == 2:
        t = (h[0] - 1, h[1])
    # tail two right
    if t[0] - h[0] == 2:
        t = (h[0] + 1, h[1])
    return h, t


def first(puzzle):
    h, t = (0, 0), (0, 0)
    tails = set()
    for l in puzzle.split("\n")[:-1]:
        direction, count = l.split(" ")
        for _ in range(int(count)):
            h, t = move(direction, h, t)
            tails.add(t)
    return len(tails)


def print_rope(rope):
    min_x = min(x[0] for x in rope)
    max_x = max(x[0] for x in rope)
    min_y = min(x[1] for x in rope)
    max_y = max(x[1] for x in rope)
    lines = []
    for i in range(min_y, max_y + 1):
        line = ""
        for j in range(min_x, max_x + 1):
            if (j, i) in rope:
                line += str(rope.index((j, i)))
            else:
                line += "."
        lines.append(line)
    for line in reversed(lines):
        print(line)
    print()


def second(puzzle):
    rope = [(0, 0) for _ in range(10)]
    tails = set()
    print_rope(rope)
    for l in puzzle.split("\n")[:-1]:
        direction, count = l.split(" ")
        for _ in range(int(count)):
            rope = move_rope(direction, rope)
            tails.add(rope[-1])
        print_rope(rope)
    return len(tails)


def move_rope(direction, rope):
    # move head
    prev_h, prev_t = rope[0], rope[1]
    h, t = move(direction, prev_h, prev_t)
    rope[0], rope[1] = h, t
    # adjust rest
    for i in range(1, len(rope) - 1):
        prev_h, prev_t = rope[i], rope[i + 1]
        h, t = move("-", prev_h, prev_t)
        rope[i], rope[i + 1] = h, t
    return rope


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
