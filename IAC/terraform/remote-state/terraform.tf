terraform {
  backend "remote" {
    organization = "ikbal312"
    workspaces {
      name = "notes-workspace"
    }
  }
}
