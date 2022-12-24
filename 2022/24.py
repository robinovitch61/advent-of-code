from collections import defaultdict

import common

PUZZLE = common.string(24)

TEST_PUZZLE = """#.######
#>>.<^<#
#.<..<<#
#>v.><>#
#<^v^^>#
######.#
"""

UP, DOWN, LEFT, RIGHT = "^", "v", "<", ">"


def parse(puzzle):
    w, h = len(puzzle.split("\n")[0]) - 2, len(puzzle.split("\n")) - 3
    blizzards = defaultdict(list)  # pos -> List[direction]
    for j, l in enumerate(puzzle.split("\n")[1:-2]):
        for i, c in enumerate(l[1:len(l) - 1]):
            if c != ".":
                blizzards[(i, j)].append(c)
    return w, h, blizzards


def adjacent(w, h, i, j):
    adjs = []
    for di, dj in ((0, 1), (0, -1), (1, 0), (-1, 0)):
        ni, nj = i + di, j + dj
        if 0 <= ni <= w - 1 and 0 <= nj <= h - 1:
            adjs.append((ni, nj))
    return adjs


def update_blizzards(w, h, blizzards):
    new_blizzards = defaultdict(list)  # pos -> List[direction]
    for (i, j), dirs in blizzards.items():
        for d in dirs:
            if d == UP:
                if j > 0:
                    new_blizzards[(i, j - 1)].append(d)
                else:
                    new_blizzards[(i, h - 1)].append(d)
            elif d == DOWN:
                if j < h - 1:
                    new_blizzards[(i, j + 1)].append(d)
                else:
                    new_blizzards[(i, 0)].append(d)
            elif d == LEFT:
                if i > 0:
                    new_blizzards[(i - 1, j)].append(d)
                else:
                    new_blizzards[(w - 1, j)].append(d)
            elif d == RIGHT:
                if i < w - 1:
                    new_blizzards[(i + 1, j)].append(d)
                else:
                    new_blizzards[(0, j)].append(d)
            else:
                raise Exception(d)
    return new_blizzards


def smallest_mins_to_end(mins, i, j, w, h, blizzards):
    positions = {(i, j)}

    while True:
        new_positions = set()
        mins += 1
        blizzards = update_blizzards(w, h, blizzards)

        for i, j in positions:
            # done
            if i == w - 1 and j == h - 1:
                return mins

            # wait
            if (i, j) not in blizzards:
                new_positions.add((i, j))

            # move
            if i is None:
                if (0, 0) not in blizzards:
                    new_positions.add((0, 0))
            else:
                for a in adjacent(w, h, i, j):
                    if a not in blizzards:
                        new_positions.add(a)

        positions = new_positions


def first(puzzle):
    w, h, blizzards = parse(puzzle)
    return smallest_mins_to_end(0, None, None, w, h, blizzards)


def second(puzzle):
    return -1


def test():
    assert first(TEST_PUZZLE) == 18
    assert second(TEST_PUZZLE) == -1


if __name__ == "__main__":
    print(first(PUZZLE))
    print(second(PUZZLE))
