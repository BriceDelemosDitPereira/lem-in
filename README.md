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

***********************************************************************************************


# Lem-in Project

## Overview

Lem-in is a project that challenges you to create a digital version of an ant farm. The goal is to guide a specified number of ants from the starting room to the ending room in a colony, using the shortest path possible. The project involves reading input from a file, processing the data, and outputting the movements of the ants as they traverse the colony.

## Table of Contents

- [How it Works](#how-it-works)
- [Input Format](#input-format)
- [Output Format](#output-format)
- [Usage](#usage)
- [Example](#example)
- [Error Handling](#error-handling)
- [Implementation Details](#implementation-details)
- [Good Practices](#good-practices)

## How it Works

1. You create an ant farm consisting of rooms and tunnels.
2. Ants are initially placed in the ##start room.
3. The objective is to guide the ants to the ##end room with the fewest moves.
4. The program finds the quickest path using the backtracking method.
5. The result is displayed, showcasing the ant movements from room to room.

## Input Format

The input file should describe the ant colony with the following format:

```
number_of_ants
the_rooms
the_links
```

Rooms are defined by "name coord_x coord_y," and links are defined by "name1-name2."

## Output Format

The program displays the results on the standard output in the following format:

```
number_of_ants
the_rooms
the_links
Lx-y Lz-w Lr-o ...
```

## Usage

Compile and run the program by providing the input file as a command-line argument:

```bash
$ go build lem-in.go
$ ./lem-in input_file
```

## Example

```plaintext
##start
1 23 3
2 16 7
#comment
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
#another comment
4-2
2-1
7-6
7-2
7-4
6-5
```

## Error Handling

The program carefully handles errors such as invalid data format, invalid number of ants, missing start or end room, duplicated rooms, and other issues. It returns informative error messages to guide the user in resolving input problems.

## Implementation Details

The project is implemented in Go and utilizes the backtracking method to find the shortest path for the ants. The code follows good programming practices to ensure readability, maintainability, and efficiency.

## Good Practices

The code adheres to best practices in Go programming, emphasizing clear code structure, meaningful variable names, and proper error handling. It aims to be efficient and well-documented for ease of understanding.