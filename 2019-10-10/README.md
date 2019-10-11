### Problem of the Day:
You are given a perfect binary tree where all leaves are on the same level, and every parent has two children. The binary tree has the following definition:

```javascript
  /**
   * // Definition for a Node.
   * function Node(val,left,right,next) {
   *    this.val = val;
   *    this.left = left;
   *    this.right = right;
   *    this.next = next;
   * };
   */
```


Populate each next pointer to point to its next right node. If there is no next right node, the next pointer should be set to NULL.

Initially, all next pointers are set to NULL.

### Example:
```shell script
  Input:
       1
     /   \
    2     3
   / \   / \
  4   5 6   7


  Output:
        1 -> Null
      /    \
     /      \
    2   ->   3 -> Null
   /  \     / \
  4 -> 5 -> 6 -> 7 -> Null
```
