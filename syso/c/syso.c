// +build ignore
#include <stdint.h>

typedef struct {
	uint64_t a;
	uint64_t b;
	uint64_t *ret;
} params;

void mul(params p) {
	*p.ret = p.a * p.b;
}
