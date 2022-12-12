import string

import common

PUZZLE = common.string(12)

heights = {s: i for i, s in enumerate(string.ascii_lowercase)}
heights.update({"S": 0})


def steps_to_finish(grid, visited, sx, sy):
    visited[sy][sx] = True
    adjacent = []
    w, h = len(grid[0]), len(grid)
    for dx, dy in ((-1, 0), (1, 0), (0, -1), (0, 1)):
        if 0 <= sx + dx < w and 0 <= sy + dy < h:
            adjacent.append((sx + dx, sy + dy))

    to_finish = []
    for ax, ay in adjacent:
        if visited[ay][ax]:
            continue
        if grid[ay][ax] == "E":
            return 1
        if heights[grid[ay][ax]] - heights[grid[sy][sx]] <= 1:
            to_finish.append(1 + steps_to_finish(grid, visited, ax, ay))
    return min(to_finish) if len(to_finish) else 1e9


def first(puzzle):
    grid = [list(l) for l in puzzle.split("\n")[:-1]]
    for i in range(len(grid[0])):
        for j in range(len(grid)):
            if grid[j][i] == "S":
                sx, sy = i, j
    visited = [[False for _ in range(len(grid[0]))] for _ in range(len(grid))]
    return steps_to_finish(grid, visited, sx, sy)


def second(puzzle):
    return -1


# `pytest *`
def test():
    assert first(TEST_PUZZLE) == 31
    assert second(TEST_PUZZLE) == -1


if __name__ == "__main__":
    print(first(PUZZLE))
    print(second(PUZZLE))

TEST_PUZZLE = """Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi
"""
