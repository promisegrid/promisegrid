# PromiseGrid Grid CBOR Tag Specification

## IANA Registration Status

This document is the intended public semantics document for an IANA CBOR
Tags registry request.

## Terminology

- **Concise Binary Object Representation (CBOR):** the binary data
  format whose semantic-tag registry this document targets.
- **Content Identifier (CID):** a content-addressed identifier in
  CIDv1 format for bytes or structured content.  In this document,
  CIDs are carried using CBOR tag 42.
- **InterPlanetary Linked Data (IPLD):** the content-addressed data
  model and tooling ecosystem whose CID representation is registered as
  CBOR tag 42.
- **Protocol CID (pCID):** a CID whose target content is a protocol
  specification.  It is roughly analogous to a port number, except the
  protocol specification's content hash is the protocol identifier
  rather than a centrally assigned number.
- **slot:** a zero-based position in the tagged CBOR array.  Slot 0 is
  the first array element.
- **grid envelope:** the CBOR value tagged by the `grid` tag.  The
  envelope carries a pCID in slot 0 and protocol-defined slots after
  that.

## Background

More information about PromiseGrid and the `grid` tag is available at
[https://github.com/promisegrid/promisegrid](https://github.com/promisegrid/promisegrid).

## Purpose

This document specifies the PromiseGrid `grid` CBOR tag for a
pCID-selected grid envelope.  The tag identifies the outer envelope;
the pCID in slot 0 selects the protocol specification that owns the
remaining slots.

## Tag Value

- Decimal: `1735551332`
- Hexadecimal: `0x67726964`
- CBOR tag header: `da 67 72 69 64`
- ASCII rendering of the four tag-number bytes: `grid`

The leading `da` byte is the CBOR tag header for a 32-bit tag number.
The bytes that encode the tag number itself are `67 72 69 64`, which
spell `grid`.

## Data Item

The tagged data item is a CBOR array.

```text
1735551332([42(pCID), ...protocol-defined-slots])
```

Diagnostic notation may use `grid(...)` as an alias for tag
`1735551332`:

```text
grid([42(pCID), ...protocol-defined-slots])
```

## Protocol Selection

The pCID selects how to parse and interpret the slots that follow slot
0.  It is the CID of a protocol spec document.

## Envelope Semantics

Slot 0 MUST be CBOR tag 42 applied to a byte string containing the
pCID.  The pCID is the content identifier for the protocol specification
that defines how to parse and interpret the remaining slots.

The pCID is wrapped in `42(...)` because tag 42 is the IPLD content
identifier tag for CID byte strings.  The outer `grid` tag identifies
the whole PromiseGrid envelope; the inner IPLD CID tag preserves the
standard CID wire representation and lets CID-aware CBOR tooling
recognize slot 0 as a CID before the receiver applies the narrower pCID
meaning to that CID.

After slot 0, the meaning and format of the remaining slots are
defined by the protocol specification identified by the pCID.

## Processing Requirements

A decoder that recognizes the `grid` tag MUST verify that the tagged
item is an array and that slot 0 is a tag 42 byte-string pCID before
applying any pCID-selected interpretation to the remaining slots.

An implementation MUST NOT execute, trust, or authorize a
pCID-selected protocol payload solely because a `grid` tag is present.
Trust, authorization, replay handling, signatures, proofs, and
resource limits are defined by the pCID-selected protocol.

If the pCID is unknown or unsupported, the implementation SHOULD reject
the envelope or preserve it as opaque data rather than guessing slot
semantics.

## Appendix A: IANA CBOR Tag Registration Request

```text
Registry: Concise Binary Object Representation (CBOR) Tags
Tag: 1735551332
Data item: array
Semantics: PromiseGrid grid envelope
Point of contact: Steve Traugott <ietf-poc@t7a.org>
Description of semantics: https://github.com/promisegrid/promisegrid/blob/main/docs/grid-cbor-tag-spec.md
```
