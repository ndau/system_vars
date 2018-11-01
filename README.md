## System Variables

The Ndau chain has the concept of system variables: values which live on the
Chaos chain, which control the behavior of the Ndau chain. These variables must
additionally have the ability to change their namespace and/or key over time,
as their permission levels change. This implies some sort of mapping from a
canonical name to a namespace-key pair. It's also easiest, for consistency
purposes, if the mapping itself also lives on the chaos chain.

We solve the chicken-and-egg problem of how to find the SVI map itself by putting
a namespace and key as mandatory data in the configuration file that every ndau
node must have in order to run.

## `pkg/system_vars`

This package contains canonical names and some helper types for various ndau
system variables

## `pkg/svi`

This package contains the fundamental struct and method definitons for the SVI
datatype.

## FAQ

### Why is this its own repo, instead of living in the ndau repo?

Circular dependencies.

The `genesisfile` package lives in the chaos repo because it fundamentally
defines the initial state of the chaos chain at genesis. It's sensible to put
it there.

The `genesisfile` package necessarily depends on the `svi` package, because one
of its more important responsibilities is to generate the SVI map and insert it
in the appropriate place on load.

We can't therefore put `svi` inside the `ndau` repo, because the `ndau` repo
depends on `genesisfile` to use as a mockfile when running without the chaos
chain; that would be a circular dependency, and we hates those.

In theory we could have put `svi` and `system_vars` into the `chaos` repo, but
it would have been silly: both of them are fundamentally ndau concepts.

We might have put `svi` into a common util repo such as `ndaumath`, but it
would have been a poor fit. It made more sense to start a new repo.

