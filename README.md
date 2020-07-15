# envs
Smol package for envs parsing in golang

```go get github.com/jancimertel/envs```

Usage: 
```golang
import github.com/jancimertel/envs

...

type EnvsBundle struct {
    Var1 bool
    Var2 string
    Var3 int
    Var4 float32
}

bundle := EnvsBundle{}
if err := envs.MustParse(&bundle); err == nil {
    fmt.Println("%v, %v, %v, %v", bundle.Var1, bundle.Var2, bundle.Var3, bundle.Var4)
}
```