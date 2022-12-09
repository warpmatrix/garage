#!/bin/bash
nicknamefile=res/nickname.log
if test -r $nicknamefile; then
	source $nicknamefile
else
	echo "nickname.log doesn't exists" >&2
	exit -1
fi

function showNickNames() {
	echo ${#nicknames[*]}
	for i in `seq 0 ${#nicknames[*]}`; do
		echo "${nicknames[i]}" "${realnames[i]}"
	done
}

function createFinList() {
	# for line in $(cat data.log | grep -e ".*:" | sed -E 's/(.*):(.*)/\1x/g'); do 
	charfile=chat.log
	cat $charfile | while read line; do
		line="$(echo $line | grep -e ".*:" | sed -E 's/(.*):.*/\1:/g')"
		if [ -z "$line" ]; then
			continue
		fi
		if [ "$line" == ":" ]; then
			line="$specname:"
		fi
		for i in `seq 0 ${#nicknames[*]}`; do
			if [ "${line:0:${#line}-1}" == "${nicknames[i]}" ]; then
				line="${realnames[i]}:"
			fi
		done
		echo "${line:0:${#line}-1}"
	done
}

function substNickName() {
    if [ "$1" == ":" ]; then
        echo "$specname:"
        return
    fi
    for i in `seq 0 ${#nicknames[*]}`; do
        if [ "${1:0:${#1}-1}" == "${nicknames[i]}" ]; then
            echo "${realnames[i]}:"
            return
        fi
    done
    echo "$1"
}

for name in `echo $(createFinList) | sort`; do
    echo $name >> finList.log
done
# cat finList.log | sort | uniq > sorted.log
cat finList.log | sort > sorted.log
# cat finList.log | sort | uniq -c | sort | awk '{ print $2 " " $1 }' > count.log
OS=`uname`
listfile=res/list.log
if [ ! -r sortedList-$OS.log ]; then
	cat $listfile | sort > sortedList-$OS.log
fi
diff --strip-trailing-cr sorted.log sortedList-$OS.log
rm -f sorted.log finList.log
