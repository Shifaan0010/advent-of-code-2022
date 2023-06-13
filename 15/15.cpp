#include <iostream>
#include <regex>
#include <string>
#include <vector>

typedef std::pair<int, int> Position;
typedef std::pair<Position, Position> SensorBeacon;

int manhattenDistance(Position a, Position b) {
	return abs(a.first - b.first) + abs(a.second - b.second);
}

std::pair<std::vector<int>, std::vector<int>> lineXIntercepts(std::vector<SensorBeacon>& positions) {
	std::vector<int> slope1;
	std::vector<int> slope2;

	for (auto [sensor, beacon] : positions) {
		int distance = manhattenDistance(sensor, beacon);

		Position top = Position(sensor.first, sensor.second + distance + 1);
		Position bottom = Position(sensor.first, sensor.second - distance - 1);

		slope1.emplace_back(top.first - top.second);
		slope1.emplace_back(bottom.first - bottom.second);

		slope2.emplace_back(top.first + top.second);
		slope2.emplace_back(bottom.first + bottom.second);
	}

	return std::pair(slope1, slope2);
}

bool notInRange(Position point, std::vector<SensorBeacon>& positions) {
	for (auto [sensor, beacon] : positions) {
		if (manhattenDistance(point, sensor) <= manhattenDistance(sensor, beacon)) {
			return false;
		}
	}
	return true;
}

bool isInMap(Position point) {
	return 0 <= point.first && point.first <= 4000000 && 0 <= point.second && point.second <= 4000000;
}

Position findPosition(std::vector<SensorBeacon>& positions) {
	auto [slope1, slope2] = lineXIntercepts(positions);

	for (auto x1 : slope1) {
		for (auto x2 : slope2) {
			if ((x1 + x2) % 2 == 0) { // lines intersect
				Position point = Position((x1 + x2) / 2, abs(x1 - x2) / 2);

				if (isInMap(point) && notInRange(point, positions)) {
					printf("(%d %d)\n", point.first, point.second);
				}
			}
		}
	}

	return Position(0, 0);
}

int main() {
	std::regex r = std::regex("Sensor at x=([+-]?\\d+), y=([+-]?\\d+): closest beacon is at x=([+-]?\\d+), y=([+-]?\\d+)");
	std::smatch match = std::smatch();

	std::vector<SensorBeacon> positions;

	std::string line;
	while (getline(std::cin, line)) {
		if (std::regex_match(line, match, r)) {
			int sx = std::atoi(match[1].str().c_str());
			int sy = std::atoi(match[2].str().c_str());
			int bx = std::atoi(match[3].str().c_str());
			int by = std::atoi(match[4].str().c_str());

			positions.emplace_back(SensorBeacon(Position(sx, sy), Position(bx, by)));
		}
	}

	// for (auto p : positions) {
	// 	auto [sensor, beacon] = p;
	// 	printf("(%d %d) (%d %d)\n", sensor.first, sensor.second, beacon.first, beacon.second);
	// }

	findPosition(positions);
}
