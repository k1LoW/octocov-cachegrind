# octocov-cachegrind

Generate [octocov custom metrics JSON](https://github.com/k1LoW/octocov#custom-metrics) from the output of `valgrind --tool=cachegrind`.

## Usage

```console
$ cat cachegrind.out | octocov-cachegrind
```
