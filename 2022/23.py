from collections import defaultdict

import common

PUZZLE = common.string(23)

TEST_PUZZLE = """....#..
..###.#
#...#.#
.#...##
#.###..
##.#.##
.#..#..
"""

DIFFS = ((0, 1), (0, -1), (1, 0), (-1, 0), (1, 1), (1, -1), (-1, 1), (-1, -1))
DIR_DIFFS = {
    "N": (((-1, -1), (0, -1), (1, -1)), (0, -1)),
    "E": (((1, 1), (1, 0), (1, -1)), (1, 0)),
    "S": (((-1, 1), (0, 1), (1, 1)), (0, 1)),
    "W": (((-1, 1), (-1, 0), (-1, -1)), (-1, 0)),
}


def get_elves(puzzle):
    elves = set()
    for j, l in enumerate(puzzle.split("\n")[:-1]):
        for i, c in enumerate(l):
            if c == "#":
                elves.add((i, j))
    return elves


def get_adjacent(i, j):
    adj = []
    for di, dj in DIFFS:
        adj.append((i + di, j + dj))
    return adj


def move_elves(directions, elves):
    # first half
    proposals = defaultdict(list)  # proposal -> List[elf]
    for i, j in elves:
        adj = get_adjacent(i, j)
        if not any(a in elves for a in adj):
            continue
        else:
            for d in directions:
                check, (dpi, dpj) = DIR_DIFFS[d]
                if not any((i + ci, j + cj) in elves for ci, cj in check):
                    # print(f"elf {(i, j)} proposing {(i + dpi, j + dpj)}")
                    proposals[(i + dpi, j + dpj)].append((i, j))
                    break

    # second half
    elf_moved = False
    for p, pelves in proposals.items():
        if len(pelves) == 1:
            elf_moved = True
            elves.remove(pelves[0])
            elves.add(p)
    return elf_moved


def first(puzzle):
    elves = get_elves(puzzle)
    directions = ["N", "S", "W", "E"]
    for _ in range(10):
        move_elves(directions, elves)
        directions = directions[1:] + [directions[0]]
    all_i, all_j = zip(*elves)
    empty_tiles = 0
    for i in range(min(all_i), max(all_i) + 1):
        for j in range(min(all_j), max(all_j) + 1):
            if (i, j) not in elves:
                empty_tiles += 1
    return empty_tiles


def second(puzzle):
    elves = get_elves(puzzle)
    directions = ["N", "S", "W", "E"]
    round = 0
    while True:
        round += 1
        if not move_elves(directions, elves):
            return round
        directions = directions[1:] + [directions[0]]


def test():
    assert first(TEST_PUZZLE) == 110
    assert second(TEST_PUZZLE) == 20


if __name__ == "__main__":
    print(first(PUZZLE))
    print(second(PUZZLE))
