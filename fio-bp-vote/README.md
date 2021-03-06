# fio-bp-vote

This is an automated voting tool that ranks producers based upon criteria (that can be validated on-chain) in the
[BP Code of Conduct](https://developers.fioprotocol.io/fio-chain/bp#code-of-conduct). The goal is to reward producers
that actively contribute to keeping the chain healthy, therefore actions like participating in proposals for setcode,
and performing fee maintenance are strongly favored. Sending requests that should be seen frequently are also rewarded.

Current ranks for mainnet can be viewed [here](https://snap.blockpane.com/proxy.html)

It allows setting the number of producers to vote for, how often to recalculate votes, allows using a linked-auth key
(highly recommended!), and can restrict to only include a list of allowed producers. It will check every minute for
producers missing (full) rounds, and retracts their votes. It will not vote for them again for three times the
vote calculation frequency (3 days by default.)

## Options

```
  -a string
        actor
  -address string
        fio address
  -allowed string
        plaintext file of producers eligible for votes: FIO address, 1 per line
  -dry-run
        don't push transactions, only print what would have been done.
  -h int
        how often (hours) to run (default 24)
  -k string
        wif key
  -n int
        how many (max) producers to vote for (default 30)
  -p string
        permission, if not using 'active'
  -u string
        url for connect
  -v
        verbose logging
```

## Scoring Criteria:

```
participation in msig (propose, approve, exec)              + 10 points each, last 30 days
fee votes (bundlevote, setfeevote, setfeemult)              + 1 points each, last 30 days
maintenance (bpclaim, tpidclaim, burnexpired, computefees)  + 1 point each, last 24 hours

url in producers table is reachable            + 1 point
url has a bp.json or chain json                + 1 point
permissive CORS for bp.json                    + 1 point

Penalties:
----------

average CPU for transactions, last 48 hours    - neg 3 points, per each 1ms over 5ms avg
have not performed bpclaim in last 30 days     - neg 100 points
does not have any existing votes               - neg 200 points
missed round                                   - will not get a vote for next 3 cycles, triggers immediate re-calculation
```

The CPU penalty is admittedly not an objective measurement, but does seem to be effective as slower nodes tend to take
significantly longer to process a transaction, often as much as 10x longer for an underpowered node.

