target:
	g++ -o functions.o -c functions.cpp
	g++ -o functions-test.o -c functions-test.cpp
	g++ -o test-all.o -c test-all.cpp
	g++ -o test *.o -L lib -lgtest -lpthread

clean:
	rm *.o
	rm test
