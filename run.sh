#!/bin/bash

echo "#################################################"

case $1 in

  env)
    echo "All ENV vars: "
    echo "branch: " ${BRANCH_NAME}
    printenv
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
