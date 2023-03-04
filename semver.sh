#! /bin/bash
#
# Gets the last (alphabetically) git tag that is "SemVer/*", increments the patch
# number, and creates a new git tag with that patch number.
#
# Increment the major or minor by specifiying "semver major" or "semver minor"
#
# Original version of this from https://jon.sprig.gs/blog/post/1175, but modified
# to use tags labelled "SemVer/*" and to create a new label.
#

RE='[^0-9]*\([0-9]*\)[.]\([0-9]*\)[.]\([0-9]*\)\([0-9A-Za-z-]*\)'

step="$1"
if [ -z "$1" ]
then
  step=patch
fi

base="$2"
if [ -z "$2" ]
then
  base=$(cat CURRENT_VERSION)
  if [ -z "$base" ]
  then
    base=0.0.0
  fi
fi

MAJOR=`echo $base | sed -e "s#$RE#\1#"`
MINOR=`echo $base | sed -e "s#$RE#\2#"`
PATCH=`echo $base | sed -e "s#$RE#\3#"`

echo "Current version: $MAJOR.$MINOR.$PATCH"

case "$step" in
  major)
    let MAJOR+=1
    ;;
  minor)
    let MINOR+=1
    ;;
  patch)
    let PATCH+=1
    ;;
esac

echo "New version: $MAJOR.$MINOR.$PATCH"
echo $MAJOR.$MINOR.$PATCH > CURRENT_VERSION

echo "package version" > version.go
echo "var CurrentVersion = \"$MAJOR.$MINOR.$PATCH\"" >> version.go

