/*************************************************************************
	> File Name: main.c
	> Author: 
	> Mail: 
	> Created Time: Sun 28 Jun 2015 08:33:10 PM CST
 ************************************************************************/

#include <stdio.h>

#include <unistd.h>
#include <signal.h>


void handler(int num)
{
    printf("recv = %d\n",num);
}

int main(int argc,char **argv)
{
    unsigned int n = 100;

    printf("pid = %d\n",getpid());

    signal(SIGINT,handler);

    while(n > 0)
    {
        n = sleep(n);
        printf("n = %d\n",n);
    }
    return 0;
}
