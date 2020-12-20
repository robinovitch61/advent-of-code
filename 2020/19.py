
import itertools
import re

with open('./19_input.txt', 'r') as f:
    INPUT_TEXT = f.read()

TEST_INPUT_TEXT = """0: 4 1 5
1: 2 3 | 3 2
2: 4 4 | 5 5
3: 4 5 | 5 4
4: "a"
5: "b"

ababbb
bababa
abbbab
aaabbb
aaaabbb"""

TEST_INPUT_2_TEXT = """42: 9 14 | 10 1
9: 14 27 | 1 26
10: 23 14 | 28 1
1: "a"
11: 42 31
5: 1 14 | 15 1
19: 14 1 | 14 14
12: 24 14 | 19 1
16: 15 1 | 14 14
31: 14 17 | 1 13
6: 14 14 | 1 14
2: 1 24 | 14 4
0: 8 11
13: 14 3 | 1 12
15: 1 | 14
17: 14 2 | 1 7
23: 25 1 | 22 14
28: 16 1
4: 1 1
20: 14 14 | 1 15
3: 5 14 | 16 1
27: 1 6 | 14 18
14: "b"
21: 14 1 | 1 14
25: 1 1 | 1 14
22: 14 14
8: 42
26: 14 22 | 1 20
18: 15 15
7: 14 5 | 1 21
24: 14 1

abbbbbabbbaaaababbaabbbbabababbbabbbbbbabaaaa
bbabbbbaabaabba
babbbbaabbbbbabbbbbbaabaaabaaa
aaabbbbbbaaaabaababaabababbabaaabbababababaaa
bbbbbbbaaaabbbbaaabbabaaa
bbbababbbbaaaaaaaabbababaaababaabab
ababaaaaaabaaab
ababaaaaabbbaba
baabbaaaabbaaaababbaababb
abbbbabbbbaaaababbbbbbaaaababb
aaaaabbaabaaaaababaa
aaaabbaaaabbaaa
aaaabbaabbaaaaaaabbbabbbaaabbaabaaa
babaaabbbaaabaababbaabababaaab
aabbbbbaabbbaaaaaabbbbbababaaaaabbaaabba"""


def combine(*args):
    return [''.join(c) for c in itertools.product(*args)]


def parse_rule(num, rule_set):
    def helper(num, rule_set, memo):
        if memo is None:
            memo = {}
        elif num in memo:
            return memo[num]

        rule = rule_set[num]

        if rule in ("a", "b"):
            memo[num] = [rule]
            return memo[num]

        combos = [[int(n) for n in v.split(' ')] for v in rule.split(' | ')]
        final = []
        for combo in combos:
            rule_result = []
            for num in combo:
                if num in memo:
                    rule_result += memo[num]
                else:
                    memo[num] = [helper(num, rule_set, memo)]
                    rule_result += memo[num]
            final += combine(*rule_result)
        return final
    return set(helper(num, rule_set, None))



def test_first():
    rule_text, messages_text = TEST_INPUT_TEXT.split("\n\n")
    rule_set = {int(k): v.replace('"', '') for k, v in [l.split(': ') for l in rule_text.split('\n')]}
    options = parse_rule(0, rule_set)
    assert(len([m for m in messages_text.split('\n') if m in options]) == 2)


def first():
    rule_text, messages_text = INPUT_TEXT.split("\n\n")
    rule_set = {int(k): v.replace('"', '') for k, v in [l.split(': ') for l in rule_text.split('\n')]}
    options = parse_rule(0, rule_set)
    print(len([m for m in messages_text.split('\n') if m in options]))


#  The above method - generate every possible string specified in the rules - breaks down in the second part of the
#  question, as there are just too many combinations. Looking around at other's solutions, I see that compiling to regex
#  and a depth-first search through the rules are more scaleable. I implement the regex below.


def gen_regex(rule_set, rule_num):
    def helper(rule_set, rule_num, memo):
        regex = ''
        if memo is None:
            memo = {}
        elif rule_num in memo:
            return memo[rule_num]
        rule = rule_set[rule_num]
        if rule in ('a', 'b'):
            memo[rule_num] = rule
            return memo[rule_num]
        else:
            for num in rule.split(' '):
                if num == '|':
                    regex += '|'
                else:
                    if num in memo:
                        regex += memo[num]
                    else:
                        res = helper(rule_set, int(num), memo)
                        to_add = res if res in ('a', 'b') else f"({res})"
                        memo[num] = to_add
                        regex += memo[num]
            return regex
    return helper(rule_set, rule_num, None)



def test_first():
    rule_text, messages_text = TEST_INPUT_TEXT.split("\n\n")
    rule_set = {int(k): v.replace('"', '') for k, v in [l.split(': ') for l in rule_text.split('\n')]}
    regex = gen_regex(rule_set, 0)
    assert(regex == "a((aa|bb)(ab|ba)|(ab|ba)(aa|bb))b")
    assert(sum([1 if re.fullmatch(regex, m) is not None else 0 for m in messages_text.split('\n')]) == 2)


def first():
    rule_text, messages_text = INPUT_TEXT.split("\n\n")
    rule_set = {int(k): v.replace('"', '') for k, v in [l.split(': ') for l in rule_text.split('\n')]}
    regex = gen_regex(rule_set, 0)
    print(sum([1 if re.fullmatch(regex, m) is not None else 0 for m in messages_text.split('\n')]))


def test_second():
    rule_text, messages_text = TEST_INPUT_2_TEXT.split("\n\n")
    rule_set = {int(k): v.replace('"', '') for k, v in [l.split(': ') for l in rule_text.split('\n')]}
    rule_set[8] = '42 | 42 42 | 42 42 42 | 42 42 42 42 | 42 42 42 42 42'
    rule_set[11] = '42 31 | 42 42 31 31 | 42 42 42 31 31 31 | 42 42 42 42 31 31 31 31 | 42 42 42 42 42 31 31 31 31 31'
    regex = gen_regex(rule_set, 0)
    assert(sum([1 if re.fullmatch(regex, m) is not None else 0 for m in messages_text.split('\n')]) == 12)


def second():
    rule_text, messages_text = INPUT_TEXT.split("\n\n")
    rule_set = {int(k): v.replace('"', '') for k, v in [l.split(': ') for l in rule_text.split('\n')]}
    rule_set[8] = '42 | 42 42 | 42 42 42 | 42 42 42 42 | 42 42 42 42 42'
    rule_set[11] = '42 31 | 42 42 31 31 | 42 42 42 31 31 31 | 42 42 42 42 31 31 31 31 | 42 42 42 42 42 31 31 31 31 31'
    regex = gen_regex(rule_set, 0)
    print(sum([1 if re.fullmatch(regex, m) is not None else 0 for m in messages_text.split('\n')]))


if __name__ == "__main__":
    test_first()
    first()
    test_second()
    second()
