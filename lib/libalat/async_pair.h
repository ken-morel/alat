#ifndef ASYNC_PAIR_H
#define ASYNC_PAIR_H

#include <stdbool.h>
#include <stdlib.h>

typedef void (*async_pair_request_callback)(int handle, const char *request_id,
                                            const char *token_json,
                                            const char *device_details_json);

void call_async_callback_bridge(async_pair_request_callback cb, int handle,
                                const char *request_id, const char *token,
                                const char *details);

#endif // ASYNC_PAIR_H
