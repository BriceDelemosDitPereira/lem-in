# lem-in

## Objectives

The objective of this project is to create a digital version of an ant farm in Go. The program, `lem-in`, reads from a file describing the ants and the colony provided as arguments. It then finds the quickest path for the ants to traverse the colony and displays the process of each move they make from room to room.

## How It Works

- You create an ant farm with tunnels and rooms.
- Ants are placed on one side, and the program observes how they find the exit.
- The goal is to find the quickest way to get n ants across a colony, composed of rooms and tunnels.
- Initially, all ants are placed in the room `##start`, and the objective is to bring them to the room `##end` with the fewest moves possible.
- The program handles various scenarios, such as colonies with many rooms and links, no path between `##start` and `##end`, rooms linking to themselves, and other invalid or poorly-formatted inputs.

## Output Format

The results are displayed on the standard output in the following format:

```
number_of_ants
the_rooms
the_links

Lx-y Lz-w Lr-o ...
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

Example 1:

```
$ go run . test0.txt
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
```

Example 2:

```
$ go run . test1.txt
3
##start
0 1 0
##end
1 5 0
2 9 0
3 13 0
0-2
2-3
3-1

L1-2
L1-3 L2-2
L1-1 L2-3 L3-2
L2-1 L3-3
L3-1
$
```

Example 3:

```
$ go run . test1.txt
3
2 5 0
##start
0 1 2
##end
1 9 2
3 5 4
0-2
0-3
2-1
3-1
2-3

L1-2 L2-3
L1-1 L2-1 L3-2
L3-1
$
```