Simple Log Structured merge tree based key value store.  This server exposes a
REST API that accepts three operations, PUT, GET, DELETE.

PUT - adds a new key
DELETE - Removes a key
GET - Obtain the value for a key.

Clearly, this code is for educational purposes only.  Please don't use it for
anything aside from that purpose.

Design
* memtree is a homebrewed Red/Black tree
* Deleted nodes are marked tombstoned

* sstables will be formatted as JSON as the goal with this is education, not
  actual performance

References:

The following were used heavily:

* https://static.googleusercontent.com/media/research.google.com/en//archive/bigtable-osdi06.pdf
* https://github.com/facebook/rocksdb/
* https://www.cs.umb.edu/~poneil/lsmtree.pdf
* https://www.geeksforgeeks.org/red-black-tree-set-1-introduction-2/
