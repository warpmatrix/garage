#include <signal.h>
#include <sys/syscall.h>
#include <unistd.h>

int main() {
    pid_t pid = syscall(SYS_getpid);
    syscall(SYS_tgkill, getpid(), pid, SIGHUP);
    return 0;
}
