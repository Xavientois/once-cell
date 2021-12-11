#!/bin/bash

SCRIPT_DIR="$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )"
GEN_SCRIPT="$SCRIPT_DIR/once-cell-gen.sh"

# Generate oncecell types for all primitive types
"$GEN_SCRIPT" bool          true
"$GEN_SCRIPT" string        "\"test_value\""
"$GEN_SCRIPT" int           -4
"$GEN_SCRIPT" int8          -8
"$GEN_SCRIPT" int16         -16
"$GEN_SCRIPT" int32         -32
"$GEN_SCRIPT" int64         -64
"$GEN_SCRIPT" uint          4
"$GEN_SCRIPT" uint8         8
"$GEN_SCRIPT" uint16        16
"$GEN_SCRIPT" uint32        32
"$GEN_SCRIPT" uint64        64
"$GEN_SCRIPT" float32       0.32
"$GEN_SCRIPT" float64       0.64
"$GEN_SCRIPT" complex64     "complex(6,4)"
"$GEN_SCRIPT" complex128    "complex(1,28)"
