# internal/npcGen — Code Quality, Test Coverage & Scalability TODO

---

## 1. Architecture & Design

- [ ] **Extract named sub-structs from `NPCBase`** — The anonymous inline structs (`OCEAN`, `MICE`, `CS`, `REI`, `Enneagram`, etc.) inside `npcBase.go` are not reusable, not independently testable, and cannot be passed around without dragging the entire `NPCBase`. Define each as a named, exported type (e.g. `OCEANData`, `EnneagramData`, `CSData`) and compose them into `NPCBase`. This is the most important structural change in the file.

- [ ] **Introduce a `Generator` interface and dependency-injectable struct** — Every `Create*` function currently re-opens a DB connection and re-reads a config file independently. There is no shared state, no coordination, and no way to inject test doubles. Define a `Generator` struct that holds a `*sql.DB` and config, and make all `Create*` functions methods on it. `CreateNPC` becomes the orchestrator that calls its own methods rather than free functions.

- [ ] **Define a `Repository` interface for data access** — The DB query logic is entangled directly in generation logic throughout `createCSData.go`, `createMICE.go`, `createOCEAN.go`, `createREIData.go`, `createEnneagram.go`, and `readNPC.go`. Extract a `Repository` interface (e.g. `NPCRepository`) with methods like `GetCSData()`, `GetOCEANData()`, `GetEnneagramByID()`, `GetNPC()`, `SaveNPC()`. This decouples generation logic from persistence and makes unit testing possible without a live DB.

- [ ] **Stop treating `NPCBase` as a write-through scratchpad** — Every `Create*` function mutates `npc_object *NPCBase` directly. This creates implicit ordering dependencies (`CreateMICEData` and `CreateOCEANData` must run after `CreateCSData` because they read `npc_object.CS.Coords`). Make these dependencies explicit: either pass `CSData` as a parameter to functions that need it, or document the required call order in a single orchestrator and enforce it structurally.

- [ ] **Delete or finish the commented-out code in `createEnneagram.go`** — There are ~80 lines of commented-out functions. Either restore them as part of the implementation or delete them. Dead code in version control is noise; that's what git history is for.

---

## 2. Database & Connection Management

- [ ] **Open the DB connection once per generation, not once per trait** — `CreateCSData`, `CreateREIData`, `CreateOCEANData`, `CreateEnneagram`, `CreateMICEData`, and `ReadNPC` each independently call `config.ReadConfig(...)` and `db.ConnectDatabase(...)`. A single NPC generation opens and closes the DB at least 5 times. Fix this by opening one connection at the `CreateNPC` level and passing it (or the `Generator` struct) down.

- [ ] **Read config once at startup, not per function call** — `config.ReadConfig("configs/dbconf.yml")` is called with a hardcoded relative path inside every `Create*` function. This is fragile (path depends on working directory), slow, and untestable. Config should be read once at application startup and injected.

- [ ] **Replace `log.Fatal` inside library functions** — `createCSData.go`, `createEnneagram.go`, `createMICE.go`, `createOCEAN.go`, `createREIData.go`, and `readNPC.go` all call `log.Fatal(err)` on config or DB errors inside functions that return `error`. `log.Fatal` calls `os.Exit(1)` — it bypasses all deferred cleanup, makes the caller unable to handle the error, and makes the functions untestable. Replace every `log.Fatal` with `return err` (or a wrapped error).

- [ ] **Fix SQL injection vulnerabilities in `readNPC.go` and `createEnneagram.go`** — `fmt.Sprintf` is used to build every query string with user-controlled or runtime values interpolated directly (e.g. `"SELECT * FROM ... WHERE id='%s'"`, `fmt.Sprintf("SELECT * FROM generator.enneagram WHERE id='%d'", enneagram_id)`). Use parameterized queries (`db.QueryRow("SELECT * FROM ... WHERE id = $1", id)`) for all queries.

- [ ] **Remove the double `defer rows.Close()` in `getCSData`, `getMICEData`, `getREIData`, and `getOCEANData`** — Each of these functions calls `defer rows.Close()` twice. The second call is a no-op but signals that the author wasn't sure the first one fired. Remove the duplicate.

- [ ] **Replace `SELECT *` with explicit column lists in all queries** — `SELECT *` couples the query to the DB schema column order. If a column is added, removed, or reordered, `Scan()` silently breaks or panics. List columns explicitly.

---

## 3. Correctness Bugs

- [x] **Fix `makeBMI` boundary logic in `createNPCBody.go`** — The conditions use `||` where `&&` is needed:
  - `18.5 < BMI || BMI <= 24.9` should be `18.5 < BMI && BMI <= 24.9`
  - `25 < BMI || BMI <= 29.29` should be `25 < BMI && BMI <= 29.29`
  As written, every BMI value matches the second condition because `||` short-circuits incorrectly. This means the BMI bucket is almost always wrong.

- [x] **Fix `CreateCSData` — coords set after they are used** — `coordsToSelection` is called with `npc_object.CS.Coords` before those coords are populated, meaning `Coords` is always `[0, 0]` on first generation. The coord generation (`utilities.RandomRange(...)`) happens *after* the selection. Swap the order.

- [x] **Fix `CreateOrientationType` — writes UUID to wrong field** — In `createNPCSexOri.go`, the UUID is written to `npc_object.NPCType.UUID` instead of `npc_object.SexualOrientationType.UUID`. This silently corrupts the NPC type UUID on every generation.

- [x] **Fix `BodTransition` copy-paste bug in `bodyTypeEnum.go`** — The cases for `"LARGE"`, `"REEDY"`, `"SOFT"`, and `"PLUMP"` all return `FAT` instead of their correct enum values. This is clearly a copy-paste error.

- [x] **Fix `SelectEnneagram` range in `createEnneagram.go`** — `rand.Intn(8) + 1` produces values 1–8. The Enneagram has 9 types. Type 9 is never generated. Change to `rand.Intn(9) + 1`.

- [x] **Fix `CreateEnneaLODLevel` range** — Same issue: `rand.Intn(8) + 1` produces 1–8 but `enn_LOD` is a `[9]string` (indices 0–8). Index 0 is never selected, and the intent of the LOD system is unclear. Decide on 0-indexed or 1-indexed and be consistent.

- [x] **Fix `CreatePronouns` random selection** — `rand.Intn(len(enums.Pronouns)) + 1` can produce index 4 when `Pronouns` has 3 entries (indices 1–3), causing a map miss and returning a nil slice. Clamp to the actual map size.

- [ ] **Fix `coordsToSelection` missing case** — When both coords are exactly `0`, the function returns `0` via the `<= 0 && <= 0` branch. When one is exactly `0` it could fall into multiple branches depending on order. Define explicit behavior for the zero boundary.

---

## 4. Code Quality & Maintainability

- [ ] **Adopt Go naming conventions throughout** — Snake_case (`npc_object`, `q_str`, `cs_data`, `r_val`, `bod_state`, `gen_state`) is not idiomatic Go. Use camelCase (`npcObject`, `qStr`, `csData`, `rVal`, `bodState`, `genState`). This applies to every file in the package.

- [ ] **Replace magic numbers with named constants** — `rand.Intn(3) + 1` in `createNPCSex.go`, `health_min := 5`, `health_max := 7`, `lbs_min := 110`, `lbs_max := 250` in `createNPCBody.go`, the `110.0` loop bound and `10.0` step in `createOCEAN.go`, etc. Every unexplained literal should be a named constant with a comment explaining its origin.

- [ ] **Remove the `ToJSON` method explosion on `NPCBase`** — There are 13 near-identical `ToJSON` methods in `npcBase.go` (`DataToJSON`, `NameToJSON`, `OCEANToJSON`, `MICEToJSON`, etc.). This is not scalable — every new sub-struct requires a new method. Replace with a single generic serialization approach or expose the sub-structs directly and let the caller marshal what they need.

- [ ] **Remove `Transition` functions from enums or fix them** — `BodTransition`, `GenTransition`, `NPCTransition`, `OriTransition`, `SexTransition` all perform a string-to-enum lookup via a switch. None of them are called anywhere in the codebase. If the intent is to allow string-based lookup, replace with a `map[string]EnumType` lookup function. If they're not needed, delete them.

- [ ] **Standardize error handling — no mixing of `log.Fatal`, `log.Fatalf`, and `return err`** — The codebase mixes all three inconsistently. Library-level code should never call `log.Fatal`. Pick one strategy: return errors and let the top-level caller decide how to log and exit.

- [ ] **Remove the `NpcTypeEnum` `panic` in `NPCTransition`** — `npcTypeEnum.go` panics on an unknown state while every other enum just logs. Inconsistent. All enum transitions should handle the default case the same way.

- [ ] **Make the `OCEAN` traits and descriptions data-driven, not hardcoded** — In `createOCEAN.go`, the traits lists and description strings are hardcoded inline in the generation function. These belong in the database or in an embedded data file, not in application logic.

- [ ] **Clarify `MakeSizeImperial` — height and weight are independent of each other** — A person's weight is generated completely independently of their height (same `lbs_min`/`lbs_max` range regardless of whether they're 2ft or 7ft tall). This will generate nonsensical BMI values routinely. Correlate the weight range to the height range.

- [ ] **Replace `npc_object.NPCAppearance.Height_In` field — it stores total inches, not inches component** — The field is named `Height_In` but `MakeSizeImperial` stores `(ft * 12) + inch` in it, i.e. total inches. `MakeSizeMetric` then passes it to `imperialToMetric` as `inches`, which is correct but the field name is misleading. Rename to `Height_TotalInches` or store feet and inches separately.

---

## 5. Test Coverage

- [ ] **Add unit tests for all enum packages** — `bodyTypeEnum.go`, `genderTypeEnum.go`, `npcTypeEnum.go`, `orientationTypeEnum.go`, `sexTypeEnum.go` contain pure logic with no external dependencies and zero test coverage. These should be the easiest tests to write and should be done first.

- [ ] **Add unit tests for `createUUID.go`** — Verify format, uniqueness across N iterations, version bits (`[6] & 0xf0 == 0x40`), and variant bits.

- [ ] **Add unit tests for `makeBMI`** — This is a pure function with clear boundary conditions. After fixing the `||`/`&&` bug, write table-driven tests covering every BMI bracket boundary.

- [ ] **Add unit tests for `coordsToSelection`** — Table-driven tests for all four quadrants, and explicit tests for the boundary cases where one or both coordinates are exactly `0`.

- [ ] **Add unit tests for `CreateOCEANText`** — This is pure logic (no DB dependency) and is the most complex function in the package. It needs tests verifying that each OCEAN dimension produces the correct trait slice, attribute string, and adjective list for representative degree values.

- [ ] **Add unit tests for `imperialToMetric`** — Pure math function, trivial to test.

- [ ] **Add integration tests for all `Create*` functions using a test DB** — Use SQLite in-memory (`:memory:`) with the same schema as the production DB seeded with minimal fixture data. Each `Create*` function should have at least one integration test that verifies it populates the expected fields without error.

- [ ] **Add a test for the full `CreateNPC` pipeline** — After fixing the DB connection architecture, write an integration test that runs the full pipeline and asserts that the returned `NPCBase` has no zero-value UUIDs, all required fields populated, and no contradictory data (e.g. UUID collision across sub-structs).

- [ ] **Add tests for `ReadNPC`** — Verify round-trip: generate an NPC, write it to a test DB, read it back, assert equality.

- [ ] **Add regression tests for the bugs identified in Section 3** — Each correctness bug fixed should have a corresponding test that would have caught it. Add them before or alongside the fix.

---

## 6. Scalability & Future-Proofing

- [ ] **Make `CreateNPC` concurrent where safe** — `CreateCSData` must run first (others depend on `CS.Coords`). After that, `CreateREIData`, `CreateOCEANData`, `CreateEnneagram`, `CreateMICEData`, and all the type/body/sex/gender functions are independent of each other and can run concurrently. Use a `sync.WaitGroup` and `errgroup` to parallelize the independent stage once the DB connection is shared (not per-function).

- [ ] **Replace the OCEAN degree loop with a lookup table or map** — The nested loop in `CreateOCEANText` that iterates `j` from `0` to `110` in steps of `10` to find a bucket is an O(n) linear scan over a fixed domain. Replace with a direct index calculation: `bucket := int(degree[i] / 10.0)` clamped to bounds.

- [ ] **Design a `GenerationOptions` struct for user overrides** — The two `TODO(wholesomeow)` comments in `CreateNPC` about "user-driven configuration overrides" are correct instincts. Define a `GenerationOptions` struct now (even if most fields are ignored initially) so the function signature accommodates it: `CreateNPC(opts GenerationOptions) (NPCBase, error)`. Retrofitting this later will break every caller.

- [ ] **Plan for `NPCBase` versioning** — The comments in `npcBase.go` (`// v1.1`, `// v2.0`, `// v2.1`) indicate planned schema evolution. Before those fields are added, decide how serialized NPCs saved to the DB will be migrated when the struct changes. Document the migration strategy now.

- [ ] **Remove `github.com/google/uuid` and keep your own implementation, or do the opposite** — `go.mod` lists `google/uuid` as an indirect dependency while `createUUID.go` reimplements UUID v4 from scratch with a comment saying it was "appropriated" from that library. Pick one: use the library directly, or remove the dependency and own the implementation fully. The current state is the worst of both worlds.