#!/bin/bash

echo "#################################################"

case $1 in

  env)
    #echo "All ENV vars: "
    echo "Branch: " $BRANCH_NAME
    echo "Git Tag: " $TAG_NAME
    # echo "All ENV vars: "
    # printenv
    ;;

  ls)
    echo "Files in workspace:"
    ls -ltr
    ;;

  *)
    echo "Unknown Argument!"
    exit 1
    ;;
esac

echo "#################################################"
