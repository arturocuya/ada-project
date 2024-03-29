package fdct

import (
  consts "./../../consts"
)
/*
var C = float64[8] = {
  1.0,
  0.9807852804032304,
  0.9238795325112867,
  0.8314696123025452,
  0.7071067811865476,
  0.5555702330196023,
  0.38268343236508984,
  0.19509032201612833,
}
*/

var S = [8]float64{
  0.35355339059327373,
  0.2548977895520796,
  0.2705980500730985,
  0.30067244346752264,
  0.35355339059327373,
  0.4499881115682078,
  0.6532814824381882,
  1.2814577238707527,
}

var A = [8]float64{
  0.0,
  0.7071067811865476,
  0.5411961001461969,
  0.7071067811865476,
  1.3065629648763766,
  0.38268343236508984,
}
// :=
func Fdct(block *consts.Block){
  b := block

  v0 := float64(b[0] + b[7])
  v1 := float64(b[1] + b[6])
  v2 := float64(b[2] + b[5])
  v3 := float64(b[3] + b[4])
  v4 := float64(b[3] - b[4])
  v5 := float64(b[2] - b[5])
  v6 := float64(b[1] - b[6])
  v7 := float64(b[0] - b[7])

  v8 := v0 + v3
  v9 := v1 + v2
  v10 := v1 - v2
  v11 := v0 - v3
  v12 := -v4 - v5
  v13 := (v5 + v6) * A[3]
  v14 := v6 + v7

  v15 := v8 + v9
  v16 := v8 - v9
  v17 := (v10 + v11) * A[1]
  v18 := (v12 + v14) * A[5]

  v19 := -v12 * A[2] - v18
  v20 := v14 * A[4] - v18

  v21 := v17 + v11
  v22 := v11 - v17
  v23 := v13 + v7
  v24 := v7 - v13

  v25 := v19 + v24
  v26 := v23 + v20
  v27 := v23 - v20
  v28 := v24 - v19

  block[0] = int32(S[0] * v15)
  block[1] = int32(S[1] * v26)
  block[2] = int32(S[2] * v21)
  block[3] = int32(S[3] * v28)
  block[4] = int32(S[4] * v16)
  block[5] = int32(S[5] * v25)
  block[6] = int32(S[6] * v22)
  block[7] = int32(S[7] * v27)
}

func Idct(block *consts.Block) {
  b := block

  v15 := float64(b[0]) / S[0]
  v26 := float64(b[1]) / S[1]
  v21 := float64(b[2]) / S[2]
  v28 := float64(b[3]) / S[3]
  v16 := float64(b[4]) / S[4]
  v25 := float64(b[5]) / S[5]
  v22 := float64(b[6]) / S[6]
  v27 := float64(b[7]) / S[7]

  v19 := (v25 - v28) / 2.0
  v20 := (v26 - v27) / 2.0
  v23 := (v26 + v27) / 2.0
  v24 := (v25 + v28) / 2.0

  v7  := (v23 + v24) / 2.0
  v11 := (v21 + v22) / 2.0
  v13 := (v23 - v24) / 2.0
  v17 := (v21 - v22) / 2.0

  v8 := (v15 + v16) / 2.0
  v9 := (v15 - v16) / 2.0

  v18 := (v19 - v20) * A[5]
  v12 := (v19 * A[4] - v18) / (A[2] * A[5] - A[2] * A[4] - A[4] * A[5])
  v14 := (v18 - v20 * A[2]) / (A[2] * A[5] - A[2] * A[4] - A[4] * A[5])

  v6 := v14 - v7
  v5 := v13 / A[3] - v6
  v4 := -v5 - v12
  v10 := v17 / A[1] - v11

  v0 := (v8 + v11) / 2.0
  v1 := (v9 + v10) / 2.0
  v2 := (v9 - v10) / 2.0
  v3 := (v8 - v11) / 2.0

  block[0] = int32((v0 + v7) / 2.0)
  block[1] = int32((v1 + v6) / 2.0)
  block[2] = int32((v2 + v5) / 2.0)
  block[3] = int32((v3 + v4) / 2.0)
  block[4] = int32((v3 - v4) / 2.0)
  block[5] = int32((v2 - v5) / 2.0)
  block[6] = int32((v1 - v6) / 2.0)
  block[7] = int32((v0 - v7) / 2.0)
}
