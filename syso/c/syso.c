// +build ignore
#include <stdint.h>

typedef struct {
	uint64_t a0;
	uint64_t a1;
	uint64_t a2;
	uint64_t a3;
	uint64_t a4;
	uint64_t *ret;
} params;

void mul(params p) {
		*p.ret = p.a0 * p.a1 * p.a2 * p.a3 * p.a4;
}
