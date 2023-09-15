#!/bin/bash
File="dic.txt"
Lines=$(cat $File)
for L in $Lines

do
	if [ ${#L} -eq 5 ]
	then
		echo "$L" >> out_dic.txt
	fi
done
