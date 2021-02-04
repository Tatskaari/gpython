#!/bin/sh
set -e
rm -rf build
PYTHON=/opt/python3.8/bin/python3.8
if [ ! -f "$PYTHON" ]; then
    PYTHON="$(which python3.8)"
fi
"$PYTHON" setup.py build_ext --inplace
cp -av *.so *.dll ..
