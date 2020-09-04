maybe
=====

Provides optional values in go.

Representing optional values as pointers can create many small chunks on the
heap. This package represents optional values without heap allocations.

This API is stable. Newly added API that is not yet stable will be clearly
marked as draft API.

## Using

	import "github.com/keep94/maybe"
