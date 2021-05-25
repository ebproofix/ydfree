#/bin/bash

curDir="$(pwd)/.."

yd_file="./data/yd.txt"
yd_dir="./testing/yd_dir"

if [ ! -d $yd_dir ]; then
    mkdir $yd_dir
fi

while read p; do 
    if [ ! -d "$yd_dir/$p" ]; then
        mkdir "$yd_dir/$p" ;
    fi

done < $yd_file
###########################################################################################
rc_file="./data/rec.txt"
rc_dir="./testing/rc_dir"

if [ ! -d $rc_dir ]; then
    mkdir $rc_dir
fi

while read p; do 
    if [ ! -d "$rc_dir/$p" ]; then
        mkdir "$rc_dir/$p" ;
    fi

done < $rc_file


echo "DONE!"