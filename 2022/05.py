import re

import common

PUZZLE = common.string(5)

TEST_PUZZLE = (
    "    [D]    \n"
    + "[N] [C]    \n"
    + "[Z] [M] [P]\n"
    + " 1   2   3 \n"
    + """
move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2
""")


def parse(puzzle):
    stack, moves = puzzle.split("\n\n")

    stack_lines = stack.split("\n")
    stacks = []
    for i in range(1, len(stack_lines[-1]), 4):
        new_stack = list(reversed([l[i] for l in stack_lines[:-1]]))
        stacks.append([s for s in new_stack if s.strip() != ""])

    parsed_moves = [
        tuple(int(n) for n in re.findall(r"\d+", m))
        for m in moves.split("\n")[:-1]
    ]

    return stacks, parsed_moves


def solve(puzzle, move):
    stacks, moves = parse(puzzle)
    for m in moves:
        stacks = move(stacks, *m)
    return "".join(s[-1] for s in stacks)


def move_first(stacks, cnt, src, dest):
    for _ in range(cnt):
        val = stacks[src - 1].pop()
        stacks[dest - 1].append(val)
    return stacks


def first(puzzle):
    return solve(puzzle, move_first)


def move_second(stacks, cnt, src, dest):
    val = stacks[src - 1][-cnt:]
    stacks[src - 1] = stacks[src - 1][:-cnt]
    stacks[dest - 1] += val
    return stacks


def second(puzzle):
    return solve(puzzle, move_second)


# `pytest *`
def test():
    assert first(TEST_PUZZLE) == "CMZ"
    assert first(PUZZLE) == "DHBJQJCCW"
    assert second(TEST_PUZZLE) == "MCD"
    assert second(PUZZLE) == "WJVRLSJJT"


if __name__ == "__main__":
    print(first(PUZZLE))
    print(second(PUZZLE))
