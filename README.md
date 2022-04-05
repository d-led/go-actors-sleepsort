# go-actors-sleepsort

 trying out various golang actors implementations to: [sleepsort](https://rosettacode.org/wiki/Sorting_algorithms/Sleep_sort#Go)

## Input for the Quest

 [language:go actor](https://github.com/search?q=language%3Ago+actor)

## Demo

- [run.sh](run.sh)
- [![Demo](https://github.com/d-led/go-actors-sleepsort/actions/workflows/demo.yml/badge.svg)](https://github.com/d-led/go-actors-sleepsort/actions/workflows/demo.yml)

## Proto Actor Go

- [github.com/asynkron/protoactor-go](https://github.com/asynkron/protoactor-go)
- [protoactor-sleepsort](protoactor-sleepsort)
- details
  - distribution
  - [multi-language](https://github.com/asynkron)

## Ergo

- [github.com/ergo-services/ergo](https://github.com/ergo-services/ergo)
- [ergo-sleepsort](ergo-sleepsort)
- details
  - distribution compatible with Erlang/OTP nodes
  - [sagas](https://github.com/ergo-services/ergo/tree/master/examples/gensaga)

## Molizen

- [github.com/sanposhiho/molizen](https://github.com/sanposhiho/molizen)
  - the first spotted using code generation
- details
  - doesn't have a native scheduled send ('after')
  - uses a code generator for type-safe actor proxies

## Phony

- [github.com/Arceliar/phony](https://github.com/Arceliar/phony)
- details
  - minimalistic, runtime-local implementation
  - embedding provides actor functionality on top of state structs

## Queue

- [Azer0s/quacktors](https://github.com/Azer0s/quacktors)
  - requires [qpmd?](https://github.com/Azer0s/qpmd)
