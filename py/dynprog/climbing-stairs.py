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

def climbStairs2(n):
  '''
  1 = 1
  2 = 1+1, 2
  3 = 1+1+1, 2+1, 1+2
  4 = 1+1+1+1, 2+1+1, 1+2+1, 1+1+2, 2+2

  ways(i) = ways(i-1) + ways(i-2)
  '''
  m = [1, 2]
  if n <= 2:
    return m[n-1]
  for i in range (3,n+1):
    m[0], m[1] = m[1], m[0] + m[1]
  return m[1]

assert(climbStairs(1) == 1)
assert(climbStairs(2) == 2)
assert(climbStairs(3) == 3)
assert(climbStairs(4) == 5)

assert(climbStairs2(1) == 1)
assert(climbStairs2(2) == 2)
assert(climbStairs2(3) == 3)
assert(climbStairs2(4) == 5)
