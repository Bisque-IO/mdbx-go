package main

import (
	"os"
	"strings"
)

const usage = `Bisque libmdbx B-Tree Storage Engine Utilities

chk [-V] [-v] [-q] [-c] [-0|1|2] [-w] [-d] [-i] [-s subdb] dbpath
	-V         print version and exit
	-v         more verbose, could be used multiple times
	-q         be quiet
	-c         force cooperative mode (don't try exclusive)
	-w         write-mode checking
	-d         disable page-by-page traversal of B-tree
	-i         ignore wrong order errors (for custom comparators case)
	-s subdb   process a specific subdatabase only
	-0|1|2     force using specific meta-page 0, or 2 for checking
	-t         turn to a specified meta-page on successful check
	-T         turn to a specified meta-page EVEN ON UNSUCCESSFUL CHECK!


stat [-V] [-q] [-e] [-f[f[f]]] [-r[r]] [-a|-s name] dbpath
	-V         print version and exit
	-q         be quiet
	-p         show statistics of page operations for current session
	-e         show whole DB info
	-f         show GC info
	-r         show readers
	-a         print stat of main DB and all subDBs
	-s name    print stat of only the specified named subDB
                   by default print stat of only the main DB

copy [-V] [-q] [-c] [-u|U] src_path [dest_path]
	-V         print version and exit
	-q         be quiet
	-c         enable compactification (skip unused pages)
	-u         warmup database before copying
	-U         warmup and try lock database pages in memory before copying
	src_path   source database
	dest_path  destination (stdout if not specified)


dump [-V] [-q] [-f file] [-l] [-p] [-r] [-a|-s subdb] [-u|U] dbpath
	-V         print version and exit
	-q         be quiet
	-f         write to file instead of stdout
	-l         list subDBs and exit
	-p         use printable characters
	-r         rescue mode (ignore errors to dump corrupted DB)
	-a         dump main DB and all subDBs
	-s name    dump only the specified named subDB
	-u         warmup database before dumping
	-U         warmup and try lock database pages in memory before dumping
                   by default dump only the main DB


drop %s [-V] [-q] [-d] [-s name] dbpath
	-V	   print version and exit"
	-q	   be quiet"
	-d	   delete the specified database, don't just empty it"
	-s	   name\tdrop the specified named subDB"
		   by default empty the main DB


load %s [-V] [-q] [-a] [-f file] [-s name] [-N] [-p] [-T] [-r] [-n] dbpath
	-V        print version and exit
	-q        be quiet
	-a        append records in input order (required for custom comparators)
	-f file   read from file instead of stdin
	-s name   load into specified named subDB
	-N        don't overwrite existing records when loading, just skip ones
	-p        purge subDB before loading
	-T        read plaintext
	-r        rescue mode (ignore errors to load corrupted DB dump)
	-n        don't use subdirectory for newly created database (MDBX_NOSUBDIR)
`

const usageChk = `Bisque libmdbx B-Tree Storage Engine Utilities

mdbx chk [-V] [-v] [-q] [-c] [-0|1|2] [-w] [-d] [-i] [-s subdb] dbpath
	-V           print version and exit
	-v           more verbose, could be used multiple times
	-q           be quiet
	-c           force cooperative mode (don't try exclusive)
	-w           write-mode checking
	-d           disable page-by-page traversal of B-tree
	-i           ignore wrong order errors (for custom comparators case)
	-s subdb     process a specific subdatabase only
	-0|1|2       force using specific meta-page 0, or 2 for checking
	-t           turn to a specified meta-page on successful check
	-T           turn to a specified meta-page EVEN ON UNSUCCESSFUL CHECK!
`

const usageStat = `Bisque libmdbx B-Tree Storage Engine Utilities

mdbx stat [-V] [-q] [-e] [-f[f[f]]] [-r[r]] [-a|-s name] dbpath
	-V           print version and exit
	-q           be quiet
	-p           show statistics of page operations for current session
	-e           show whole DB info
	-f           show GC info
	-r           show readers
	-a           print stat of main DB and all subDBs
	-s name      print stat of only the specified named subDB
                     by default print stat of only the main DB
`

const usageCopy = `Bisque libmdbx B-Tree Storage Engine Utilities

mdbx copy [-V] [-q] [-c] [-u|U] src_path [dest_path]
	-V          print version and exit
	-q          be quiet
	-c          enable compactification (skip unused pages)
	-u          warmup database before copying
	-U          warmup and try lock database pages in memory before copying
	src_path    source database
	dest_path   destination (stdout if not specified)
`

const usageDump = `Bisque libmdbx B-Tree Storage Engine Utilities

mdbx dump [-V] [-q] [-f file] [-l] [-p] [-r] [-a|-s subdb] [-u|U] dbpath
	-V         print version and exit
	-q         be quiet
	-f         write to file instead of stdout
	-l         list subDBs and exit
	-p         use printable characters
	-r         rescue mode (ignore errors to dump corrupted DB)
	-a         dump main DB and all subDBs
	-s name    dump only the specified named subDB
	-u         warmup database before dumping
	-U         warmup and try lock database pages in memory before dumping
                   by default dump only the main DB,
`

const usageDrop = `Bisque libmdbx B-Tree Storage Engine Utilities

mdbx drop %s [-V] [-q] [-d] [-s name] dbpath
	-V	    print version and exit"
	-q	    be quiet"
	-d	    delete the specified database, don't just empty it"
	-s	    name\tdrop the specified named subDB"
		    by default empty the main DB
`

const usageLoad = `Bisque libmdbx B-Tree Storage Engine Utilities

mdbx load %s [-V] [-q] [-a] [-f file] [-s name] [-N] [-p] [-T] [-r] [-n] dbpath
  -V        print version and exit
  -q        be quiet
  -a        append records in input order (required for custom comparators)
  -f file   read from file instead of stdin
  -s name   load into specified named subDB
  -N        don't overwrite existing records when loading, just skip ones
  -p        purge subDB before loading
  -T        read plaintext
  -r        rescue mode (ignore errors to load corrupted DB dump)
  -n        don't use subdirectory for newly created database (MDBX_NOSUBDIR)
`

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		println(usage)
		os.Exit(1)
	}

	switch strings.ToLower(args[0]) {
	default:
		println(usage)
		os.Exit(1)
	}
}
