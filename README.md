RiverFS
=======

Distributed FUSE+Raft filesystem in Go (golang).

This is still experimental and doesn't work yet.

Goals
-------
Mounts via FUSE as a normall looking file directory.
Files are stored as normal files accross a cluster.
Directory tree stucture is stored in memory on each node and
kept in consensous via raft.
Effichent streaming of large files.