import re
from collections import deque
from copy import copy
from dataclasses import dataclass
from typing import Tuple, List

import common

PUZZLE = common.lines(19)

BLUEPRINT_REGEX = re.compile(
    r"Blueprint (\d+): Each ore robot costs (\d+) ore. Each clay robot costs (\d+) ore. Each obsidian robot costs (\d+) ore and (\d+) clay. Each geode robot costs (\d+) ore and (\d+) obsidian."
)

TEST_PUZZLE = [
    "Blueprint 1: Each ore robot costs 4 ore. Each clay robot costs 2 ore. Each obsidian robot costs 3 ore and 14 clay. Each geode robot costs 2 ore and 7 obsidian.",
    "Blueprint 2: Each ore robot costs 2 ore. Each clay robot costs 3 ore. Each obsidian robot costs 3 ore and 8 clay. Each geode robot costs 3 ore and 12 obsidian.",
]


@dataclass
class Blueprint:
    num: int
    ore_for_ore_robot: int
    ore_for_clay_robot: int
    ore_for_obs_robot: int
    clay_for_obs_robot: int
    ore_for_geode_robot: int
    obs_for_geode_robot: int


@dataclass
class Inventory:
    ore: int = 0
    ore_robots: int = 1
    pending_ore_robots: Tuple[int] = ()
    clay: int = 0
    clay_robots: int = 0
    pending_clay_robots: Tuple[int] = ()
    obs: int = 0
    obs_robots: int = 0
    pending_obs_robots: Tuple[int] = ()
    geodes: int = 0
    geode_robots: int = 0
    pending_geode_robots: Tuple[int] = ()


def potential_inventories(i: Inventory, bp: Blueprint) -> List[Inventory]:
    inventories = [i]

    # can buy geode robot
    if i.ore >= bp.ore_for_geode_robot and i.obs >= bp.obs_for_geode_robot:
        ni = copy(i)
        ni.ore -= bp.ore_for_geode_robot
        ni.obs -= bp.obs_for_geode_robot
        ni.pending_geode_robots += (2,)
        inventories.append(ni)

        # can buy obs robot
    if i.ore >= bp.ore_for_obs_robot and i.clay >= bp.clay_for_obs_robot:
        ni = copy(i)
        ni.ore -= bp.ore_for_obs_robot
        ni.clay -= bp.clay_for_obs_robot
        ni.pending_obs_robots += (2,)
        inventories.append(ni)

        # can buy clay robot
    if i.ore >= bp.ore_for_clay_robot:
        ni = copy(i)
        ni.ore -= bp.ore_for_clay_robot
        ni.pending_clay_robots += (2,)
        inventories.append(ni)

        # can buy ore robot
    if i.ore >= bp.ore_for_ore_robot:
        ni = copy(i)
        ni.ore -= bp.ore_for_ore_robot
        ni.pending_ore_robots += (2,)
        inventories.append(ni)

    return inventories


def update(i: Inventory) -> Inventory:
    # collect new resources
    for _ in range(i.ore_robots):
        i.ore += 1
    for _ in range(i.clay_robots):
        i.clay += 1
    for _ in range(i.obs_robots):
        i.obs += 1
    for _ in range(i.geode_robots):
        i.geodes += 1

    # update robots
    i.pending_ore_robots = tuple(r - 1 for r in i.pending_ore_robots)
    for p in i.pending_ore_robots:
        if p == 0:
            i.ore_robots += 1
    i.pending_ore_robots = tuple(r for r in i.pending_ore_robots if r != 0)

    i.pending_obs_robots = tuple(r - 1 for r in i.pending_obs_robots)
    for p in i.pending_obs_robots:
        if p == 0:
            i.obs_robots += 1
    i.pending_obs_robots = tuple(r for r in i.pending_obs_robots if r != 0)

    i.pending_clay_robots = tuple(r - 1 for r in i.pending_clay_robots)
    for p in i.pending_clay_robots:
        if p == 0:
            i.clay_robots += 1
    i.pending_clay_robots = tuple(r for r in i.pending_clay_robots if r != 0)

    i.pending_geode_robots = tuple(r - 1 for r in i.pending_geode_robots)
    for p in i.pending_geode_robots:
        if p == 0:
            i.geode_robots += 1
    i.pending_geode_robots = tuple(r for r in i.pending_geode_robots if r != 0)

    return i


def get_blueprints(puzzle):
    blueprints = []
    for bp in puzzle:
        blueprints.append(Blueprint(*(map(int, re.findall(BLUEPRINT_REGEX, bp)[0]))))
    return blueprints


def quality_level(bp: Blueprint):
    inventories = deque([Inventory()])
    for t in range(24):
        for i in range(len(inventories)):
            for ni in potential_inventories(inventories.popleft(), bp):
                inventories.append(ni)
        inventories = deque([update(i) for i in inventories])
        print(t, len(inventories))
    print(inventories)
    return bp.num * i.geodes


def first(puzzle):
    puzzle = TEST_PUZZLE
    blueprints = get_blueprints(puzzle)
    qls = []
    for bp in blueprints:
        qls.append(quality_level(bp))
    return sum(qls)


def second(puzzle):
    return -1


def test():
    assert first(TEST_PUZZLE) == -1
    assert second(TEST_PUZZLE) == -1


if __name__ == "__main__":
    print(first(PUZZLE))
    print(second(PUZZLE))
