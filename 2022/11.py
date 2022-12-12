# TODO: clean up
import math
import re

import common

PUZZLE = common.string(11)

MONKEY_PARSE = re.compile(r"""Monkey (\d+):
  Starting items: (.*)
  Operation: new = (.*)
  Test: divisible by (\d+)
    If true: throw to monkey (\d+)
    If false: throw to monkey (\d+)""")


class Monkey:
    def __init__(self, s):
        mid, items, op, div, t, f = re.findall(MONKEY_PARSE, s)[0]
        self.mid, self.div, self.t, self.f = map(int, (mid, div, t, f))
        self.items = [int(i) for i in items.split(", ")]
        self.op = op
        self.inspected = 0


def first(puzzle):
    monkeys = []
    for monkey in puzzle.split("\n\n"):
        monkeys.append(Monkey(monkey))

    for _ in range(20):
        for m in monkeys:
            for old in m.items:
                m.inspected += 1
                new = eval(m.op) // 3
                if new % m.div == 0:
                    monkeys[m.t].items.append(new)
                else:
                    monkeys[m.f].items.append(new)
            m.items = []

    return math.prod(sorted([m.inspected for m in monkeys])[-2:])


def second(puzzle):
    monkeys = []
    for monkey in puzzle.split("\n\n"):
        monkeys.append(Monkey(monkey))

    mod_by = math.prod(m.div for m in monkeys)
    for r in range(10000):
        for m in monkeys:
            for old in m.items:
                m.inspected += 1
                new = eval(m.op) % mod_by
                if new % m.div == 0:
                    monkeys[m.t].items.append(new)
                else:
                    monkeys[m.f].items.append(new)
            m.items = []

    return math.prod(sorted([m.inspected for m in monkeys])[-2:])


# `pytest *`
def test():
    assert first(TEST_PUZZLE) == 10605
    # assert first(PUZZLE) == 58056
    # assert second(TEST_PUZZLE) == 2713310158
    # assert second(PUZZLE) == 15048718170


if __name__ == "__main__":
    print(first(PUZZLE))
    print(second(PUZZLE))

TEST_PUZZLE = """Monkey 0:
  Starting items: 79, 98
  Operation: new = old * 19
  Test: divisible by 23
    If true: throw to monkey 2
    If false: throw to monkey 3

Monkey 1:
  Starting items: 54, 65, 75, 74
  Operation: new = old + 6
  Test: divisible by 19
    If true: throw to monkey 2
    If false: throw to monkey 0

Monkey 2:
  Starting items: 79, 60, 97
  Operation: new = old * old
  Test: divisible by 13
    If true: throw to monkey 1
    If false: throw to monkey 3

Monkey 3:
  Starting items: 74
  Operation: new = old + 3
  Test: divisible by 17
    If true: throw to monkey 0
    If false: throw to monkey 1
"""
