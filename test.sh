#!/bin/sh
#for some odd reason, go test won't link the .syso file

echo compiling syso version
cd syso && gcc -c -O3 -march=native -o syso.syso c/syso.c && go build && ./syso; rm -f syso; cd ../

echo compiling cgo version
cd cgo && go build && ./cgo; rm -f cgo && cd ../
