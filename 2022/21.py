import common

PUZZLE = common.string(21)


def make_results(puzzle):
    results = {}
    for l in puzzle.split("\n"):
        if not len(l):
            continue
        key, val = l.split(": ")
        try:
            val = int(val)
        except:
            val = tuple(val.split(" "))
        results[key] = val
    return results


def update_results(results):
    for k, v in results.items():
        if isinstance(v, int):
            continue
        n1, op, n2 = v
        if isinstance(rn1 := results[n1], int) and isinstance(rn2 := results[n2], int):
            if op == "+":
                results[k] = rn1 + rn2
            elif op == "*":
                results[k] = rn1 * rn2
            elif op == "/":
                results[k] = rn1 // rn2
            elif op == "-":
                results[k] = rn1 - rn2
            else:
                raise Exception(op)


def first(puzzle):
    results = make_results(puzzle)
    while isinstance(results["root"], tuple):
        update_results(results)
    return results["root"]


def diff(candidate, results):
    n1, _, n2 = results["root"]
    results["humn"] = int(candidate)
    while isinstance(results[n1], tuple) or isinstance(results[n2], tuple):
        update_results(results)
    return results[n1] - results[n2]


def second(puzzle):
    results = make_results(puzzle)
    l, r = 1, int(1e15)
    positive_and_decreasing = (diff(l, results.copy()) > 0) and (diff(r, results.copy()) < 0)
    while l <= r:
        m = l + (r - l) // 2
        at_m = diff(m, results.copy())
        if at_m == 0 and diff(m - 1, results.copy()) != 0:
            return m
        if positive_and_decreasing:
            if at_m > 0:
                l = m + 1
            else:
                r = m - 1
        else:
            if at_m < 0:
                l = m + 1
            else:
                r = m - 1


def test():
    assert first(TEST_PUZZLE) == 152
    assert first(PUZZLE) == 124765768589550
    assert second(TEST_PUZZLE) == 301
    assert second(PUZZLE) == 3059361893920


if __name__ == "__main__":
    print(first(PUZZLE))
    print(second(PUZZLE))

TEST_PUZZLE = """root: pppw + sjmn
dbpl: 5
cczh: sllz + lgvd
zczc: 2
ptdq: humn - dvpt
dvpt: 3
lfqf: 4
humn: 5
ljgn: 2
sjmn: drzm * dbpl
sllz: 4
pppw: cczh / lfqf
lgvd: ljgn * ptdq
drzm: hmdt - zczc
hmdt: 32
"""
