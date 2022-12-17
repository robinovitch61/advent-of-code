import functools
import re

import common

# BELLMAN-FORD ALGO? FLOYD-WARSHALL ALGO

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

GRAPH, FLOW_RATES = {}, {}


@functools.lru_cache(maxsize=None)
def max_pressure_released(current_valve, open_valves, release, time):
    if time > 30:
        return 0
    max_pressure_release = 0
    if FLOW_RATES[current_valve] > 0 and current_valve not in open_valves:
        max_next = release + max_pressure_released(current_valve, open_valves + (current_valve,), release + FLOW_RATES[current_valve], time + 1)
        max_pressure_release = max(max_pressure_release, max_next)
    for next_valve in GRAPH[current_valve]:
        max_next = release + max_pressure_released(next_valve, open_valves, release, time + 1)
        max_pressure_release = max(max_pressure_release, max_next)
    return max_pressure_release


def first(puzzle):
    GRAPH.clear()
    FLOW_RATES.clear()
    max_pressure_released.cache_clear()
    for line in puzzle.split("\n")[:-1]:
        valve, flow_rate, conns = \
            re.findall(r"Valve ([A-Z]{2}) has flow rate=(\d+); tunnels? leads? to valves? (.*)", line)[0]
        GRAPH[valve] = conns.split(", ")
        FLOW_RATES[valve] = int(flow_rate)
    return max_pressure_released("AA", tuple(), 0, 1)


def second(puzzle):
    return -1


# `pytest *`
def test():
    assert first(TEST_PUZZLE) == 1651
    assert first(PUZZLE) == 1584
    assert second(TEST_PUZZLE) == -1


if __name__ == "__main__":
    print(first(PUZZLE))
    print(second(PUZZLE))
