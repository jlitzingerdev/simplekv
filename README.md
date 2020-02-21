Simple Log Structured merge tree based key value store.

# API
For initial simplicity this server exposes a HTTP API that accepts three
operations, PUT, GET, DELETE.  While the underlying database apis are written
to support any byte string, the inital implementation is only going to work
well with strings.  This is a deliberate simplification meant to be built upon
as I have interest in doing so.

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
