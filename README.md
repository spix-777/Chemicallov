# Chemical lov
Overview

This program is a command-line tool that searches for a given chemical in the Norway government's list of controlled substances. It is written in Go and uses the flag, http, ioutil, os, strings, github.com/PuerkitoBio/goquery, and github.com/TwiN/go-color packages.

Requirements

The program has the following dependencies:

flag: for parsing command-line flags
http: for downloading the website containing the list of chemicals
ioutil: for reading and writing files
os: for interacting with the filesystem
strings: for manipulating strings
github.com/PuerkitoBio/goquery: for parsing HTML
github.com/TwiN/go-color: for coloring the output
Usage

To use the program, you need to compile it and then run the executable with the appropriate flags.

Copy code
go build main.go
./main -w <chemical>
The following flags are available:

-u: update the list of chemicals by downloading the latest version from the Norway government website
-w: search for the given chemical in the list of chemicals
-f: search for all the chemicals in the given file (one chemical per line)
Examples

Update the list of chemicals:

Copy code
./main -u
Search for a single chemical:

Copy code
./main -w ketamine
Search for multiple chemicals from a file:

Copy code
./main -f chemicals.txt
Limitations

The program has the following limitations:

It only searches in the Norway government's list of controlled substances.
It requires an internet connection to update the list of chemicals.
It uses a hard-coded URL for the website containing the list of chemicals. This URL may change in the future, causing the program to break.
It relies on the structure and formatting of the website and the HTML tables. If the website or the tables are changed, the program may not work correctly.
