data "composite_schema" "app" {
  # Load the fluent schema first with all tables.
  schema "public" {
    url = "fluent://fluent/schema"
  }
  # Then, load the functional indexes.
  schema "public" {
    url = "file://schema.sql"
  }
}

env "local" {
  src = data.composite_schema.app.url
  dev = "docker://postgres/15/dev?search_path=public"
}