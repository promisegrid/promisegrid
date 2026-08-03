# PromiseGrid Grid CBOR Multicodec Specification

## Status

This document specifies the proposed `grid-cbor` multicodec. The requested
draft code is `0x1027`; that number is provisional until merged into the
multicodec registry.

## Terminology

- **Concise Binary Object Representation (CBOR):** the binary data format
  defined by RFC 8949.
- **Content Identifier (CID):** an identifier that includes a codec and a hash
  of the identified bytes.
- **InterPlanetary Linked Data (IPLD):** the content-addressed data model and
  tools that use CIDs to connect stored data.
- **multicodec:** the number in a CID that tells software which format and
  decoder apply to the identified bytes.
- **multihash:** the part of a CID that identifies the hash algorithm and
  contains the hash value.
- **Protocol CID (pCID):** the CID of a protocol specification. The pCID in a
  Grid envelope tells an agent how to interpret the protocol-defined slots.
- **Grid envelope:** the CBOR value defined by the
  [Grid CBOR tag specification](grid-cbor-tag-spec.md).
- **slot:** a numbered position in the Grid envelope's tagged CBOR array.
- **visible CID link:** a CBOR tag-42 CID that a decoder reaches by walking
  arrays, maps, and tags. Encoded CBOR inside a byte string is not visible to
  that walk.

## Background

More information about PromiseGrid is available at
[https://github.com/promisegrid/promisegrid](https://github.com/promisegrid/promisegrid).
The [Grid CBOR tag specification](grid-cbor-tag-spec.md) defines the
`grid` tag, the tagged array, the slot-0 pCID, and how that pCID
selects the meaning of the remaining slots. This document extends
those concepts, specifying how to encode/decode the complete envelope
and how software finds CID links in it.

## Purpose

The `grid-cbor` multicodec identifies the CBOR encoding of a
PromiseGrid Grid message or file envelope. It lets software recognize
that Grid format checks apply before decoding the block. It also lets
software find visible CID links without understanding the application
protocol carried by the envelope.

The multicodec and pCID answer different questions:  The multicodec
tells software how to decode the stored bytes. The pCID tells software
what the decoded message fields mean.

## Encoding Requirements

Every complete Grid envelope MUST use the
[Core Deterministic Encoding Requirements in RFC 8949 Section 4.2.1](https://www.rfc-editor.org/rfc/rfc8949.html#section-4.2.1),
with these requirements made explicit:

1. Integers, lengths, tags, and floating-point values use their preferred
   shortest serialization.
2. Indefinite-length items are forbidden.
3. Every map is sorted by bytewise lexicographic order of each key's
   deterministic encoding. RFC 7049 length-first ordering is not used.
4. Duplicate map keys and encodings that do not follow these rules are invalid.
5. Every NaN is encoded as `f97e00`. Positive and negative infinity use
   `f97c00` and `f9fc00`. Negative zero remains distinct from positive zero.
6. Encoding preserves differences that CBOR can represent. It does not change
   an integral float to an integer, normalize Unicode, or replace a tagged value
   with an untagged value merely because an application treats them as equal.
7. Arbitrary valid CBOR tags and map-key types are permitted. A decoder must
   preserve unknown tags and keys rather than silently discarding them.
8. A conforming codec can decode a valid block and encode it back to the same
   bytes without fetching or understanding the pCID.
9. A pCID may further restrict slot values, define their meaning and
   signatures, and select formats stored inside byte strings.

## CID Links

Every visible CBOR tag-42 value MUST contain a valid CID using the standard
tag-42 byte-string representation.

## Processing Requirements

A `grid-cbor` decoder SHOULD reject:

- a data item that is not a valid Grid envelope under the
  [Grid CBOR tag specification](grid-cbor-tag-spec.md);
- malformed visible tag-42 values;
- duplicate map keys, indefinite-length items, trailing top-level items, or any
  encoding that does not follow this specification;
- malformed CBOR or input that exceeds the decoder's stated size, depth, or
  item-count limits.

An unknown pCID does not prevent format checking, reproducing the same bytes, or
finding visible CID links. It does prevent an agent from knowing what the
protocol-defined slots mean. As required by the Grid CBOR tag specification, an
agent MUST NOT execute, authorize, or trust a message merely because its Grid
format is valid.

## Appendix A: Multicodec Registration Proposal

```text
Name: grid-cbor
Code: 0x1027
Tag: ipld
Status: draft
Description: Grid message envelope: RFC-8949 deterministic CBOR; 'grid'-tagged; codec preserves arbitrary tags/keys while allowing traversal of tag-42 CID links; spec: https://github.com/promisegrid/promisegrid/blob/main/docs/grid-cbor-multicodec-spec.md
```
