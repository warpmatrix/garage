#include <stdio.h>
#include <unistd.h>
#include <pthread.h>

void *f(void *arg) {
    printf("arg: %p in %s\n", arg, __func__);
    sleep(3);
    return NULL;
}

int main(int argc, char const *argv[]) {
    pthread_t thd;
    int ret = pthread_create(&thd, NULL, f, NULL);
    pthread_exit(NULL);
    return 0;
}
