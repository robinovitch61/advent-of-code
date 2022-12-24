import functools
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


class HashableDefaultDict(defaultdict):
    def __hash__(self):
        tuples = []
        for pos, dirs in self.items():
            tuples.append((pos, tuple(dirs)))
        return hash(tuple(sorted(tuples)))


def parse(puzzle):
    w, h = len(puzzle.split("\n")[0]) - 2, len(puzzle.split("\n")) - 3
    blizzards = HashableDefaultDict(list)  # pos -> List[direction]
    for j, l in enumerate(puzzle.split("\n")[1:-2]):
        for i, c in enumerate(l[1:len(l) - 2]):
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
    new_blizzards = HashableDefaultDict(list)  # pos -> List[direction]
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


@functools.lru_cache(None)
def smallest_mins_to_end(mins, i, j, w, h, blizzards):
    blizzards = blizzards.copy()
    candidates = []

    # update blizzards
    blizzards = update_blizzards(w, h, blizzards)

    # done
    if i == w - 1 and j == h - 1:
        return mins + 1

    # wait
    if (i, j) not in blizzards:
        candidates.append(smallest_mins_to_end(mins + 1, i, j, w, h, blizzards))

    # move
    if i is None:
        if (0, 0) not in blizzards:
            candidates.append(smallest_mins_to_end(mins + 1, 0, 0, w, h, blizzards))
    else:
        for a in adjacent(w, h, i, j):
            if a not in blizzards:
                candidates.append(smallest_mins_to_end(mins + 1, a[0], a[1], w, h, blizzards))

    # blocked
    if len(candidates) == 0:
        return float("inf")

    return min(candidates)


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
