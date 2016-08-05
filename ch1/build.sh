#!/bin/bash
suffix=.go
for file in $(ls | grep -E "^*\.go$"); do
  target=${file%%$suffix}
  target_list="$target_list ${file%%$suffix}"
  cmd_list="$cmd_list\n$target :\n\tgo build ${file}\n"
done

echo -e ".PHONY : ALL" > Makefile
echo -e "ALL : $target_list" >> Makefile
echo -e "$cmd_list" >> Makefile
echo -e ".PHONY : clean\nclean:\n\trm -rf $target_list" >> Makefile
