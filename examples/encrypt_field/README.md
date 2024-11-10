# Encrypted field example using go.dev

### Setup 

Add the `secrets.Keeper` as a dependency to your project and enable the `intercept` feature flag.

```diff
func main() {
+	opts := []flc.Option{
+		flc.Dependency(
+			flc.DependencyType(&secrets.Keeper{}),
+		),
+		flc.FeatureNames("intercept"),
+	}
	if err := flc.Generate("./schema", &gen.Config{}, opts...); err != nil {
		log.Fatal("running fluent codegen:", err)
	}
}
```

### Generate Assets

```console
go generate ./...
```

### Update the schema with secret field.

See `ent/schema/user.go` for full example.

### Run Example

```console
go test
```
