import functools
import re
from collections import deque
from dataclasses import dataclass
from typing import Tuple

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


@dataclass
class State:
    time: int
    valve: str
    open_valves: Tuple[str, ...]
    rate: int
    released: int


def parse(puzzle):
    graph, flow_rates = {}, {}
    for line in puzzle.split("\n")[:-1]:
        valve, flow_rate, conns = \
            re.findall(r"Valve ([A-Z]{2}) has flow rate=(\d+); tunnels? leads? to valves? (.*)", line)[0]
        graph[valve] = conns.split(", ")
        flow_rates[valve] = int(flow_rate)
    return graph, flow_rates


def first(puzzle):
    graph, flow_rates = parse(puzzle)
    states = deque([State(time=1, valve="AA", open_valves=tuple(), rate=0, released=0)])
    run_for = 30
    max_states = 1000
    while not all(state.time == run_for for state in states):
        if len(states) > max_states:
            states = deque(sorted(states, key=lambda s: s.released)[-max_states:])

        for _ in range(len(states)):
            state = states.popleft()
            if state.time == run_for:
                states.append(state)
                continue
            if flow_rates[state.valve] > 0 and state.valve not in state.open_valves:
                new_rate = state.rate + flow_rates[state.valve]
                states.append(
                    State(
                        time=state.time + 1,
                        valve=state.valve,
                        open_valves=state.open_valves + (state.valve,),
                        rate=new_rate,
                        released=state.released + new_rate
                    )
                )
            for next_valve in graph[state.valve]:
                states.append(
                    State(
                        time=state.time + 1,
                        valve=next_valve,
                        open_valves=state.open_valves,
                        rate=state.rate,
                        released=state.released + state.rate
                    )
                )
    return max(s.released for s in states)


# @functools.lru_cache(maxsize=None)
# def with_elephant(me, elephant, open_valves, release, time):
#     if time > 26:
#         return 0
#
#     max_pressure_release = 0
#     me_at_closed = FLOW_RATES[me] > 0 and me not in open_valves
#     elephant_at_closed = FLOW_RATES[elephant] > 0 and elephant not in open_valves
#
#     # both of us open our valves
#     if me_at_closed and elephant_at_closed:
#         if me != elephant:
#             max_next = release + with_elephant(
#                 me, elephant, open_valves + (me, elephant), release + FLOW_RATES[me] + FLOW_RATES[elephant], time + 1
#             )
#             max_pressure_release = max(max_pressure_release, max_next)
#
#     # i open my valve, elephant moves
#     if me_at_closed:
#         for next_elephant in GRAPH[elephant]:
#             max_next = release + with_elephant(
#                 me, next_elephant, open_valves + (me,), release + FLOW_RATES[me], time + 1
#             )
#             max_pressure_release = max(max_pressure_release, max_next)
#
#     # elephant opens their valve, i move
#     if elephant_at_closed:
#         for next_me in GRAPH[me]:
#             max_next = release + with_elephant(
#                 next_me, elephant, open_valves + (elephant,), release + FLOW_RATES[elephant], time + 1
#             )
#             max_pressure_release = max(max_pressure_release, max_next)
#
#     # both of us move
#     for next_elephant in GRAPH[elephant]:
#         for next_me in GRAPH[me]:
#             max_next = release + with_elephant(next_me, next_elephant, open_valves, release, time + 1)
#             max_pressure_release = max(max_pressure_release, max_next)
#
#     return max_pressure_release


def second(puzzle):
    return -1


# `pytest *`
def test():
    assert first(TEST_PUZZLE) == 1651
    assert first(PUZZLE) == 1584
    # assert second(TEST_PUZZLE) == 1707


if __name__ == "__main__":
    print(first(PUZZLE))
    # print(second(PUZZLE))
