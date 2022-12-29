from __future__ import annotations

import functools
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


@dataclass(frozen=True)
class Blueprint:
    num: int
    ore_for_ore_robot: int
    ore_for_clay_robot: int
    ore_for_obs_robot: int
    clay_for_obs_robot: int
    ore_for_geode_robot: int
    obs_for_geode_robot: int


@functools.total_ordering
class Inventory:
    bp: Blueprint
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

    # def __hash__(self):
    #     return hash(f"{hash(self.bp)}-ore:{self.ore}-orer{self.ore_robots}-pore{self.pending_ore_robots}-clay{self.clay}-clayr{self.clay_robots}-pclay{self.pending_clay_robots}-obs{self.obs}-obsr{self.obs_robots}-pobs{self.pending_obs_robots}-geode{self.geodes}-geoder{self.geode_robots}-pgeode={self.pending_geode_robots}")

    def __init__(self, bp):
        self.bp = bp

    def _score(self):
        # if have ore and obsidian, lose score for how far away you are from buying geode robot
        return self.geodes
        # return (
        #     (self.geodes + self.geode_robots + len(self.pending_geode_robots))
        #     + (self.obs + self.obs_robots + len(self.pending_obs_robots)) / 2
        #     + (self.clay + self.clay_robots + len(self.pending_clay_robots)) / 3
        #     + (self.ore + self.ore_robots + len(self.pending_ore_robots)) / 4
        # )

    def __lt__(self, other: Inventory):
        return self._score() < other._score()

    def __eq__(self, other: Inventory):
        return self._score() == other._score()

    def can_buy_geode_robot(self):
        return self.ore >= self.bp.ore_for_geode_robot and self.obs >= self.bp.obs_for_geode_robot

    def buy_geode_robot(self):
        self.ore -= self.bp.ore_for_geode_robot
        self.obs -= self.bp.obs_for_geode_robot
        self.pending_geode_robots += (1,)

    def can_buy_obs_robot(self):
        return self.ore >= self.bp.ore_for_obs_robot and self.clay >= self.bp.clay_for_obs_robot

    def should_buy_obs_robot(self):
        return (self.obs_robots + len(self.pending_obs_robots)) < self.bp.obs_for_geode_robot

    def buy_obs_robot(self):
        self.ore -= self.bp.ore_for_obs_robot
        self.clay -= self.bp.clay_for_obs_robot
        self.pending_obs_robots += (1,)

    def can_buy_clay_robot(self):
        return self.ore >= self.bp.ore_for_clay_robot

    def should_buy_clay_robot(self):
        return (self.clay_robots + len(self.pending_clay_robots)) < self.bp.clay_for_obs_robot

    def buy_clay_robot(self):
        self.ore -= self.bp.ore_for_clay_robot
        self.pending_clay_robots += (1,)

    def can_buy_ore_robot(self):
        return self.ore >= self.bp.ore_for_ore_robot

    def should_buy_ore_robot(self):
        # return self.ore_robots < self.bp.ore_for_clay_robot + self.bp.ore_for_clay_robot + self.bp.ore_for_obs_robot
        return (self.ore_robots + len(self.pending_ore_robots)) < max(self.bp.ore_for_ore_robot,
                                                                      self.bp.ore_for_clay_robot,
                                                                      self.bp.ore_for_clay_robot,
                                                                      self.bp.ore_for_obs_robot)

    def buy_ore_robot(self):
        self.ore -= self.bp.ore_for_ore_robot
        self.pending_ore_robots += (1,)



def potential_inventories(i: Inventory) -> List[Inventory]:
    inventories = [i]

    # can buy geode robot
    if i.can_buy_geode_robot():
        ni = copy(i)
        ni.buy_geode_robot()
        inventories.append(ni)
        # return inventories

    # can buy obs robot
    if i.can_buy_obs_robot() and i.should_buy_obs_robot():
        ni = copy(i)
        ni.buy_obs_robot()
        inventories.append(ni)
        # return inventories

    # can buy clay robot
    if i.can_buy_clay_robot() and i.should_buy_clay_robot():
        ni = copy(i)
        ni.buy_clay_robot()
        inventories.append(ni)
        # return inventories

    # can buy ore robot
    if i.can_buy_ore_robot() and i.should_buy_ore_robot():
        ni = copy(i)
        ni.buy_ore_robot()
        inventories.append(ni)
        # return inventories

    # if not len(inventories):
    #     return [i]
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
    inventories = deque([Inventory(bp=bp)])
    # max_inventories = 100000
    # kill_no_clay_robot_inventories_starting_at = bp.ore_for_clay_robot + 1
    # kill_no_geode_inventories_starting_at = 20
    max_geodes = 0
    min_t_geode = None
    run_for = 24
    for t in range(24):
        print(t, len(inventories))
        for i in range(len(inventories)):
            for ni in potential_inventories(inventories.popleft()):
                inventories.append(ni)
        inventories = deque([update(i) for i in inventories])
        max_geodes = max(max_geodes, max(i.geodes for i in inventories))
        if max_geodes > 0 and min_t_geode is None:
            min_t_geode = t
        if min_t_geode is not None and t >= min_t_geode:
            inventories = deque([i for i in inventories if i.geodes > 0])
        max_possible_geodes = (run_for - (t + 1)) ** 2 // 2
        inventories = deque([i for i in inventories if (i.geodes + max_possible_geodes) >= max_geodes])
        # if len(inventories) > max_inventories:
        #     inventories = deque(sorted(inventories)[-max_inventories:])
        # if t >= kill_no_clay_robot_inventories_starting_at:
        #     inventories = deque(i for i in inventories if i.clay_robots > 0)
        # if t >= kill_no_geode_inventories_starting_at:
        #     inventories = deque(i for i in inventories if i.geodes > 0)
    print(max(i.geodes for i in inventories))
    return bp.num * max(i.geodes for i in inventories)


def first(puzzle):
    puzzle = TEST_PUZZLE
    blueprints = get_blueprints(puzzle)
    qls = []
    for bp in blueprints:
        qls.append(quality_level(bp))
        print(qls)
    return sum(qls)


def second(puzzle):
    return -1


def test():
    assert first(TEST_PUZZLE) == -1
    assert second(TEST_PUZZLE) == -1


if __name__ == "__main__":
    print(first(PUZZLE))
    # print(second(PUZZLE))
