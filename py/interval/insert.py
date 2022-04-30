hline = '\n' + ('-' * 40) + '\n'

def insert(intervals, n):
  print(f"{hline}Solving for intervals: {intervals}, {n}")
  res = []
  i = 0 # needed after loop as well
  for i, (x, y) in enumerate(intervals):
    print(f"\n-Loop i = {i}, x,y = ({x},{y})")
    # newInterval is after current one
    if n[0] > y:
      print(f"n[0] > y [{n[0]} > {y}] -- add [{x}, {y}] to result")
      res.append([x, y])
    # newInterval intersec current one, exit loop and merge pieces
    elif n[1] < x:
      print(f"n[1] < x [{n[1]} < {x}] -- decrement i to {i-1}")
      i -= 1
      break
    # set the newInterval using the current values and go on
    else:
      n[0] = min(n[0], x)
      n[1] = max(n[1], y)
      print(f"set new interval and go on...")
  # concat res so far, the new interval and the rest of the intervals
  res += [n] + intervals[i+1:]
  print(f"result = {res} + {[n]} + {intervals[i+1:]}")
  return res

assert(insert([[1, 3], [3, 5]], [2, 4]) == [[1, 5]])
assert(insert([[1, 3], [7, 8]], [4, 6]) == [[1, 3], [4, 6], [7, 8]])
assert(insert([[1, 3], [7, 8]], [-2, 0]) == [[-2, 0], [1, 3], [7, 8]])
assert(insert([[1, 3], [7, 8]], [10, 12]) == [[1, 3], [7, 8], [10, 12]])
assert(insert([[1, 3], [7, 8]], [2, 5]) == [[1, 5], [7, 8]])
assert(insert([[1, 3], [7, 8]], [0, 5]) == [[0, 5], [7, 8]])
assert(insert([[1, 3], [7, 8]], [8, 10]) == [[1, 3], [7, 10]])
assert(insert([[1, 3], [7, 8]], [0, 10]) == [[0, 10]])
