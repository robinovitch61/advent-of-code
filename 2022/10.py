import re
from collections import deque

import common

PUZZLE = common.string(10)


def get_x_vals(puzzle):
    ops = deque(l for l in puzzle.split("\n") if len(l))
    current_op = None
    x = 1
    cycle_to_x = {}
    cycle = 0
    while len(ops) or current_op is not None:
        cycle += 1
        # start
        if current_op is None and len(ops):
            op = ops.popleft()
            if op.startswith("addx"):
                n = next(int(n) for n in re.findall(r"-?\d+", op))
                current_op = (n, 1)
        # during
        cycle_to_x[cycle] = x
        # after
        if current_op:
            val, delay = current_op
            if delay == 0:
                x += val
                current_op = None
            else:
                current_op = (val, delay-1)
    return cycle_to_x


def first(puzzle):
    cycle_to_x = get_x_vals(puzzle)
    return sum(cycle * x for cycle, x in cycle_to_x.items() if (cycle - 20) % 40 == 0)


def second(puzzle):
    cycle_to_x = get_x_vals(puzzle)
    screen = [["." for _ in range(40)] for _ in range(6)]
    for cycle, x in cycle_to_x.items():
        idx = cycle - 1
        draw_x, draw_y = idx % 40, idx // 40
        if abs(draw_x - x) <= 1:
            screen[draw_y][draw_x] = "#"
    return "\n".join("".join(s) for s in screen)


# `pytest *`
def test():
    assert first(TEST_PUZZLE) == 13140
    assert first(PUZZLE) == 14360
    assert second(TEST_PUZZLE) == """##..##..##..##..##..##..##..##..##..##..
###...###...###...###...###...###...###.
####....####....####....####....####....
#####.....#####.....#####.....#####.....
######......######......######......####
#######.......#######.......#######....."""
    assert second(PUZZLE) == """###...##..#..#..##..####.###..####.####.
#..#.#..#.#.#..#..#.#....#..#.#.......#.
###..#....##...#..#.###..#..#.###....#..
#..#.#.##.#.#..####.#....###..#.....#...
#..#.#..#.#.#..#..#.#....#.#..#....#....
###...###.#..#.#..#.####.#..#.####.####."""


if __name__ == "__main__":
    print(first(PUZZLE))
    print(second(PUZZLE))

TEST_PUZZLE = """addx 15
addx -11
addx 6
addx -3
addx 5
addx -1
addx -8
addx 13
addx 4
noop
addx -1
addx 5
addx -1
addx 5
addx -1
addx 5
addx -1
addx 5
addx -1
addx -35
addx 1
addx 24
addx -19
addx 1
addx 16
addx -11
noop
noop
addx 21
addx -15
noop
noop
addx -3
addx 9
addx 1
addx -3
addx 8
addx 1
addx 5
noop
noop
noop
noop
noop
addx -36
noop
addx 1
addx 7
noop
noop
noop
addx 2
addx 6
noop
noop
noop
noop
noop
addx 1
noop
noop
addx 7
addx 1
noop
addx -13
addx 13
addx 7
noop
addx 1
addx -33
noop
noop
noop
addx 2
noop
noop
noop
addx 8
noop
addx -1
addx 2
addx 1
noop
addx 17
addx -9
addx 1
addx 1
addx -3
addx 11
noop
noop
addx 1
noop
addx 1
noop
noop
addx -13
addx -19
addx 1
addx 3
addx 26
addx -30
addx 12
addx -1
addx 3
addx 1
noop
noop
noop
addx -9
addx 18
addx 1
addx 2
noop
noop
addx 9
noop
noop
noop
addx -1
addx 2
addx -37
addx 1
addx 3
noop
addx 15
addx -21
addx 22
addx -6
addx 1
noop
addx 2
addx 1
noop
addx -10
noop
noop
addx 20
addx 1
addx 2
addx 2
addx -6
addx -11
noop
noop
noop
"""
