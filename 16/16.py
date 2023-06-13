import re
import sys
from functools import cache, reduce
from itertools import combinations

START_VALVE = 'AA'
TOTAL_TIME = 30
TOTAL_TIME_ELEPHANT = 26
VALVE_OPEN_TIME = 1
TUNNEL_TRAVEL_TIME = 1


def log(foo):
    def logged(*args):
        # print(args)
        res = foo(*args)
        print(f'{foo.__name__}({", ".join(map(str, args))}) = {res}')
        return res
    return logged


def maximize_pressure(graph, start_valve, total_time):
    @cache
    # @log
    def max_pressure(node, rem_time, valves_turned_on=frozenset()):
        if rem_time == 0:
            return 0
        else:
            flow, neighbors = graph[node]

            max_pres = max(
                max_pressure(
                    neighbor,
                    rem_time - 1,
                    valves_turned_on
                )
                for neighbor in neighbors
            )

            if node not in valves_turned_on and flow > 0:
                max_pres = max(
                    max_pres,
                    flow * (rem_time - 1) +
                    max_pressure(
                        node,
                        rem_time - 1,
                        valves_turned_on | {node}
                    )
                )

            return max_pres

    return max_pressure(start_valve, total_time)


def maximize_pressure_elephant(graph, start_valve, total_time):  # too slow
    @cache
    # @log
    def max_pressure(nodes, rem_time, valves_turned_on=frozenset(), turn=0):
        if rem_time == 0:
            return 0
        else:
            node, el_node = nodes

            cur_node = node if turn == 0 else el_node

            flow, neighbors = graph[cur_node]

            max_pres = max(
                max_pressure(
                    (neighbor, el_node) if turn == 0 else (node, neighbor),
                    rem_time - 1 if turn == 1 else rem_time,
                    valves_turned_on,
                    (turn + 1) % 2
                )
                for neighbor in neighbors
            )

            if cur_node not in valves_turned_on and flow > 0:
                max_pres = max(
                    max_pres,
                    flow * (rem_time - 1) +
                    max_pressure(
                        (node, el_node) if turn == 0 else (node, el_node),
                        rem_time - 1 if turn == 1 else rem_time,
                        valves_turned_on | {cur_node},
                        (turn + 1) % 2
                    )
                )

            return max_pres

    return max_pressure((start_valve, start_valve), total_time)


def max_dict_vals(d1, d2):
    max_dict = dict()
    for key in d1.keys() | d2.keys():
        if key not in d1:
            max_dict[key] = d2[key]
        elif key not in d2:
            max_dict[key] = d1[key]
        else:
            max_dict[key] = max(d1[key], d2[key])
    return max_dict


def dict_add(x, d):
    return {k: d[k] + x for k in d}


def maximize_pressure_elephant_2(graph, start_valve, total_time):
    # @log
    @cache
    def calc_max_pressure(node, rem_time, valves_turned_on=frozenset()):
        if rem_time == 0:
            return {valves_turned_on: 0}
        else:
            flow, neighbors = graph[node]

            max_pres = reduce(
                max_dict_vals,
                (
                    calc_max_pressure(
                        neighbor,
                        rem_time - 1,
                        valves_turned_on
                    )
                    for neighbor in neighbors
                )
            )

            if node not in valves_turned_on and flow > 0:
                pressure = dict_add(
                    flow * (rem_time - 1),
                    calc_max_pressure(
                        node,
                        rem_time - 1,
                        valves_turned_on | {node}
                    )
                )

                max_pres = max_dict_vals(
                    max_pres,
                    pressure
                )

            return max_pres

    max_pressures = calc_max_pressure(start_valve, total_time)

    max_disjoint = (None, -1)

    for s1, s2 in combinations(max_pressures.keys(), r=2):
        if s1 & s2 == frozenset():
            pressure = max_pressures[s1] + max_pressures[s2]
            if pressure > max_disjoint[1]:
                max_disjoint = ((s1, s2), pressure)

        # print(s1, s2)

    # return calc_max_pressure(start_valve, total_time)
    return max_disjoint


# def astar_maximize(graph, start_valve, total_time):
#     start_state = (start_valve, total_time, frozenset())

#     distances = defaultdict(lambda: -1)

#     p_queue = []
#     while p_queue:
#         node =


graph = dict()

for line in sys.stdin:
    if match := re.match(r'Valve (\w+) has flow rate=(\d+); tunnels? leads? to valves? ([, \w]+)', line):
        valve, flow, neighbors = match.groups()

        flow = int(flow)
        neighbors = tuple(neighbors.split(', '))

        graph[valve] = (flow, neighbors)

print(*graph.items(), sep='\n')
# print(maximize_pressure_elephant(graph, START_VALVE, TOTAL_TIME_ELEPHANT))

# print(*sorted(maximize_pressure_elephant_2(graph, START_VALVE, TOTAL_TIME).items(), key=lambda t: (len(t[0]), t[1])), sep='\n')

print(maximize_pressure_elephant_2(graph, START_VALVE, TOTAL_TIME_ELEPHANT))
