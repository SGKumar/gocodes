# How To Test

## From the commandline
From the parent of directory channelbw, just run `go test ./channelbw -v` to test this code with test cases in bandwidth.csv.

## Input
Test cases are provided in CSV file at `channelbw/bandwidth.csv`.

### CSV file format
First line contains a number T, the number of test cases.

Test cases follow, with each test case specified as below:
* First line contains a number S, the number of parallel channels/shows viewer is interested in
* The next **S** lines have 3 comma separated numbers 
  1. Start epoch of the show in channel (`int/long` value)
  2. End epoch of the show in channel (`int/long` value)
  3. Bandwidth required to watch this show (`int` value)
* Last item in the test case is the expected solution for Maximum Bandwidth required to watch parallel shows.

## Output
The output runs through each test case with the figured 3 solutions:
Representative output:
```
=== RUN   TestSolve1
--- PASS: TestSolve1 (0.00s)
=== RUN   TestSolve2
--- PASS: TestSolve2 (0.00s)
=== RUN   TestSolveFromCSV
=== RUN   TestSolveFromCSV/Solve1_1,3,9
=== RUN   TestSolveFromCSV/Solve2_1,3,9
=== RUN   TestSolveFromCSV/Solve1_2,4,11
=== RUN   TestSolveFromCSV/Solve2_2,4,11
=== RUN   TestSolveFromCSV/Solve1_3,5,12
=== RUN   TestSolveFromCSV/Solve2_3,5,12
=== RUN   TestSolveFromCSV/Solve1_4,6,14
=== RUN   TestSolveFromCSV/Solve2_4,6,14
=== RUN   TestSolveFromCSV/Solve1_5,1,5
=== RUN   TestSolveFromCSV/Solve2_5,1,5
=== RUN   TestSolveFromCSV/Solve1_6,2,13
=== RUN   TestSolveFromCSV/Solve2_6,2,13
=== RUN   TestSolveFromCSV/Solve1_7,0,0
=== RUN   TestSolveFromCSV/Solve2_7,0,0
--- PASS: TestSolveFromCSV (0.00s)
    --- PASS: TestSolveFromCSV/Solve1_1,3,9 (0.00s)
    --- PASS: TestSolveFromCSV/Solve2_1,3,9 (0.00s)
    --- PASS: TestSolveFromCSV/Solve1_2,4,11 (0.00s)
    --- PASS: TestSolveFromCSV/Solve2_2,4,11 (0.00s)
    --- PASS: TestSolveFromCSV/Solve1_3,5,12 (0.00s)
    --- PASS: TestSolveFromCSV/Solve2_3,5,12 (0.00s)
    --- PASS: TestSolveFromCSV/Solve1_4,6,14 (0.00s)
    --- PASS: TestSolveFromCSV/Solve2_4,6,14 (0.00s)
    --- PASS: TestSolveFromCSV/Solve1_5,1,5 (0.00s)
    --- PASS: TestSolveFromCSV/Solve2_5,1,5 (0.00s)
    --- PASS: TestSolveFromCSV/Solve1_6,2,13 (0.00s)
    --- PASS: TestSolveFromCSV/Solve2_6,2,13 (0.00s)
    --- PASS: TestSolveFromCSV/Solve1_7,0,0 (0.00s)
    --- PASS: TestSolveFromCSV/Solve2_7,0,0 (0.00s)
PASS
```
Each solution's output is compared with provided expected output and `Assert`ed.
Assert throws a failure/message only when the actual solution doesn't equal the expected value.

## Constraints
Normally none
Test Cases: 1 <= T <= 1000 
Parallel Channels/shows: 1 <= S <= 1000

## Sample Input
Here number of test cases T=3

Second test-case viewer is keen on 4 channels (S)
* Channels take bandwidth at 4, 6, 5, 11 respectively
* Correct solution for **maximum bandwidth** required is **`11`**
```
3
3
800, 1200, 4
600, 800, 6
900, 1400, 5
9
4
800, 1200, 4
600, 800, 6
900, 1400, 5
1500, 1600, 11
11
5
800, 1200, 4
600, 800, 6
900, 1400, 5
1500, 1600, 11
800, 900, 8
12
```
