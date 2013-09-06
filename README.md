RiverFS
=======

Goal: fuse mounted distributed filesystem in Go (golang).

This is still experimental and doesn't work yet.

Things that are done
----------------------
In memory directory tree structer mounted via fuse supporting:

ls, rm, mkdir, touch, cd, mv, cp

Place holder file code that ignores file contents.

Todo
-------
Synchronize dir tree state across nodes via raft

Reads and write file data via http (or?).



