.PHONY: wire

project_name=$(shell pwd | awk -F "/" '{print $$NF}')

wire:
	echo $(project_name)
	wire gen --recursion