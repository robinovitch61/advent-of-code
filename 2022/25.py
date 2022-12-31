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

MIN_NUM_AT_PLACE = {}
for place in range(50):
    MIN_NUM_AT_PLACE[place] = sum(-2 * 5 ** p for p in range(place))
MAX_NUM_AT_PLACE = {k: -v for k, v in MIN_NUM_AT_PLACE.items()}
NUM_TO_SNAFU = {-2: "=", -1: "-", 0: "0", 1: "1", 2: "2"}
SNAFU_TO_NUM = {v: k for k, v in NUM_TO_SNAFU.items()}


def from_snafu(snafu):
    res = 0
    for i, v in enumerate(reversed(list(snafu))):
        res += 5 ** i * SNAFU_TO_NUM[v]
    return res


def to_string(nums):
    res = ""
    for n in nums:
        res += NUM_TO_SNAFU[n]
    return res if res == "0" else res.lstrip("0")


def to_snafu(dec):
    start_place = 0
    while 2 * 5 ** start_place < dec:
        start_place += 1
    nums = deque([(i * 5 ** start_place, [i] + [0] * start_place, start_place) for i in range(-2, 3)])
    while nums:
        val, rep, place = nums.popleft()
        if val == dec:
            return to_string(rep)

        new_place = place - 1
        if new_place < 0:
            continue

        # skip if val impossible to reach dec even with all 2s/-2s in remaining spots
        if val > dec and val + MIN_NUM_AT_PLACE[place] > dec:
            continue
        if val < dec and val + MAX_NUM_AT_PLACE[place] < dec:
            continue

        for i in range(-2, 3):
            new_val = val + i * 5 ** new_place
            new_rep = rep.copy()
            new_rep[len(new_rep) - place] = i
            nums.append((new_val, new_rep, new_place))


def first(puzzle):
    res = 0
    for sn in puzzle.split("\n")[:-1]:
        res += from_snafu(sn)
    return to_snafu(res)


def test():
    assert first(TEST_PUZZLE) == "2=-1=0"
    assert to_snafu(0) == "0"
    assert to_snafu(1) == "1"
    assert to_snafu(5) == "10"
    assert to_snafu(6) == "11"
    assert to_snafu(25) == "100"
    assert to_snafu(1257) == "20012"
    assert to_snafu(353) == "1=-1="
    assert first(PUZZLE) == "2011-=2=-1020-1===-1"


if __name__ == "__main__":
    print(first(PUZZLE))
