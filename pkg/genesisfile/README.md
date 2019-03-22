# `genesisfile`

`genesisfile` describes a configuration file format used to represent the
system variables at genesis. This file is used for two purposes:
it both represents the desired initial sysvar state, and is used to mock the
state of the sysvars at genesis.

The data model in the genesisfile is more complicated than that of the actual
system variables. Internally to ndau, they are a map of names to byte slices,
which are then msgp-decoded into the appropriate types. That's not convenient
for humans, though.

Instead, a `Value` type is defined which contains an almost-arbitrary TOML value, the corresponding go type, and some metadata. This is then converted into the appropriate
binary data as necessary.

Valid `Value` types:

- primitives, in which case `Data` must be the correct primitive
- `[]byte`, in which case `Data` must be a string with the base64-encoded representation
- structs which implement `Valuable`
- lists of structs which implement `Valuable`
