BIN_DIR = bin
SRC_DIR = src
INC_DIR = include
BUILD_DIR = build

TARGET = $(BIN_DIR)/app 
SRC = $(wildcard $(SRC_DIR)/*.c)
INCLUDE = -I ./$(INC_DIR)
OBJ = $(SRC:$(SRC_DIR)/%.cpp=$(BUILD_DIR)/%.o)
# OBJ = $(patsubst $(SRC_DIR)/%.cpp, $(BUILD_DIR)/%.o, $(SRC))
# OBJ = $(BUILD_DIR)/fun1.o $(BUILD_DIR)/fun2.o $(BUILD_DIR)/main.o

TEST_BIN_DIR = test/bin
TEST_SRC_DIR = test/src
TEST_BUILD_DIR = test/build

TEST_TARGET = $(TEST_BIN_DIR)/gtest
TEST_SRC = $(wildcard $(TEST_SRC_DIR)/*.cpp)
TEST_OBJ = $(TEST_SRC:$(TEST_SRC_DIR)/%.cpp=$(TEST_BUILD_DIR)/%.o)

TARGETS = $(TARGET) $(TEST_TARGET)

CC = g++
FLAGS = -w
TEST_FLAGS = -lgtest -lpthread

all: $(TARGETS)

$(TARGET): $(OBJ)
	@mkdir -p $(BIN_DIR)
	$(CC) $^ -o $@

$(BUILD_DIR)/%.o: $(SRC_DIR)/%.cpp
	@mkdir -p $(BUILD_DIR)
	$(CC) $(INCLUDE) $(FLAGS) -c $< -o $@

$(TEST_TARGET): $(TEST_OBJ) $(filter-out $(BUILD_DIR)/main.o, $(OBJ))
	@mkdir -p $(TEST_BIN_DIR)
	$(CC) $^ -o $@ $(TEST_FLAGS)

$(TEST_BUILD_DIR)/%.o: $(TEST_SRC_DIR)/%.cpp
	@mkdir -p $(TEST_BUILD_DIR)
	$(CC) $(INCLUDE) $(FLAGS) -c $< -o $@

clean:
	@rm -rf $(BIN_DIR)
	@rm -rf $(BUILD_DIR)
	@rm -rf $(TEST_BIN_DIR)
	@rm -rf $(TEST_BUILD_DIR)

clean-obj:
	rm -rf $(BUILD_DIR)
	rm -rf $(TEST_BUILD_DIR)

# run:
# 	make all
# 	./test/bin/gtest
