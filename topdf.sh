#!/bin/bash
echo "generating pdf of comments for entry with id: $2"
echo "looking through ld $1 entries ..."
go build ldcom.go && ./ldcom --num $1 --id $2 | pandoc -o "ld${1}_${2}.pdf"
echo "done"
