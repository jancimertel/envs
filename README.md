# envs
Smol package for envs parsing in golang

```go get github.com/jancimertel/envs```

Usage: 
```golang
import "github.com/jancimertel/envs"
```

...

```golang
// search env vars named VAR1 ...
type EnvsBundle struct {
    VAR1 bool
    VAR2 string
    VAR3 int
    VAR4 float32
}

bundle := EnvsBundle{}
if err := envs.MustHave(&bundle); err == nil {
    fmt.Println(fmt.Sprintf("%v, %v, %v, %v", bundle.VAR1, bundle.VAR2, bundle.VAR3, bundle.VAR4))
}
```

```golang
// search env vars with tags TAG1 ...
type EnvsBundle struct {
    VAR1 bool    `envs:"TAG1"`
    VAR2 string  `envs:"TAG2"`
    VAR3 int     `envs:"TAG3"`
    VAR4 float32 `envs:"TAG4"`
}

bundle := EnvsBundle{}
if err := envs.MustHave(&bundle); err == nil {
    fmt.Println(fmt.Sprintf("%v, %v, %v, %v", bundle.VAR1, bundle.VAR2, bundle.VAR3, bundle.VAR4))
}
```