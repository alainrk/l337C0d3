def climbStairs(n):
  '''
  1 = 1
  2 = 1+1, 2
  3 = 1+1+1, 2+1, 1+2
  4 = 1+1+1+1, 2+1+1, 1+2+1, 1+1+2, 2+2

  ways(i) = ways(i-1) + ways(i-2)
  '''
  m = [-1, 1, 2]
  if n < 3:
    return m[n]
  for i in range (3,n+1):
    m.append(m[i-1] + m[i-2])
  return m[-1]

assert(climbStairs(1) == 1)
assert(climbStairs(2) == 2)
assert(climbStairs(3) == 3)
assert(climbStairs(4) == 5)
