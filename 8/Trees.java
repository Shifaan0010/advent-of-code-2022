import java.util.Scanner;
import java.io.File;
import java.io.FileNotFoundException;
import java.util.ArrayList;

class Tree {
	enum Direction {
		Right, Down, Left, Up
	}

	public static int[] getRowCol(ArrayList<ArrayList<Integer>> grid, int i, int j, Direction d) {
		int row = 0;
		int col = 0;

		switch (d) {
			case Right:
				row = i;
				col = j;
				break;
			case Down:
				row = j;
				col = i;
				break;
			case Left:
				row = i;
				col = grid.get(row).size() - j - 1;
				break;
			case Up:
				row = grid.size() - j - 1;
				col = i;
				break;
			default:
				System.out.println("Error");
		}

		return new int[] { row, col };
	}

	public static void resizeGrid(ArrayList<ArrayList<Boolean>> a, int rows, int cols) {
		for (int r = 0; r < rows; r += 1) {
			ArrayList<Boolean> row = new ArrayList<Boolean>();

			for (int c = 0; c < cols; c += 1) {
				row.add(false);
			}

			a.add(row);
		}
	}

	public static boolean inRange(int n, int max) {
		return n >= 0 && n < max;
	}

	public static ArrayList<ArrayList<Boolean>> visibility(ArrayList<ArrayList<Integer>> grid) {
		ArrayList<ArrayList<Boolean>> visible = new ArrayList<ArrayList<Boolean>>();
		resizeGrid(visible, grid.size(), grid.get(0).size());

		for (Direction direction : Direction.values()) {
			// System.out.println(direction.toString());

			for (int i = 0; i < grid.size(); i += 1) {
				int maxHeight = 0;

				for (int j = 0; j < grid.get(i).size(); j += 1) {
					int[] rc = getRowCol(grid, i, j, direction);
					int row = rc[0];
					int col = rc[1];

					int height = grid.get(row).get(col);

					if (height > maxHeight) {
						maxHeight = height;
						visible.get(row).set(col, true);
					}

					// System.out.format("(%d, %d, %d) %d\n", row, col, height, maxHeight);

					if (row == 0 || col == 0 || row >= grid.size() - 1 || col >= grid.get(row).size() - 1) {
						visible.get(row).set(col, true);
					}
				}
			}
		}

		return visible;
	}

	public static int scenic_score(ArrayList<ArrayList<Integer>> grid, int row, int col) {
		int score = 1;

		for (Direction dir : Direction.values()) {
			int height = grid.get(row).get(col);

			int len = 1;

			while (true) {
				int r = row;
				int c = col;

				switch (dir) {
					case Right:
						r += len;
						break;
					case Down:
						c += len;
						break;
					case Left:
						r -= len;
						break;
					case Up:
						c -= len;
						break;
					default:
						System.out.println("Error");
				}

				
				if (!inRange(r, grid.size()) || !inRange(c, grid.get(0).size())) {
					break;
				}
				
				int curHeight = grid.get(r).get(c);
				
				len += 1;

				// System.out.printf("(%d, %d)\n", r, c);

				if (curHeight >= height) {
					break;
				}

			}

			score *= (len - 1);
		}

		return score;
	}

	public static int maxScenicScore(ArrayList<ArrayList<Integer>> grid) {
	// public static int maxScenicScore(ArrayList<ArrayList<Integer>> grid, ArrayList<ArrayList<Boolean>> visible) {
		int maxScore = 0;
		for (int r = 0; r < grid.size(); r += 1) {
			for (int c = 0; c < grid.get(r).size(); c += 1) {
				// if (visible.get(r).get(c)) {
				int score = scenic_score(grid, r, c);

				if (score > maxScore) {
					maxScore = score;
				}
				// }
			}
		}
		return maxScore;
	}

	public static void main(String[] args) throws FileNotFoundException {
		Scanner scanner = new Scanner(new File("8.txt"));

		ArrayList<ArrayList<Integer>> grid = new ArrayList<ArrayList<Integer>>();

		while (scanner.hasNextLine()) {
			ArrayList<Integer> row = new ArrayList<Integer>();

			String line = scanner.nextLine();
			for (int i = 0; i < line.length(); i += 1) {
				row.add(line.charAt(i) - '0');
			}

			grid.add(row);
		}

		// System.out.println(grid.toString());

		ArrayList<ArrayList<Boolean>> visible = visibility(grid);

		int visibleCount = 0;

		for (ArrayList<Boolean> row : visible) {
			for (boolean v : row) {
				if (v) {
					visibleCount += 1;
				}

				System.out.print(v ? 1 : 0);
			}
			System.out.println();
		}

		System.out.format("Visible Count = %d\n", visibleCount);

		System.out.format("Max Scenic score = %d\n", maxScenicScore(grid));
	}
}
