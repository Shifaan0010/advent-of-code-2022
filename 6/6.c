#include <stdio.h>
#include <memory.h>
#include <inttypes.h>

int unique_count(int len, uint8_t counts[])
{
    int unique_count = 0;

    for (int i = 0; i < len; i += 1)
    {
        if (counts[i] > 1) {
            return -1;
        } else if (counts[i] == 1) {
            unique_count += 1;
        }
    }

    return unique_count;
}

int count_ones(uint64_t n)
{
    int count = 0;
    for (int i = 0; i < 64; i += 1, n >>= 1)
    {
        if (n & 1) {
            count += 1;
        }
    }
    return count;
}

void print_array(int len, uint8_t arr[])
{
    printf("[");
    for (int i = 0; i < len; i += 1)
    {
        printf("%d ", arr[i]);
    }
    printf("]\n");
}

#define BUFFER_SIZE 14

#define BIT_FLIPS

int main(int argc, char *argv[])
{
    FILE *file = fopen(argv[1], "r");

#ifdef BIT_FLIPS
    uint64_t filter = 0;
#else
    uint8_t counts[256];
    memset(counts, 0, 256);
#endif

    char buf[BUFFER_SIZE];
    int i = 0;

    while (1)
    {
        char ch = fgetc(file);

        if (ch == EOF) {
            break;
        }

#ifdef BIT_FLIPS
        filter ^= 1 << (ch % 64);
#else
        counts[ch] += 1;
#endif

        // must update before modifying buffer
        if (i >= BUFFER_SIZE) {
#ifdef BIT_FLIPS
            filter ^= 1 << (buf[(i - BUFFER_SIZE) % BUFFER_SIZE] % 64);
#else
            counts[buf[(i - BUFFER_SIZE) % BUFFER_SIZE]] -= 1;
#endif
        }

        buf[i % BUFFER_SIZE] = ch;

        // printf("%c ", ch);
        // print_array(256, counts);

#ifdef BIT_FLIPS
        if (count_ones(filter) == BUFFER_SIZE) {
#else
        if (unique_count(256, counts) == BUFFER_SIZE) {
#endif
            printf("%d\n", i);
            break;
        }

        i += 1;
    }
}
