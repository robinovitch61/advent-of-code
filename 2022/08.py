import math

import common

PUZZLE = common.string(8)

TEST_PUZZLE = """30373
25512
65332
33549
35390
"""


def parse(puzzle):
    return [[int(n) for n in line] for line in puzzle.split("\n") if len(line)]


def first(puzzle):
    trees = parse(puzzle)
    w, h = len(trees[0]), len(trees)
    visible = 0
    for x in range(0, w):
        for y in range(0, h):
            tree_height = trees[y][x]
            if any([
                all(trees[y][v] < tree_height for v in range(0, x)),  # left
                all(trees[y][v] < tree_height for v in range(x + 1, w)),  # right
                all(trees[v][x] < tree_height for v in range(0, y)),  # top
                all(trees[v][x] < tree_height for v in range(y + 1, h)),  # bottom
            ]):
                visible += 1
    return visible


def view_dist(trees, x, y, dx, dy):
    tree_height = trees[y][x]
    w, h = len(trees[0]), len(trees)
    d = 0
    x, y = x + dx, y + dy
    while 0 <= x < w and 0 <= y < h:
        d += 1
        if trees[y][x] >= tree_height:
            break
        x, y = x + dx, y + dy
    return d


def scenic_score(trees, x, y):
    return math.prod(
        view_dist(trees, x, y, dx, dy) for dx, dy in ((1, 0), (-1, 0), (0, 1), (0, -1))
    )


def second(puzzle):
    trees = parse(puzzle)
    w, h = len(trees[0]), len(trees)
    scenic_scores = [scenic_score(trees, x, y) for x in range(1, w - 1) for y in range(1, h - 1)]
    return max(scenic_scores)


# `pytest *`
def test():
    assert first(TEST_PUZZLE) == 21
    assert first(PUZZLE) == 1801
    assert second(TEST_PUZZLE) == 8
    assert second(PUZZLE) == 209880


if __name__ == "__main__":
    print(first(PUZZLE))
    print(second(PUZZLE))
