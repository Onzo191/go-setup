env "local" {
  src = data.external_schema.gorm.url
  dev = "docker://postgres/15/dev?search_path=public"
  
  migration {
    dir = "file://migrations"
    format = atlas
    revisions_schema = "public"
  }
}

data "external_schema" "gorm" {
  program = [
    "go", 
    "run", 
    "./cmd/migrate/main.go",
  ]
}