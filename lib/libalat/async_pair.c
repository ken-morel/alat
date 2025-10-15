#include "async_pair.h"
#include <stddef.h>

// Implementation of the C helper function.
// It checks for a NULL callback before attempting to call it.
void call_async_callback_bridge(async_pair_request_callback cb, int handle,
                                const char *request_id, const char *token,
                                const char *details) {
  if (cb != NULL) {
    cb(handle, request_id, token, details);
  }
}
