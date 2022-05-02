import queue

class TreeNode:
  def __init__(self, val=0, left=None, right=None):
    self.val = val
    self.left = left
    self.right = right

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

t1 = TreeNode(1,
  TreeNode(2,
    TreeNode(3, None),
    None
  ),
  TreeNode(2,
    TreeNode(3, None),
    TreeNode(3,
      TreeNode(4, None),
      TreeNode(4, None),
    ),
  )
)

t2 = TreeNode(1, TreeNode(2, TreeNode(3, TreeNode(4, TreeNode(5, None), None), None), None), None)

assert(maxDepthRec(None) == maxDepthIter(None) == 0)
assert(maxDepthRec(t1) == maxDepthIter(t1) == 4)
assert(maxDepthRec(t2) == maxDepthIter(t2) == 5)
