### Problem of the Day:
A robot is located at the top-left corner of a m x n grid.

The robot can only move either down or right at any point in time. The robot is trying to reach the bottom-right corner of the grid.

How many possible unique paths are there?

### Example:
```shell script
  m = 3, n = 2, S= Start, F=Finish
  [
    [S,0,0],
    [0,0,F]
  ]
```

### Example:
```shell script
  Input: m = 3, n = 2
  Output: 3
  Why?:
  From the top-left corner, there are a total of 3 ways to reach the bottom-right corner:
    1. Right -> Right -> Down
    2. Right -> Down -> Right
    3. Down -> Right -> Right

  Input: m = 7, n = 3
  Output: 28
```
