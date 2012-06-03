gofmtcss
========

A command line helper for formatting .css files.

I built this in order to keep my css rules in alphabetical order.

It is used as a command line tool that, when passed in a path to a .css file, 
it will overwrite the file with better formatting and the rule sets in
alphabetical order.

<code>go run gofmtcss.go your.css</code>

Or compile it if you'd like.

It has some known issues with detecting comments, duplicate rules, and the like.

