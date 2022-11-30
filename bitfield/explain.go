package bitfield

// One of the most interesting types of message is the bitfield,
// which is a data structure that peers use to efficiently encode
//  which pieces they are able to send us. A bitfield looks like a byte array,
//  and to check which pieces they have, we just need to look at the positions
//  of the bits set to 1. You can think of it like the digital equivalent of a coffee shop loyalty card.
// We start with a blank card of all 0, and flip bits to 1 to mark their positions as stamped.

// By working with bits instead of bytes, this data structure is super compact.
// We can stuff information about eight pieces in the space of a single byte—the
// size of a bool. The tradeoff is that accessing values becomes a little more tricky.
//  The smallest unit of memory that computers can address are bytes, so to get to our bits,
//   we have to do some bitwise manipulation:
