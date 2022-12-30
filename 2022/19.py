# heavily inspired by https://topaz.github.io/paste/#XQAAAQBSCgAAAAAAAAA0m0pnuFI8c/T1e7Ae/cusMcunC50HRo4qEuS8gTduVPKHvqa+158jZSEjRPjwhsmfq/WpYLa7b2cLmgfAzlwMkcrv/uDj0oOUznPoxUe1DItVRQVkbgiRARMXmWVIZRd6C14IZf/QA8CGrpnCgd/3qvat+o0WhEWnGLb1CKhc9pGtZY7ao6byl15GEJB2vvnvzPAYRkPb2SAFmIB11zXraLd1DX91D5LKcMVlAX1qkP9salg6v1Z0pc/qtUHdvUzfNWCaJ2FK/bkfWE7QBTfleIOe4z7D0YRnQgh8HytVLx04UN2VC3g+1Npe/6UHaN+8Xn1pi3rYThE+qKgsGjjdPTceyQD9cRTDj3gBgfX1UAhpvyzqI9Dk8F3nCYwFJl0IaniVDIOwKIF31VakghYVl88evTGIMpC9MWTwFhxRK6lsq5Pk0+0Qihl5yqn+3hqedoVBuyWi1Ykufqg5ANdff4OfadPci/CcnSmjIi0NPp6LH0sfrGDWWvV41rCDo8C0+lt3HN0idff+LtVB0BkHXz7uxIEZ7WY83NLOFgaN6eTA2ato5Ps7Ak1KyHNvTQPkEeGRVS+8IPqpSBjgA8kNKlb/XzcJ4V5PzlzWFSnjM52Nt/vrtMeJnTa/C0lSOQhV9PRGBcfVacSU3MY/cSWe1wdThcriGw6AJxCHq5/HPFKOPww8nz2tWREY9Vtb9MfPJUhaUs4cfsCQYcmR2OzdBoWJs8+Qh9j4IUtwBN7bJmb6zse7KdJEhDx78pZPDdVKeTBofuxSzh7sCcSaMkh3RylV0dnl4feg+hstKdt/HYCcgysNdWrkO2npBOxu2z8uCCA6IWUvnO8oRzyp3sepNtjRyhaj1zPRgG/LRbL6jG5AJRlkHc0TwOnE32THbqhFAICQCAorCzcW481ZnB/OTpO3p76iufyuExC3tzHqIVC7h3e2CkccBowIV+K98wL/keFd8SQhhqgoJE9LSkyPrTWfoi6lY7tqU83dNHHmFJqePsel9SmC3LCmRYLwwL1ssOWzkiu1YYCG0xItxCewpoRAqJ6gP2b3Sbw0OEGFpfotpHwFFMtQfHR4Cf9sX4RYHABknaO2JjjVfXVujcnRXtlp94VNJNuxybM1NvWSyRQEF8kuLIhltAaGK2f6eCldLIjqVeDUR/C+Hr59d1MMOJtNHZcpB/NL2Fcn/zh7DwA=

import math
import re
from collections import deque, defaultdict

import common

PUZZLE = common.lines(19)

BLUEPRINT_REGEX = re.compile(
    r"Blueprint (\d+): Each ore robot costs (\d+) ore. Each clay robot costs (\d+) ore. Each obsidian robot costs (\d+) ore and (\d+) clay. Each geode robot costs (\d+) ore and (\d+) obsidian."
)

TEST_PUZZLE = [
    "Blueprint 1: Each ore robot costs 4 ore. Each clay robot costs 2 ore. Each obsidian robot costs 3 ore and 14 clay. Each geode robot costs 2 ore and 7 obsidian.",
    "Blueprint 2: Each ore robot costs 2 ore. Each clay robot costs 3 ore. Each obsidian robot costs 3 ore and 8 clay. Each geode robot costs 3 ore and 12 obsidian.",
]

GEODE, OBS, CLAY, ORE = "geode", "obs", "clay", "ore"


def get_blueprints(puzzle):
    bps = {}
    for line in puzzle:
        vals = list(map(int, re.findall(BLUEPRINT_REGEX, line)[0]))
        bps[vals[0]] = {
            ORE: {ORE: vals[1]},
            CLAY: {ORE: vals[2]},
            OBS: {
                ORE: vals[3],
                CLAY: vals[4],
            },
            GEODE: {
                ORE: vals[5],
                OBS: vals[6],
            },
        }
    return bps


def get_robot_limits(bp):
    limits = defaultdict(float)
    limits[GEODE] = float("inf")
    for costs in bp.values():
        for mat, cost in costs.items():
            limits[mat] = max(limits[mat], cost)
    return limits


def get_buildable_robots(bp, resources):
    buildable = set()
    for robot, costs in bp.items():
        if all(resources[mat] >= cost for mat, cost in costs.items()):
            buildable.add(robot)
    return buildable


def update_resources(resources, robots):
    res = resources.copy()
    for robot, count in robots.items():
        res[robot] += count
    return res


def build_robot(robot, bp, resources, robots):
    new_resources, new_robots = resources.copy(), robots.copy()
    new_robots[robot] += 1
    for mat, cost in bp[robot].items():
        new_resources[mat] -= cost
    return new_resources, new_robots


def max_geodes(bp, total_time):
    start_resources = {mat: 0 for mat in bp.keys()}
    start_robots = start_resources.copy()
    start_robots[ORE] = 1
    states = deque([(0, start_resources, start_robots, set())])
    max_geodes_at_time = defaultdict(int)
    robot_limits = get_robot_limits(bp)
    while states:
        time, resources, robots, skipped = states.popleft()
        max_geodes_at_time[time] = max(max_geodes_at_time[time], resources[GEODE])
        if time <= total_time:
            # huge optimization - only consider the leader states
            if resources[GEODE] < max_geodes_at_time[time]:
                continue

            buildable_robots = get_buildable_robots(bp, resources)
            buildable_robots = set(r for r in buildable_robots if robots[r] < robot_limits[r])

            # if geode buildable, only build it
            if GEODE in buildable_robots:
                new_resources, new_robots = build_robot(GEODE, bp, resources, robots)
                new_resources = update_resources(new_resources, robots)  # exclude new robot
                states.append((time + 1, new_resources, new_robots, set()))
            else:
                # don't build a robot
                new_resources = update_resources(resources, robots)
                states.append((time + 1, new_resources, robots, buildable_robots))

                # build buildable robots
                for buildable_robot in buildable_robots:
                    # huge optimization - if you could have built a robot and didn't
                    # build it last time, don't build it this time
                    if buildable_robot in skipped:
                        continue
                    new_resources, new_robots = build_robot(buildable_robot, bp, resources, robots)
                    new_resources = update_resources(new_resources, robots)  # exclude new robot
                    states.append((time + 1, new_resources, new_robots, set()))
    return max_geodes_at_time[total_time]


def first(puzzle):
    bps = get_blueprints(puzzle)
    return sum(n * max_geodes(bp, 24) for n, bp in bps.items())


def second(puzzle):
    bps = get_blueprints(puzzle[:3])
    return math.prod(max_geodes(bp, 32) for _, bp in bps.items())


def test():
    assert first(TEST_PUZZLE) == 33
    assert first(PUZZLE) == 1147
    assert second(TEST_PUZZLE) == 56 * 62
    assert second(PUZZLE) == 3080


if __name__ == "__main__":
    print(first(PUZZLE))
    print(second(PUZZLE))
