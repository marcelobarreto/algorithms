class Node {
  constructor(val) {
    this.val = val;
    this.next = null;
  }
}

class SinglyLinkedList {
  constructor() {
    this.head = null;
    this.tail = null;
    this.length = 0;
  }

  push(val) {
    // Add value at the end of the SLL
    var newNode = new Node(val);
    if (!this.head) {
      this.head = newNode;
      this.tail = this.head;
    } else {
      this.tail.next = newNode;
      this.tail = newNode;
    }

    this.length += 1;
    return this;
  }

  pop() {
    // Remove value from the end of the SLL
    if (!this.head) return undefined;
    var current = this.head;
    var newTail = current;
    while (current.next) {
      newTail = current;
      current = current.next;
    }
    this.tail = newTail;
    this.tail.next = null;
    this.length--;

    if (this.length === 0) {
      this.head = null;
      this.tail = null;
    }

    return current;
  }

  shift() {
    // Remove value from the start of the SLL
    if (!this.head) return undefined;
    var currentHead = this.head;
    this.head = currentHead.next;
    this.length--;
    if (this.length === 0) {
      this.tail = null;
    }
    return currentHead;
  }

  unshift(val) {
    // Add value at the start of the SLL
    var node = new Node(val);
    if (!this.head) {
      this.head = node;
      this.tail = node;
    } else {
      node.next = this.head;
      this.head = node;
    }
    this.length++;
    return node;
  }

  get(index) {
    // Get node at the index of the SLL
    if (index < 0 || index > this.length || this.length === 0) return null;

    if (index === 0) {
      return this.head;
    } else if (index === this.length - 1) {
      return this.tail;
    }

    var currentPosition = 0;
    var currentNode = this.head;

    while (currentPosition <= index) {
      currentPosition++;
      currentNode = currentNode.next;
      if (currentPosition >= index) {
        return currentNode;
      }
    }
  }

  set(index, value) {
    // Set node at the index of the SLL
    var node = this.get(index);
    if (!node) return false;
    node.val = value;
    return true;
  }

  insert(index, val) {
    // Insert node at the index of the SLL
    if (index < 0 || index > this.length) return false;
    if (index === this.length) return !!this.push(val);
    if (index === 0) return !!this.unshift(val);

    var newNode = new Node(val);
    var prev = this.get(index - 1);
    var temp = prev.next;
    prev.next = newNode;
    newNode.next = temp;
    this.length++;

    return true;
  }

  remove(index) {
    if (index < 0 || index >= this.length) return undefined;
    if (index === 0) return this.shift();
    if (index === this.length - 1) return this.pop();
    var previousNode = this.get(index - 1);
    var removed = previousNode.next;
    previousNode.next = removed.next;
    this.length--;
    return removed;
  }

  reverse() {
    var node = this.head;
    this.head = this.tail;
    this.tail = node;
    var next;
    var prev = null;
    for (var i = 0; i < this.length; i++) {
      next = node.next;
      node.next = prev;
      prev = node;
      node = next;
    }
    return this;
  }

  print() {
    var arr = [];
    var current = this.head
    while (current) {
      arr.push(current.val)
      current = current.next
    }
    console.log(arr);
  }
}

var list = new SinglyLinkedList();
list.push("Hello")
list.push("World")
list.push("!")

list.pop()
list.pop()

list.unshift("foo")
list.unshift("head!")

list.set(0, "set")

list.insert(1, "inserted")
list.insert(1, "inserted again")

list.remove(1)

var currentNode = list.head;

while (currentNode.next) {
  console.log(currentNode.val);
  if (!currentNode.next) {
    console.log(currentNode.next);
  }
  currentNode = currentNode.next;
}
console.log(list.tail.val);