resource "local_file" "foo" {
  content  = "new foo!"
  filename = "foo.bar"
}
