# lem-in

## Objective

The objective of this project is to create a digital version of an ant farm in Go. The program, `lem-in`, reads from a file describing the ants and the colony provided as arguments. It then finds the quickest path for the ants to traverse the colony and displays the process of each move they make from room to room.


## How It Works

- You create an ant farm with tunnels and rooms.
- Ants are placed on one side, and the program observes how they find the exit.
- The goal is to find the quickest way to get n ants across a colony, composed of rooms and tunnels.
- Initially, all ants are placed in the room `##start`, and the objective is to bring them to the room `##end` with the fewest moves possible.
- The program handles various scenarios, such as colonies with many rooms and links, no path between `##start` and `##end`, rooms linking to themselves, and other invalid or poorly-formatted inputs.


## Usage

Compile and run the program by providing the input file as a command-line argument:

Executable :

```bash
$ go build lem-in.go
$ ./lem-in input_file
```

Or without executable :
```bash
$ go run . input_file
```

Each audit test can be done with : (in this file each test is run with `time go run` for check the speed of the code)
```bash
$ ./audit.sh
```


## Output Format

The results are displayed on the standard output in the following format:

```
number_of_ants
the_rooms
the_links

Lx-y Lz-w Lr-o ...

Number of turns: ...
```

- `x, z, r` represent the ants' numbers (from 1 to number_of_ants).
- `y, w, o` represent the rooms' names.
- A room is defined by "name coord_x coord_y".
- The links are defined by "name1-name2".


## Instructions

- Create tunnels and rooms.
- A room must never start with the letter L or # and must have no spaces.
- Join rooms together with tunnels.
- A tunnel joins only two rooms together, never more.
- Each room can only contain one ant at a time.
- Ants must take the shortest paths and avoid traffic jams.
- Display only the ants that moved at each turn.
- The program must handle errors carefully and not quit unexpectedly.
- Rooms' coordinates will always be integers.
- Write the project in Go, respecting good practices.
- Recommended to have test files for unit testing.


## Usage

Example 1

```
$ go run . test/autre/test0.txt
3
##start
1 23 3
2 16 7
3 16 3
4 16 5
5 9 3
6 1 5
7 4 8
##end
0 9 5
0-4
0-6
1-3
4-3
5-2
3-5
4-2
2-1
7-6
7-2
7-4
6-5

L1-3 L2-2
L1-4 L2-5 L3-3
L1-0 L2-6 L3-4
L2-0 L3-0
$
Number of turns: 4
```


## Error Handling

The program carefully handles errors such as invalid data format, invalid number of ants, missing start or end room, duplicated rooms, and other issues. It returns informative error messages to guide the user in resolving input problems.


## Implementation Details

The project is implemented in Go and utilizes the backtracking method to find the shortest path for the ants. The code follows good programming practices to ensure readability, maintainability, and efficiency.


## Co-developers

- [Delestre Thomas](https://github.com/Thomas-Delestre)
- [Marchais Mickael](https://github.com/Jeancrock)