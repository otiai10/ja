# `ja`: Yet another Japanese sentence processor

[![Go](https://github.com/otiai10/ja/actions/workflows/go.yaml/badge.svg)](https://github.com/otiai10/ja/actions/workflows/go.yaml)
[![codecov](https://codecov.io/gh/otiai10/ja/branch/main/graph/badge.svg?token=V75ubSechE)](https://codecov.io/gh/otiai10/ja)

```go
var s = `沖縄本島地方と先島諸島では、２１日夜のはじめ頃にかけて、台風の接近と大潮の時期が重なるため潮位が高くなる見込みです。海岸や河口付近の低地での高潮による浸水や冠水に注意してください。`


lines := ja.Linebreak(s, Within(60))
```
