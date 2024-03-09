SRC = $(wildcard ./*.c)
OBJ = $(SRC:$(SRC_DIR)/%.cpp=$(BUILD_DIR)/%.o)

pthread: pthread.o

syscall: syscall.o

%.o: %.c
	gcc -o $@ $^

