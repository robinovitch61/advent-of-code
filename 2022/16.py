import functools
import re

import common

PUZZLE = common.string(16)

TEST_PUZZLE = """Valve AA has flow rate=0; tunnels lead to valves DD, II, BB
Valve BB has flow rate=13; tunnels lead to valves CC, AA
Valve CC has flow rate=2; tunnels lead to valves DD, BB
Valve DD has flow rate=20; tunnels lead to valves CC, AA, EE
Valve EE has flow rate=3; tunnels lead to valves FF, DD
Valve FF has flow rate=0; tunnels lead to valves EE, GG
Valve GG has flow rate=0; tunnels lead to valves FF, HH
Valve HH has flow rate=22; tunnel leads to valve GG
Valve II has flow rate=0; tunnels lead to valves AA, JJ
Valve JJ has flow rate=21; tunnel leads to valve II
"""

puzzle = TEST_PUZZLE
graph = {}
flow_rates = {}
for line in puzzle.split("\n")[:-1]:
    valve, flow_rate, conns = \
        re.findall(r"Valve ([A-Z]{2}) has flow rate=(\d+); tunnels? leads? to valves? (.*)", line)[0]
    graph[valve] = conns.split(", ")
    flow_rates[valve] = int(flow_rate)


cache = {}


def max_pressure_released(current_valve, open_valves, time, path):
    val, cached_path = cache.get((current_valve, open_valves, time), (None, None))
    if (val, cached_path) != (None, None):
        return val, cached_path
    current_pressure_release = sum(flow_rates[v] for v in open_valves)
    if time >= 30:
        return 0, path
    if flow_rates[current_valve] > 0 and current_valve not in open_valves:
        current_pressure_release *= 2
        open_valves = open_valves + (current_valve,)
        time += 1
        path = path + (f"{time} ({current_valve}: open {current_valve}",)
    max_pressure_release = 0
    next_path = None
    for next_valve in graph[current_valve]:
        max_next, path = max_pressure_released(next_valve, open_valves, time + 1, path)
        if (current_pressure_release + max_next) > max_pressure_release:
            max_pressure_release = current_pressure_release + max_next
            next_path = next_valve
    path = path + (f"{time} ({current_valve}): move to {next_path}",)
    cache[(current_valve, open_valves, time)] = max_pressure_release, path
    return max_pressure_release, path


def first(puzzle):
    # puzzle = TEST_PUZZLE
    # graph = {}
    # flow_rates = {}
    # for line in puzzle.split("\n")[:-1]:
    #     valve, flow_rate, conns = \
    #         re.findall(r"Valve ([A-Z]{2}) has flow rate=(\d+); tunnels? leads? to valves? (.*)", line)[0]
    #     graph[valve] = conns.split(", ")
    #     flow_rates[valve] = int(flow_rate)
    res, path = max_pressure_released("AA", tuple(), 1, tuple())
    print(path)
    return res


def second(puzzle):
    return -1


# `pytest *`
def test():
    assert first(TEST_PUZZLE) == 1651
    assert second(TEST_PUZZLE) == -1


if __name__ == "__main__":
    print(first(PUZZLE))
    print(second(PUZZLE))
