import queue

def maxDepthRec(root):
  if root is None:
    return 0
  return 1 + max(maxDepthRec(root.left), maxDepthRec(root.right))

def maxDepthIter(root):
  if not root:
    return 0
  maxd = 0
  q = queue.Queue()
  q.put({ "node": root, "h": 0 })
  while q.qsize():
    c = q.get()
    maxd = max(c["h"] + 1, maxd)
    if c["node"].left:
      q.put({ "node": c["node"].left, "h": c["h"] + 1 })
    if c["node"].right:
      q.put({ "node": c["node"].right, "h": c["h"] + 1 })
  return maxd