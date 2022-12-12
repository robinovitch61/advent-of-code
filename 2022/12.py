import string
from collections import deque

import common

PUZZLE = common.string(12)

heights = {s: i for i, s in enumerate(string.ascii_lowercase)}
heights.update({"S": 0})
heights.update({"E": 25})


def get_adjacent(grid, sx, sy):
    adjacent = []
    w, h = len(grid[0]), len(grid)
    for dx, dy in ((-1, 0), (1, 0), (0, -1), (0, 1)):
        if 0 <= sx + dx < w and 0 <= sy + dy < h:
            adjacent.append((sx + dx, sy + dy))
    return adjacent


def steps_to_finish(grid, sx, sy):
    min_steps_to = [[float('inf') for _ in range(len(grid[0]))] for _ in range(len(grid))]
    min_steps_to[sy][sx] = 0
    q = deque()
    q.append((sx, sy))
    while len(q):
        x, y = q.popleft()
        curr_steps = min_steps_to[y][x]
        for ax, ay in get_adjacent(grid, x, y):
            if heights[grid[ay][ax]] - heights[grid[y][x]] <= 1:
                if grid[ay][ax] == "E":
                    return curr_steps + 1
                if curr_steps + 1 < min_steps_to[ay][ax]:
                    min_steps_to[ay][ax] = curr_steps + 1
                    q.append((ax, ay))


def first(puzzle):
    grid = [list(l) for l in puzzle.split("\n")[:-1]]
    for i in range(len(grid[0])):
        for j in range(len(grid)):
            if grid[j][i] == "S":
                sx, sy = i, j
    return steps_to_finish(grid, sx, sy)


def second(puzzle):
    grid = [list(l) for l in puzzle.split("\n")[:-1]]
    all_a = []
    for i in range(len(grid[0])):
        for j in range(len(grid)):
            if grid[j][i] == "a":
                all_a.append((i, j))
    return min(x for x in [steps_to_finish(grid, a[0], a[1]) for a in all_a] if x)


# `pytest *`
def test():
    assert first(TEST_PUZZLE) == 31
    assert second(TEST_PUZZLE) ==29


if __name__ == "__main__":
    print(first(PUZZLE))
    print(second(PUZZLE))

TEST_PUZZLE = """Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi
"""
