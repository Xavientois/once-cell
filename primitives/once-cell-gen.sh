#!/bin/bash

SCRIPT_DIR="$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )"
TEMPLATE_PATH="$SCRIPT_DIR/once-cell.template.go"
TEST_TEMPLATE_PATH="$SCRIPT_DIR/once-cell-test.template_test.go"


if [ -z "$1" ]; then
    echo "Missing type parameter"
    exit 1
fi


if [ ! -z "$3" ]; then
    echo "Too many arguments passed"
    exit 1
fi

TYPE_NAME="$1"
CAMEL_TYPE="$(tr '[:lower:]' '[:upper:]' <<< ${TYPE_NAME:0:1})${TYPE_NAME:1}"
TEST_VALUE="$2"

if [ ! -f "$TEMPLATE_PATH" ]; then
    echo "Missing template file: $TEMPLATE_PATH"
    exit 2
fi

if [ ! -f "$TEST_TEMPLATE_PATH" ]; then
    echo "Missing test template file: $TEST_TEMPLATE_PATH"
    exit 2
fi

GEN_FILE="$SCRIPT_DIR/$TYPE_NAME.go"
GEN_TEST_FILE="$SCRIPT_DIR/${TYPE_NAME}_test.go"

sed "s/myType/$TYPE_NAME/g" "$TEMPLATE_PATH" |
sed "s/MyType/$CAMEL_TYPE/g" > "$GEN_FILE"

if [ ! -z "$TEST_VALUE" ]; then
    sed "s/myType{}/$TEST_VALUE/g" "$TEST_TEMPLATE_PATH" |
    sed "s/myType/$TYPE_NAME/g" |
    sed "s/MyType/$CAMEL_TYPE/g" > "$GEN_TEST_FILE"
fi 
