import sys
from collections import defaultdict


def has_neighbors(pos, elf_pos):
    return any((r, c) != pos and (r, c) in elf_pos
               for r in range(pos[0] - 1, pos[0] + 2)
               for c in range(pos[1] - 1, pos[1] + 2))


def has_neighbors_in_dir(pos, elf_pos, direction):
    r, c = pos
    dr, dc = direction

    if dr == 0:
        new_c = c + dc
        return any((new_r, new_c) in elf_pos for new_r in range(r - 1, r + 2))
    elif dc == 0:
        new_r = r + dr
        return any((new_r, new_c) in elf_pos for new_c in range(c - 1, c + 2))
    else:
        raise ValueError('Invalid Direction')


class ElfPos:
    def __init__(self, elf_pos):
        self.elf_positions = set(elf_pos)
        self.round = 0
        self.directions = ((-1, 0), (1, 0), (0, -1), (0, 1))  # N S W E

    def next_state(self):
        proposed_moves = dict()
        move_positions = defaultdict(list)

        for pos in self.elf_positions:
            if has_neighbors(pos, self.elf_positions):
                # print(pos)

                for direction in self.directions:
                    if not has_neighbors_in_dir(pos, self.elf_positions, direction):
                        # print(pos, direction)
                        moved_pos = (pos[0] + direction[0],
                                     pos[1] + direction[1])
                        proposed_moves[pos] = moved_pos
                        move_positions[moved_pos].append(pos)
                        break

        for pos, elfs in move_positions.items():
            if len(elfs) >= 2:
                for pos in elfs:
                    del proposed_moves[pos]

        # print(proposed_moves)

        self.elf_positions = {proposed_moves[pos] if pos in proposed_moves else pos
                              for pos in self.elf_positions}

        self.round += 1
        self.directions = self.directions[1:] + self.directions[:1]

        return len(proposed_moves)

    def bounds(self):
        return ((min(r for r, c in self.elf_positions), max(r for r, c in self.elf_positions)+1),
                (min(c for r, c in self.elf_positions), max(c for r, c in self.elf_positions)+1))

    def ground_tiles_in_bounds(self):
        (min_r, max_r), (min_c, max_c) = self.bounds()

        return (max_r - min_r) * (max_c - min_c) - len(self.elf_positions)

    def __str__(self):
        row_bounds, col_bounds = self.bounds()

        return '\n'.join(
            ''.join('#' if (r, c) in self.elf_positions else '.'
                    for c in range(*col_bounds))
            for r in range(*row_bounds)
        )


elfs = ElfPos([(r, c) for r, line in enumerate(sys.stdin)
               for c, tile in enumerate(line) if tile == '#'])


print('Part 1')
for i in range(10):
    # print(f'Round {elfs.round}')
    # print(elfs)
    # print()

    elfs.next_state()

# print(f'Round {elfs.round}')
# print(elfs)

print(f'Empty ground tiles in bounds = {elfs.ground_tiles_in_bounds()}')

print()

print('Part 2')

while elfs.next_state() > 0:
    pass

print(f'First round where no elf moves = {elfs.round}')