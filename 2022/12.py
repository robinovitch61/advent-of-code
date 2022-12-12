import string
from collections import deque

import common

PUZZLE = common.string(12)

HEIGHTS = {s: i for i, s in enumerate(string.ascii_lowercase)}
HEIGHTS.update({"S": 0, "E": 25})


def get_adjacent(grid, sx, sy):
    adjacent = []
    w, h = len(grid[0]), len(grid)
    for dx, dy in ((-1, 0), (1, 0), (0, -1), (0, 1)):
        if 0 <= sx + dx < w and 0 <= sy + dy < h:
            adjacent.append((sx + dx, sy + dy))
    return adjacent


def get_min_steps_to(grid, sx, sy, char):
    min_steps_to = [[float("inf") for _ in range(len(grid[0]))] for _ in range(len(grid))]
    min_steps_to[sy][sx] = 0
    q = deque()
    q.append((sx, sy))
    while len(q):
        x, y = q.popleft()
        curr_steps = min_steps_to[y][x]
        for ax, ay in get_adjacent(grid, x, y):
            if HEIGHTS[grid[y][x]] - HEIGHTS[grid[ay][ax]] <= 1:
                if grid[ay][ax] == char:
                    return curr_steps + 1
                elif curr_steps + 1 < min_steps_to[ay][ax]:
                    min_steps_to[ay][ax] = curr_steps + 1
                    q.append((ax, ay))


def solve(puzzle, char):
    grid = [list(l) for l in puzzle.split("\n")[:-1]]
    for i in range(len(grid[0])):
        for j in range(len(grid)):
            if grid[j][i] == "E":
                sx, sy = i, j
    return get_min_steps_to(grid, sx, sy, char)


def first(puzzle):
    return solve(puzzle, "S")


def second(puzzle):
    return solve(puzzle, "a")


# `pytest *`
def test():
    assert first(TEST_PUZZLE) == 31
    assert first(PUZZLE) == 472
    assert second(TEST_PUZZLE) == 29
    assert second(PUZZLE) == 465


if __name__ == "__main__":
    print(first(PUZZLE))
    print(second(PUZZLE))

TEST_PUZZLE = """Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi
"""
