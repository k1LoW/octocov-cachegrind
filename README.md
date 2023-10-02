# octocov-cachegrind

Generate [octocov custom metrics JSON](https://github.com/k1LoW/octocov#custom-metrics) from the output of `valgrind --tool=cachegrind`.

## Usage

```console
$ cat cachegrind.out | octocov-cachegrind
```

## Install

**go install:**

```console
$ go install github.com/k1LoW/octocov-cachegrind/cmd/octocov-cachegrind@latest
```

**deb:**

``` console
$ export OCTOCOV_CACHEGRIND_VERSION=X.X.X
$ curl -o octocov-cachegrind.deb -L https://github.com/k1LoW/octocov-cachegrind/releases/download/v$OCTOCOV_CACHEGRIND_VERSION/octocov-cachegrind_$OCTOCOV_CACHEGRIND_VERSION-1_amd64.deb
$ dpkg -i octocov-cachegrind.deb
```

**RPM:**

``` console
$ export OCTOCOV_CACHEGRIND_VERSION=X.X.X
$ yum install https://github.com/k1LoW/octocov-cachegrind/releases/download/v$OCTOCOV_CACHEGRIND_VERSION/octocov-cachegrind_$OCTOCOV_CACHEGRIND_VERSION-1_amd64.rpm
```

**apk:**

``` console
$ export OCTOCOV_CACHEGRIND_VERSION=X.X.X
$ curl -o octocov-cachegrind.apk -L https://github.com/k1LoW/octocov-cachegrind/releases/download/v$OCTOCOV_CACHEGRIND_VERSION/octocov-cachegrind_$OCTOCOV_CACHEGRIND_VERSION-1_amd64.apk
$ apk add octocov-cachegrind.apk
```

**homebrew tap:**

```console
$ brew install k1LoW/tap/octocov-cachegrind
```

**manually:**

Download binary from [releases page](https://github.com/k1LoW/octocov-cachegrind/releases)
