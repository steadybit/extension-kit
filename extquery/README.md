# extquery

Parses and evaluates Steadybit's target query language — the language used for environment scopes,
experiment blast radii, action target selections and advice assessment queries:

```
k8s.namespace="prod" AND NOT k8s.label.tier="canary"
```

Two things it is for:

- **`Validate`** — check the query strings an extension ships in its own action and advice
  definitions, at build or start-up time rather than at platform registration.
- **`Parse`** — evaluate a query against a target's attributes inside the extension, so targets an
  operator does not want can be dropped before they ever reach the agent.

## Why validate

Extensions embed query strings in their action and advice definitions:

| Kit        | Field                                        |
|------------|----------------------------------------------|
| advice-kit | `AdviceDefinition.AssessmentQueryApplicable` |
| advice-kit | `AdviceDefinition.AssessmentQueryExclude`    |
| advice-kit | `AdviceDefinitionStatus…AssessmentQuery`     |
| action-kit | `TargetSelection.TargetQuery`                |
| action-kit | `TargetSelectionTemplate.Query`              |

A malformed query in any of them is only discovered by the platform at registration time — against
an already-released extension. `extquery` lets an extension catch it in its own tests instead.

## Validating

```go
import "github.com/steadybit/extension-kit/extquery"

func TestAdviceQueriesAreValid(t *testing.T) {
    definition := getAdviceDefinition()
    require.NoError(t, extquery.ValidateAll(
        definition.AssessmentQueryApplicable,
        extutil.ToString(definition.AssessmentQueryExclude),
    ))
}
```

`Validate` returns a `*ParseError` carrying `Message`, `Line` (1-based) and `Column` (0-based):

```go
err := extquery.Validate("k8s.namespace=")
// failed to parse query at line 1 column 14: missing {INTEGER, QUOTED, VARIABLE, PLACEHOLDER, TERM} at '<EOF>'
```

`ValidateAll` joins the failures of several queries so a whole definition reports all of its
problems at once.

### What is and is not checked

- An **empty query is valid** and means "match everything", consistent with the platform.
- **Line comments are supported**: `//` to the end of the line, on its own line or trailing a query.
  They are lexed onto the hidden channel, so they never affect what a query selects. Note a query
  made up of nothing but comments has no tokens on the default channel, which makes it the empty
  query — and so match-all, not match-nothing.
- **Only syntax** is checked. Whether the referenced attributes exist, or whether the query selects
  anything, is not knowable here.
- **Variable (`{{key}}`) and template-placeholder (`[[key]]`) markers are rejected**, quoted or
  not. The grammar accepts them, but the platform resolves them against experiment and environment
  state an extension has no access to, so a query shipped inside an extension can never
  legitimately contain one.

## Evaluating

`Parse` compiles a query into a `Predicate` that matches against a target's attributes. Parsing is
the expensive part and matching is not, so parse once at start up and reuse the predicate for every
target — predicates are immutable and safe for concurrent use.

```go
filter, err := extquery.Parse(os.Getenv("STEADYBIT_DISCOVERY_EXCLUDE_QUERY"))
if err != nil {
    log.Fatal().Err(err).Msg("invalid discovery exclude query")
}
log.Info().Msgf("discovery filter active: %s", filter)

// ... per target
if filter.Matches(extquery.MapAttributes(target.Attributes)) {
    continue // excluded
}
```

`Attributes` is the only thing a predicate needs — a multi-valued attribute lookup:

```go
type Attributes interface {
    Values(key string) []string
}
```

Two adapters ship with the package: `MapAttributes` for a plain `map[string][]string`, and
`AttributesFunc` for a lookup function — useful to inject keys that are not in the discovered
attribute map, such as the target type or label, without copying the map per target:

```go
attributes := extquery.AttributesFunc(func(key string) []string {
    switch key {
    case "target.type":
        return []string{target.TargetType}
    case "steadybit.label":
        return []string{target.Label}
    }
    return target.Attributes[key]
})
```

There is no `discovery_kit_api.Target` adapter here on purpose: discovery-kit depends on
extension-kit, so the dependency cannot run the other way. That adapter belongs in discovery-kit.

### Matching semantics worth knowing

These match the platform exactly and are pinned by the shared conformance corpus (below):

- An attribute is a **set of values**. `=` means *any* value equals the given one; `!=` means *no*
  value equals it. An absent key therefore makes `=` false and `!=` true.
- `~` is substring containment; a trailing `*` on any operator makes it case-insensitive.
- `count(key) op n` is **false when the key is absent** — an absent attribute is not a count of
  zero, so `count(k) < 1` does not match a target without `k`.
- An empty query, and a composite with no children, match everything — including an empty `OR`.
- `AND` binds tighter than `OR`; `NOT` tighter still.

## Conformance corpus

`testdata/conformance.json` holds the shared behavioural corpus for the language: `query`,
`attributes`, `matches`. A copy lives in the platform repository and is driven as a table test
there too, so a semantic change landing in one implementation fails the others.

Add a case there — not only to a Go test — whenever you change or clarify matching behaviour.

## The grammar is a copy

`grammar/*.g4` is a **copy** of the platform's source of truth. A drift check in the platform
repository fails if the two diverge.

Never edit the copy to change the language — change it in the platform, copy it back, and
regenerate. Otherwise the two parsers disagree, and a query that validates here gets rejected
there (or worse, the reverse).

To regenerate `internal/gen/` after the grammar changed:

```shell
./extquery/generate.sh
```

It needs Java on the `PATH` and downloads the pinned ANTLR tool jar on first use. The generated
code is checked in and CI fails if it is stale.
