#!/bin/bash

suffix=.go
source_file_list=$(ls *.go)

for file in $source_file_list; do
  target_list="$target_list ${file%$suffix}"
  command_list="$command_list\n${file%$suffix}:\n\tgo build $file"
done

echo -e ".PHONY : ALL\nALL : $target_list" > Makefile
echo -e "$command_list" >> Makefile
echo -e "\n.PHONY : clean\nclean :\n\t rm -rf $target_list" >> Makefile

echo -e "done successfully!"
