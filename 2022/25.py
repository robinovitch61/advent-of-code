from __future__ import annotations

from collections import deque

import common

PUZZLE = common.string(25)

TEST_PUZZLE = """1=-0-2
12111
2=0=
21
2=01
111
20012
112
1=-1=
1-12
12
1=
122
"""


def from_snafu(snafu):
    res = 0
    for i, v in enumerate(reversed(list(snafu))):
        if v == "-":
            m = -1
        elif v == "=":
            m = -2
        else:
            m = int(v)
        res += 5 ** i * m
    return res


def to_string(nums):
    res = ""
    for n in nums:
        if n == -2:
            res += "="
        elif n == -1:
            res += "-"
        else:
            res += str(n)
    return res if res == "0" else res.lstrip("0")


min_num_at_place = {}
for place in range(50):
    min_num_at_place[place] = sum(-2 * 5 ** p for p in range(place))
max_num_at_place = {k: -v for k, v in min_num_at_place.items()}


def to_snafu(dec):
    start_place = 0
    while 2 * 5 ** start_place < dec:
        start_place += 1
    # print(dec, start_place, 2 * 5 ** start_place)
    nums = deque([(i * 5 ** start_place, [i] + [0] * start_place, start_place) for i in range(-2, 3)])
    while nums:
        val, rep, place = nums.popleft()
        # print(val, place, dec)
        if val == dec:
            return to_string(rep)

        new_place = place - 1
        if new_place < 0:
            continue

        if val > dec and val + min_num_at_place[place] > dec:
            continue
        if val < dec and (val + max_num_at_place[place]) < dec:
            continue

        for i in range(-2, 3):
            new_val = val + i * 5 ** new_place
            new_rep = rep.copy()
            new_rep[len(new_rep) - place] = i
            nums.append((new_val, new_rep, new_place))

    # @dataclass
    # class Node:
    #     s: str
    #     v: int
    #     children: List[Node]
    #
    # root = Node("", 0, [])
    # q = deque([root])
    # levels = 11  # lol this is so dumb
    # for i in range(levels):
    #     for _ in range(len(q)):
    #         v = q.popleft()
    #         for child in (
    #             Node("=", -2 * 5 ** i, []),
    #             Node("-", -1 * 5 ** i, []),
    #             Node("0", 0 * 5 ** i, []),
    #             Node("1", 1 * 5 ** i, []),
    #             Node("2", 2 * 5 ** i, []),
    #         ):
    #             v.children.append(child)
    #             q.append(child)
    #
    # def find_sum(v, n: Node):
    #     if not len(n.children):
    #         if v == n.v:
    #             return [n.s]
    #         else:
    #             return None
    #     r = v - n.v
    #     for child in n.children:
    #         if (p := find_sum(r, child)) is not None:
    #             return p + [n.s]
    # res = "".join(find_sum(dec, root)).lstrip("0")
    # if res == "":
    #     return "0"
    # return -1


def first(puzzle):
    res = 0
    for sn in puzzle.split("\n")[:-1]:
        res += from_snafu(sn)
    # print(res)
    return to_snafu(res)


def second(puzzle):
    return -1


def test():
    assert first(TEST_PUZZLE) == "2=-1=0"
    assert to_snafu(0) == "0"
    assert to_snafu(1) == "1"
    assert to_snafu(5) == "10"
    assert to_snafu(6) == "11"
    assert to_snafu(25) == "100"
    assert to_snafu(1257) == "20012"
    assert to_snafu(353) == "1=-1="
    # assert second(TEST_PUZZLE) == -1


if __name__ == "__main__":
    print(first(PUZZLE))
    print(second(PUZZLE))
