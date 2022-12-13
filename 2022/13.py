import functools

import common

PUZZLE = common.string(13)


def compare(left, right):
    for l, r in zip(left, right):
        if isinstance(l, int) and isinstance(r, int):
            if l < r:
                return 1
            elif l > r:
                return -1
        elif isinstance(l, list) and isinstance(r, list):
            res = compare(l, r)
            if res is not None:
                return res
        elif isinstance(l, list) and isinstance(r, int):
            res = compare(l, [r])
            if res is not None:
                return res
        elif isinstance(l, int) and isinstance(r, list):
            res = compare([l], r)
            if res is not None:
                return res
    if len(left) != len(right):
        return 1 if len(left) < len(right) else -1


def first(puzzle):
    res = 0
    for i, pair in enumerate(puzzle.split("\n\n")):
        split = pair.split("\n")
        left, right = eval(split[0]), eval(split[1])
        comp = compare(left, right)
        if comp == 1:
            res += i + 1
    return res


def second(puzzle):
    packets = []
    for l in puzzle.split("\n"):
        if len(l):
            packets.append(eval(l))
    packets.append([[2]])
    packets.append([[6]])
    packets.sort(key=functools.cmp_to_key(compare))
    packets = reversed(packets)
    res = 1
    for i, p in enumerate(packets):
        if p == [[2]] or p == [[6]]:
            res *= (i + 1)
    return res


# `pytest *`
def test():
    assert first(TEST_PUZZLE) == 13
    assert first(PUZZLE) > 4390
    assert first(PUZZLE) < 5396
    assert first(PUZZLE) > 5217
    assert second(TEST_PUZZLE) == 140


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
