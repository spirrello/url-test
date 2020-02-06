#!/bin/bash

echo "#################################################"

case $1 in

  env)
    echo "All ENV vars: "
    printenv
    ;;

  ls)
    echo "Files in workspace:"
    ls -ltr
    ;;

  *)
    echo "No valid args provided"
    ;;
esac

echo "#################################################"
