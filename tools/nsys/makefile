src := $(wildcard ./*.nsys-rep)
target := $(src:./%.nsys-rep=./%.sqlite)

all: $(target)

report1.nsys-rep: demo.py
	nsys profile python $< -o $@ -f true

%.sqlite: %.nsys-rep
	nsys export --type sqlite -f true -o $@ $<
