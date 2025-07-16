## ✅ Combined TODO List for Linked List Implementation in Go

### 📦 Core Linked List Structure

* [ ] Define `Node` struct with `data int` and `next *Node`.
* [ ] Define `LinkedList` struct with `head *Node`.

---

### ➕ Insertion Operations

* [ ] `insertAtBack(data int)`

  * [ ] Insert a new node at the **end** of the linked list.
  * [ ] Handle case where list is empty (head is `nil`).
  * [ ] Traverse to the last node and link new node at the end.

* [ ] `insertAtFront(data int)`

  * [ ] Insert a new node at the **beginning** of the linked list.
  * [ ] Handle empty list by setting head to new node.
  * [ ] If not empty, point new node to current head and update head.

* [ ] `insertAfterValue(afterValue, data int)`

  * [ ] Insert a new node **after** the first node with value `afterValue`.
  * [ ] Traverse list to find `afterValue`.
  * [ ] Insert new node after the matching node.
  * [ ] Print an error if `afterValue` not found.

* [ ] `insertBeforeValue(beforeValue, data int)`

  * [ ] Insert a new node **before** the first node with value `beforeValue`.
  * [ ] Handle case where `beforeValue` is at the head.
  * [ ] Traverse list and find the node before `beforeValue`.
  * [ ] Insert new node before matching node.
  * [ ] Print an error if `beforeValue` not found.

---

### ❌ Deletion Operations

* [ ] `deleteFront()`

  * [ ] Delete the **first node** of the linked list.
  * [ ] Handle empty list safely.
  * [ ] Update head to point to next node.

* [ ] `deleteLast()`

  * [ ] Delete the **last node** of the linked list.
  * [ ] Handle empty list.
  * [ ] Handle case with only one node.
  * [ ] Traverse to second-last node and set its `next` to `nil`.

* [ ] `deleteAfterValue(afterValue int)`

  * [ ] Delete the node **after** the node with value `afterValue`.
  * [ ] Handle case where `afterValue` is not found.
  * [ ] Handle case where `afterValue` is the last node.
  * [ ] Update links to skip the node after `afterValue`.

* [ ] `deleteBeforeValue(beforeValue int)`

  * [ ] Delete the node **before** the node with value `beforeValue`.
  * [ ] Handle empty list or single-node list.
  * [ ] Handle deletion of head if it's the one before `beforeValue`.
  * [ ] Traverse and update links to remove node before match.
  * [ ] Print an error if `beforeValue` not found or deletion is invalid.

---

### 🔧 Utility Operations

* [ ] `countNodes()`

  * [ ] Traverse the list and count total number of nodes.
  * [ ] Return the count.

* [ ] `findNodeAt(index int) *Node`

  * [ ] Find and return the node at the specified 1-based index.
  * [ ] Validate the index (should be > 0 and ≤ total nodes).
  * [ ] Traverse list to reach desired index.

* [ ] `print()`

  * [ ] Traverse the list and print all node values in `data ->` format.
  * [ ] End with a newline after printing the list.

---

### 🏁 Main Function

* [ ] Create an empty linked list.
* [ ] Insert multiple values at the front.
* [ ] Print list contents.
* [ ] Count and print the number of nodes.
* [ ] Retrieve and print data at a specific index.
* [ ] Delete the last node.
* [ ] Print updated list.

---

Let me know if you want this exported to a Markdown checklist or added directly as `// TODO:` comments in the actual Go code.
