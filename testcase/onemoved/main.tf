resource "local_file" "moved_foo" {
  content  = "foo!"
  filename = "foo.bar"
}
