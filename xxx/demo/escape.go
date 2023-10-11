package demotest

import (
)

type S struct {}

func Escapemain() {
    var x S
    _ = identity(x)
}

func identity(x S) S {
    return x
}


func Escapemain2() {
    var x S
    y := &x
    _ = *identity2(y)
}

func identity2(z *S) *S {
    return z
}

func Escapemain3() {
  var x S
  _ = *ref(x)
}

func ref(z S) *S {
  return &z
}
