env "local" {
  src = "ent://fluent/schema"
  dev = "docker://postgres/16/dev?search_path=public"
}