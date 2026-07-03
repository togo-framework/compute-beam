<!-- togo-brand -->
<p align="center"><img src=".github/assets/togo-mark.svg" width="96" alt="togo" /></p>
<h1 align="center">compute-beam</h1>
<p align="center"><sub>part of the <a href="https://github.com/togo-framework">togo-framework</a></sub></p>

A togo **compute** backend that runs **Apache Beam** pipelines (runs the pipeline
program with `--runner <BEAM_RUNNER>` appended).

```bash
togo install togo-framework/compute-beam
togo provider:use compute beam
togo config:set BEAM_RUNNER DataflowRunner
```

| Key | Default | Purpose |
|---|---|---|
| `BEAM_RUNNER` | `DirectRunner` | Beam runner (Dataflow/Flink/Spark…) |

MIT © fadymondy
