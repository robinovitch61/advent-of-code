import re

import common

PUZZLE = common.string(22)

TEST_PUZZLE = """        ...#
        .#..
        #...
        ....
...#.......#
........#...
..#....#....
..........#.
        ...#....
        .....#..
        .#......
        ......#.

10R5L5R10L4R5L5
"""

WALL = "#"
OPEN = "."
VOID = " "
CLOCKWISE = {
    ">": "v",
    "v": "<",
    "<": "^",
    "^": ">",
}
COUNTER = {v: k for k, v in CLOCKWISE.items()}


def parse(puzzle):
    board, commands = puzzle.split("\n\n")
    commands = re.findall(r"(\d+|R|L)", commands)
    commands = [int(c) if c not in ("L", "R") else c for c in commands]
    board = board.split("\n")
    max_width = max(len(b) for b in board)
    board = [" " + b.ljust(max_width) + " " for b in board]
    void_row = " " * (max_width + 2)
    board = [void_row] + board + [void_row]
    rows = tuple(tuple(b) for b in board)
    cols = []
    for i in range(max_width + 2):
        col = []
        for j in range(len(board)):
            col.append(board[j][i])
        cols.append(tuple(col))
    cols = tuple(cols)
    return rows, cols, commands


def move(r, c, d, rows, cols):
    if d == ">":
        if rows[r][c + 1] == OPEN:
            return r, c + 1
        elif rows[r][c + 1] == WALL:
            return r, c
        elif rows[r][c + 1] == VOID:
            cc = 0
            while cc < len(rows[r]):
                if rows[r][cc] == WALL:
                    return r, c
                elif rows[r][cc] == OPEN:
                    return r, cc
                cc += 1
    elif d == "<":
        if rows[r][c - 1] == OPEN:
            return r, c - 1
        elif rows[r][c - 1] == WALL:
            return r, c
        elif rows[r][c - 1] == VOID:
            cc = len(rows[r]) - 1
            while cc >= 0:
                if rows[r][cc] == WALL:
                    return r, c
                elif rows[r][cc] == OPEN:
                    return r, cc
                cc -= 1
    elif d == "^":
        if cols[c][r - 1] == OPEN:
            return r - 1, c
        elif cols[c][r - 1] == WALL:
            return r, c
        elif cols[c][r - 1] == VOID:
            rr = len(cols[c]) - 1
            while rr >= 0:
                if cols[c][rr] == WALL:
                    return r, c
                elif cols[c][rr] == OPEN:
                    return rr, c
                rr -= 1
    elif d == "v":
        if cols[c][r + 1] == OPEN:
            return r + 1, c
        elif cols[c][r + 1] == WALL:
            return r, c
        elif cols[c][r + 1] == VOID:
            rr = 0
            while rr < len(cols[c]):
                if cols[c][rr] == WALL:
                    return r, c
                elif cols[c][rr] == OPEN:
                    return rr, c
                rr += 1
    else:
        raise Exception(d)
    return r, c


def first(puzzle):
    rows, cols, commands = parse(puzzle)
    for i, row in enumerate(rows):
        if OPEN in row:
            r = i
            break
    for j, col in enumerate(cols):
        if col[r] == OPEN:
            c = j
            break
    d = ">"
    for command in commands:
        if command == "L":
            d = COUNTER[d]
        elif command == "R":
            d = CLOCKWISE[d]
        else:
            for _ in range(command):
                r, c = move(r, c, d, rows, cols)  # can check if didn't move and break
    return 1000 * r + 4 * c + list(CLOCKWISE.keys()).index(d)


def second(puzzle):
    return -1


def test():
    assert first(TEST_PUZZLE) == 6032
    assert second(TEST_PUZZLE) == -1


if __name__ == "__main__":
    print(first(PUZZLE))
    print(second(PUZZLE))
