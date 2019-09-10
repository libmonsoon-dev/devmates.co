Problem of the Day:
Implement an iterator over a Binary Search Tree (BST). Your iterator will be initialized with the root node of a BST.
Calling next() will return the next smallest number in the BST.

```javascript

    /**
     * Definition for a binary tree node.
     * function TreeNode(val) {
     *     this.val = val;
     *     this.left = this.right = null;
     * }
     */
    /**
     * @param {TreeNode} root
     */
    var BSTIterator = function(root) {
        // Put your Code here
    };

    /**
     * @return the next smallest number
     * @return {number}
     */
    BSTIterator.prototype.next = function() {
        // Put your Code here
    };

    /**
     * @return whether we have a next smallest number
     * @return {boolean}
     */
    BSTIterator.prototype.hasNext = function() {
        // Put your Code here
    };

    /**
     * Your BSTIterator object will be
     * instantiated and called as such:
     * var obj = new BSTIterator(root)
     * var param_1 = obj.next()
     * var param_2 = obj.hasNext()
     */
```

Example:

```text
      7
    /   \
   3    15
        / \
       9   20
```

```javascript
  BSTIterator iterator = new BSTIterator(root);
  iterator.next();    // return 3
  iterator.next();    // return 7
  iterator.hasNext(); // return true
  iterator.next();    // return 9
  iterator.hasNext(); // return true
  iterator.next();    // return 15
  iterator.hasNext(); // return true
  iterator.next();    // return 20
  iterator.hasNext(); // return false
```
