#!/bin/sh

# initializing variable via arguement value
var=$@

# Printing arguement
echo $var

# Length of arg
lenVar=${#var}

echo "Length = $lenVar"

# Print reverse of String 
revVar=$(echo $var | rev)
echo Reversed = $revVar
