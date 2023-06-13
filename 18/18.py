import sys


def neighbors(pos):
    return (
        (pos[0] + 1, pos[1], pos[2]),
        (pos[0] - 1, pos[1], pos[2]),
        (pos[0], pos[1] + 1, pos[2]),
        (pos[0], pos[1] - 1, pos[2]),
        (pos[0], pos[1], pos[2] + 1),
        (pos[0], pos[1], pos[2] - 1),
    )


def manhatten_distance(pos1, pos2):
    return abs(pos1[0] - pos2[0]) + abs(pos1[1] - pos2[1]) + abs(pos1[2] - pos2[2])


def surface_area(cube_positions: set):
    area = 6 * len(cube_positions)

    for pos in cube_positions:
        for neighbor in neighbors(pos):
            if neighbor in cube_positions:
                area -= 1

    return area


def in_exterior(pos, cube_positions, heuristic):
    dest = (0, 0, 0)

    distances = {pos: 0}
    # in_queue = set([pos])

    prio_q = [pos]
    while prio_q:
        node = min(
            prio_q, key=lambda pos: distances[pos] + heuristic(pos, dest))
        prio_q.remove(node)

        if node == dest:
            return True

        node_neighbors = set(neighbors(node)) - cube_positions

        for neighbor in node_neighbors:
            if neighbor not in distances:
                distances[neighbor] = distances[node] + 1
                prio_q.append(neighbor)
            elif distances[node] + 1 < distances[neighbor]:
                distances[neighbor] = distances[node] + 1

    return False


def exterior_surface_area(cube_positions: set):
    area = 6 * len(cube_positions)

    for pos in cube_positions:
        for neighbor in neighbors(pos):
            if neighbor in cube_positions:
                area -= 1
            elif not in_exterior(neighbor, cube_positions, manhatten_distance):
                area -= 1
        print(pos)

    return area


def in_bounds(pos, min_xyz, max_xyz):
    return min_xyz[0] <= pos[0] <= max_xyz[0] and \
        min_xyz[1] <= pos[1] <= max_xyz[1] and \
        min_xyz[2] <= pos[2] <= max_xyz[2]


def exterior_surface_area_2(cube_positions):
    area = 0

    min_xyz = tuple(min(vals) - 1 for vals in zip(*cube_positions))
    max_xyz = tuple(max(vals) + 1 for vals in zip(*cube_positions))

    print(min_xyz, max_xyz)

    enqueued = set()
    queue = [min_xyz]

    while queue:
        pos = queue.pop(0)

        for neighbor in neighbors(pos):
            if neighbor in cube_positions:
                area += 1
            elif in_bounds(neighbor, min_xyz, max_xyz) and \
                    neighbor not in enqueued:
                queue.append(neighbor)
                enqueued.add(neighbor)

    return area


droplet = {tuple(map(int, line.strip().split(','))) for line in sys.stdin}


# print(min(droplet, key=lambda p: p[0]))
# print(min(droplet, key=lambda p: p[1]))
# print(min(droplet, key=lambda p: p[2]))

# print(max(droplet, key=lambda p: p[0]))
# print(max(droplet, key=lambda p: p[1]))
# print(max(droplet, key=lambda p: p[2]))

print(exterior_surface_area_2(droplet))
