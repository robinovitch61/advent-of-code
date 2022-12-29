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
    states = deque([State(valve="AA", open_valves=tuple(), rate=0, released=0)])
    run_for = 30
    max_states = 1000
    time = 1
    while time < run_for:
        time += 1

        if len(states) > max_states:
            states = deque(sorted(states, key=lambda s: s.released)[-max_states:])

        for _ in range(len(states)):
            state = states.popleft()
            if flow_rates[state.valve] > 0 and state.valve not in state.open_valves:
                new_rate = state.rate + flow_rates[state.valve]
                states.append(
                    State(
                        valve=state.valve,
                        open_valves=state.open_valves + (state.valve,),
                        rate=new_rate,
                        released=state.released + new_rate
                    )
                )
            for next_valve in graph[state.valve]:
                states.append(
                    State(
                        valve=next_valve,
                        open_valves=state.open_valves,
                        rate=state.rate,
                        released=state.released + state.rate
                    )
                )
    return max(s.released for s in states)


@dataclass
class ElephantState(State):
    elephant_valve: str


def second(puzzle):
    graph, flow_rates = parse(puzzle)
    states = deque([ElephantState(valve="AA", elephant_valve="AA", open_valves=tuple(), rate=0, released=0)])
    run_for = 26
    max_states = 100000
    time = 1
    while time < run_for:
        time += 1

        if len(states) > max_states:
            states = deque(sorted(states, key=lambda s: s.released)[-max_states:])

        for _ in range(len(states)):
            state = states.popleft()
            me_at_closed = flow_rates[state.valve] > 0 and state.valve not in state.open_valves
            elephant_at_closed = flow_rates[state.elephant_valve] > 0 and state.elephant_valve not in state.open_valves

            # both open different valves
            if me_at_closed and elephant_at_closed and state.valve != state.elephant_valve:
                new_rate = state.rate + flow_rates[state.valve] + flow_rates[state.elephant_valve]
                states.append(
                    ElephantState(
                        valve=state.valve,
                        elephant_valve=state.elephant_valve,
                        open_valves=state.open_valves + (state.valve, state.elephant_valve),
                        rate=new_rate,
                        released=state.released + new_rate
                    )
                )

            # i open valve, elephant moves
            if me_at_closed:
                new_rate = state.rate + flow_rates[state.valve]
                for next_elephant in graph[state.elephant_valve]:
                    states.append(
                        ElephantState(
                            valve=state.valve,
                            elephant_valve=next_elephant,
                            open_valves=state.open_valves + (state.valve,),
                            rate=new_rate,
                            released=state.released + new_rate
                        )
                    )

            # elephant opens valve, i move
            if elephant_at_closed:
                new_rate = state.rate + flow_rates[state.elephant_valve]
                for next_valve in graph[state.valve]:
                    states.append(
                        ElephantState(
                            valve=next_valve,
                            elephant_valve=state.elephant_valve,
                            open_valves=state.open_valves + (state.elephant_valve,),
                            rate=new_rate,
                            released=state.released + new_rate
                        )
                    )

            # both move
            for next_valve in graph[state.valve]:
                for next_elephant in graph[state.elephant_valve]:
                    states.append(
                        ElephantState(
                            valve=next_valve,
                            elephant_valve=next_elephant,
                            open_valves=state.open_valves,
                            rate=state.rate,
                            released=state.released + state.rate
                        )
                    )
    return max(s.released for s in states)


# `pytest *`
def test():
    assert first(TEST_PUZZLE) == 1651
    assert first(PUZZLE) == 1584
    assert second(TEST_PUZZLE) == 1707
    assert second(PUZZLE) == 2052


if __name__ == "__main__":
    print(first(PUZZLE))
    print(second(PUZZLE))
