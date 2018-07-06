sprig-math
==========

Extends the excellent [sprig](https://github.com/Masterminds/sprig) library with
more complex math functions and error handling.

The Go language comes with a [built-in template language](http://golang.org/pkg/text/template/),
but not very many template functions. Sprig and this library provides a group
of commonly used template functions.

Sprig as-is does not return errors when executing mathematical conversions and
functions, and tries to allow execution to continue instead. This isn't ideal for
some applications, and that's where this library comes in.

Author
======

Dustin Spicuzza (dustin@virtualroadside.com)