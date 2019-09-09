Problem of the Day:
Given an array of N sorted linked lists. Merge them and return it as one sorted linked list.

```javascript
    /**
    * Definition for singly-linked list.
    */
    function ListNode(val) {
        this.val = val;
        this.next = null;
    }
    /**
     * @param {ListNode[]} lists
     * @return {ListNode}
     */
    var mergeNLists = function(lists) {

    }
```

```text
Example:

  Input:
  [
    1 -> 10 -> 20,
    4 -> 11 -> 13,
    3 -> 8 -> 9
  ]

  Output: 1->3->4->8->9->10->11->13->20

  Input:
  [
    1->4->5,
    2->6
  ]

  Output: 1->2->4->5->6
```