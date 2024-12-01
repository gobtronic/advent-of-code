#!/bin/bash

re='^[0-9]+$'
if ! [[ $1 =~ $re ]] ; then
    echo "error: expected a day number, got "$1 >&2; exit 1
fi

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
DAY=$1
DAY_DIR=$SCRIPT_DIR/day$DAY
TEMPLATE_DIR=$SCRIPT_DIR/day_template

if [ -d "$DAY_DIR" ]; then
    echo "$DAY_DIR directory already exists"
else
    mkdir $DAY_DIR
    cp $TEMPLATE_DIR/* $DAY_DIR
fi

