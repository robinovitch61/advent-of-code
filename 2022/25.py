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
        if v == "=":
            m = -2
        elif v == "-":
            m = -1
        else:
            m = int(v)
        res += 5 ** i * m
    return res


def _to_snafu(dec):
    if dec == 0:
        return ""
    remainder, quotient, carry = dec % 5, dec // 5, (dec + 5) // 5
    if remainder == 2:
        return _to_snafu(quotient) + "2"
    elif remainder == 1:
        return _to_snafu(quotient) + "1"
    elif remainder == 0:
        return _to_snafu(quotient) + "0"
    elif remainder == 3:
        return _to_snafu(carry) + "="
    elif remainder == 4:
        return _to_snafu(carry) + "-"


def to_snafu(dec):
    return _to_snafu(dec) or "0"


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
