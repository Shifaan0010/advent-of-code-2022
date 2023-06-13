import sys
import re


def manhatten_distance(pos1, pos2):
    return abs(pos1[0] - pos2[0]) + abs(pos1[1] - pos2[1])


def non_overlapping(ranges):
    left, right = 0, -2
    for l, r in sorted(ranges):
        # print(l, r)
        if left <= l <= right + 1:
            right = max(right, r)
        else:
            if left <= right:
                yield (left, right)

            left, right = l, r

    yield (left, right)


positions = []

for line in sys.stdin:
    if match := re.match(r'Sensor at x=([+-]?\d+), y=([+-]?\d+): closest beacon is at x=([+-]?\d+), y=([+-]?\d+)', line):
        sx, sy, bx, by = map(int, match.groups())

        sensor = (sx, sy)
        beacon = (bx, by)

        positions.append((sensor, beacon))


# for y in range(0, 21):
for y in range(0, 4000001):
    position_ranges = []

    for sensor, beacon in positions:
        beacon_distance = manhatten_distance(sensor, beacon)

        # print(sensor, beacon, abs(y - sy), distance)

        y_distance = abs(y - sensor[1])

        if y_distance <= beacon_distance:
            width = beacon_distance - y_distance

            left, right = sensor[0] - width, sensor[0] + width

            # if beacon[1] == y:
            #     if left == beacon[0]:
            #         left += 1
            #     if right == beacon[0]:
            #         right -= 1

            if left <= right:
                pos_range = (left, right)
                position_ranges.append(pos_range)

    ranges = list(non_overlapping(position_ranges))
    if len(ranges) >= 2:
        # y = 2766584 (-864133, 3135799) (3135801, 4102797)
        # 3135800 * 4000000 + 2766584 = 12543202766584
        print(f'y = {y:2}', *ranges)

    if y % 400000 == 0:
        print(f'{100 * y / 4000000}% done')

    # print(position_ranges)
    # print(sum(r - l + 1 for l, r in non_overlapping(position_ranges)))
