#include <iostream>
#include <string>
#include <unordered_set>
#include <array>

enum class Direction {
	Right, Left, Up, Down
};

Direction str_to_dir(std::string &&s) {
	Direction dir = Direction::Right;

	if (s == "R") {
		dir = Direction::Right;
	} else if (s == "L") {
		dir = Direction::Left;
	} else if (s == "U") {
		dir = Direction::Up;
	} else if (s == "D") {
		dir = Direction::Down;
	}

	return dir;
}

std::string dir_to_str(Direction dir) {
	switch (dir) {
		case Direction::Right:
			return "R";
		case Direction::Left:
			return "L";
		case Direction::Up:
			return "U";
		case Direction::Down:
			return "D";
	}
}

class KnotPosition {
public:
	std::pair<int, int> position;

	KnotPosition *follower;
public:
	KnotPosition(std::pair<int, int> position, KnotPosition *follower): position(position), follower(follower) {}

	void move(Direction dir) {
		switch (dir) {
			case Direction::Right:
				position.first += 1;
				break;
			case Direction::Left:
				position.first -= 1;
				break;
			case Direction::Up:
				position.second += 1;
				break;
			case Direction::Down:
				position.second -= 1;
				break;
		}

		if (follower != nullptr) {
			follower->follow(*this);
		}
	}
private:
	void follow(KnotPosition& head) {
		// std::printf("Follow: (%d %d) (%d %d)\n", position.first, position.second, head.position.first, head.position.second);

		if (abs(head.position.first - position.first) >= 2 || abs(head.position.second - position.second) >= 2) {
			if (head.position.first - position.first >= 1) {
				position.first += 1;
			} else if (position.first - head.position.first >= 1) {
				position.first -= 1;
			}

			if (head.position.second - position.second >= 1) {
				position.second += 1;
			} else if (position.second - head.position.second >= 1) {
				position.second -= 1;
			}

		}
		if (follower != nullptr) {
			follower->follow(*this);
		}
	}
};

struct pair_hash {
	inline std::size_t operator()(const std::pair<int, int> &v) const {
		return v.first * 31 + v.second;
	}
};

int main() {
	std::array<KnotPosition*, 10> snake = std::array<KnotPosition*, 10>();

	for (int i = 0; i < snake.size(); i += 1) {
		snake[i] = new KnotPosition(std::pair<int, int>(0, 0), (i > 0) ? snake[i - 1] : nullptr);
	}

	KnotPosition *head = snake[snake.size() - 1];
	KnotPosition *tail = snake[0];

	std::unordered_set<std::pair<int, int>, pair_hash> visited = std::unordered_set<std::pair<int, int>, pair_hash>();
	visited.insert(tail->position);

	std::string line;

	while (std::getline(std::cin, line)) {
		Direction dir = str_to_dir(line.substr(0, line.find(' ')));
		int steps = std::stoi(line.substr(line.find(' ') + 1));

		for (int i = 0; i < steps; i += 1) {
			head->move(dir);

			visited.insert(tail->position);
		}

		for (int i = 0; i < snake.size(); i += 1) {
			std::printf("(%d %d) ", snake[i]->position.first, snake[i]->position.second);
		}
		std::printf("\n");

		// std::printf("Read: %s %d\n", dir_to_str(dir).c_str(), steps);
	}

	std::printf("%lu\n", visited.size());
}
