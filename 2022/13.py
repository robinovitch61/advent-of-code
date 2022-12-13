import ast
import functools

import common

PUZZLE = common.string(13)


def compare(left, right):
    for l, r in zip(left, right):
        if isinstance(l, int) and isinstance(r, int):
            if l != r:
                return -1 if l < r else 1
        elif isinstance(l, list) and isinstance(r, list):
            if (res := compare(l, r)) is not None:
                return res
        elif isinstance(l, list) and isinstance(r, int):
            if (res := compare(l, [r])) is not None:
                return res
        elif isinstance(l, int) and isinstance(r, list):
            if (res := compare([l], r)) is not None:
                return res
    if len(left) != len(right):
        return -1 if len(left) < len(right) else 1


def first(puzzle):
    res = 0
    for i, pair in enumerate(puzzle.split("\n\n")):
        split = pair.split("\n")
        if compare(ast.literal_eval(split[0]), ast.literal_eval(split[1])) == -1:
            res += i + 1
    return res


def second(puzzle):
    packets = [[[2]], [[6]]]
    for l in puzzle.split("\n"):
        if len(l):
            packets.append(ast.literal_eval(l))
    sorted_packets = sorted(packets, key=functools.cmp_to_key(compare))
    return (sorted_packets.index([[2]]) + 1) * (sorted_packets.index([[6]]) + 1)


# `pytest *`
def test():
    assert first(TEST_PUZZLE) == 13
    assert first(PUZZLE) == 5366
    assert second(TEST_PUZZLE) == 140
    assert second(PUZZLE) == 23391


if __name__ == "__main__":
    print(first(PUZZLE))
    print(second(PUZZLE))

TEST_PUZZLE = """[1,1,3,1,1]
[1,1,5,1,1]

[[1],[2,3,4]]
[[1],4]

[9]
[[8,7,6]]

[[4,4],4,4]
[[4,4],4,4,4]

[7,7,7,7]
[7,7,7]

[]
[3]

[[[]]]
[[]]

[1,[2,[3,[4,[5,6,7]]]],8,9]
[1,[2,[3,[4,[5,6,0]]]],8,9]
"""
