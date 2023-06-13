import re


def get_dir(file_tree, path):
    cur_dir = file_tree
    for dir_name in path:
        cur_dir = cur_dir[dir_name]
    return cur_dir


def iter_dirs(file_tree, path=tuple()):
    for key in file_tree:
        if type(file_tree[key]) == dict:
            yield from iter_dirs(file_tree[key], path=path + (key,))

    yield file_tree, path
    # yield path


def print_tree(file_tree, depth=0):
    for key in file_tree:
        if type(file_tree[key]) == dict:
            print(f'{"    " * depth}|-- {key}')
            print_tree(file_tree[key], depth=depth+1)
        else:
            print(f'{"    " * depth}|-- {key} {file_tree[key]}')


def dir_size(file_tree):
    size = 0
    for key in file_tree:
        if type(file_tree[key]) == dict:
            size += dir_size(file_tree[key])
        else:
            size += file_tree[key]

    return size


with open('7.txt', 'r') as file:
    file_tree = dict()

    command = ''
    cur_path = []
    for line in file:
        if match := re.match(r'\$ cd (.+)', line):
            command = 'cd'

            arg = match.groups()[0]

            if arg == '..':
                cur_path.pop()
            elif arg == '.':
                pass
            elif arg == '/':
                cur_path = []
            else:
                cur_path.append(arg)

        elif match := re.match(r'\$ ls', line):
            command = 'ls'

        elif command == 'ls' and (match := re.match(r'dir (.+)', line)):
            dir_name = match.groups()[0]
            cur_dir = get_dir(file_tree, cur_path)
            cur_dir[dir_name] = dict()

        elif command == 'ls' and (match := re.match(r'(\d+) (.+)', line)):
            file_size, file_name = match.groups()
            cur_dir = get_dir(file_tree, cur_path)
            cur_dir[file_name] = int(file_size)

        else:
            print('Unmatched Line:', line)

    # print_tree(file_tree)

    # print(sum([dir_size(d) for d in iter_dirs(file_tree) if dir_size(d) <= 100000]))

    space_needed = 30000000 - (70000000 - dir_size(file_tree))

    # print(sorted([dir_size(d) for d in iter_dirs(file_tree)]))

    print(min([(dir_size(d), p) for d, p in iter_dirs(file_tree) if dir_size(d) >= space_needed]))
