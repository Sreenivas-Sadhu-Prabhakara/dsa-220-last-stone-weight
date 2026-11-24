# Last Stone Weight

> NeetCode 150 · Heap — a direct simulation where the data structure is the whole answer: "the two heaviest" on demand is a max-heap's one job. Go, stdlib only, tested.

## Problem

Repeatedly smash the two heaviest stones: equal weights annihilate both, unequal leaves their difference. What weight remains at the end (0 if nothing)?

## Approach

No insight to find — the problem *is* "give me the two largest, repeatedly," and that's a max-heap's defining operation. Initialize via `heap.Init` (O(n), cheaper than n pushes), pop two, push back the difference when nonzero, stop below two stones. The pop order hands the pair already sorted, so `a−b` needs no abs. Worth knowing the contrast: the stones shrink, so this terminates in at most n−1 smashes; and the deceptively similar Last Stone Weight **II** is an entirely different beast (partition DP, dsa-238's family), a classic trap of problem-name adjacency.

**Complexity:** O(n log n) time, O(n) heap.

## Run

    go test ./...   # verify
    go run .        # demo

---

Part of a 275-repo DSA series in Go — Striver's SDE Sheet (dsa-001–191) plus the NeetCode 150 / Blind 75 additions (dsa-192–275), stdlib only, every solution tested. Browse: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
